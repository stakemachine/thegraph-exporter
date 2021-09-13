// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewards

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RewardsMetaData contains all meta data concerning the Rewards contract.
var RewardsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"param\",\"type\":\"string\"}],\"name\":\"ParameterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"RewardsDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sinceBlock\",\"type\":\"uint256\"}],\"name\":\"RewardsDenylistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"SetController\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accRewardsPerSignal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accRewardsPerSignalLastBlockUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGraphProxy\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"acceptProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGraphProxy\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"acceptProxyAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addressCache\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"denylist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"getAccRewardsForSubgraph\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"getAccRewardsPerAllocatedToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccRewardsPerSignal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewRewardsPerSignal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_issuanceRate\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"isDenied\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuanceRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"onSubgraphAllocationUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"onSubgraphSignalUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_deny\",\"type\":\"bool\"}],\"name\":\"setDenied\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"_deny\",\"type\":\"bool[]\"}],\"name\":\"setDeniedMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_issuanceRate\",\"type\":\"uint256\"}],\"name\":\"setIssuanceRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subgraphAvailabilityOracle\",\"type\":\"address\"}],\"name\":\"setSubgraphAvailabilityOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subgraphAvailabilityOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subgraphs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accRewardsForSubgraph\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsForSubgraphSnapshot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsPerSignalSnapshot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsPerAllocatedToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"takeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAccRewardsPerSignal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsMetaData.ABI instead.
var RewardsABI = RewardsMetaData.ABI

// Rewards is an auto generated Go binding around an Ethereum contract.
type Rewards struct {
	RewardsCaller     // Read-only binding to the contract
	RewardsTransactor // Write-only binding to the contract
	RewardsFilterer   // Log filterer for contract events
}

// RewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsSession struct {
	Contract     *Rewards          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCallerSession struct {
	Contract *RewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsTransactorSession struct {
	Contract     *RewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsRaw struct {
	Contract *Rewards // Generic contract binding to access the raw methods on
}

// RewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCallerRaw struct {
	Contract *RewardsCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsTransactorRaw struct {
	Contract *RewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewards creates a new instance of Rewards, bound to a specific deployed contract.
func NewRewards(address common.Address, backend bind.ContractBackend) (*Rewards, error) {
	contract, err := bindRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// NewRewardsCaller creates a new read-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsCaller(address common.Address, caller bind.ContractCaller) (*RewardsCaller, error) {
	contract, err := bindRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCaller{contract: contract}, nil
}

// NewRewardsTransactor creates a new write-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsTransactor, error) {
	contract, err := bindRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsTransactor{contract: contract}, nil
}

// NewRewardsFilterer creates a new log filterer instance of Rewards, bound to a specific deployed contract.
func NewRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsFilterer, error) {
	contract, err := bindRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsFilterer{contract: contract}, nil
}

// bindRewards binds a generic wrapper to an already deployed contract.
func bindRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.RewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transact(opts, method, params...)
}

// AccRewardsPerSignal is a free data retrieval call binding the contract method 0xe242cf1e.
//
// Solidity: function accRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsCaller) AccRewardsPerSignal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "accRewardsPerSignal")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// AccRewardsPerSignal is a free data retrieval call binding the contract method 0xe242cf1e.
//
// Solidity: function accRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsSession) AccRewardsPerSignal() (*big.Int, error) {
	return _Rewards.Contract.AccRewardsPerSignal(&_Rewards.CallOpts)
}

// AccRewardsPerSignal is a free data retrieval call binding the contract method 0xe242cf1e.
//
// Solidity: function accRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsCallerSession) AccRewardsPerSignal() (*big.Int, error) {
	return _Rewards.Contract.AccRewardsPerSignal(&_Rewards.CallOpts)
}

// AccRewardsPerSignalLastBlockUpdated is a free data retrieval call binding the contract method 0x9006ce8b.
//
// Solidity: function accRewardsPerSignalLastBlockUpdated() view returns(uint256)
func (_Rewards *RewardsCaller) AccRewardsPerSignalLastBlockUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "accRewardsPerSignalLastBlockUpdated")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// AccRewardsPerSignalLastBlockUpdated is a free data retrieval call binding the contract method 0x9006ce8b.
//
// Solidity: function accRewardsPerSignalLastBlockUpdated() view returns(uint256)
func (_Rewards *RewardsSession) AccRewardsPerSignalLastBlockUpdated() (*big.Int, error) {
	return _Rewards.Contract.AccRewardsPerSignalLastBlockUpdated(&_Rewards.CallOpts)
}

// AccRewardsPerSignalLastBlockUpdated is a free data retrieval call binding the contract method 0x9006ce8b.
//
// Solidity: function accRewardsPerSignalLastBlockUpdated() view returns(uint256)
func (_Rewards *RewardsCallerSession) AccRewardsPerSignalLastBlockUpdated() (*big.Int, error) {
	return _Rewards.Contract.AccRewardsPerSignalLastBlockUpdated(&_Rewards.CallOpts)
}

// AddressCache is a free data retrieval call binding the contract method 0xdc675a65.
//
// Solidity: function addressCache(bytes32 ) view returns(address)
func (_Rewards *RewardsCaller) AddressCache(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "addressCache", arg0)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// AddressCache is a free data retrieval call binding the contract method 0xdc675a65.
//
// Solidity: function addressCache(bytes32 ) view returns(address)
func (_Rewards *RewardsSession) AddressCache(arg0 [32]byte) (common.Address, error) {
	return _Rewards.Contract.AddressCache(&_Rewards.CallOpts, arg0)
}

// AddressCache is a free data retrieval call binding the contract method 0xdc675a65.
//
// Solidity: function addressCache(bytes32 ) view returns(address)
func (_Rewards *RewardsCallerSession) AddressCache(arg0 [32]byte) (common.Address, error) {
	return _Rewards.Contract.AddressCache(&_Rewards.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Rewards *RewardsCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "controller")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Rewards *RewardsSession) Controller() (common.Address, error) {
	return _Rewards.Contract.Controller(&_Rewards.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Rewards *RewardsCallerSession) Controller() (common.Address, error) {
	return _Rewards.Contract.Controller(&_Rewards.CallOpts)
}

// Denylist is a free data retrieval call binding the contract method 0x16a84ab2.
//
// Solidity: function denylist(bytes32 ) view returns(uint256)
func (_Rewards *RewardsCaller) Denylist(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "denylist", arg0)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Denylist is a free data retrieval call binding the contract method 0x16a84ab2.
//
// Solidity: function denylist(bytes32 ) view returns(uint256)
func (_Rewards *RewardsSession) Denylist(arg0 [32]byte) (*big.Int, error) {
	return _Rewards.Contract.Denylist(&_Rewards.CallOpts, arg0)
}

// Denylist is a free data retrieval call binding the contract method 0x16a84ab2.
//
// Solidity: function denylist(bytes32 ) view returns(uint256)
func (_Rewards *RewardsCallerSession) Denylist(arg0 [32]byte) (*big.Int, error) {
	return _Rewards.Contract.Denylist(&_Rewards.CallOpts, arg0)
}

// GetAccRewardsForSubgraph is a free data retrieval call binding the contract method 0x5c6cbd59.
//
// Solidity: function getAccRewardsForSubgraph(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_Rewards *RewardsCaller) GetAccRewardsForSubgraph(opts *bind.CallOpts, _subgraphDeploymentID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getAccRewardsForSubgraph", _subgraphDeploymentID)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetAccRewardsForSubgraph is a free data retrieval call binding the contract method 0x5c6cbd59.
//
// Solidity: function getAccRewardsForSubgraph(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_Rewards *RewardsSession) GetAccRewardsForSubgraph(_subgraphDeploymentID [32]byte) (*big.Int, error) {
	return _Rewards.Contract.GetAccRewardsForSubgraph(&_Rewards.CallOpts, _subgraphDeploymentID)
}

// GetAccRewardsForSubgraph is a free data retrieval call binding the contract method 0x5c6cbd59.
//
// Solidity: function getAccRewardsForSubgraph(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_Rewards *RewardsCallerSession) GetAccRewardsForSubgraph(_subgraphDeploymentID [32]byte) (*big.Int, error) {
	return _Rewards.Contract.GetAccRewardsForSubgraph(&_Rewards.CallOpts, _subgraphDeploymentID)
}

// GetAccRewardsPerAllocatedToken is a free data retrieval call binding the contract method 0x702a280e.
//
// Solidity: function getAccRewardsPerAllocatedToken(bytes32 _subgraphDeploymentID) view returns(uint256, uint256)
func (_Rewards *RewardsCaller) GetAccRewardsPerAllocatedToken(opts *bind.CallOpts, _subgraphDeploymentID [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getAccRewardsPerAllocatedToken", _subgraphDeploymentID)
	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err
}

// GetAccRewardsPerAllocatedToken is a free data retrieval call binding the contract method 0x702a280e.
//
// Solidity: function getAccRewardsPerAllocatedToken(bytes32 _subgraphDeploymentID) view returns(uint256, uint256)
func (_Rewards *RewardsSession) GetAccRewardsPerAllocatedToken(_subgraphDeploymentID [32]byte) (*big.Int, *big.Int, error) {
	return _Rewards.Contract.GetAccRewardsPerAllocatedToken(&_Rewards.CallOpts, _subgraphDeploymentID)
}

// GetAccRewardsPerAllocatedToken is a free data retrieval call binding the contract method 0x702a280e.
//
// Solidity: function getAccRewardsPerAllocatedToken(bytes32 _subgraphDeploymentID) view returns(uint256, uint256)
func (_Rewards *RewardsCallerSession) GetAccRewardsPerAllocatedToken(_subgraphDeploymentID [32]byte) (*big.Int, *big.Int, error) {
	return _Rewards.Contract.GetAccRewardsPerAllocatedToken(&_Rewards.CallOpts, _subgraphDeploymentID)
}

// GetAccRewardsPerSignal is a free data retrieval call binding the contract method 0xa8cc0ee2.
//
// Solidity: function getAccRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsCaller) GetAccRewardsPerSignal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getAccRewardsPerSignal")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetAccRewardsPerSignal is a free data retrieval call binding the contract method 0xa8cc0ee2.
//
// Solidity: function getAccRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsSession) GetAccRewardsPerSignal() (*big.Int, error) {
	return _Rewards.Contract.GetAccRewardsPerSignal(&_Rewards.CallOpts)
}

// GetAccRewardsPerSignal is a free data retrieval call binding the contract method 0xa8cc0ee2.
//
// Solidity: function getAccRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsCallerSession) GetAccRewardsPerSignal() (*big.Int, error) {
	return _Rewards.Contract.GetAccRewardsPerSignal(&_Rewards.CallOpts)
}

// GetNewRewardsPerSignal is a free data retrieval call binding the contract method 0xe284f848.
//
// Solidity: function getNewRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsCaller) GetNewRewardsPerSignal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getNewRewardsPerSignal")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetNewRewardsPerSignal is a free data retrieval call binding the contract method 0xe284f848.
//
// Solidity: function getNewRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsSession) GetNewRewardsPerSignal() (*big.Int, error) {
	return _Rewards.Contract.GetNewRewardsPerSignal(&_Rewards.CallOpts)
}

// GetNewRewardsPerSignal is a free data retrieval call binding the contract method 0xe284f848.
//
// Solidity: function getNewRewardsPerSignal() view returns(uint256)
func (_Rewards *RewardsCallerSession) GetNewRewardsPerSignal() (*big.Int, error) {
	return _Rewards.Contract.GetNewRewardsPerSignal(&_Rewards.CallOpts)
}

// GetRewards is a free data retrieval call binding the contract method 0x79ee54f7.
//
// Solidity: function getRewards(address _allocationID) view returns(uint256)
func (_Rewards *RewardsCaller) GetRewards(opts *bind.CallOpts, _allocationID common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "getRewards", _allocationID)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// GetRewards is a free data retrieval call binding the contract method 0x79ee54f7.
//
// Solidity: function getRewards(address _allocationID) view returns(uint256)
func (_Rewards *RewardsSession) GetRewards(_allocationID common.Address) (*big.Int, error) {
	return _Rewards.Contract.GetRewards(&_Rewards.CallOpts, _allocationID)
}

// GetRewards is a free data retrieval call binding the contract method 0x79ee54f7.
//
// Solidity: function getRewards(address _allocationID) view returns(uint256)
func (_Rewards *RewardsCallerSession) GetRewards(_allocationID common.Address) (*big.Int, error) {
	return _Rewards.Contract.GetRewards(&_Rewards.CallOpts, _allocationID)
}

// IsDenied is a free data retrieval call binding the contract method 0xe820e284.
//
// Solidity: function isDenied(bytes32 _subgraphDeploymentID) view returns(bool)
func (_Rewards *RewardsCaller) IsDenied(opts *bind.CallOpts, _subgraphDeploymentID [32]byte) (bool, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "isDenied", _subgraphDeploymentID)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsDenied is a free data retrieval call binding the contract method 0xe820e284.
//
// Solidity: function isDenied(bytes32 _subgraphDeploymentID) view returns(bool)
func (_Rewards *RewardsSession) IsDenied(_subgraphDeploymentID [32]byte) (bool, error) {
	return _Rewards.Contract.IsDenied(&_Rewards.CallOpts, _subgraphDeploymentID)
}

// IsDenied is a free data retrieval call binding the contract method 0xe820e284.
//
// Solidity: function isDenied(bytes32 _subgraphDeploymentID) view returns(bool)
func (_Rewards *RewardsCallerSession) IsDenied(_subgraphDeploymentID [32]byte) (bool, error) {
	return _Rewards.Contract.IsDenied(&_Rewards.CallOpts, _subgraphDeploymentID)
}

// IssuanceRate is a free data retrieval call binding the contract method 0x3c9ae2ba.
//
// Solidity: function issuanceRate() view returns(uint256)
func (_Rewards *RewardsCaller) IssuanceRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "issuanceRate")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// IssuanceRate is a free data retrieval call binding the contract method 0x3c9ae2ba.
//
// Solidity: function issuanceRate() view returns(uint256)
func (_Rewards *RewardsSession) IssuanceRate() (*big.Int, error) {
	return _Rewards.Contract.IssuanceRate(&_Rewards.CallOpts)
}

// IssuanceRate is a free data retrieval call binding the contract method 0x3c9ae2ba.
//
// Solidity: function issuanceRate() view returns(uint256)
func (_Rewards *RewardsCallerSession) IssuanceRate() (*big.Int, error) {
	return _Rewards.Contract.IssuanceRate(&_Rewards.CallOpts)
}

// SubgraphAvailabilityOracle is a free data retrieval call binding the contract method 0x05bb8c6b.
//
// Solidity: function subgraphAvailabilityOracle() view returns(address)
func (_Rewards *RewardsCaller) SubgraphAvailabilityOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "subgraphAvailabilityOracle")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// SubgraphAvailabilityOracle is a free data retrieval call binding the contract method 0x05bb8c6b.
//
// Solidity: function subgraphAvailabilityOracle() view returns(address)
func (_Rewards *RewardsSession) SubgraphAvailabilityOracle() (common.Address, error) {
	return _Rewards.Contract.SubgraphAvailabilityOracle(&_Rewards.CallOpts)
}

// SubgraphAvailabilityOracle is a free data retrieval call binding the contract method 0x05bb8c6b.
//
// Solidity: function subgraphAvailabilityOracle() view returns(address)
func (_Rewards *RewardsCallerSession) SubgraphAvailabilityOracle() (common.Address, error) {
	return _Rewards.Contract.SubgraphAvailabilityOracle(&_Rewards.CallOpts)
}

// Subgraphs is a free data retrieval call binding the contract method 0x4986594f.
//
// Solidity: function subgraphs(bytes32 ) view returns(uint256 accRewardsForSubgraph, uint256 accRewardsForSubgraphSnapshot, uint256 accRewardsPerSignalSnapshot, uint256 accRewardsPerAllocatedToken)
func (_Rewards *RewardsCaller) Subgraphs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	AccRewardsForSubgraph         *big.Int
	AccRewardsForSubgraphSnapshot *big.Int
	AccRewardsPerSignalSnapshot   *big.Int
	AccRewardsPerAllocatedToken   *big.Int
}, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "subgraphs", arg0)

	outstruct := new(struct {
		AccRewardsForSubgraph         *big.Int
		AccRewardsForSubgraphSnapshot *big.Int
		AccRewardsPerSignalSnapshot   *big.Int
		AccRewardsPerAllocatedToken   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccRewardsForSubgraph = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccRewardsForSubgraphSnapshot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccRewardsPerSignalSnapshot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccRewardsPerAllocatedToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err
}

// Subgraphs is a free data retrieval call binding the contract method 0x4986594f.
//
// Solidity: function subgraphs(bytes32 ) view returns(uint256 accRewardsForSubgraph, uint256 accRewardsForSubgraphSnapshot, uint256 accRewardsPerSignalSnapshot, uint256 accRewardsPerAllocatedToken)
func (_Rewards *RewardsSession) Subgraphs(arg0 [32]byte) (struct {
	AccRewardsForSubgraph         *big.Int
	AccRewardsForSubgraphSnapshot *big.Int
	AccRewardsPerSignalSnapshot   *big.Int
	AccRewardsPerAllocatedToken   *big.Int
}, error) {
	return _Rewards.Contract.Subgraphs(&_Rewards.CallOpts, arg0)
}

// Subgraphs is a free data retrieval call binding the contract method 0x4986594f.
//
// Solidity: function subgraphs(bytes32 ) view returns(uint256 accRewardsForSubgraph, uint256 accRewardsForSubgraphSnapshot, uint256 accRewardsPerSignalSnapshot, uint256 accRewardsPerAllocatedToken)
func (_Rewards *RewardsCallerSession) Subgraphs(arg0 [32]byte) (struct {
	AccRewardsForSubgraph         *big.Int
	AccRewardsForSubgraphSnapshot *big.Int
	AccRewardsPerSignalSnapshot   *big.Int
	AccRewardsPerAllocatedToken   *big.Int
}, error) {
	return _Rewards.Contract.Subgraphs(&_Rewards.CallOpts, arg0)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_Rewards *RewardsTransactor) AcceptProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "acceptProxy", _proxy)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_Rewards *RewardsSession) AcceptProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.AcceptProxy(&_Rewards.TransactOpts, _proxy)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_Rewards *RewardsTransactorSession) AcceptProxy(_proxy common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.AcceptProxy(&_Rewards.TransactOpts, _proxy)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_Rewards *RewardsTransactor) AcceptProxyAndCall(opts *bind.TransactOpts, _proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "acceptProxyAndCall", _proxy, _data)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_Rewards *RewardsSession) AcceptProxyAndCall(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _Rewards.Contract.AcceptProxyAndCall(&_Rewards.TransactOpts, _proxy, _data)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_Rewards *RewardsTransactorSession) AcceptProxyAndCall(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _Rewards.Contract.AcceptProxyAndCall(&_Rewards.TransactOpts, _proxy, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _controller, uint256 _issuanceRate) returns()
func (_Rewards *RewardsTransactor) Initialize(opts *bind.TransactOpts, _controller common.Address, _issuanceRate *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "initialize", _controller, _issuanceRate)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _controller, uint256 _issuanceRate) returns()
func (_Rewards *RewardsSession) Initialize(_controller common.Address, _issuanceRate *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Initialize(&_Rewards.TransactOpts, _controller, _issuanceRate)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _controller, uint256 _issuanceRate) returns()
func (_Rewards *RewardsTransactorSession) Initialize(_controller common.Address, _issuanceRate *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Initialize(&_Rewards.TransactOpts, _controller, _issuanceRate)
}

// OnSubgraphAllocationUpdate is a paid mutator transaction binding the contract method 0xeeac3e0e.
//
// Solidity: function onSubgraphAllocationUpdate(bytes32 _subgraphDeploymentID) returns(uint256)
func (_Rewards *RewardsTransactor) OnSubgraphAllocationUpdate(opts *bind.TransactOpts, _subgraphDeploymentID [32]byte) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "onSubgraphAllocationUpdate", _subgraphDeploymentID)
}

// OnSubgraphAllocationUpdate is a paid mutator transaction binding the contract method 0xeeac3e0e.
//
// Solidity: function onSubgraphAllocationUpdate(bytes32 _subgraphDeploymentID) returns(uint256)
func (_Rewards *RewardsSession) OnSubgraphAllocationUpdate(_subgraphDeploymentID [32]byte) (*types.Transaction, error) {
	return _Rewards.Contract.OnSubgraphAllocationUpdate(&_Rewards.TransactOpts, _subgraphDeploymentID)
}

// OnSubgraphAllocationUpdate is a paid mutator transaction binding the contract method 0xeeac3e0e.
//
// Solidity: function onSubgraphAllocationUpdate(bytes32 _subgraphDeploymentID) returns(uint256)
func (_Rewards *RewardsTransactorSession) OnSubgraphAllocationUpdate(_subgraphDeploymentID [32]byte) (*types.Transaction, error) {
	return _Rewards.Contract.OnSubgraphAllocationUpdate(&_Rewards.TransactOpts, _subgraphDeploymentID)
}

// OnSubgraphSignalUpdate is a paid mutator transaction binding the contract method 0x1d1c2fec.
//
// Solidity: function onSubgraphSignalUpdate(bytes32 _subgraphDeploymentID) returns(uint256)
func (_Rewards *RewardsTransactor) OnSubgraphSignalUpdate(opts *bind.TransactOpts, _subgraphDeploymentID [32]byte) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "onSubgraphSignalUpdate", _subgraphDeploymentID)
}

// OnSubgraphSignalUpdate is a paid mutator transaction binding the contract method 0x1d1c2fec.
//
// Solidity: function onSubgraphSignalUpdate(bytes32 _subgraphDeploymentID) returns(uint256)
func (_Rewards *RewardsSession) OnSubgraphSignalUpdate(_subgraphDeploymentID [32]byte) (*types.Transaction, error) {
	return _Rewards.Contract.OnSubgraphSignalUpdate(&_Rewards.TransactOpts, _subgraphDeploymentID)
}

// OnSubgraphSignalUpdate is a paid mutator transaction binding the contract method 0x1d1c2fec.
//
// Solidity: function onSubgraphSignalUpdate(bytes32 _subgraphDeploymentID) returns(uint256)
func (_Rewards *RewardsTransactorSession) OnSubgraphSignalUpdate(_subgraphDeploymentID [32]byte) (*types.Transaction, error) {
	return _Rewards.Contract.OnSubgraphSignalUpdate(&_Rewards.TransactOpts, _subgraphDeploymentID)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Rewards *RewardsTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Rewards *RewardsSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetController(&_Rewards.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Rewards *RewardsTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetController(&_Rewards.TransactOpts, _controller)
}

// SetDenied is a paid mutator transaction binding the contract method 0x1324a506.
//
// Solidity: function setDenied(bytes32 _subgraphDeploymentID, bool _deny) returns()
func (_Rewards *RewardsTransactor) SetDenied(opts *bind.TransactOpts, _subgraphDeploymentID [32]byte, _deny bool) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setDenied", _subgraphDeploymentID, _deny)
}

// SetDenied is a paid mutator transaction binding the contract method 0x1324a506.
//
// Solidity: function setDenied(bytes32 _subgraphDeploymentID, bool _deny) returns()
func (_Rewards *RewardsSession) SetDenied(_subgraphDeploymentID [32]byte, _deny bool) (*types.Transaction, error) {
	return _Rewards.Contract.SetDenied(&_Rewards.TransactOpts, _subgraphDeploymentID, _deny)
}

// SetDenied is a paid mutator transaction binding the contract method 0x1324a506.
//
// Solidity: function setDenied(bytes32 _subgraphDeploymentID, bool _deny) returns()
func (_Rewards *RewardsTransactorSession) SetDenied(_subgraphDeploymentID [32]byte, _deny bool) (*types.Transaction, error) {
	return _Rewards.Contract.SetDenied(&_Rewards.TransactOpts, _subgraphDeploymentID, _deny)
}

// SetDeniedMany is a paid mutator transaction binding the contract method 0x1debaded.
//
// Solidity: function setDeniedMany(bytes32[] _subgraphDeploymentID, bool[] _deny) returns()
func (_Rewards *RewardsTransactor) SetDeniedMany(opts *bind.TransactOpts, _subgraphDeploymentID [][32]byte, _deny []bool) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setDeniedMany", _subgraphDeploymentID, _deny)
}

// SetDeniedMany is a paid mutator transaction binding the contract method 0x1debaded.
//
// Solidity: function setDeniedMany(bytes32[] _subgraphDeploymentID, bool[] _deny) returns()
func (_Rewards *RewardsSession) SetDeniedMany(_subgraphDeploymentID [][32]byte, _deny []bool) (*types.Transaction, error) {
	return _Rewards.Contract.SetDeniedMany(&_Rewards.TransactOpts, _subgraphDeploymentID, _deny)
}

// SetDeniedMany is a paid mutator transaction binding the contract method 0x1debaded.
//
// Solidity: function setDeniedMany(bytes32[] _subgraphDeploymentID, bool[] _deny) returns()
func (_Rewards *RewardsTransactorSession) SetDeniedMany(_subgraphDeploymentID [][32]byte, _deny []bool) (*types.Transaction, error) {
	return _Rewards.Contract.SetDeniedMany(&_Rewards.TransactOpts, _subgraphDeploymentID, _deny)
}

// SetIssuanceRate is a paid mutator transaction binding the contract method 0xfc24ffdf.
//
// Solidity: function setIssuanceRate(uint256 _issuanceRate) returns()
func (_Rewards *RewardsTransactor) SetIssuanceRate(opts *bind.TransactOpts, _issuanceRate *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setIssuanceRate", _issuanceRate)
}

// SetIssuanceRate is a paid mutator transaction binding the contract method 0xfc24ffdf.
//
// Solidity: function setIssuanceRate(uint256 _issuanceRate) returns()
func (_Rewards *RewardsSession) SetIssuanceRate(_issuanceRate *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetIssuanceRate(&_Rewards.TransactOpts, _issuanceRate)
}

// SetIssuanceRate is a paid mutator transaction binding the contract method 0xfc24ffdf.
//
// Solidity: function setIssuanceRate(uint256 _issuanceRate) returns()
func (_Rewards *RewardsTransactorSession) SetIssuanceRate(_issuanceRate *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetIssuanceRate(&_Rewards.TransactOpts, _issuanceRate)
}

// SetSubgraphAvailabilityOracle is a paid mutator transaction binding the contract method 0x0903c094.
//
// Solidity: function setSubgraphAvailabilityOracle(address _subgraphAvailabilityOracle) returns()
func (_Rewards *RewardsTransactor) SetSubgraphAvailabilityOracle(opts *bind.TransactOpts, _subgraphAvailabilityOracle common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setSubgraphAvailabilityOracle", _subgraphAvailabilityOracle)
}

// SetSubgraphAvailabilityOracle is a paid mutator transaction binding the contract method 0x0903c094.
//
// Solidity: function setSubgraphAvailabilityOracle(address _subgraphAvailabilityOracle) returns()
func (_Rewards *RewardsSession) SetSubgraphAvailabilityOracle(_subgraphAvailabilityOracle common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetSubgraphAvailabilityOracle(&_Rewards.TransactOpts, _subgraphAvailabilityOracle)
}

// SetSubgraphAvailabilityOracle is a paid mutator transaction binding the contract method 0x0903c094.
//
// Solidity: function setSubgraphAvailabilityOracle(address _subgraphAvailabilityOracle) returns()
func (_Rewards *RewardsTransactorSession) SetSubgraphAvailabilityOracle(_subgraphAvailabilityOracle common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetSubgraphAvailabilityOracle(&_Rewards.TransactOpts, _subgraphAvailabilityOracle)
}

// TakeRewards is a paid mutator transaction binding the contract method 0xdb750926.
//
// Solidity: function takeRewards(address _allocationID) returns(uint256)
func (_Rewards *RewardsTransactor) TakeRewards(opts *bind.TransactOpts, _allocationID common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "takeRewards", _allocationID)
}

// TakeRewards is a paid mutator transaction binding the contract method 0xdb750926.
//
// Solidity: function takeRewards(address _allocationID) returns(uint256)
func (_Rewards *RewardsSession) TakeRewards(_allocationID common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TakeRewards(&_Rewards.TransactOpts, _allocationID)
}

// TakeRewards is a paid mutator transaction binding the contract method 0xdb750926.
//
// Solidity: function takeRewards(address _allocationID) returns(uint256)
func (_Rewards *RewardsTransactorSession) TakeRewards(_allocationID common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TakeRewards(&_Rewards.TransactOpts, _allocationID)
}

// UpdateAccRewardsPerSignal is a paid mutator transaction binding the contract method 0xc7d1117d.
//
// Solidity: function updateAccRewardsPerSignal() returns(uint256)
func (_Rewards *RewardsTransactor) UpdateAccRewardsPerSignal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "updateAccRewardsPerSignal")
}

// UpdateAccRewardsPerSignal is a paid mutator transaction binding the contract method 0xc7d1117d.
//
// Solidity: function updateAccRewardsPerSignal() returns(uint256)
func (_Rewards *RewardsSession) UpdateAccRewardsPerSignal() (*types.Transaction, error) {
	return _Rewards.Contract.UpdateAccRewardsPerSignal(&_Rewards.TransactOpts)
}

// UpdateAccRewardsPerSignal is a paid mutator transaction binding the contract method 0xc7d1117d.
//
// Solidity: function updateAccRewardsPerSignal() returns(uint256)
func (_Rewards *RewardsTransactorSession) UpdateAccRewardsPerSignal() (*types.Transaction, error) {
	return _Rewards.Contract.UpdateAccRewardsPerSignal(&_Rewards.TransactOpts)
}

// RewardsParameterUpdatedIterator is returned from FilterParameterUpdated and is used to iterate over the raw logs and unpacked data for ParameterUpdated events raised by the Rewards contract.
type RewardsParameterUpdatedIterator struct {
	Event *RewardsParameterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsParameterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsParameterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsParameterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsParameterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsParameterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsParameterUpdated represents a ParameterUpdated event raised by the Rewards contract.
type RewardsParameterUpdated struct {
	Param string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParameterUpdated is a free log retrieval operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_Rewards *RewardsFilterer) FilterParameterUpdated(opts *bind.FilterOpts) (*RewardsParameterUpdatedIterator, error) {
	logs, sub, err := _Rewards.contract.FilterLogs(opts, "ParameterUpdated")
	if err != nil {
		return nil, err
	}
	return &RewardsParameterUpdatedIterator{contract: _Rewards.contract, event: "ParameterUpdated", logs: logs, sub: sub}, nil
}

// WatchParameterUpdated is a free log subscription operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_Rewards *RewardsFilterer) WatchParameterUpdated(opts *bind.WatchOpts, sink chan<- *RewardsParameterUpdated) (event.Subscription, error) {
	logs, sub, err := _Rewards.contract.WatchLogs(opts, "ParameterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsParameterUpdated)
				if err := _Rewards.contract.UnpackLog(event, "ParameterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParameterUpdated is a log parse operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_Rewards *RewardsFilterer) ParseParameterUpdated(log types.Log) (*RewardsParameterUpdated, error) {
	event := new(RewardsParameterUpdated)
	if err := _Rewards.contract.UnpackLog(event, "ParameterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardsAssignedIterator is returned from FilterRewardsAssigned and is used to iterate over the raw logs and unpacked data for RewardsAssigned events raised by the Rewards contract.
type RewardsRewardsAssignedIterator struct {
	Event *RewardsRewardsAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsRewardsAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardsAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsRewardsAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsRewardsAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardsAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardsAssigned represents a RewardsAssigned event raised by the Rewards contract.
type RewardsRewardsAssigned struct {
	Indexer      common.Address
	AllocationID common.Address
	Epoch        *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardsAssigned is a free log retrieval operation binding the contract event 0x315d9cdbc182c9118c140c78a121ebb9f24bf73f841339a8a41cdc3586c34e18.
//
// Solidity: event RewardsAssigned(address indexed indexer, address indexed allocationID, uint256 epoch, uint256 amount)
func (_Rewards *RewardsFilterer) FilterRewardsAssigned(opts *bind.FilterOpts, indexer []common.Address, allocationID []common.Address) (*RewardsRewardsAssignedIterator, error) {
	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardsAssigned", indexerRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &RewardsRewardsAssignedIterator{contract: _Rewards.contract, event: "RewardsAssigned", logs: logs, sub: sub}, nil
}

// WatchRewardsAssigned is a free log subscription operation binding the contract event 0x315d9cdbc182c9118c140c78a121ebb9f24bf73f841339a8a41cdc3586c34e18.
//
// Solidity: event RewardsAssigned(address indexed indexer, address indexed allocationID, uint256 epoch, uint256 amount)
func (_Rewards *RewardsFilterer) WatchRewardsAssigned(opts *bind.WatchOpts, sink chan<- *RewardsRewardsAssigned, indexer []common.Address, allocationID []common.Address) (event.Subscription, error) {
	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardsAssigned", indexerRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardsAssigned)
				if err := _Rewards.contract.UnpackLog(event, "RewardsAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsAssigned is a log parse operation binding the contract event 0x315d9cdbc182c9118c140c78a121ebb9f24bf73f841339a8a41cdc3586c34e18.
//
// Solidity: event RewardsAssigned(address indexed indexer, address indexed allocationID, uint256 epoch, uint256 amount)
func (_Rewards *RewardsFilterer) ParseRewardsAssigned(log types.Log) (*RewardsRewardsAssigned, error) {
	event := new(RewardsRewardsAssigned)
	if err := _Rewards.contract.UnpackLog(event, "RewardsAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardsDeniedIterator is returned from FilterRewardsDenied and is used to iterate over the raw logs and unpacked data for RewardsDenied events raised by the Rewards contract.
type RewardsRewardsDeniedIterator struct {
	Event *RewardsRewardsDenied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsRewardsDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardsDenied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsRewardsDenied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsRewardsDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardsDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardsDenied represents a RewardsDenied event raised by the Rewards contract.
type RewardsRewardsDenied struct {
	Indexer      common.Address
	AllocationID common.Address
	Epoch        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardsDenied is a free log retrieval operation binding the contract event 0x07b6d8dafdf1323cc892f303125d1308c0edc07b15b14b6de6dcb4a5eea818fa.
//
// Solidity: event RewardsDenied(address indexed indexer, address indexed allocationID, uint256 epoch)
func (_Rewards *RewardsFilterer) FilterRewardsDenied(opts *bind.FilterOpts, indexer []common.Address, allocationID []common.Address) (*RewardsRewardsDeniedIterator, error) {
	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardsDenied", indexerRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &RewardsRewardsDeniedIterator{contract: _Rewards.contract, event: "RewardsDenied", logs: logs, sub: sub}, nil
}

// WatchRewardsDenied is a free log subscription operation binding the contract event 0x07b6d8dafdf1323cc892f303125d1308c0edc07b15b14b6de6dcb4a5eea818fa.
//
// Solidity: event RewardsDenied(address indexed indexer, address indexed allocationID, uint256 epoch)
func (_Rewards *RewardsFilterer) WatchRewardsDenied(opts *bind.WatchOpts, sink chan<- *RewardsRewardsDenied, indexer []common.Address, allocationID []common.Address) (event.Subscription, error) {
	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardsDenied", indexerRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardsDenied)
				if err := _Rewards.contract.UnpackLog(event, "RewardsDenied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsDenied is a log parse operation binding the contract event 0x07b6d8dafdf1323cc892f303125d1308c0edc07b15b14b6de6dcb4a5eea818fa.
//
// Solidity: event RewardsDenied(address indexed indexer, address indexed allocationID, uint256 epoch)
func (_Rewards *RewardsFilterer) ParseRewardsDenied(log types.Log) (*RewardsRewardsDenied, error) {
	event := new(RewardsRewardsDenied)
	if err := _Rewards.contract.UnpackLog(event, "RewardsDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardsDenylistUpdatedIterator is returned from FilterRewardsDenylistUpdated and is used to iterate over the raw logs and unpacked data for RewardsDenylistUpdated events raised by the Rewards contract.
type RewardsRewardsDenylistUpdatedIterator struct {
	Event *RewardsRewardsDenylistUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsRewardsDenylistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardsDenylistUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsRewardsDenylistUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsRewardsDenylistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardsDenylistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardsDenylistUpdated represents a RewardsDenylistUpdated event raised by the Rewards contract.
type RewardsRewardsDenylistUpdated struct {
	SubgraphDeploymentID [32]byte
	SinceBlock           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsDenylistUpdated is a free log retrieval operation binding the contract event 0xe016102b339c3889f4967b491f3381f2c352c8fe3d4f880007807d45b124065a.
//
// Solidity: event RewardsDenylistUpdated(bytes32 indexed subgraphDeploymentID, uint256 sinceBlock)
func (_Rewards *RewardsFilterer) FilterRewardsDenylistUpdated(opts *bind.FilterOpts, subgraphDeploymentID [][32]byte) (*RewardsRewardsDenylistUpdatedIterator, error) {
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardsDenylistUpdated", subgraphDeploymentIDRule)
	if err != nil {
		return nil, err
	}
	return &RewardsRewardsDenylistUpdatedIterator{contract: _Rewards.contract, event: "RewardsDenylistUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDenylistUpdated is a free log subscription operation binding the contract event 0xe016102b339c3889f4967b491f3381f2c352c8fe3d4f880007807d45b124065a.
//
// Solidity: event RewardsDenylistUpdated(bytes32 indexed subgraphDeploymentID, uint256 sinceBlock)
func (_Rewards *RewardsFilterer) WatchRewardsDenylistUpdated(opts *bind.WatchOpts, sink chan<- *RewardsRewardsDenylistUpdated, subgraphDeploymentID [][32]byte) (event.Subscription, error) {
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardsDenylistUpdated", subgraphDeploymentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardsDenylistUpdated)
				if err := _Rewards.contract.UnpackLog(event, "RewardsDenylistUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsDenylistUpdated is a log parse operation binding the contract event 0xe016102b339c3889f4967b491f3381f2c352c8fe3d4f880007807d45b124065a.
//
// Solidity: event RewardsDenylistUpdated(bytes32 indexed subgraphDeploymentID, uint256 sinceBlock)
func (_Rewards *RewardsFilterer) ParseRewardsDenylistUpdated(log types.Log) (*RewardsRewardsDenylistUpdated, error) {
	event := new(RewardsRewardsDenylistUpdated)
	if err := _Rewards.contract.UnpackLog(event, "RewardsDenylistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsSetControllerIterator is returned from FilterSetController and is used to iterate over the raw logs and unpacked data for SetController events raised by the Rewards contract.
type RewardsSetControllerIterator struct {
	Event *RewardsSetController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardsSetControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsSetController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardsSetController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardsSetControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsSetControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsSetController represents a SetController event raised by the Rewards contract.
type RewardsSetController struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetController is a free log retrieval operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_Rewards *RewardsFilterer) FilterSetController(opts *bind.FilterOpts) (*RewardsSetControllerIterator, error) {
	logs, sub, err := _Rewards.contract.FilterLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return &RewardsSetControllerIterator{contract: _Rewards.contract, event: "SetController", logs: logs, sub: sub}, nil
}

// WatchSetController is a free log subscription operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_Rewards *RewardsFilterer) WatchSetController(opts *bind.WatchOpts, sink chan<- *RewardsSetController) (event.Subscription, error) {
	logs, sub, err := _Rewards.contract.WatchLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsSetController)
				if err := _Rewards.contract.UnpackLog(event, "SetController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetController is a log parse operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_Rewards *RewardsFilterer) ParseSetController(log types.Log) (*RewardsSetController, error) {
	event := new(RewardsSetController)
	if err := _Rewards.contract.UnpackLog(event, "SetController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
