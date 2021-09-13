package main

import (
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/akme/thegraph-exporter/ethereum"
	thegraph "github.com/akme/thegraph-exporter/graphql"
	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/peterbourgon/ff"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/graphql"
)

var (
	// Following vars are injected through the build flags (see Makefile)
	GitCommit string
	GitBranch string
	GitTag    string
)

// Service clients store
type Service struct {
	GqlMainnet   *thegraph.GraphService
	GqlTokenLock *thegraph.GraphService
	Eth          *ethereum.EthService
	Metrics      *metrics.Set
}

func strFloatToBigInt(s string) (big.Int, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return big.Int{}, err
	}
	bigval := new(big.Float)
	bigval.SetFloat64(f)
	result := new(big.Int)
	bigval.Int(result)

	return *result, nil
}

func stringToGRT(s string) (big.Int, error) {
	bi := new(big.Int)
	bi, ok := bi.SetString(s, 10)
	if !ok {
		return big.Int{}, fmt.Errorf("SetString failed")
	}
	dec := big.NewInt(1000000000000000000)
	bi.Div(bi, dec)
	return *bi, nil
}

// old formula: unrealizedReturn = (delegationExchangeRate / personalExchangeRate - 1) * shareAmount
// new formula: (delegationExchangeRate - personalExchangeRate) * shareAmount
func calcDelegatorRewards(delegationExchangeRate, personalExchangeRate, shareAmount string) (*big.Int, error) {
	delegationExchangeRateFloat, err := strconv.ParseFloat(delegationExchangeRate, 64)
	if err != nil {
		return &big.Int{}, err
	}
	personalExchangeRateFloat, err := strconv.ParseFloat(personalExchangeRate, 64)
	if err != nil {
		return &big.Int{}, err
	}
	shareAmountFloat, err := strconv.ParseFloat(shareAmount, 64)
	if err != nil {
		return &big.Int{}, err
	}
	// rewards := (delegationExchangeRateFloat/personalExchangeRateFloat - 1) * shareAmountFloat / 1000000000000000000
	rewards := (delegationExchangeRateFloat - personalExchangeRateFloat) * shareAmountFloat / 1000000000000000000
	bigval := new(big.Float)
	bigval.SetFloat64(rewards)
	result := new(big.Int)
	bigval.Int(result)
	return result, nil
}

func (service Service) metricsPage(w http.ResponseWriter, r *http.Request) {
	service.Metrics.WritePrometheus(w)
}

func handle(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func main() {
	fs := flag.NewFlagSet("thegraph-exporter", flag.ExitOnError)
	var (
		ethNode                   = fs.String("eth-node", "http://geth-mainnet:8545", "Ethereum RPC API URL")
		networkSubgraph           = fs.String("network-subgraph", "https://gateway.thegraph.com/network", "The Graph Network Subgraph URL")
		tokenDistributionSubgraph = fs.String("distribution-subgraph", "https://api.thegraph.com/subgraphs/name/graphprotocol/token-distribution", "The Graph Token Distribution Subgraph URL")
		listenPort                = fs.String("listen", ":8080", "Listen HTTP Port ")
		logLevelFlag              = fs.Int("log-level", 0, "Log level")
		httpTimeout               = fs.Int("http-timeout", 15, "HTTP request timeout")
	)
	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("THEGRAPH_EXPORTER"),
	)
	if err != nil {
		log.Fatal().Msgf("Incorrect flags")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logLevel := zerolog.Level(*logLevelFlag)
	zerolog.SetGlobalLevel(logLevel)

	log.Info().Msgf("Build info: %s/%s/%s built with %s", GitBranch, GitCommit, GitTag, runtime.Version())

	if ok := govalidator.IsURL(*ethNode); !ok {
		log.Fatal().Msgf("Incorrect Ethereum Node URL: %s", *ethNode)
	}

	if ok := govalidator.IsURL(*networkSubgraph); !ok {
		log.Fatal().Msgf("Incorrect Network Subgraph URL: %s", *networkSubgraph)
	}

	if ok := govalidator.IsURL(*tokenDistributionSubgraph); !ok {
		log.Fatal().Msgf("Incorrect Token Distribution Subgraph URL: %s", *tokenDistributionSubgraph)
	}

	ethClient, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal().Msgf("Oops! Ethereum node connection error: %s", err)
	}

	gqlMainnetClient := graphql.NewClient(*networkSubgraph, &http.Client{Timeout: time.Duration(*httpTimeout) * time.Second})
	gqlTokenLockClient := graphql.NewClient(*tokenDistributionSubgraph, &http.Client{Timeout: time.Duration(*httpTimeout) * time.Second})
	metricSet := metrics.NewSet()

	clients := Service{
		Eth:          &ethereum.EthService{Client: ethClient},
		GqlMainnet:   &thegraph.GraphService{Client: gqlMainnetClient},
		GqlTokenLock: &thegraph.GraphService{Client: gqlTokenLockClient},
		Metrics:      metricSet,
	}

	go func() {
		for {
			err := clients.GetAndSetSignals()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetSignals")
			}

			err = clients.GetAndSetNameSignals()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetNameSignals")
			}

			err = clients.GetAndSetVestingBalanceMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetVestingBalanceMetrics")
			}

			err = clients.GetAndSetGraphNetworkMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetGraphNetworkMetrics")
			}
			clients.GetAndSetGRTRate()
			err = clients.GetAndSetCurators()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetCuratorMetrics")
			}
			err = clients.GetAndSetSubgraphDeploymentMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetSubgraphDeploymentMetrics")
			}
			err = clients.GetAndSetDelegatorMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetDelegatorMetrics")
			}
			err = clients.GetAndSetDelegatedStakesMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetDelegatedStakesMetrics")
			}
			err = clients.GetAndSetIndexerMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetIndexerMetrics")
			}
			err = clients.GetAndSetActiveAllocationsRewardsMetrics()
			if err != nil {
				log.Error().Err(err).Msg("GetAndSetActiveAllocationsRewardsMetrics")
			}

			log.Info().Msgf("Metrics collected: %d", len(clients.Metrics.ListMetricNames()))
			// time.Sleep(30 * time.Second)
		}
	}()
	http.HandleFunc("/metrics", handle(clients.metricsPage))
	log.Fatal().Msgf("failed to listen %v", http.ListenAndServe(*listenPort, nil))
}
