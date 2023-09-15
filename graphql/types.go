package graphql

import (
	"math/big"

	"github.com/shurcooL/graphql"
)

// GraphService stores graphql client
type GraphService struct {
	Client *graphql.Client
}

// GraphNetwork info
type GraphNetwork struct {
	ID                            string   // 	ID is set to 1
	Controller                    string   // Bytes! // Controller address
	GraphToken                    string   // Bytes! // Graph token address
	EpochManager                  string   // Bytes! // Epoch manager address
	EpochManagerImplementations   []string // [Bytes!]! // Epoch Manager implementations. Last in the array is current
	Curation                      string   // Bytes! // Curation address
	CurationImplementations       []string // [Bytes!]! // Curation implementations. Last in the array is current
	Staking                       string   // Bytes! // Staking address
	StakingImplementations        []string // [Bytes!]! // Graph token implementations. Last in the array is current
	DisputeManager                string   // Dispute manager address
	Gns                           string   // Bytes! // GNS address
	ServiceRegistry               string   // Bytes! // Service registry address
	RewardsManager                string   // Bytes! // Rewards manager address
	RewardsManagerImplementations []string //  [Bytes!]! // Rewards Manager implementations. Last in the array is current
	IsPaused                      bool     // True if the protocol is paused
	IsPartialPaused               bool     // True if the protocol is partially paused
	Governor                      string   // Bytes! // Governor of the controller i.e. the whole protocol
	PauseGuardian                 string   // Bytes! // Pause guardian address
	CurationPercentage            int      // Percentage of fees going to curators
	ProtocolFeePercentage         int      // Percentage of fees burn as protocol fee
	DelegationRatio               int      // Ratio of max staked delegation tokens to indexers stake that earns rewards
	ChannelDisputeEpochs          int      // Period to wait before fees can be claimed in rebate pool
	MaxAllocationEpochs           int      // Period to wait before delegators can settle
	ThawingPeriod                 int      // Time in blocks needed to wait to unstake
	DelegationParametersCooldown  int      // Minimum time an Indexer must use for resetting their Delegation parameters
	// WAS REMOVED 28.01.2021 IndexingRewardsPerEpoch        int      // The number of times indexing rewards are handed out each epoch
	MinimumIndexerStake                 string   // big.Int  // Minimum GRT an indexer must stake
	Slashers                            []string // [Bytes!] // Contracts that have been approved to be a slasher
	DelegationUnbondingPeriod           int      // Time in epochs a delegator needs to wait to withdraw delegated stake
	RebateRatio                         string   // big.Int  // BigDecimal! // Alpha in the cobbs douglas formula
	DelegationTaxPercentage             int      // Tax that delegators pay to deposit
	AssetHolders                        []string // [Bytes!] // Asset holder for the protocol
	TotalTokensStaked                   string   // big.Int  // The total amount of GRT staked in the staking contract
	TotalTokensClaimable                string   // big.Int  // Total tokens that are settled and waiting to be claimed
	TotalUnstakedTokensLocked           string   // big.Int  // Total tokens that are currently locked or withdrawable in the network from unstaking
	TotalTokensAllocated                string   // big.Int  // Total GRT currently in allocation
	TotalDelegatedTokens                string   // big.Int  // Total delegated tokens in the protocol
	TotalQueryFees                      string   // big.Int  // Total query fees generated in the network
	TotalIndexerQueryFeesCollected      string   // big.Int  // Total query fees collected by indexers
	TotalIndexerQueryFeeRebates         string   // big.Int  // Total query fees rebates claimed by indexers
	TotalDelegatorQueryFeeRebates       string   // big.Int  // Total query fees rebates claimed by delegators
	TotalCuratorQueryFees               string   // big.Int  // Total query fees payed to curators
	TotalTaxedQueryFees                 string   // big.Int  // Total query fees taxed by the protocol
	TotalUnclaimedQueryFeeRebates       string   // big.Int  // Total query fees taxed by the protocol
	TotalIndexingRewards                string   // big.Int  // Total indexing rewards earned. Including delegators and indexers parts
	TotalIndexingDelegatorRewards       string   // big.Int  // Total indexing rewards earned by Delegators
	TotalIndexingIndexerRewards         string   // big.Int  // Total indexing rewards earned by Indexers
	NetworkGRTIssuance                  string   `graphql:"networkGRTIssuance"` // big.Int  // The issuance rate that GRT is minted at for rewards to staked Indexers
	SubgraphAvailabilityOracle          string   // Bytes! // Address of the availability oracle
	DefaultReserveRatio                 int      // Default reserve ratio for all subgraphs
	MinimumCurationDeposit              string   // big.Int // Minimum amount of tokens needed to start curating
	CurationTaxPercentage               int      // The fee charged when a curator withdraws signal
	TotalTokensSignalled                string   // big.Int // The total amount of GRT signalled in the Curation contract
	TotalSupply                         string   // big.Int // Graph Token supply
	GRTinUSD                            string   `graphql:"GRTinUSD"`       // big.Int // BigDecimal! // Price of one GRT in USD
	GRTinETH                            string   `graphql:"GRTinETH"`       // big.Int // BigDecimal // Price of one GRT in ETH
	TotalGRTMinted                      string   `graphql:"totalGRTMinted"` // big.Int // Total amount of GRT minted
	TotalGRTBurned                      string   `graphql:"totalGRTBurned"` // big.Int // Total amount of GRT burned
	EpochLength                         int      // Epoch Length in blocks
	LastRunEpoch                        int      // Epoch that was last run
	LastLengthUpdateEpoch               int      // Epoch when epoch length was last updated
	LastLengthUpdateBlock               int      // Block when epoch length was last updated
	CurrentEpoch                        int      // Current epoch the protocol is in
	IndexerCount                        int      // Total indexers
	StakedIndexersCount                 int      // Number of indexers that currently have some stake in the protocol
	DelegatorCount                      int      // Total delegators
	CuratorCount                        int      // Total curators
	SubgraphCount                       int      // Total subgraphs
	SubgraphDeploymentCount             int      // Total subgraphs
	EpochCount                          int      // Total epochs
	Arbitrator                          string   // Bytes! // Dispute arbitrator
	SlashingPercentage                  int      // Penalty to Indexer on successful disputes
	MinimumDisputeDeposit               string   // big.Int // Minimum deposit to create a dispute
	FishermanRewardPercentage           int      // Reward to Fisherman on successful disputes
	TotalTokensStakedTransferredToL2    string
	TotalDelegatedTokensTransferredToL2 string
	TotalSignalledTokensTransferredToL2 string
}

// Indexer info
type Indexer struct {
	ID                               string                                    // Eth address of Indexer
	CreatedAt                        int                                       // Time this indexer was created
	Account                          struct{ ID string }                       // GraphAccount // Graph account of this indexer
	URL                              string                                    // Service registry URL string
	GeoHash                          string                                    // Geohash of the indexer
	DefaultDisplayName               string                                    // Default display name is the current default name. Used for filtered queries
	StakedTokens                     string                                    // Current tokens staked in the protocol
	AllocatedTokens                  string                                    // Total tokens allocated on all subgraphs
	UnstakedTokens                   string                                    // Tokens that have been unstaked and withdrawn
	LockedTokens                     string                                    // Current tokens locked
	TokensLockedUntil                int                                       // The block when Indexers tokens unlock
	Allocations                      []struct{ ID string }                     // []Allocation        // Active allocations of stake for this Indexer
	TotalAllocations                 []struct{ ID string }                     // []Allocation     // Closed and active allocations of stake for this Indexer
	AllocationCount                  int                                       // Number of active allocations of stake for this Indexer
	TotalAllocationCount             string                                    // Number of active and closed allocations for this Indexer
	QueryFeesCollected               string                                    // Total query fees collected. Includes the indexer and delegator parts
	QueryFeeRebates                  string                                    // Query fee rebate amount claimed from the protocol
	RewardsEarned                    string                                    // Total indexing rewards earned by this indexer from inflation. Including delegation rewards
	IndexerIndexingRewards           string                                    // The total amount of indexing rewards the indexer kept
	DelegatorIndexingRewards         string                                    // The total amount of indexing rewards given to delegators
	IndexerRewardsOwnGenerationRatio string                                    // Amount of delegated tokens that can be eligible for rewards
	DelegatedCapacity                string                                    // Amount of delegated tokens that can be eligible for rewards
	TokenCapacity                    string                                    // Total token capacity = delegatedCapacity + stakedTokens
	AvailableStake                   string                                    // Stake available to earn rewards. tokenCapacity - allocationTokens - lockedTokens
	Delegators                       []struct{ Delegator struct{ ID string } } // Delegators to this Indexer
	DelegatedTokens                  string                                    // Total tokens delegated to the indexer
	OwnStakeRatio                    string                                    // Ratio between the amount of delegated stake over the total usable stake.
	DelegatedStakeRatio              string                                    // Total shares of the delegator pool
	DelegatorShares                  string                                    // Total shares of the delegator pool
	DelegationExchangeRate           string                                    // bigDecimal (?) Exchange rate of of tokens received for each share
	IndexingRewardCut                int                                       // The percent the Indexer agrees to share with delegators. In parts per million
	IndexingRewardEffectiveCut       string                                    // The percent of reward dilution delegators experience because of overdelegation. Overdelegated stake can't be used to generate rewards but still gets accounted while distributing the generated rewards. This causes dilution of the rewards for the rest of the pool.
	OverDelegationDilution           string                                    // The total amount of query fees given to delegators
	DelegatorQueryFees               string                                    // The total amount of query fees given to delegators
	QueryFeeCut                      int                                       // The percent of query rebate rewards the Indexer agrees to share with delegators. In parts per million
	DelegatorParameterCooldown       int                                       // Amount of blocks a delegator chooses for the waiting period for changing their params
	LastDelegationParameterUpdate    int                                       // Block number for the last time the delegator updated their parameters
	ForcedClosures                   int                                       // Count of how many times this indexer has been forced to close an allocation
	TotalReturn                      string                                    // BigDecimal // Total return this indexer has earned
	AnnualizedReturn                 string                                    // BigDecimal! // Annualized rate of return for the indexer
	StakingEfficiency                string                                    // BigDecimal! // Staking efficiency of the indexer
	StakedTokensTransferredToL2      string
}

// Allocation info
type Allocation struct {
	ID string // 	Channel Address
	// WAS REMOVED 28.01.2021 Price              string              // Price of the queries
	Indexer            struct{ ID string } // Indexer of this allocation
	ActiveForIndexer   struct{ ID string } // Indexer            // Indexer where this allocation is active
	SubgraphDeployment struct {
		ID           string
		OriginalName string
	} // Subgraph deployment that is being allocated
	AllocatedTokens          string              // Tokens associated with the allocation
	EffectiveAllocation      string              // Effective allocation that is realized upon closing
	CreatedAtEpoch           int                 // Epoch this allocation was created
	CreatedAtBlockHash       string              // Bytes! // Block at which this allocation was created
	ClosedAtEpoch            int                 // Epoch this allocation was closed in
	ClosedAtBlockHash        string              // Bytes Block at which this allocation was closed
	QueryFeesCollected       string              // Fees this allocation collected from query fees upon closing. Excludes curator reward and rebate claimed
	QueryFeeRebates          string              // Rebate amount claimed from the protocol
	CuratorRewards           string              // Curator rewards deposited to the curating bonding curve
	IndexingRewards          string              // Indexing rewards earned by this allocation. Includes delegator and indexer rewards
	IndexingIndexerRewards   string              // Indexing rewards earned by this allocation by indexers
	IndexingDelegatorRewards string              // Indexing rewards earned by this allocation by delegators
	PoolClosedIn             struct{ ID string } // Pool               // Pool in which this allocation was closed
	DelegationFees           string              // Fees paid out to delegators
	Status                   string              // AllocationStatus   // Status of the allocation
	CreatedAt                int                 // Timestamp this allocation was created at
	Poi                      string              // Bytes // POI submitted with a closed allocation
	TotalReturn              string              // BigDecimal! // Return for this allocation
	AnnualizedReturn         string              // BigDecimal! // Yearly annualzied return
}

// Pool info
type Pool struct {
	ID                string       // Epoch number of the pool
	Allocation        big.Int      // Total allocations closed in this epoch - includes all effective allocation
	TotalQueryFees    big.Int      // Total query fees collected in this epoch
	ClaimedFees       big.Int      // Total query fees claimed in this epoch. Can be smaller than totalFees because of cobbs douglas function
	CuratorRewards    big.Int      // Total rewards deposited to all curator bonding curves during the epoch
	ClosedAllocations []Allocation // Allocations that were closed during this epoch
}

// Subgraph info
type Subgraph struct {
	ID                 string            // Subgraph ID - which is derived from the Organization/Individual graph accountID
	Owner              GraphAccount      // Graph account that owns this subgraph
	CurrentVersion     SubgraphVersion   // Current version. Null if the subgraph is deprecated
	PastVersions       []SubgraphVersion // Past versions
	CreatedAt          int               // Creation timestamp
	UpdatedAt          int               // Updated timestamp
	SignalledTokens    string            // Tokens signalled for the name all time
	UnsignalledTokens  string            // Tokens signalled for the name
	NameSignalAmount   string            // Current name signal amount for this subgraph
	ReserveRatio       int               // Reserve ratio of the name curation curve
	WithdrawableTokens string            // Tokens that can be withdrawn once the name signal is deprecated
	WithdrawnTokens    string            // Tokens the curators have withdrawn from the deprecated name signal
	NameSignals        []NameSignal      // Curators of this subgraph deployment
	MetadataHash       string            // Bytes Subgraph metadata
	Description        string            // Short description of the subgraph
	Image              string            // Image in string format
	CodeRepository     string            // Location of the code for this project
	Website            string            // Projects website
	DisplayName        string            // Display name
	// WAS REMOVED 28.01.2021 TotalIndexingRewards    big.Int           // Total inflation rewards earned all time by this subgraph
	// WAS REMOVED 28.01.2021 TotalQueryFeesCollected big.Int // Total query fees earned by this subgraph
}

// SubgraphVersion info
type SubgraphVersion struct {
	ID                 string             // 	Concatenation of subgraph, subgraph deployment, and version ID
	Subgraph           string             // Subgraph           // Subgraph of this version
	SubgraphDeployment SubgraphDeployment // Subgraph deployments of this version
	Version            int                // Version number
	CreatedAt          int                // Creation timestamp
	MetadataHash       string             // Bytes Subgraph version metadata content address
	Description        string             // Short description of the version
	Label              string             // A human readable name
}

// SubgraphDeployment info
type SubgraphDeployment struct {
	ID string // 	Subgraph Deployment ID. The IPFS hash with Qm removed to fit into 32 bytes
	// Versions                      struct{ ID string } //[]SubgraphVersion // The versions this subgraph deployment relates to
	CreatedAt    int    // Creation timestamp
	DeniedAt     int    // The block at which this deployment was denied for rewards. Null if not denied
	OriginalName string // The original Subgraph that was deployed through GNS. Can be null if never created through GNS
	StakedTokens string // Total stake of all indexers on this Subgraph deployment
	// IndexerAllocations            struct{ ID string } //[]Allocation        // Allocations created by indexers for this subgraph
	IndexingRewardAmount          string // Total rewards accrued all time by this subgraph deployment. Includes delegation and indexer rewards
	IndexingIndexerRewardAmount   string // Total rewards accrued all time by indexers by this subgraph deployment
	IndexingDelegatorRewardAmount string // Total rewards accrued all time by delegators by this subgraph deployment
	QueryFeesAmount               string // Total query fees earned by this subgraph deployment
	QueryFeeRebates               string // Total query fee rebates earned from the protocol by this subgraph deployment
	CuratorFeeRewards             string // Total curator rewards from fees
	SignalledTokens               string // Signalled tokens
	UnsignalledTokens             string // Unsignalled tokens
	SignalAmount                  string // Current curation signal for this subgraph deployment
	// CuratorSignals                []Signal            // Curators of this subgraph deployment
	ReserveRatio        int // Bonding curve reserve ratio
	IpfsHash            string
	ActiveSubgraphCount int // Amount of active Subgraph entities that are currently using this deployment. Deprecated subgraph entities are not counted
}

// NameSignal info
type NameSignal struct {
	ID       string  // Eth address + subgraph ID
	Curator  Curator // Eth address of the curator
	Subgraph struct {
		ID             string
		CurrentVersion struct {
			SubgraphDeployment SubgraphDeployment
		}
	} // Subgraph being signalled
	SignalledTokens           string // Cumulative number of tokens the curator has signalled
	UnsignalledTokens         string // Cumulative number of tokens the curator has unsignalled
	WithdrawnTokens           string // Tokens the curator has withdrawn from a deprecated name curve
	NameSignal                string // Signal that the curator has from signaling their GRT
	Signal                    string // Actual signal shares that the name pool minted with the GRT provided by the curator
	LastNameSignalChange      int    // Block for which the curator last entered or exited the curve
	RealizedRewards           string // Summation of realized rewards from before the last time the curator entered the curation curve
	AverageCostBasis          string // BigDecimal! // Curator average cost basis for this name signal on this subgraph
	AverageCostBasisPerSignal string // BigDecimal! // averageCostBasis / nameSignal
}

// Signal info
type Signal struct {
	ID                 string             // Eth address + subgraph deployment ID
	Curator            Curator            // Eth address of the curator
	SubgraphDeployment SubgraphDeployment // Subgraph being signalled
	SignalledTokens    string             // Cumulative number of tokens the curator has signalled
	UnsignalledTokens  string             // Cumulative number of tokens the curator has unsignalled
	Signal             string             // Signal that the curator has from signaling their GRT
	LastSignalChange   int                // Block for which the curator last entered or exited the curve
	RealizedRewards    string             // Summation of realized rewards from before the last time the curator entered the curation curve
}

// Delegator info
type Delegator struct {
	ID      string              // 	Delegator address
	Account struct{ ID string } // GraphAccount   // Graph account of the delegator
	// Stakes               {struct //DelegatedStake      // Stakes of this delegator
	TotalStakedTokens    string // Summation of all staked tokens in DelegatorStakes of this Delegator. Not equivalent to balance
	TotalUnstakedTokens  string // Summation of all unstaked tokens in DelegatorStakes of this Delegator
	CreatedAt            int    // Time created at
	TotalRealizedRewards string // float BigDecimal // Total realized rewards on all delegated stakes
}

// DelegatedStake store delegations info
type DelegatedStake struct {
	ID      string // Concatenation of Delegator address and Indexer address
	Indexer struct {
		ID                     string
		DelegationExchangeRate string
		DelegatorShares        string
	} // Indexer // Index the stake is delegated to
	Delegator            struct{ ID string } // Delegator // Delegator
	StakedTokens         string              // Amount delegated all time (static)
	UnstakedTokens       string              // Amount un-delegated all time (static)
	LockedTokens         string              // Amount currently locked
	LockedUntil          int                 // Epoch the locked tokens are unlocked
	ShareAmount          string              // Shares owned in the delegator pool. Used to calculate total amount delegated
	PersonalExchangeRate string              // BigDecimal // The rate this delegator paid for their shares (calculated using average cost basis). Used for rewards calculations
	RealizedRewards      string              // BigDecimal! // Realized rewards from selling delegated stake that accumulated rewards
	CreatedAt            int                 // Time this delegator first delegated to an indexer
}

// GraphAccountName info
type GraphAccountName struct {
	ID         string // Name system concatenated with the unique ID of the name system
	NameSystem string // NameSystem! // Name system for this name
	Name       string // Name from the system
	// GraphAccount GraphAccount // The graph account that owned the name when it was linked in the graph network
}

// GraphAccount info
type GraphAccount struct {
	ID string // Graph account ID
	//	Names              []GraphAccountName // All names this graph account has claimed from all name systems
	//	DefaultName        GraphAccountName   // Default name the graph account has chosen
	CreatedAt          int      // Time the account was created
	DefaultDisplayName string   // Default display name is the current default name. Used for filtered queries
	IsOrganization     bool     // True if it is an organization. False if it is an individual
	MetadataHash       string   // Bytes IPFS hash with account metadata details
	CodeRepository     string   // Main repository of code for the graph account
	Description        string   // Description of the graph account
	Image              string   // Image URL
	Website            string   // Website URL
	DisplayName        string   // Display name. Not unique
	OperatorOf         []string // Operator of other Graph Accounts
	Operators          []string // Operators of this Graph Accounts
	Balance            big.Int  // Graph token balance
	CurationApproval   big.Int  // Amount this account has approved staking to transfer their GRT
	StakingApproval    big.Int  // Amount this account has approved curation to transfer their GRT
	GnsApproval        big.Int  // Amount this account has approved the GNS to transfer their GRT
	//	Subgraphs          []Subgraph         // Subgraphs the graph account owns
	DeveloperCreatedAt int // Time that this graph account became a developer
	//	SubgraphQueryFees  big.Int            // Total query fees the subgraphs created by this account have accumulated in GRT
	//	CreatedDisputes    Dispute            // Disputes this graph account has created
	//	DisputesAgainst    Dispute            // Disputes against this graph account
	//	Curator            Curator            // Subgraphs this account curates
	// Indexer                Indexer             //Indexer            // Subgraphs this account indexes
	//	Delegator              Delegator // Subgraphs this account has delegated to
	//	NameSignalTransactions []NameSignalTransaction
}

// NameSignalTransaction info
type NameSignalTransaction struct {
	ID            string
	BlockNumber   int
	Timestamp     int
	Signer        GraphAccount
	Type          string // TransactionType
	NameSignal    big.Int
	VersionSignal big.Int
	Tokens        big.Int
	Subgraph      Subgraph
}

// Dispute info
type Dispute struct {
	ID                 string             // 	Dispute ID
	SubgraphDeployment SubgraphDeployment // Subgraph deployment being disputed
	Fisherman          string             // GraphAccount // Fisherman address
	Deposit            big.Int            // Fisherman deposit
	CreatedAt          int                // Time dispute was created
	ClosedAt           int                // Time dispute was closed at
	Status             string             // DisputeStatus      // Status of the dispute. Accepted means the Indexer was slashed.
	TokensSlashed      big.Int            // BigDecimal! // Tokens slashed
	TokensRewarded     big.Int            // Tokens rewarded
	Type               string             // DisputeType        // Type of dispute
	Indexer            string             // GraphAccount // Indexer(s) disputed. Conflicting disputes have 2 indexers. Others have 1
	Attestation        Attestation        // Only for single query and conflicting attestation
	LinkedDispute      string             // Dispute // Only for conflicting attestation
	Allocation         Allocation         // Allocation ID. Only for Indexing Disputes
}

// Attestation info
type Attestation struct {
	ID                 string             // 	Concatenation of the requestCID and responseCID
	SubgraphDeployment SubgraphDeployment // Subgraph deployment being disputed
	RequestCID         string             // RequestCID
	ResponseCID        string             // ResponseCID
	GasUsed            big.Int            // Gas used by the attested query
	ResponseNumBytes   big.Int            // Bytes of attested query
	V                  int                // V of the indexers signature
	R                  string             // R of the indexers signature
	S                  string             // S of the indexers signature
}

// Curator info
type Curator struct {
	ID                     string              // 	Eth address of the Curator
	CreatedAt              int                 // Time this curator was created
	Account                struct{ ID string } // GraphAccount // Graph account of this curator
	TotalSignalledTokens   string              // Total tokens signalled on all the subgraphs
	TotalUnsignalledTokens string              // Total tokens unsignalled on all the subgraphs
	// Signals                            []Signal     // Subgraphs the curator is curating
	TotalNameSignalledTokens   string // Total tokens signalled on all names
	TotalNameUnsignalledTokens string // Total tokens unsignalled on all names
	TotalWithdrawnTokens       string // Total withdrawn tokens from deprecated subgraphs
	// NameSignals                        []NameSignal // Subgraphs the curator is curating
	RealizedRewards                    string // Summation of realized rewards from all Signals
	AnnualizedReturn                   string // BigDecimal! // Annualized rate of return on curator signal
	TotalReturn                        string // BigDecimal! // Total return of the curator
	SignalingEfficiency                string // BigDecimal! // Signaling efficiency of the curator
	TotalNameSignal                    string // BigDecimal! // Total name signal summed for all bonding curves
	TotalNameSignalAverageCostBasis    string // BigDecimal! // Total curator cost basis of all shares purchased on all bonding curves
	TotalAverageCostBasisPerNameSignal string // BigDecimal! // totalNameSignalCostBasis / totalNameSignal
}

// TokenLockWallet info
type TokenLockWallet struct {
	ID                        string
	Manager                   string
	InitHash                  string
	Beneficiary               string
	Token                     string
	ManagedAmount             string
	StartTime                 string
	EndTime                   string
	Periods                   string
	ReleaseStartTime          string
	VestingCliffTime          string
	Revocable                 string
	TokenDestinationsApproved graphql.Boolean
	TokensReleased            string
	TokensWithdrawn           string
	TokensRevoked             string
	BlockNumberCreated        string
	TxHash                    string
}

type IndexerDataPoint struct {
	ID                            string
	AvgIndexerBlocksBehind        string `graphql:"avg_indexer_blocks_behind"`
	AvgIndexerLatencyMs           string `graphql:"avg_indexer_latency_ms"`
	AvgQueryFee                   string `graphql:"avg_query_fee"`
	StartEpoch                    string `graphql:"start_epoch"`
	EndEpoch                      string `graphql:"end_epoch"`
	MaxIndexerBlocksBehind        string `graphql:"max_indexer_blocks_behind"`
	MaxIndexerLatencyMs           string `graphql:"max_indexer_latency_ms"`
	MaxQueryFee                   string `graphql:"max_query_fee"`
	NumUbdexer200Responses        string `graphql:"num_indexer_200_responses"`
	ProportionIndexer200Responses string `graphql:"proportion_indexer_200_responses"`
	QueryCount                    string `graphql:"query_count"`
	StdevIndexerLatencyMs         string `graphql:"stdev_indexer_latency_ms"`
	SubgraphDeploymentIpshHash    string `graphql:"subgraph_deployment_ipfs_hash"`
	TotalQueryFees                string `graphql:"total_query_fees"`
	IndexerWallet                 string `graphql:"indexer_wallet"`
}
