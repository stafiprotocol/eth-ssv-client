// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package distributor

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

// DistributorMetaData contains all meta data concerning the Distributor contract.
var DistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableDeposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dealedHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformAmount\",\"type\":\"uint256\"}],\"name\":\"DistributeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dealedHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"DistributeSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dealedHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformAmount\",\"type\":\"uint256\"}],\"name\":\"DistributeSuperNodeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dealedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"SetMerkleRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteProposal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalExitDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealedHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformAmount\",\"type\":\"uint256\"}],\"name\":\"distributeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealedHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"distributeSlashAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealedHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformAmount\",\"type\":\"uint256\"}],\"name\":\"distributeSuperNodeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeWithdrawals\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentNodeDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributeFeeDealedHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributeSlashDealedHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributeSuperNodeFeeDealedHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleDealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getTotalClaimedDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getTotalClaimedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveEtherWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"updateMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributorMetaData.ABI instead.
var DistributorABI = DistributorMetaData.ABI

// Distributor is an auto generated Go binding around an Ethereum contract.
type Distributor struct {
	DistributorCaller     // Read-only binding to the contract
	DistributorTransactor // Write-only binding to the contract
	DistributorFilterer   // Log filterer for contract events
}

// DistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributorSession struct {
	Contract     *Distributor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributorCallerSession struct {
	Contract *DistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributorTransactorSession struct {
	Contract     *DistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributorRaw struct {
	Contract *Distributor // Generic contract binding to access the raw methods on
}

// DistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributorCallerRaw struct {
	Contract *DistributorCaller // Generic read-only contract binding to access the raw methods on
}

// DistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributorTransactorRaw struct {
	Contract *DistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributor creates a new instance of Distributor, bound to a specific deployed contract.
func NewDistributor(address common.Address, backend bind.ContractBackend) (*Distributor, error) {
	contract, err := bindDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distributor{DistributorCaller: DistributorCaller{contract: contract}, DistributorTransactor: DistributorTransactor{contract: contract}, DistributorFilterer: DistributorFilterer{contract: contract}}, nil
}

// NewDistributorCaller creates a new read-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorCaller(address common.Address, caller bind.ContractCaller) (*DistributorCaller, error) {
	contract, err := bindDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorCaller{contract: contract}, nil
}

// NewDistributorTransactor creates a new write-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributorTransactor, error) {
	contract, err := bindDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorTransactor{contract: contract}, nil
}

// NewDistributorFilterer creates a new log filterer instance of Distributor, bound to a specific deployed contract.
func NewDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributorFilterer, error) {
	contract, err := bindDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributorFilterer{contract: contract}, nil
}

// bindDistributor binds a generic wrapper to an already deployed contract.
func bindDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.DistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_Distributor *DistributorCaller) GetCurrentNodeDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getCurrentNodeDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_Distributor *DistributorSession) GetCurrentNodeDepositAmount() (*big.Int, error) {
	return _Distributor.Contract.GetCurrentNodeDepositAmount(&_Distributor.CallOpts)
}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_Distributor *DistributorCallerSession) GetCurrentNodeDepositAmount() (*big.Int, error) {
	return _Distributor.Contract.GetCurrentNodeDepositAmount(&_Distributor.CallOpts)
}

// GetDistributeFeeDealedHeight is a free data retrieval call binding the contract method 0xf1e26b0d.
//
// Solidity: function getDistributeFeeDealedHeight() view returns(uint256)
func (_Distributor *DistributorCaller) GetDistributeFeeDealedHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getDistributeFeeDealedHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributeFeeDealedHeight is a free data retrieval call binding the contract method 0xf1e26b0d.
//
// Solidity: function getDistributeFeeDealedHeight() view returns(uint256)
func (_Distributor *DistributorSession) GetDistributeFeeDealedHeight() (*big.Int, error) {
	return _Distributor.Contract.GetDistributeFeeDealedHeight(&_Distributor.CallOpts)
}

// GetDistributeFeeDealedHeight is a free data retrieval call binding the contract method 0xf1e26b0d.
//
// Solidity: function getDistributeFeeDealedHeight() view returns(uint256)
func (_Distributor *DistributorCallerSession) GetDistributeFeeDealedHeight() (*big.Int, error) {
	return _Distributor.Contract.GetDistributeFeeDealedHeight(&_Distributor.CallOpts)
}

// GetDistributeSlashDealedHeight is a free data retrieval call binding the contract method 0xb56fab8b.
//
// Solidity: function getDistributeSlashDealedHeight() view returns(uint256)
func (_Distributor *DistributorCaller) GetDistributeSlashDealedHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getDistributeSlashDealedHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributeSlashDealedHeight is a free data retrieval call binding the contract method 0xb56fab8b.
//
// Solidity: function getDistributeSlashDealedHeight() view returns(uint256)
func (_Distributor *DistributorSession) GetDistributeSlashDealedHeight() (*big.Int, error) {
	return _Distributor.Contract.GetDistributeSlashDealedHeight(&_Distributor.CallOpts)
}

// GetDistributeSlashDealedHeight is a free data retrieval call binding the contract method 0xb56fab8b.
//
// Solidity: function getDistributeSlashDealedHeight() view returns(uint256)
func (_Distributor *DistributorCallerSession) GetDistributeSlashDealedHeight() (*big.Int, error) {
	return _Distributor.Contract.GetDistributeSlashDealedHeight(&_Distributor.CallOpts)
}

// GetDistributeSuperNodeFeeDealedHeight is a free data retrieval call binding the contract method 0xad4a473a.
//
// Solidity: function getDistributeSuperNodeFeeDealedHeight() view returns(uint256)
func (_Distributor *DistributorCaller) GetDistributeSuperNodeFeeDealedHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getDistributeSuperNodeFeeDealedHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributeSuperNodeFeeDealedHeight is a free data retrieval call binding the contract method 0xad4a473a.
//
// Solidity: function getDistributeSuperNodeFeeDealedHeight() view returns(uint256)
func (_Distributor *DistributorSession) GetDistributeSuperNodeFeeDealedHeight() (*big.Int, error) {
	return _Distributor.Contract.GetDistributeSuperNodeFeeDealedHeight(&_Distributor.CallOpts)
}

// GetDistributeSuperNodeFeeDealedHeight is a free data retrieval call binding the contract method 0xad4a473a.
//
// Solidity: function getDistributeSuperNodeFeeDealedHeight() view returns(uint256)
func (_Distributor *DistributorCallerSession) GetDistributeSuperNodeFeeDealedHeight() (*big.Int, error) {
	return _Distributor.Contract.GetDistributeSuperNodeFeeDealedHeight(&_Distributor.CallOpts)
}

// GetMerkleDealedEpoch is a free data retrieval call binding the contract method 0x7cd1fc9d.
//
// Solidity: function getMerkleDealedEpoch() view returns(uint256)
func (_Distributor *DistributorCaller) GetMerkleDealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getMerkleDealedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleDealedEpoch is a free data retrieval call binding the contract method 0x7cd1fc9d.
//
// Solidity: function getMerkleDealedEpoch() view returns(uint256)
func (_Distributor *DistributorSession) GetMerkleDealedEpoch() (*big.Int, error) {
	return _Distributor.Contract.GetMerkleDealedEpoch(&_Distributor.CallOpts)
}

// GetMerkleDealedEpoch is a free data retrieval call binding the contract method 0x7cd1fc9d.
//
// Solidity: function getMerkleDealedEpoch() view returns(uint256)
func (_Distributor *DistributorCallerSession) GetMerkleDealedEpoch() (*big.Int, error) {
	return _Distributor.Contract.GetMerkleDealedEpoch(&_Distributor.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_Distributor *DistributorCaller) GetMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_Distributor *DistributorSession) GetMerkleRoot() ([32]byte, error) {
	return _Distributor.Contract.GetMerkleRoot(&_Distributor.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_Distributor *DistributorCallerSession) GetMerkleRoot() ([32]byte, error) {
	return _Distributor.Contract.GetMerkleRoot(&_Distributor.CallOpts)
}

// GetTotalClaimedDeposit is a free data retrieval call binding the contract method 0x1a94fbdf.
//
// Solidity: function getTotalClaimedDeposit(address _account) view returns(uint256)
func (_Distributor *DistributorCaller) GetTotalClaimedDeposit(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getTotalClaimedDeposit", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalClaimedDeposit is a free data retrieval call binding the contract method 0x1a94fbdf.
//
// Solidity: function getTotalClaimedDeposit(address _account) view returns(uint256)
func (_Distributor *DistributorSession) GetTotalClaimedDeposit(_account common.Address) (*big.Int, error) {
	return _Distributor.Contract.GetTotalClaimedDeposit(&_Distributor.CallOpts, _account)
}

// GetTotalClaimedDeposit is a free data retrieval call binding the contract method 0x1a94fbdf.
//
// Solidity: function getTotalClaimedDeposit(address _account) view returns(uint256)
func (_Distributor *DistributorCallerSession) GetTotalClaimedDeposit(_account common.Address) (*big.Int, error) {
	return _Distributor.Contract.GetTotalClaimedDeposit(&_Distributor.CallOpts, _account)
}

// GetTotalClaimedReward is a free data retrieval call binding the contract method 0x1f6c9fc0.
//
// Solidity: function getTotalClaimedReward(address _account) view returns(uint256)
func (_Distributor *DistributorCaller) GetTotalClaimedReward(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "getTotalClaimedReward", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalClaimedReward is a free data retrieval call binding the contract method 0x1f6c9fc0.
//
// Solidity: function getTotalClaimedReward(address _account) view returns(uint256)
func (_Distributor *DistributorSession) GetTotalClaimedReward(_account common.Address) (*big.Int, error) {
	return _Distributor.Contract.GetTotalClaimedReward(&_Distributor.CallOpts, _account)
}

// GetTotalClaimedReward is a free data retrieval call binding the contract method 0x1f6c9fc0.
//
// Solidity: function getTotalClaimedReward(address _account) view returns(uint256)
func (_Distributor *DistributorCallerSession) GetTotalClaimedReward(_account common.Address) (*big.Int, error) {
	return _Distributor.Contract.GetTotalClaimedReward(&_Distributor.CallOpts, _account)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Distributor *DistributorCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Distributor *DistributorSession) Version() (uint8, error) {
	return _Distributor.Contract.Version(&_Distributor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Distributor *DistributorCallerSession) Version() (uint8, error) {
	return _Distributor.Contract.Version(&_Distributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x92333d90.
//
// Solidity: function claim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_Distributor *DistributorTransactor) Claim(opts *bind.TransactOpts, _index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "claim", _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// Claim is a paid mutator transaction binding the contract method 0x92333d90.
//
// Solidity: function claim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_Distributor *DistributorSession) Claim(_index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _Distributor.Contract.Claim(&_Distributor.TransactOpts, _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// Claim is a paid mutator transaction binding the contract method 0x92333d90.
//
// Solidity: function claim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_Distributor *DistributorTransactorSession) Claim(_index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _Distributor.Contract.Claim(&_Distributor.TransactOpts, _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// DistributeFee is a paid mutator transaction binding the contract method 0x897a980c.
//
// Solidity: function distributeFee(uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount) returns()
func (_Distributor *DistributorTransactor) DistributeFee(opts *bind.TransactOpts, _dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeFee", _dealedHeight, _userAmount, _nodeAmount, _platformAmount)
}

// DistributeFee is a paid mutator transaction binding the contract method 0x897a980c.
//
// Solidity: function distributeFee(uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount) returns()
func (_Distributor *DistributorSession) DistributeFee(_dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeFee(&_Distributor.TransactOpts, _dealedHeight, _userAmount, _nodeAmount, _platformAmount)
}

// DistributeFee is a paid mutator transaction binding the contract method 0x897a980c.
//
// Solidity: function distributeFee(uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount) returns()
func (_Distributor *DistributorTransactorSession) DistributeFee(_dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeFee(&_Distributor.TransactOpts, _dealedHeight, _userAmount, _nodeAmount, _platformAmount)
}

// DistributeSlashAmount is a paid mutator transaction binding the contract method 0x535c47a2.
//
// Solidity: function distributeSlashAmount(uint256 _dealedHeight, uint256 _amount) returns()
func (_Distributor *DistributorTransactor) DistributeSlashAmount(opts *bind.TransactOpts, _dealedHeight *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeSlashAmount", _dealedHeight, _amount)
}

// DistributeSlashAmount is a paid mutator transaction binding the contract method 0x535c47a2.
//
// Solidity: function distributeSlashAmount(uint256 _dealedHeight, uint256 _amount) returns()
func (_Distributor *DistributorSession) DistributeSlashAmount(_dealedHeight *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeSlashAmount(&_Distributor.TransactOpts, _dealedHeight, _amount)
}

// DistributeSlashAmount is a paid mutator transaction binding the contract method 0x535c47a2.
//
// Solidity: function distributeSlashAmount(uint256 _dealedHeight, uint256 _amount) returns()
func (_Distributor *DistributorTransactorSession) DistributeSlashAmount(_dealedHeight *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeSlashAmount(&_Distributor.TransactOpts, _dealedHeight, _amount)
}

// DistributeSuperNodeFee is a paid mutator transaction binding the contract method 0x39eddf16.
//
// Solidity: function distributeSuperNodeFee(uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount) returns()
func (_Distributor *DistributorTransactor) DistributeSuperNodeFee(opts *bind.TransactOpts, _dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeSuperNodeFee", _dealedHeight, _userAmount, _nodeAmount, _platformAmount)
}

// DistributeSuperNodeFee is a paid mutator transaction binding the contract method 0x39eddf16.
//
// Solidity: function distributeSuperNodeFee(uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount) returns()
func (_Distributor *DistributorSession) DistributeSuperNodeFee(_dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeSuperNodeFee(&_Distributor.TransactOpts, _dealedHeight, _userAmount, _nodeAmount, _platformAmount)
}

// DistributeSuperNodeFee is a paid mutator transaction binding the contract method 0x39eddf16.
//
// Solidity: function distributeSuperNodeFee(uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount) returns()
func (_Distributor *DistributorTransactorSession) DistributeSuperNodeFee(_dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeSuperNodeFee(&_Distributor.TransactOpts, _dealedHeight, _userAmount, _nodeAmount, _platformAmount)
}

// DistributeWithdrawals is a paid mutator transaction binding the contract method 0x67e2c718.
//
// Solidity: function distributeWithdrawals() payable returns()
func (_Distributor *DistributorTransactor) DistributeWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeWithdrawals")
}

// DistributeWithdrawals is a paid mutator transaction binding the contract method 0x67e2c718.
//
// Solidity: function distributeWithdrawals() payable returns()
func (_Distributor *DistributorSession) DistributeWithdrawals() (*types.Transaction, error) {
	return _Distributor.Contract.DistributeWithdrawals(&_Distributor.TransactOpts)
}

// DistributeWithdrawals is a paid mutator transaction binding the contract method 0x67e2c718.
//
// Solidity: function distributeWithdrawals() payable returns()
func (_Distributor *DistributorTransactorSession) DistributeWithdrawals() (*types.Transaction, error) {
	return _Distributor.Contract.DistributeWithdrawals(&_Distributor.TransactOpts)
}

// ReceiveEtherWithdrawal is a paid mutator transaction binding the contract method 0x0a019eaf.
//
// Solidity: function receiveEtherWithdrawal() payable returns()
func (_Distributor *DistributorTransactor) ReceiveEtherWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "receiveEtherWithdrawal")
}

// ReceiveEtherWithdrawal is a paid mutator transaction binding the contract method 0x0a019eaf.
//
// Solidity: function receiveEtherWithdrawal() payable returns()
func (_Distributor *DistributorSession) ReceiveEtherWithdrawal() (*types.Transaction, error) {
	return _Distributor.Contract.ReceiveEtherWithdrawal(&_Distributor.TransactOpts)
}

// ReceiveEtherWithdrawal is a paid mutator transaction binding the contract method 0x0a019eaf.
//
// Solidity: function receiveEtherWithdrawal() payable returns()
func (_Distributor *DistributorTransactorSession) ReceiveEtherWithdrawal() (*types.Transaction, error) {
	return _Distributor.Contract.ReceiveEtherWithdrawal(&_Distributor.TransactOpts)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x18712c21.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot) returns()
func (_Distributor *DistributorTransactor) SetMerkleRoot(opts *bind.TransactOpts, _dealedEpoch *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "setMerkleRoot", _dealedEpoch, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x18712c21.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot) returns()
func (_Distributor *DistributorSession) SetMerkleRoot(_dealedEpoch *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Distributor.Contract.SetMerkleRoot(&_Distributor.TransactOpts, _dealedEpoch, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x18712c21.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot) returns()
func (_Distributor *DistributorTransactorSession) SetMerkleRoot(_dealedEpoch *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Distributor.Contract.SetMerkleRoot(&_Distributor.TransactOpts, _dealedEpoch, _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _merkleRoot) returns()
func (_Distributor *DistributorTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "updateMerkleRoot", _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _merkleRoot) returns()
func (_Distributor *DistributorSession) UpdateMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _Distributor.Contract.UpdateMerkleRoot(&_Distributor.TransactOpts, _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _merkleRoot) returns()
func (_Distributor *DistributorTransactorSession) UpdateMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _Distributor.Contract.UpdateMerkleRoot(&_Distributor.TransactOpts, _merkleRoot)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Distributor *DistributorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Distributor *DistributorSession) Receive() (*types.Transaction, error) {
	return _Distributor.Contract.Receive(&_Distributor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Distributor *DistributorTransactorSession) Receive() (*types.Transaction, error) {
	return _Distributor.Contract.Receive(&_Distributor.TransactOpts)
}

// DistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Distributor contract.
type DistributorClaimedIterator struct {
	Event *DistributorClaimed // Event containing the contract specifics and raw log

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
func (it *DistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorClaimed)
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
		it.Event = new(DistributorClaimed)
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
func (it *DistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorClaimed represents a Claimed event raised by the Distributor contract.
type DistributorClaimed struct {
	Index            *big.Int
	Account          common.Address
	ClaimableReward  *big.Int
	ClaimableDeposit *big.Int
	ClaimType        uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x805593669516ad95e9aa092bd501707436d9563eac1ef7c017cd6639eb29f8ee.
//
// Solidity: event Claimed(uint256 index, address account, uint256 claimableReward, uint256 claimableDeposit, uint8 claimType)
func (_Distributor *DistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*DistributorClaimedIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &DistributorClaimedIterator{contract: _Distributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x805593669516ad95e9aa092bd501707436d9563eac1ef7c017cd6639eb29f8ee.
//
// Solidity: event Claimed(uint256 index, address account, uint256 claimableReward, uint256 claimableDeposit, uint8 claimType)
func (_Distributor *DistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *DistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorClaimed)
				if err := _Distributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x805593669516ad95e9aa092bd501707436d9563eac1ef7c017cd6639eb29f8ee.
//
// Solidity: event Claimed(uint256 index, address account, uint256 claimableReward, uint256 claimableDeposit, uint8 claimType)
func (_Distributor *DistributorFilterer) ParseClaimed(log types.Log) (*DistributorClaimed, error) {
	event := new(DistributorClaimed)
	if err := _Distributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorDistributeFeeIterator is returned from FilterDistributeFee and is used to iterate over the raw logs and unpacked data for DistributeFee events raised by the Distributor contract.
type DistributorDistributeFeeIterator struct {
	Event *DistributorDistributeFee // Event containing the contract specifics and raw log

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
func (it *DistributorDistributeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorDistributeFee)
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
		it.Event = new(DistributorDistributeFee)
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
func (it *DistributorDistributeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorDistributeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorDistributeFee represents a DistributeFee event raised by the Distributor contract.
type DistributorDistributeFee struct {
	DealedHeight   *big.Int
	UserAmount     *big.Int
	NodeAmount     *big.Int
	PlatformAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributeFee is a free log retrieval operation binding the contract event 0x21c58bec2fee2b8fa31ce6802a99242adf6226b7ca63d69966c2036047374382.
//
// Solidity: event DistributeFee(uint256 dealedHeight, uint256 userAmount, uint256 nodeAmount, uint256 platformAmount)
func (_Distributor *DistributorFilterer) FilterDistributeFee(opts *bind.FilterOpts) (*DistributorDistributeFeeIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "DistributeFee")
	if err != nil {
		return nil, err
	}
	return &DistributorDistributeFeeIterator{contract: _Distributor.contract, event: "DistributeFee", logs: logs, sub: sub}, nil
}

// WatchDistributeFee is a free log subscription operation binding the contract event 0x21c58bec2fee2b8fa31ce6802a99242adf6226b7ca63d69966c2036047374382.
//
// Solidity: event DistributeFee(uint256 dealedHeight, uint256 userAmount, uint256 nodeAmount, uint256 platformAmount)
func (_Distributor *DistributorFilterer) WatchDistributeFee(opts *bind.WatchOpts, sink chan<- *DistributorDistributeFee) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "DistributeFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorDistributeFee)
				if err := _Distributor.contract.UnpackLog(event, "DistributeFee", log); err != nil {
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

// ParseDistributeFee is a log parse operation binding the contract event 0x21c58bec2fee2b8fa31ce6802a99242adf6226b7ca63d69966c2036047374382.
//
// Solidity: event DistributeFee(uint256 dealedHeight, uint256 userAmount, uint256 nodeAmount, uint256 platformAmount)
func (_Distributor *DistributorFilterer) ParseDistributeFee(log types.Log) (*DistributorDistributeFee, error) {
	event := new(DistributorDistributeFee)
	if err := _Distributor.contract.UnpackLog(event, "DistributeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorDistributeSlashIterator is returned from FilterDistributeSlash and is used to iterate over the raw logs and unpacked data for DistributeSlash events raised by the Distributor contract.
type DistributorDistributeSlashIterator struct {
	Event *DistributorDistributeSlash // Event containing the contract specifics and raw log

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
func (it *DistributorDistributeSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorDistributeSlash)
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
		it.Event = new(DistributorDistributeSlash)
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
func (it *DistributorDistributeSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorDistributeSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorDistributeSlash represents a DistributeSlash event raised by the Distributor contract.
type DistributorDistributeSlash struct {
	DealedHeight *big.Int
	SlashAmount  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistributeSlash is a free log retrieval operation binding the contract event 0x92cfd04ead0f32488c7f490c6afce1b145630dfcc9c3ed06b2ab0d0696a4bcef.
//
// Solidity: event DistributeSlash(uint256 dealedHeight, uint256 slashAmount)
func (_Distributor *DistributorFilterer) FilterDistributeSlash(opts *bind.FilterOpts) (*DistributorDistributeSlashIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "DistributeSlash")
	if err != nil {
		return nil, err
	}
	return &DistributorDistributeSlashIterator{contract: _Distributor.contract, event: "DistributeSlash", logs: logs, sub: sub}, nil
}

// WatchDistributeSlash is a free log subscription operation binding the contract event 0x92cfd04ead0f32488c7f490c6afce1b145630dfcc9c3ed06b2ab0d0696a4bcef.
//
// Solidity: event DistributeSlash(uint256 dealedHeight, uint256 slashAmount)
func (_Distributor *DistributorFilterer) WatchDistributeSlash(opts *bind.WatchOpts, sink chan<- *DistributorDistributeSlash) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "DistributeSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorDistributeSlash)
				if err := _Distributor.contract.UnpackLog(event, "DistributeSlash", log); err != nil {
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

// ParseDistributeSlash is a log parse operation binding the contract event 0x92cfd04ead0f32488c7f490c6afce1b145630dfcc9c3ed06b2ab0d0696a4bcef.
//
// Solidity: event DistributeSlash(uint256 dealedHeight, uint256 slashAmount)
func (_Distributor *DistributorFilterer) ParseDistributeSlash(log types.Log) (*DistributorDistributeSlash, error) {
	event := new(DistributorDistributeSlash)
	if err := _Distributor.contract.UnpackLog(event, "DistributeSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorDistributeSuperNodeFeeIterator is returned from FilterDistributeSuperNodeFee and is used to iterate over the raw logs and unpacked data for DistributeSuperNodeFee events raised by the Distributor contract.
type DistributorDistributeSuperNodeFeeIterator struct {
	Event *DistributorDistributeSuperNodeFee // Event containing the contract specifics and raw log

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
func (it *DistributorDistributeSuperNodeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorDistributeSuperNodeFee)
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
		it.Event = new(DistributorDistributeSuperNodeFee)
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
func (it *DistributorDistributeSuperNodeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorDistributeSuperNodeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorDistributeSuperNodeFee represents a DistributeSuperNodeFee event raised by the Distributor contract.
type DistributorDistributeSuperNodeFee struct {
	DealedHeight   *big.Int
	UserAmount     *big.Int
	NodeAmount     *big.Int
	PlatformAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributeSuperNodeFee is a free log retrieval operation binding the contract event 0x4b6a75e8741caebe2695484108d8f5df71d33d430e453b66f68aa802f172edf7.
//
// Solidity: event DistributeSuperNodeFee(uint256 dealedHeight, uint256 userAmount, uint256 nodeAmount, uint256 platformAmount)
func (_Distributor *DistributorFilterer) FilterDistributeSuperNodeFee(opts *bind.FilterOpts) (*DistributorDistributeSuperNodeFeeIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "DistributeSuperNodeFee")
	if err != nil {
		return nil, err
	}
	return &DistributorDistributeSuperNodeFeeIterator{contract: _Distributor.contract, event: "DistributeSuperNodeFee", logs: logs, sub: sub}, nil
}

// WatchDistributeSuperNodeFee is a free log subscription operation binding the contract event 0x4b6a75e8741caebe2695484108d8f5df71d33d430e453b66f68aa802f172edf7.
//
// Solidity: event DistributeSuperNodeFee(uint256 dealedHeight, uint256 userAmount, uint256 nodeAmount, uint256 platformAmount)
func (_Distributor *DistributorFilterer) WatchDistributeSuperNodeFee(opts *bind.WatchOpts, sink chan<- *DistributorDistributeSuperNodeFee) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "DistributeSuperNodeFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorDistributeSuperNodeFee)
				if err := _Distributor.contract.UnpackLog(event, "DistributeSuperNodeFee", log); err != nil {
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

// ParseDistributeSuperNodeFee is a log parse operation binding the contract event 0x4b6a75e8741caebe2695484108d8f5df71d33d430e453b66f68aa802f172edf7.
//
// Solidity: event DistributeSuperNodeFee(uint256 dealedHeight, uint256 userAmount, uint256 nodeAmount, uint256 platformAmount)
func (_Distributor *DistributorFilterer) ParseDistributeSuperNodeFee(log types.Log) (*DistributorDistributeSuperNodeFee, error) {
	event := new(DistributorDistributeSuperNodeFee)
	if err := _Distributor.contract.UnpackLog(event, "DistributeSuperNodeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Distributor contract.
type DistributorProposalExecutedIterator struct {
	Event *DistributorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *DistributorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorProposalExecuted)
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
		it.Event = new(DistributorProposalExecuted)
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
func (it *DistributorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorProposalExecuted represents a ProposalExecuted event raised by the Distributor contract.
type DistributorProposalExecuted struct {
	ProposalId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_Distributor *DistributorFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId [][32]byte) (*DistributorProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &DistributorProposalExecutedIterator{contract: _Distributor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_Distributor *DistributorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *DistributorProposalExecuted, proposalId [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorProposalExecuted)
				if err := _Distributor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_Distributor *DistributorFilterer) ParseProposalExecuted(log types.Log) (*DistributorProposalExecuted, error) {
	event := new(DistributorProposalExecuted)
	if err := _Distributor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorSetMerkleRootIterator is returned from FilterSetMerkleRoot and is used to iterate over the raw logs and unpacked data for SetMerkleRoot events raised by the Distributor contract.
type DistributorSetMerkleRootIterator struct {
	Event *DistributorSetMerkleRoot // Event containing the contract specifics and raw log

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
func (it *DistributorSetMerkleRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorSetMerkleRoot)
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
		it.Event = new(DistributorSetMerkleRoot)
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
func (it *DistributorSetMerkleRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorSetMerkleRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorSetMerkleRoot represents a SetMerkleRoot event raised by the Distributor contract.
type DistributorSetMerkleRoot struct {
	DealedEpoch *big.Int
	MerkleRoot  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMerkleRoot is a free log retrieval operation binding the contract event 0x06300242dd77c603f351a5b575426a902d850335fbeb7c1aa2aaa83e64df395b.
//
// Solidity: event SetMerkleRoot(uint256 dealedEpoch, bytes32 merkleRoot)
func (_Distributor *DistributorFilterer) FilterSetMerkleRoot(opts *bind.FilterOpts) (*DistributorSetMerkleRootIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "SetMerkleRoot")
	if err != nil {
		return nil, err
	}
	return &DistributorSetMerkleRootIterator{contract: _Distributor.contract, event: "SetMerkleRoot", logs: logs, sub: sub}, nil
}

// WatchSetMerkleRoot is a free log subscription operation binding the contract event 0x06300242dd77c603f351a5b575426a902d850335fbeb7c1aa2aaa83e64df395b.
//
// Solidity: event SetMerkleRoot(uint256 dealedEpoch, bytes32 merkleRoot)
func (_Distributor *DistributorFilterer) WatchSetMerkleRoot(opts *bind.WatchOpts, sink chan<- *DistributorSetMerkleRoot) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "SetMerkleRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorSetMerkleRoot)
				if err := _Distributor.contract.UnpackLog(event, "SetMerkleRoot", log); err != nil {
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

// ParseSetMerkleRoot is a log parse operation binding the contract event 0x06300242dd77c603f351a5b575426a902d850335fbeb7c1aa2aaa83e64df395b.
//
// Solidity: event SetMerkleRoot(uint256 dealedEpoch, bytes32 merkleRoot)
func (_Distributor *DistributorFilterer) ParseSetMerkleRoot(log types.Log) (*DistributorSetMerkleRoot, error) {
	event := new(DistributorSetMerkleRoot)
	if err := _Distributor.contract.UnpackLog(event, "SetMerkleRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorVoteProposalIterator is returned from FilterVoteProposal and is used to iterate over the raw logs and unpacked data for VoteProposal events raised by the Distributor contract.
type DistributorVoteProposalIterator struct {
	Event *DistributorVoteProposal // Event containing the contract specifics and raw log

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
func (it *DistributorVoteProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorVoteProposal)
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
		it.Event = new(DistributorVoteProposal)
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
func (it *DistributorVoteProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorVoteProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorVoteProposal represents a VoteProposal event raised by the Distributor contract.
type DistributorVoteProposal struct {
	ProposalId [32]byte
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteProposal is a free log retrieval operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed proposalId, address voter)
func (_Distributor *DistributorFilterer) FilterVoteProposal(opts *bind.FilterOpts, proposalId [][32]byte) (*DistributorVoteProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "VoteProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &DistributorVoteProposalIterator{contract: _Distributor.contract, event: "VoteProposal", logs: logs, sub: sub}, nil
}

// WatchVoteProposal is a free log subscription operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed proposalId, address voter)
func (_Distributor *DistributorFilterer) WatchVoteProposal(opts *bind.WatchOpts, sink chan<- *DistributorVoteProposal, proposalId [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "VoteProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorVoteProposal)
				if err := _Distributor.contract.UnpackLog(event, "VoteProposal", log); err != nil {
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

// ParseVoteProposal is a log parse operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed proposalId, address voter)
func (_Distributor *DistributorFilterer) ParseVoteProposal(log types.Log) (*DistributorVoteProposal, error) {
	event := new(DistributorVoteProposal)
	if err := _Distributor.contract.UnpackLog(event, "VoteProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
