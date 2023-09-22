package ethereum

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stakemachine/thegraph-exporter/ethereum/contracts/rewards"
	"github.com/stakemachine/thegraph-exporter/ethereum/erc20"
)

// GetRewards get rewards
func (es *EthService) GetRewards(alloID, rewardsManagerContractAddress string) (*hexutil.Big, error) {
	contractAddress := common.HexToAddress(rewardsManagerContractAddress)
	allocationID := common.HexToAddress(alloID)
	instance, err := rewards.NewRewards(contractAddress, es.Client)
	if err != nil {
		return nil, errors.New("allocation ID: " + allocationID.String() + " " + err.Error())
	}
	balance, err := instance.GetRewards(&bind.CallOpts{}, allocationID)
	if err != nil {
		return nil, errors.New("allocation ID: " + allocationID.String() + " " + err.Error())
	}
	return (*hexutil.Big)(balance), nil
}

// GetTokenBalance gets ERC20 balance // 0xc944e90c64b2c07662a292be6244bdf05cda44a7
func (es *EthService) GetTokenBalance(tokenHex, addressHex string) (*hexutil.Big, error) {
	tokenAddress := common.HexToAddress(tokenHex)
	address := common.HexToAddress(addressHex)
	instance, err := erc20.NewErc20(tokenAddress, es.Client)
	if err != nil {
		return nil, err
	}
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(balance), nil
}
