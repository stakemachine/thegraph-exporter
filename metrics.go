package main

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
	"github.com/stakemachine/graph-indexer-cli/utils"
	"github.com/tidwall/gjson"
)

func metricExists(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (service Service) clearOldMetrics(oldList, newList []string, metricPrefix string) {
	for _, om := range oldList {
		if !metricExists(om, newList) {
			if strings.HasPrefix(om, metricPrefix) {
				if service.Metrics.UnregisterMetric(om) {
					log.Debug().Msgf("unregistered metric: %s", om)
				} else {
					log.Debug().Msgf("failed to unregister metric: %s", om)
				}
			}
		}
	}
}

// GetAndSetActiveAllocationsRewardsMetrics get rewards and set metrics
func (service Service) GetAndSetActiveAllocationsRewardsMetrics() error {
	allos, err := service.GqlMainnet.GetAllocations()
	if err != nil {
		return err
	}
	totalActiveAllocations := len(allos)
	totalActiveRewards := big.NewInt(0)

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, allo := range allos {
		rewards, err := service.Eth.GetRewards(allo.ID)
		if err != nil {
			log.Error().Err(err).Msg("GetRewards error")
			continue // fmt.Println(err)
		}
		grtDec := big.NewInt(1000000000000000000)
		normalizedReward := rewards.ToInt().Quo(rewards.ToInt(), grtDec)
		totalActiveRewards.Add(totalActiveRewards, normalizedReward)
		// fmt.Printf("#%d Indexer: %s allocation: %s reward: %d GRT\n", i, allo.Indexer, allo.ID, normalizedReward)
		subgraphDeploymentID, err := utils.SubgraphHexToHash(allo.SubgraphDeployment.ID)
		if err != nil {
			return err
		}
		s := fmt.Sprintf(`thegraph_allocations_active{allocation="%s",indexer="%s",ipfs_hash="%s"}`, allo.ID, allo.Indexer.ID, subgraphDeploymentID)
		service.Metrics.GetOrCreateCounter(s).Set(normalizedReward.Uint64())
		metricsList = append(metricsList, s)
	}

	log.Info().Msgf("Total active not requested rewards: %d GRT", totalActiveRewards)
	totalActiveAllocationsMetric := "thegraph_allocations_active_count"
	log.Info().Msgf("Number of active allocations: %d", totalActiveAllocations)
	service.Metrics.GetOrCreateCounter(totalActiveAllocationsMetric).Set(uint64(totalActiveAllocations))
	metricsList = append(metricsList, totalActiveAllocationsMetric)
	service.clearOldMetrics(oldList, metricsList, "thegraph_allocations_active")
	return nil
}

// GetAndSetVestingBalanceMetrics get vesting balances and convert to metrics
func (service Service) GetAndSetVestingBalanceMetrics() error {
	tokenLockWallets, err := service.GqlTokenLock.GetAllTokenLockWallets()
	if err != nil {
		return err
	}
	var addresses []string
	for _, wallet := range tokenLockWallets {
		// log.Debug().Msgf("wallet: %s\trevokable: %s", wallet.ID, wallet.Revocable)
		if wallet.Revocable == "Disabled" {
			addresses = append(addresses, wallet.ID)
		}
	}

	for _, address := range addresses {
		balance, err := service.Eth.GetTokenBalance("0xc944e90c64b2c07662a292be6244bdf05cda44a7", address)
		if err != nil {
			log.Error().Err(err).Msg("GetTokenBalance error")
			continue
		}
		intbalance, err := hexutil.DecodeBig(balance.String())
		if err != nil {
			log.Error().Err(err).Msg("Decode GetTokenBalance error")
			continue
		}
		dec := big.NewInt(1000000000000000000)
		nb := new(big.Int).Div(intbalance, dec)
		if len(nb.Bits()) == 0 {
			continue
		}
		vestedBalanceMetric := fmt.Sprintf(`thegraph_vested_tokens{address="%s"}`, address)
		log.Trace().Msgf("%s %d", vestedBalanceMetric, nb)
		service.Metrics.GetOrCreateGauge(vestedBalanceMetric, func() float64 {
			return float64(nb.Int64())
		})
	}

	return nil
}

// GetAndSetGraphNetworkMetrics requests graphql endpoint for data and split response to metrics
func (service Service) GetAndSetGraphNetworkMetrics() error {
	graphNetwork, err := service.GqlMainnet.GetGraphNetwork()
	if err != nil {
		return err
	}
	contractsMetric := fmt.Sprintf(`thegraph_graphnetwork_contracts{controller="%s", token="%s", epochmanager="%s", curation="%s",staking="%s", disputemanager="%s",gns="%s",serviceregistry="%s",rewardsmanager="%s", governor="%s",pauseguardian="%s",subgraphavailabilityoracle="%s",arbitrator="%s"}`, graphNetwork.Controller, graphNetwork.GraphToken, graphNetwork.EpochManager, graphNetwork.Curation, graphNetwork.Staking, graphNetwork.DisputeManager, graphNetwork.Gns, graphNetwork.ServiceRegistry, graphNetwork.RewardsManager, graphNetwork.Governor, graphNetwork.PauseGuardian, graphNetwork.SubgraphAvailabilityOracle, graphNetwork.Arbitrator)
	service.Metrics.GetOrCreateCounter(contractsMetric).Set(0)
	// graphnetwork_slashers
	// graphnetwork_asset_holders
	// service.Metrics.GetOrCreateCounter(`graphnetwork_`).Set(uint64(graphNetwork))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_delegationratio`).Set(uint64(graphNetwork.DelegationRatio))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_curationpercentage`).Set(uint64(graphNetwork.CurationPercentage))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_protocolfeepercentage`).Set(uint64(graphNetwork.ProtocolFeePercentage))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_channeldisputeepochs`).Set(uint64(graphNetwork.ChannelDisputeEpochs))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_maxallocationepochs`).Set(uint64(graphNetwork.MaxAllocationEpochs))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_thawingperiod`).Set(uint64(graphNetwork.ThawingPeriod))
	totalTokensStaked, err := stringToGRT(graphNetwork.TotalTokensStaked)
	if err != nil {
		return err
	}
	//  service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaltokensstaked`).Set(totalTokensStaked.Uint64())
	// #nosec G101
	totalTokensStakedMetric := "thegraph_graphnetwork_totaltokensstaked"
	service.Metrics.GetOrCreateGauge(totalTokensStakedMetric, func() float64 {
		return float64(totalTokensStaked.Int64())
	})

	totalTokensClaimable, err := stringToGRT(graphNetwork.TotalTokensClaimable)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaltokensclaimable`).Set(totalTokensClaimable.Uint64())
	totalUnstakedTokensLocked, err := stringToGRT(graphNetwork.TotalUnstakedTokensLocked)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalunstakedtokenslocked`).Set(totalUnstakedTokensLocked.Uint64())
	totalTokensAllocated, err := stringToGRT(graphNetwork.TotalTokensAllocated)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaltokensallocated`).Set(totalTokensAllocated.Uint64())

	totalDelegatedTokens, err := stringToGRT(graphNetwork.TotalDelegatedTokens)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaldelegatedtokens`).Set(totalDelegatedTokens.Uint64())

	totalQueryFees, err := stringToGRT(graphNetwork.TotalQueryFees)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalqueryfees`).Set(totalQueryFees.Uint64())
	totalIndexerQueryFeesCollected, err := stringToGRT(graphNetwork.TotalIndexerQueryFeesCollected)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalindexerqueryfeescollected`).Set(totalIndexerQueryFeesCollected.Uint64())
	totalIndexerQueryFeeRebates, err := stringToGRT(graphNetwork.TotalIndexerQueryFeeRebates)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalindexerqueryfeerebates`).Set(totalIndexerQueryFeeRebates.Uint64())
	totalDelegatorQueryFeeRebates, err := stringToGRT(graphNetwork.TotalDelegatorQueryFeeRebates)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaldelegatorqueryfeerebates`).Set(totalDelegatorQueryFeeRebates.Uint64())
	totalCuratorQueryFees, err := stringToGRT(graphNetwork.TotalCuratorQueryFees)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalcuratorqueryfees`).Set(totalCuratorQueryFees.Uint64())
	totalTaxedQueryFees, err := stringToGRT(graphNetwork.TotalTaxedQueryFees)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaltaxedqueryfees`).Set(totalTaxedQueryFees.Uint64())
	totalUnclaimedQueryFeeRebates, err := stringToGRT(graphNetwork.TotalUnclaimedQueryFeeRebates)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalunclaimedqueryfeerebates`).Set(totalUnclaimedQueryFeeRebates.Uint64())
	totalIndexingRewards, err := stringToGRT(graphNetwork.TotalIndexingRewards)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalindexingrewards`).Set(totalIndexingRewards.Uint64())
	totalIndexingDelegatorRewards, err := stringToGRT(graphNetwork.TotalIndexingDelegatorRewards)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalindexingdelegatorrewards`).Set(totalIndexingDelegatorRewards.Uint64())
	totalIndexingIndexerRewards, err := stringToGRT(graphNetwork.TotalIndexingIndexerRewards)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalindexingindexerrewards`).Set(totalIndexingIndexerRewards.Uint64())
	totalSupply, err := stringToGRT(graphNetwork.TotalSupply)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalsupply`).Set(totalSupply.Uint64())
	totalTokensSignalled, err := stringToGRT(graphNetwork.TotalTokensSignalled)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaltokensignalled`).Set(totalTokensSignalled.Uint64())
	totalGRTMinted, err := stringToGRT(graphNetwork.TotalGRTMinted)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalgrtminted`).Set(totalGRTMinted.Uint64())
	totalGRTBurned, err := stringToGRT(graphNetwork.TotalGRTBurned)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalgrtburned`).Set(totalGRTBurned.Uint64())
	networkGRTIssuance, err := stringToGRT(graphNetwork.NetworkGRTIssuance)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_networkgrtissuance`).Set(networkGRTIssuance.Uint64())

	totalDelegatedTokensTransferredToL2, err := stringToGRT(graphNetwork.TotalDelegatedTokensTransferredToL2)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totaldelegatedtokenstransferredtol2`).Set(totalDelegatedTokensTransferredToL2.Uint64())

	totalTokensStakedTransferredToL2, err := stringToGRT(graphNetwork.TotalTokensStakedTransferredToL2)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalstakedtokenstransferredtol2`).Set(totalTokensStakedTransferredToL2.Uint64())

	totalSignalledTokensTransferredToL2, err := stringToGRT(graphNetwork.TotalSignalledTokensTransferredToL2)
	if err != nil {
		return err
	}
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_totalsignalledtokenstransferredtol2`).Set(totalSignalledTokensTransferredToL2.Uint64())

	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_currentepoch`).Set(uint64(graphNetwork.CurrentEpoch))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_indexercount`).Set(uint64(graphNetwork.IndexerCount))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_stakedindexerscount`).Set(uint64(graphNetwork.StakedIndexersCount))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_delegatorcount`).Set(uint64(graphNetwork.DelegatorCount))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_curatorcount`).Set(uint64(graphNetwork.CuratorCount))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_subgraphcount`).Set(uint64(graphNetwork.SubgraphCount))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_subgraphdeploymentcount`).Set(uint64(graphNetwork.SubgraphDeploymentCount))
	service.Metrics.GetOrCreateCounter(`thegraph_graphnetwork_epochcount`).Set(uint64(graphNetwork.EpochCount))
	return nil
}

// GetAndSetIndexerMetrics creates metrics from graphql response
func (service Service) GetAndSetIndexerMetrics() error {
	indexers, err := service.GqlMainnet.GetIndexers()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, indexer := range indexers {

		indexerInfoMetric := fmt.Sprintf(`thegraph_indexers_info{indexer="%s",geohash="%s"}`, indexer.ID, indexer.GeoHash)
		service.Metrics.GetOrCreateCounter(indexerInfoMetric).Set(1)

		metricsList = append(metricsList, indexerInfoMetric)

		stakedTokens, err := stringToGRT(indexer.StakedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		stakedTokensMetric := fmt.Sprintf(`thegraph_indexer_tokens_staked{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(stakedTokensMetric).Set(stakedTokens.Uint64())

		metricsList = append(metricsList, stakedTokensMetric)

		unstakedTokens, err := stringToGRT(indexer.UnstakedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		unstakedTokensMetric := fmt.Sprintf(`thegraph_indexer_tokens_unstaked{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(unstakedTokensMetric).Set(unstakedTokens.Uint64())

		metricsList = append(metricsList, unstakedTokensMetric)

		allocatedTokens, err := stringToGRT(indexer.AllocatedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		allocatedTokensMetric := fmt.Sprintf(`thegraph_indexer_tokens_allocated{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(allocatedTokensMetric).Set(allocatedTokens.Uint64())

		metricsList = append(metricsList, allocatedTokensMetric)

		lockedTokens, err := stringToGRT(indexer.LockedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		lockedTokensMetric := fmt.Sprintf(`thegraph_indexer_tokens_locked{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(lockedTokensMetric).Set(lockedTokens.Uint64())

		metricsList = append(metricsList, lockedTokensMetric)

		delegatedTokens, err := stringToGRT(indexer.DelegatedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		delegatedTokensMetric := fmt.Sprintf(`thegraph_indexer_tokens_delegated{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(delegatedTokensMetric).Set(delegatedTokens.Uint64())

		metricsList = append(metricsList, delegatedTokensMetric)

		allocationsActiveCountMetric := fmt.Sprintf(`thegraph_indexer_allocations_active_count{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(allocationsActiveCountMetric).Set(uint64(indexer.AllocationCount))

		metricsList = append(metricsList, allocationsActiveCountMetric)

		allocationsTotalCount, err := stringToGRT(indexer.TotalAllocationCount)
		if err != nil {
			log.Error().Err(err)
		}
		allocationsTotalCountMetric := fmt.Sprintf(`thegraph_indexer_allocations_total_count{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(allocationsTotalCountMetric).Set(allocationsTotalCount.Uint64())

		metricsList = append(metricsList, allocationsTotalCountMetric)

		queryFeesCollected, err := stringToGRT(indexer.QueryFeesCollected)
		if err != nil {
			log.Error().Err(err)
		}
		queryFeesCollectedMetric := fmt.Sprintf(`thegraph_indexer_fees_query_collected{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(queryFeesCollectedMetric).Set(queryFeesCollected.Uint64())

		metricsList = append(metricsList, queryFeesCollectedMetric)

		queryFeesRebates, err := stringToGRT(indexer.QueryFeeRebates)
		if err != nil {
			log.Error().Err(err)
		}
		queryFeesRebatesMetric := fmt.Sprintf(`thegraph_indexer_fees_query_rebates{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(queryFeesRebatesMetric).Set(queryFeesRebates.Uint64())

		metricsList = append(metricsList, queryFeesRebatesMetric)

		delegatorQueryFees, err := stringToGRT(indexer.DelegatorQueryFees)
		if err != nil {
			log.Error().Err(err)
		}
		delegatorQueryFeesMetric := fmt.Sprintf(`thegraph_indexer_fees_query_delegator{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(delegatorQueryFeesMetric).Set(delegatorQueryFees.Uint64())

		metricsList = append(metricsList, delegatorQueryFeesMetric)

		rewardsEarned, err := stringToGRT(indexer.RewardsEarned)
		if err != nil {
			log.Error().Err(err)
		}
		rewardsEarnedMetric := fmt.Sprintf(`thegraph_indexer_rewards_earned{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(rewardsEarnedMetric).Set(rewardsEarned.Uint64())

		metricsList = append(metricsList, rewardsEarnedMetric)

		indexerIndexingRewards, err := stringToGRT(indexer.IndexerIndexingRewards)
		if err != nil {
			log.Error().Err(err)
		}
		indexerIndexingRewardsMetric := fmt.Sprintf(`thegraph_indexer_rewards_indexing{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(indexerIndexingRewardsMetric).Set(indexerIndexingRewards.Uint64())

		metricsList = append(metricsList, indexerIndexingRewardsMetric)

		delegatorIndexingRewards, err := stringToGRT(indexer.DelegatorIndexingRewards)
		if err != nil {
			log.Error().Err(err)
		}
		delegatorIndexingRewardsMetric := fmt.Sprintf(`thegraph_indexer_rewards_delegator_indexing{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(delegatorIndexingRewardsMetric).Set(delegatorIndexingRewards.Uint64())

		metricsList = append(metricsList, delegatorIndexingRewardsMetric)

		delegatedCapacity, err := stringToGRT(indexer.DelegatedCapacity)
		if err != nil {
			log.Error().Err(err)
		}
		delegatedCapacityMetric := fmt.Sprintf(`thegraph_indexer_capacity_delegated{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(delegatedCapacityMetric).Set(delegatedCapacity.Uint64())

		metricsList = append(metricsList, delegatedCapacityMetric)

		tokenCapacity, err := stringToGRT(indexer.TokenCapacity)
		if err != nil {
			log.Error().Err(err)
		}
		tokenCapacityMetric := fmt.Sprintf(`thegraph_indexer_capacity_token{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(tokenCapacityMetric).Set(tokenCapacity.Uint64())

		metricsList = append(metricsList, tokenCapacityMetric)

		availableStake, err := stringToGRT(indexer.AvailableStake)
		if err != nil {
			log.Error().Err(err)
		}
		availableStakeMetric := fmt.Sprintf(`thegraph_indexer_available_stake{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(availableStakeMetric).Set(availableStake.Uint64())

		metricsList = append(metricsList, availableStakeMetric)

		indexingRewardCutMetric := fmt.Sprintf(`thegraph_indexer_fees_rewards_cut{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(indexingRewardCutMetric).Set(uint64(indexer.IndexingRewardCut))

		metricsList = append(metricsList, indexingRewardCutMetric)

		queryFeeCutMetric := fmt.Sprintf(`thegraph_indexer_fees_query_cut{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(queryFeeCutMetric).Set(uint64(indexer.QueryFeeCut))

		metricsList = append(metricsList, queryFeeCutMetric)

		forcedClosuresMetric := fmt.Sprintf(`thegraph_indexer_forced_closures{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateCounter(forcedClosuresMetric).Set(uint64(indexer.ForcedClosures))

		metricsList = append(metricsList, forcedClosuresMetric)

		indexerEffectiveCut, err := strconv.ParseFloat(indexer.IndexingRewardEffectiveCut, 64)
		if err != nil {
			log.Error().Err(err)
		}
		indexerEffectiveCutMetric := fmt.Sprintf(`thegraph_indexer_indexing_rewards_effective_cut{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateGauge(indexerEffectiveCutMetric, func() float64 {
			return indexerEffectiveCut
		})

		metricsList = append(metricsList, indexerEffectiveCutMetric)

		overDelegationDilution, err := strconv.ParseFloat(indexer.OverDelegationDilution, 64)
		if err != nil {
			log.Error().Err(err)
		}
		overDelegationDilutionMetric := fmt.Sprintf(`thegraph_indexer_overdelegationdilution{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateGauge(overDelegationDilutionMetric, func() float64 {
			return overDelegationDilution
		})

		metricsList = append(metricsList, overDelegationDilutionMetric)

		delegatedStakeRatio, err := strconv.ParseFloat(indexer.DelegatedStakeRatio, 64)
		if err != nil {
			log.Error().Err(err)
		}
		delegatedStakeRatioMetric := fmt.Sprintf(`thegraph_indexer_delegatedstakeratio{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateGauge(delegatedStakeRatioMetric, func() float64 {
			return delegatedStakeRatio
		})

		metricsList = append(metricsList, delegatedStakeRatioMetric)

		indexerRewardsOwnGenerationRatio, err := strconv.ParseFloat(indexer.IndexerRewardsOwnGenerationRatio, 64)
		if err != nil {
			log.Error().Err(err)
		}
		indexerRewardsOwnGenerationRatioMetric := fmt.Sprintf(`thegraph_indexer_rewardsgenerationratio{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateGauge(indexerRewardsOwnGenerationRatioMetric, func() float64 {
			return indexerRewardsOwnGenerationRatio
		})

		metricsList = append(metricsList, indexerRewardsOwnGenerationRatioMetric)

		ownStakeRatio, err := strconv.ParseFloat(indexer.OwnStakeRatio, 64)
		if err != nil {
			log.Error().Err(err)
		}
		ownStakeRatioMetric := fmt.Sprintf(`thegraph_indexer_ownstakeratio{indexer="%s"}`, indexer.ID)
		service.Metrics.GetOrCreateGauge(ownStakeRatioMetric, func() float64 {
			return ownStakeRatio
		})

		metricsList = append(metricsList, ownStakeRatioMetric)

	}
	service.clearOldMetrics(oldList, metricsList, "thegraph_indexer")

	return nil
}

// GetAndSetDelegatedStakesMetrics converts graphql response as metrics
func (service Service) GetAndSetDelegatedStakesMetrics() error {
	delegatedStakes, err := service.GqlMainnet.GetDelegatedStakes()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, ds := range delegatedStakes {
		stakedTokens, err := stringToGRT(ds.StakedTokens)
		if err != nil {
			log.Error().Err(err).Msg("stakedTokens")
		}
		stakedTokensMetric := fmt.Sprintf(`thegraph_delegatedstake_tokens_staked{indexer="%s",delegator="%s"}`, ds.Indexer.ID, ds.Delegator.ID)
		service.Metrics.GetOrCreateCounter(stakedTokensMetric).Set(stakedTokens.Uint64())

		metricsList = append(metricsList, stakedTokensMetric)

		unstakedTokens, err := stringToGRT(ds.UnstakedTokens)
		if err != nil {
			fmt.Println(err)
		}
		unstakedTokensMetric := fmt.Sprintf(`thegraph_delegatedstake_tokens_unstaked{indexer="%s",delegator="%s"}`, ds.Indexer.ID, ds.Delegator.ID)
		service.Metrics.GetOrCreateCounter(unstakedTokensMetric).Set(unstakedTokens.Uint64())

		metricsList = append(metricsList, unstakedTokensMetric)

		lockedTokens, err := stringToGRT(ds.LockedTokens)
		if err != nil {
			fmt.Println(err)
		}
		lockedTokensMetric := fmt.Sprintf(`thegraph_delegatedstake_tokens_locked{indexer="%s",delegator="%s"}`, ds.Indexer.ID, ds.Delegator.ID)
		service.Metrics.GetOrCreateCounter(lockedTokensMetric).Set(lockedTokens.Uint64())

		metricsList = append(metricsList, lockedTokensMetric)

		realizedRewards, err := strFloatToBigInt(ds.RealizedRewards)
		if err != nil {
			fmt.Println(err)
		}
		dec := big.NewInt(1000000000000000000) // NEED REWORK
		realizedRewards.Div(&realizedRewards, dec)
		realizedRewardsMetric := fmt.Sprintf(`thegraph_delegatedstake_rewards_realized{indexer="%s",delegator="%s"}`, ds.Indexer.ID, ds.Delegator.ID)
		service.Metrics.GetOrCreateCounter(realizedRewardsMetric).Set(realizedRewards.Uint64())

		metricsList = append(metricsList, realizedRewardsMetric)

		// ShareAmount TODO
		// shareAmount * personalExchangeRate * delegationExchangeRate / delegatorShares

		shareAmount, err := strconv.ParseFloat(ds.ShareAmount, 64)
		if err != nil {
			log.Error().Err(err)
		}
		/* personalExchangeRate, err := strconv.ParseFloat(ds.PersonalExchangeRate, 64)
		if err != nil {
			log.Println(err)
		} */
		/* indexerDelegationExchangeRate, err := strconv.ParseFloat(ds.Indexer.DelegationExchangeRate, 64)
		if err != nil {
			log.Println(err)
		} */
		indexerDelegatorShares, err := strconv.ParseFloat(ds.Indexer.DelegatorShares, 64)
		if err != nil {
			log.Error().Err(err)
		}
		delegatorPercent := shareAmount / indexerDelegatorShares
		// fmt.Println(ds.Indexer.ID, ds.Delegator.ID, delegatorPercent)
		indexerDelegatorPercentMetric := fmt.Sprintf(`thegraph_delegatedstake_indexer_delegator_percent{indexer="%s",delegator="%s"}`, ds.Indexer.ID, ds.Delegator.ID)
		service.Metrics.GetOrCreateGauge(indexerDelegatorPercentMetric, func() float64 {
			dp := delegatorPercent
			/* dp, err := strconv.ParseFloat(delegatorPercent, 64)
			if err != nil {
				log.Println("failed to convert string to float:", err)
				return float64(0)
			} */
			return dp
		})

		metricsList = append(metricsList, indexerDelegatorPercentMetric)

		unrealizedRewards, err := calcDelegatorRewards(ds.Indexer.DelegationExchangeRate, ds.PersonalExchangeRate, ds.ShareAmount)
		if err != nil {
			log.Error().Err(err)
		}
		unrealizedRewardsMetric := fmt.Sprintf(`thegraph_delegatedstake_rewards_unrealized{indexer="%s",delegator="%s"}`, ds.Indexer.ID, ds.Delegator.ID)
		service.Metrics.GetOrCreateCounter(unrealizedRewardsMetric).Set(unrealizedRewards.Uint64())

		metricsList = append(metricsList, unrealizedRewardsMetric)

	}
	service.clearOldMetrics(oldList, metricsList, "thegraph_delegatedstake")

	return nil
}

// GetAndSetDelegatorMetrics convert subgraph data to metrics
func (service Service) GetAndSetDelegatorMetrics() error {
	delegators, err := service.GqlMainnet.GetDelegators()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, delegator := range delegators {
		stakedTokens, err := stringToGRT(delegator.TotalStakedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		stakedTokensMetric := fmt.Sprintf(`thegraph_delegator_tokens_staked{delegator="%s"}`, delegator.ID)
		service.Metrics.GetOrCreateCounter(stakedTokensMetric).Set(stakedTokens.Uint64())

		metricsList = append(metricsList, stakedTokensMetric)

		unstakedTokens, err := stringToGRT(delegator.TotalUnstakedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		unstakedTokensMetric := fmt.Sprintf(`thegraph_delegator_tokens_unstaked{delegator="%s"}`, delegator.ID)
		service.Metrics.GetOrCreateCounter(unstakedTokensMetric).Set(unstakedTokens.Uint64())

		metricsList = append(metricsList, unstakedTokensMetric)

		realizedRewards, err := strFloatToBigInt(delegator.TotalRealizedRewards)
		if err != nil {
			log.Error().Err(err)
		}
		dec := big.NewInt(1000000000000000000) // NEED REWORK
		realizedRewards.Div(&realizedRewards, dec)
		realizedRewardsMetric := fmt.Sprintf(`thegraph_delegator_rewards_realized{delegator="%s"}`, delegator.ID)
		service.Metrics.GetOrCreateCounter(realizedRewardsMetric).Set(realizedRewards.Uint64())

		metricsList = append(metricsList, realizedRewardsMetric)

	}

	service.clearOldMetrics(oldList, metricsList, "thegraph_delegator")

	return nil
}

// GetAndSetSubgraphDeploymentMetrics convert subgraph data to metrics
func (service Service) GetAndSetSubgraphDeploymentMetrics() error {
	subgraphDeployments, err := service.GqlMainnet.GetSubgraphDeployments()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, subgraphDeployment := range subgraphDeployments {
		subgraphName := subgraphDeployment.OriginalName
		subgraphHex := subgraphDeployment.ID
		stakedTokens, err := stringToGRT(subgraphDeployment.StakedTokens)
		if err != nil {
			log.Error().Err(err)
		}
		subgraphHash, err := utils.SubgraphHexToHash(subgraphHex)
		if err != nil {
			log.Error().Err(err)
		}
		stakedTokensMetric := fmt.Sprintf(`thegraph_subgraph_deployment_tokens_staked{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(stakedTokensMetric).Set(stakedTokens.Uint64())

		metricsList = append(metricsList, stakedTokensMetric)

		indexingRewardAmount, err := stringToGRT(subgraphDeployment.IndexingRewardAmount)
		if err != nil {
			log.Error().Err(err)
		}
		indexingRewardAmountMetric := fmt.Sprintf(`thegraph_subgraph_deployment_rewards_total{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(indexingRewardAmountMetric).Set(indexingRewardAmount.Uint64())

		metricsList = append(metricsList, indexingRewardAmountMetric)

		indexingIndexerRewardAmount, err := stringToGRT(subgraphDeployment.IndexingIndexerRewardAmount)
		if err != nil {
			log.Error().Err(err)
		}
		indexingIndexerRewardAmountMetric := fmt.Sprintf(`thegraph_subgraph_deployment_rewards_indexer{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(indexingIndexerRewardAmountMetric).Set(indexingIndexerRewardAmount.Uint64())

		metricsList = append(metricsList, indexingIndexerRewardAmountMetric)

		indexingDelegatorRewardAmount, err := stringToGRT(subgraphDeployment.IndexingDelegatorRewardAmount)
		if err != nil {
			log.Error().Err(err)
		}
		indexingDelegatorRewardAmountMetric := fmt.Sprintf(`thegraph_subgraph_deployment_rewards_delegator{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(indexingDelegatorRewardAmountMetric).Set(indexingDelegatorRewardAmount.Uint64())

		metricsList = append(metricsList, indexingDelegatorRewardAmountMetric)

		queryFeesAmount, err := stringToGRT(subgraphDeployment.QueryFeesAmount)
		if err != nil {
			log.Error().Err(err)
		}
		queryFeesAmountMetric := fmt.Sprintf(`thegraph_subgraph_deployment_query_fees{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(queryFeesAmountMetric).Set(queryFeesAmount.Uint64())

		metricsList = append(metricsList, queryFeesAmountMetric)

		queryFeeRebates, err := stringToGRT(subgraphDeployment.QueryFeeRebates)
		if err != nil {
			log.Error().Err(err)
		}
		queryFeeRebatesMetric := fmt.Sprintf(`thegraph_subgraph_deployment_query_fees_rebates{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(queryFeeRebatesMetric).Set(queryFeeRebates.Uint64())

		metricsList = append(metricsList, queryFeeRebatesMetric)

		curatorFeeRewards, err := stringToGRT(subgraphDeployment.CuratorFeeRewards)
		if err != nil {
			log.Error().Err(err)
		}
		curatorFeeRewardsMetric := fmt.Sprintf(`thegraph_subgraph_deployment_query_fees_curators{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(curatorFeeRewardsMetric).Set(curatorFeeRewards.Uint64())

		metricsList = append(metricsList, curatorFeeRewardsMetric)

		signalledTokens, err := stringToGRT(subgraphDeployment.SignalledTokens)
		if err != nil {
			log.Error().Err(err)
		}
		signalledTokensMetric := fmt.Sprintf(`thegraph_subgraph_deployment_tokens_signalled{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(signalledTokensMetric).Set(signalledTokens.Uint64())

		metricsList = append(metricsList, signalledTokensMetric)

		unsignalledTokens, err := stringToGRT(subgraphDeployment.UnsignalledTokens)
		if err != nil {
			log.Error().Err(err)
		}
		unsignalledTokensMetric := fmt.Sprintf(`thegraph_subgraph_deployment_tokens_unsignalled{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(unsignalledTokensMetric).Set(unsignalledTokens.Uint64())

		metricsList = append(metricsList, unsignalledTokensMetric)

		signalAmount := new(big.Int)
		signalAmount, ok := signalAmount.SetString(subgraphDeployment.SignalAmount, 10)
		if !ok {
			return fmt.Errorf("SetString failed")
		}

		signalAmountMetric := fmt.Sprintf(`thegraph_subgraph_deployment_signal_amount{name="%s",id="%s",ipfs_hash="%s"}`, subgraphName, subgraphHex, subgraphHash)
		service.Metrics.GetOrCreateCounter(signalAmountMetric).Set(signalAmount.Uint64())

		metricsList = append(metricsList, signalAmountMetric)

	}

	service.clearOldMetrics(oldList, metricsList, "thegraph_subgraph")

	return nil
}

// GetAndSetCurators convert subgraph data to metrics
func (service Service) GetAndSetCurators() error {
	curators, err := service.GqlMainnet.GetCurators()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, curator := range curators {
		curatorID := curator.ID
		totalSignalledTokens, err := stringToGRT(curator.TotalSignalledTokens)
		if err != nil {
			return err
		}
		totalSignalledTokensMetric := fmt.Sprintf(`thegraph_curator_tokens_signalled{id="%s"}`, curatorID)
		// service.Metrics.GetOrCreateCounter(totalSignalledTokensMetric).Set(totalSignalledTokens.Uint64())
		service.Metrics.GetOrCreateGauge(totalSignalledTokensMetric, func() float64 {
			return float64(totalSignalledTokens.Int64())
		})

		metricsList = append(metricsList, totalSignalledTokensMetric)

		totalUnsignalledTokens, err := stringToGRT(curator.TotalUnsignalledTokens)
		if err != nil {
			log.Error().Err(err)
		}
		totalUnsignalledTokensMetric := fmt.Sprintf(`thegraph_curator_tokens_unsignalled{id="%s"}`, curatorID)
		// service.Metrics.GetOrCreateCounter(totalUnsignalledTokensMetric).Set(totalUnsignalledTokens.Uint64())
		service.Metrics.GetOrCreateGauge(totalUnsignalledTokensMetric, func() float64 {
			return float64(totalUnsignalledTokens.Int64())
		})

		metricsList = append(metricsList, totalUnsignalledTokensMetric)

		totalNameSignalledTokens, err := stringToGRT(curator.TotalNameSignalledTokens)
		if err != nil {
			log.Error().Err(err)
		}
		totalNameSignalledTokensMetric := fmt.Sprintf(`thegraph_curator_tokens_namesignalled{id="%s"}`, curatorID)
		service.Metrics.GetOrCreateGauge(totalNameSignalledTokensMetric, func() float64 {
			return float64(totalNameSignalledTokens.Int64())
		})

		metricsList = append(metricsList, totalNameSignalledTokensMetric)

		totalNameUnsignalledTokens, err := stringToGRT(curator.TotalNameUnsignalledTokens)
		if err != nil {
			log.Error().Err(err)
		}
		totalNameUnsignalledTokensMetric := fmt.Sprintf(`thegraph_curator_tokens_nameunsignalled{id="%s"}`, curatorID)
		service.Metrics.GetOrCreateGauge(totalNameUnsignalledTokensMetric, func() float64 {
			return float64(totalNameUnsignalledTokens.Int64())
		})

		metricsList = append(metricsList, totalNameUnsignalledTokensMetric)

		totalWithdrawnTokens, err := stringToGRT(curator.TotalWithdrawnTokens)
		if err != nil {
			log.Error().Err(err)
		}
		totalWithdrawnTokensMetric := fmt.Sprintf(`thegraph_curator_tokens_withdrawn{id="%s"}`, curatorID)
		service.Metrics.GetOrCreateCounter(totalWithdrawnTokensMetric).Set(totalWithdrawnTokens.Uint64())

		metricsList = append(metricsList, totalWithdrawnTokensMetric)

		realizedRewards, err := stringToGRT(curator.RealizedRewards)
		if err != nil {
			log.Error().Err(err)
		}
		realizedRewardsMetric := fmt.Sprintf(`thegraph_curator_rewards_realized{id="%s"}`, curatorID)
		service.Metrics.GetOrCreateCounter(realizedRewardsMetric).Set(realizedRewards.Uint64())

		metricsList = append(metricsList, realizedRewardsMetric)

		totalReturn, err := stringToGRT(curator.TotalReturn)
		if err != nil {
			log.Error().Err(err)
		}
		totalReturnMetric := fmt.Sprintf(`thegraph_curator_return_total{id="%s"}`, curatorID)
		service.Metrics.GetOrCreateCounter(totalReturnMetric).Set(totalReturn.Uint64())

		metricsList = append(metricsList, totalReturnMetric)

		totalNameSignal, err := stringToGRT(curator.TotalNameSignal)
		if err != nil {
			log.Error().Err(err)
		}
		totalNameSignalMetric := fmt.Sprintf(`thegraph_curator_namesignal_total{id="%s"}`, curatorID)
		service.Metrics.GetOrCreateCounter(totalNameSignalMetric).Set(totalNameSignal.Uint64())

		metricsList = append(metricsList, totalNameSignalMetric)

	}

	service.clearOldMetrics(oldList, metricsList, "thegraph_curator")

	return nil
}

func (service Service) GetAndSetGRTRate() {
	service.Metrics.GetOrCreateGauge("thegraph_grt_usd", func() float64 {
		url := "https://api.coingecko.com/api/v3/coins/the-graph?localization=false&tickers=false&community_data=false&developer_data=false&sparkline=false"
		resp, err := http.Get(url)
		if err != nil {
			log.Error().Err(err).Msg("failed to get GRT rate") // log.Println("failed to get GRT rate:", err)
		}
		defer resp.Body.Close()
		if resp.Status == "200 OK" {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Error().Err(err) // log.Println("failed to read body", err)
			}
			rate, err := strconv.ParseFloat(gjson.Get(string(body), "market_data.current_price.usd").String(), 64)
			if err != nil {
				log.Error().Err(err) // log.Println("failed to conver rate to float", err)
			}
			return rate
		}
		return float64(0)
	})
}

func (service Service) GetAndSetSignals() error {
	signals, err := service.GqlMainnet.GetSignals()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, signal := range signals {
		subgraphHash, err := utils.SubgraphHexToHash(signal.SubgraphDeployment.ID)
		if err != nil {
			log.Error().Err(err)
		}
		signalledTokens, err := utils.ToDecimal(signal.SignalledTokens, 18)
		if err != nil {
			return err
		}
		signalSignalledTokens := fmt.Sprintf(`thegraph_signal_signalled_tokens{id="%s",curator="%s", subgraphDeployment="%s"}`, signal.ID, signal.Curator.ID, subgraphHash)
		service.Metrics.GetOrCreateGauge(signalSignalledTokens, func() float64 {
			res, _ := signalledTokens.Float64()
			return res
		})

		metricsList = append(metricsList, signalSignalledTokens)

		unsignalledTokens, err := utils.ToDecimal(signal.UnsignalledTokens, 18)
		if err != nil {
			return err
		}
		signalUnsignalledTokens := fmt.Sprintf(`thegraph_signal_unsignalled_tokens{id="%s",curator="%s", subgraphDeployment="%s"}`, signal.ID, signal.Curator.ID, subgraphHash)
		service.Metrics.GetOrCreateGauge(signalUnsignalledTokens, func() float64 {
			res, _ := unsignalledTokens.Float64()
			return res
		})

		metricsList = append(metricsList, signalUnsignalledTokens)

		/* 	signalAmount, err := utils.ToDecimal(signal.Signal, 0)
		if err != nil {
			return err
		} */

		signalAmount, err := strconv.ParseFloat(signal.Signal, 64)
		if err != nil {
			return err
		}
		signalSignalAmount := fmt.Sprintf(`thegraph_signal_amount{id="%s",curator="%s", subgraphDeployment="%s"}`, signal.ID, signal.Curator.ID, subgraphHash)
		service.Metrics.GetOrCreateGauge(signalSignalAmount, func() float64 {
			// res, _ := signalAmount.Float64()
			return signalAmount
		})

		metricsList = append(metricsList, signalSignalAmount)

	}
	service.clearOldMetrics(oldList, metricsList, "thegraph_signal")

	return nil
}

func (service Service) GetAndSetNameSignals() error {
	nameSignals, err := service.GqlMainnet.GetNameSignals()
	if err != nil {
		return err
	}

	oldList := service.Metrics.ListMetricNames()
	var metricsList []string

	for _, nameSignal := range nameSignals {
		subgraphHash, err := utils.SubgraphHexToHash(nameSignal.Subgraph.CurrentVersion.SubgraphDeployment.ID)
		if err != nil {
			log.Error().Err(err)
		}
		nameSignalledTokens, err := utils.ToDecimal(nameSignal.SignalledTokens, 18)
		if err != nil {
			return err
		}
		nameSignalSignalledTokens := fmt.Sprintf(`thegraph_namesignal_signalled_tokens{id="%s",curator="%s", subgraphDeployment="%s"}`, nameSignal.ID, nameSignal.Curator.ID, subgraphHash)
		service.Metrics.GetOrCreateGauge(nameSignalSignalledTokens, func() float64 {
			res, _ := nameSignalledTokens.Float64()
			return res
		})

		metricsList = append(metricsList, nameSignalSignalledTokens)

		nameUnsignalledTokens, err := utils.ToDecimal(nameSignal.UnsignalledTokens, 18)
		if err != nil {
			return err
		}
		nameSignalUnsignalledTokens := fmt.Sprintf(`thegraph_namesignal_unsignalled_tokens{id="%s",curator="%s", subgraphDeployment="%s"}`, nameSignal.ID, nameSignal.Curator.ID, subgraphHash)
		service.Metrics.GetOrCreateGauge(nameSignalUnsignalledTokens, func() float64 {
			res, _ := nameUnsignalledTokens.Float64()
			return res
		})

		metricsList = append(metricsList, nameSignalUnsignalledTokens)

		nameSignalAmount, err := strconv.ParseFloat(nameSignal.NameSignal, 64)
		if err != nil {
			return err
		}
		nameSignalSignalAmount := fmt.Sprintf(`thegraph_namesignal_amount{id="%s",curator="%s", subgraphDeployment="%s"}`, nameSignal.ID, nameSignal.Curator.ID, subgraphHash)
		service.Metrics.GetOrCreateGauge(nameSignalSignalAmount, func() float64 {
			return nameSignalAmount
		})

		metricsList = append(metricsList, nameSignalSignalAmount)

	}
	service.clearOldMetrics(oldList, metricsList, "thegraph_namesignal")

	return nil
}

func (service Service) GetAndSetIndexerDataPoints() error {
	indexerDataPoints, err := service.QoSOracle.GetIndexerDataPoints()
	if err != nil {
		return err
	}
	oldList := service.Metrics.ListMetricNames()
	var metricsList []string
	for _, indexerDataPoint := range indexerDataPoints {
		floatAvgQueryFee, err := strconv.ParseFloat(indexerDataPoint.AvgQueryFee, 64)
		if err != nil {
			return err
		}
		avgQueryFee := fmt.Sprintf(`thegraph_indexerdatapoint_avg_query_fee{indexer="%s",subgraphDeployment="%s"}`, indexerDataPoint.IndexerWallet, indexerDataPoint.SubgraphDeploymentIpshHash)
		service.Metrics.GetOrCreateGauge(avgQueryFee, func() float64 {
			return floatAvgQueryFee
		})

		metricsList = append(metricsList, avgQueryFee)

		floatAvgIndexerLatencyMs, err := strconv.ParseFloat(indexerDataPoint.AvgIndexerLatencyMs, 64)
		if err != nil {
			return err
		}
		avgIndexerLatencyMs := fmt.Sprintf(`thegraph_indexerdatapoint_avg_indexer_latency_ms{indexer="%s",subgraphDeployment="%s"}`, indexerDataPoint.IndexerWallet, indexerDataPoint.SubgraphDeploymentIpshHash)
		service.Metrics.GetOrCreateGauge(avgIndexerLatencyMs, func() float64 {
			return floatAvgIndexerLatencyMs
		})

		metricsList = append(metricsList, avgIndexerLatencyMs)

		intQueryCount, err := strconv.Atoi(indexerDataPoint.QueryCount)
		if err != nil {
			return err
		}
		queryCount := fmt.Sprintf(`thegraph_indexerdatapoint_query_count{indexer="%s",subgraphDeployment="%s"}`, indexerDataPoint.IndexerWallet, indexerDataPoint.SubgraphDeploymentIpshHash)
		service.Metrics.GetOrCreateCounter(queryCount).Set(uint64(intQueryCount))
		metricsList = append(metricsList, queryCount)

	}
	service.clearOldMetrics(oldList, metricsList, "thegraph_indexerdatapoint")
	return nil
}
