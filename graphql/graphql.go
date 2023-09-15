package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"

	"github.com/shurcooL/graphql"
)

// Print response in json
func Print(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

// GetIndexers get Indexers
func (gs *GraphService) GetIndexers() ([]Indexer, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		Indexers []Indexer `graphql:"indexers(first: $limit, where: { stakedTokens_gt:0, id_gt: $lastID  })"`
	}
	var indexers []Indexer
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []Indexer{}, err
		}
		skip += 1000
		indexers = append(indexers, q.Indexers...)
		variables["lastID"] = graphql.String(indexers[len(indexers)-1].ID)
		log.Debug().Msgf("Indexers pagination: %d", skip)
		if len(q.Indexers) == 0 || len(q.Indexers) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []Indexer{}, err
	}
	return indexers, nil
}

// GetAllocations get allocations info
func (gs *GraphService) GetAllocations() ([]Allocation, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		Allocations []Allocation `graphql:"allocations(first: $limit, where: {status: Active, id_gt: $lastID})"`
	}
	var allocations []Allocation
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []Allocation{}, err
		}
		skip += 1000
		allocations = append(allocations, q.Allocations...)
		variables["lastID"] = graphql.String(allocations[len(allocations)-1].ID)
		log.Debug().Msgf("Allocations pagination: %d", skip)
		if len(q.Allocations) == 0 || len(q.Allocations) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []Allocation{}, err
	}
	return allocations, nil
}

// GetEpoches get epoches info
func (gs *GraphService) GetEpoches() {
	var q struct {
		Epoches []struct {
			ID           graphql.String
			TotalRewards graphql.String
		} `graphql:"epoches(first: 10)"`
	}
	fmt.Println(q)
}

// GetGraphNetwork info
func (gs *GraphService) GetGraphNetwork() (GraphNetwork, error) {
	var q struct {
		GraphNetwork GraphNetwork `graphql:"graphNetwork(id: 1)"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return GraphNetwork{}, err
	}
	/* 	v := reflect.ValueOf(q.GraphNetwork)
	   	typeOfQ := v.Type()

	   	for i := 0; i < v.NumField(); i++ {
	   		fmt.Printf("Field: %s\tValue: %v\n", typeOfQ.Field(i).Name, v.Field(i).Interface())
	   	} */
	return q.GraphNetwork, nil
}

// GetDelegatedStakes info
func (gs *GraphService) GetDelegatedStakes() ([]DelegatedStake, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		DelegatedStake []DelegatedStake `graphql:"delegatedStakes(first: $limit,where: { shareAmount_gt:0, id_gt: $lastID  })"`
	}

	var DelegatedStakes []DelegatedStake
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []DelegatedStake{}, err
		}
		skip += 1000
		DelegatedStakes = append(DelegatedStakes, q.DelegatedStake...)
		variables["lastID"] = graphql.String(DelegatedStakes[len(DelegatedStakes)-1].ID)
		log.Debug().Msgf("Delegated stake pagination: %d", skip)
		if len(q.DelegatedStake) == 0 || len(q.DelegatedStake) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []DelegatedStake{}, err
	}
	return DelegatedStakes, nil
}

// GetDelegators get delegatros from graphql endpoint
func (gs *GraphService) GetDelegators() ([]Delegator, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		Delegators []Delegator `graphql:"delegators(first: $limit,where: { activeStakesCount_gt: 0, id_gt: $lastID  })"`
	}
	var delegators []Delegator
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []Delegator{}, err
		}
		skip += 1000
		delegators = append(delegators, q.Delegators...)
		variables["lastID"] = graphql.String(delegators[len(delegators)-1].ID)
		log.Debug().Msgf("Delegators pagination: %d", skip)
		// log.Println(skip)
		if len(q.Delegators) == 0 || len(q.Delegators) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []Delegator{}, err
	}
	return delegators, nil
}

// GetSubgraphDeployments gets subgraph deployments from graphql endpoint
func (gs *GraphService) GetSubgraphDeployments() ([]SubgraphDeployment, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		SubgraphDeployments []SubgraphDeployment `graphql:"subgraphDeployments(first: $limit, where: { activeSubgraphCount_gt: 0, id_gt: $lastID  })"`
	}
	var subgraphDeployments []SubgraphDeployment
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []SubgraphDeployment{}, err
		}
		skip += 1000
		subgraphDeployments = append(subgraphDeployments, q.SubgraphDeployments...)
		variables["lastID"] = graphql.String(subgraphDeployments[len(subgraphDeployments)-1].ID)
		log.Debug().Msgf("SubgraphDeployments pagination: %d", skip)
		// log.Println(skip)
		if len(q.SubgraphDeployments) == 0 || len(q.SubgraphDeployments) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []SubgraphDeployment{}, err
	}
	return subgraphDeployments, nil
}

// GetCurators gets curators info from graphql endpoint
func (gs *GraphService) GetCurators() ([]Curator, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		Curators []Curator `graphql:"curators(first: $limit, where: { activeCombinedSignalCount_gt: 0, id_gt: $lastID  })"`
	}
	var curators []Curator
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []Curator{}, err
		}
		skip += 1000
		curators = append(curators, q.Curators...)
		variables["lastID"] = graphql.String(curators[len(curators)-1].ID)
		log.Debug().Msgf("Curators pagination: %d", skip)
		// log.Println(skip)
		if len(q.Curators) == 0 || len(q.Curators) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []Curator{}, err
	}
	return curators, nil
}

// GetAllTokenLockWallets gets info from tokenlock subgraph only
func (gs *GraphService) GetAllTokenLockWallets() ([]TokenLockWallet, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}

	var q struct {
		TokenLockWallets []TokenLockWallet `graphql:"tokenLockWallets(first: $limit, where: { id_gt: $lastID  })"`
	}
	var tokenLockWallets []TokenLockWallet
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []TokenLockWallet{}, err
		}
		skip += 1000
		tokenLockWallets = append(tokenLockWallets, q.TokenLockWallets...)
		variables["lastID"] = graphql.String(tokenLockWallets[len(tokenLockWallets)-1].ID)
		log.Debug().Msgf("TokenLockWallets pagination: %d", skip)
		if len(q.TokenLockWallets) == 0 || len(q.TokenLockWallets) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []TokenLockWallet{}, err
	}
	return tokenLockWallets, nil
}

// GetSignals gets all signals from subgraph
func (gs *GraphService) GetSignals() ([]Signal, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		Signals []Signal `graphql:"signals(first: $limit, where: { id_gt: $lastID  })"`
	}
	var signals []Signal
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []Signal{}, err
		}
		skip += 1000
		signals = append(signals, q.Signals...)
		variables["lastID"] = graphql.String(signals[len(signals)-1].ID)
		log.Debug().Msgf("Signals pagination: %d", skip)
		if len(q.Signals) == 0 || len(q.Signals) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []Signal{}, err
	}
	return signals, nil
}

// GetNameSignals gets all nameSignals from subgraph
func (gs *GraphService) GetNameSignals() ([]NameSignal, error) {
	variables := map[string]interface{}{
		"limit":  graphql.Int(1000),
		"lastID": graphql.String(""),
	}
	var q struct {
		NameSignals []NameSignal `graphql:"nameSignals(first: $limit, where: { id_gt: $lastID  })"`
	}
	var nameSignals []NameSignal
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []NameSignal{}, err
		}
		skip += 1000
		nameSignals = append(nameSignals, q.NameSignals...)
		variables["lastID"] = graphql.String(nameSignals[len(nameSignals)-1].ID)
		log.Debug().Msgf("Name Signals pagination: %d", skip)
		if len(q.NameSignals) == 0 || len(q.NameSignals) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []NameSignal{}, err
	}
	return nameSignals, nil
}

func (gs *GraphService) GetIndexerDataPoints() ([]IndexerDataPoint, error) {
	variables := map[string]interface{}{
		"limit":      graphql.Int(1000),
		"lastID":     graphql.String(""),
		"epochShift": graphql.Int(time.Now().Unix() - 1800),
	}
	var q struct {
		IndexerDataPoints []IndexerDataPoint `graphql:"indexerDataPoints(first:$limit, orderBy: end_epoch, orderDirection: asc, where:{end_epoch_gte: $epochShift, id_gt: $lastID})"`
	}
	spew.Dump(variables)
	var indexerDataPoints []IndexerDataPoint
	var err error
	skip := 0
	for {
		err = gs.Client.Query(context.Background(), &q, variables)
		if err != nil {
			return []IndexerDataPoint{}, err
		}
		skip += 1000
		indexerDataPoints = append(indexerDataPoints, q.IndexerDataPoints...)
		// TODO check if empty response
		variables["lastID"] = graphql.String(indexerDataPoints[len(indexerDataPoints)-1].ID)
		log.Debug().Msgf("IndexerDataPoints pagination: %d", skip)
		if len(q.IndexerDataPoints) == 0 || len(q.IndexerDataPoints) < 1000 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return []IndexerDataPoint{}, err
	}
	return indexerDataPoints, nil
}
