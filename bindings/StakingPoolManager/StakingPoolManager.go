// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_pool_manager

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

// StakingPoolMangerMetaData contains all meta data concerning the StakingPoolManger contract.
var StakingPoolMangerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingPool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StakingPoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingPool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StakingPoolDestroyed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AddressSetStorage\",\"outputs\":[{\"internalType\":\"contractIAddressSetStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"enumDepositType\",\"name\":\"_depositType\",\"type\":\"uint8\"}],\"name\":\"createStakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyStakingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeStakingPoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeStakingPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeValidatingStakingPoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeValidatingStakingPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getPrelaunchStakingpools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getStakingPoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getStakingPoolByPubkey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPoolAddress\",\"type\":\"address\"}],\"name\":\"getStakingPoolExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPoolAddress\",\"type\":\"address\"}],\"name\":\"getStakingPoolPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPoolAddress\",\"type\":\"address\"}],\"name\":\"getStakingPoolWithdrawalProcessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"setStakingPoolPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingPoolAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_processed\",\"type\":\"bool\"}],\"name\":\"setStakingPoolWithdrawalProcessed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingPoolMangerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolMangerMetaData.ABI instead.
var StakingPoolMangerABI = StakingPoolMangerMetaData.ABI

// StakingPoolManger is an auto generated Go binding around an Ethereum contract.
type StakingPoolManger struct {
	StakingPoolMangerCaller     // Read-only binding to the contract
	StakingPoolMangerTransactor // Write-only binding to the contract
	StakingPoolMangerFilterer   // Log filterer for contract events
}

// StakingPoolMangerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolMangerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolMangerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolMangerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolMangerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolMangerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolMangerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolMangerSession struct {
	Contract     *StakingPoolManger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingPoolMangerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolMangerCallerSession struct {
	Contract *StakingPoolMangerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StakingPoolMangerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolMangerTransactorSession struct {
	Contract     *StakingPoolMangerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StakingPoolMangerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolMangerRaw struct {
	Contract *StakingPoolManger // Generic contract binding to access the raw methods on
}

// StakingPoolMangerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolMangerCallerRaw struct {
	Contract *StakingPoolMangerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolMangerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolMangerTransactorRaw struct {
	Contract *StakingPoolMangerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolManger creates a new instance of StakingPoolManger, bound to a specific deployed contract.
func NewStakingPoolManger(address common.Address, backend bind.ContractBackend) (*StakingPoolManger, error) {
	contract, err := bindStakingPoolManger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolManger{StakingPoolMangerCaller: StakingPoolMangerCaller{contract: contract}, StakingPoolMangerTransactor: StakingPoolMangerTransactor{contract: contract}, StakingPoolMangerFilterer: StakingPoolMangerFilterer{contract: contract}}, nil
}

// NewStakingPoolMangerCaller creates a new read-only instance of StakingPoolManger, bound to a specific deployed contract.
func NewStakingPoolMangerCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolMangerCaller, error) {
	contract, err := bindStakingPoolManger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolMangerCaller{contract: contract}, nil
}

// NewStakingPoolMangerTransactor creates a new write-only instance of StakingPoolManger, bound to a specific deployed contract.
func NewStakingPoolMangerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolMangerTransactor, error) {
	contract, err := bindStakingPoolManger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolMangerTransactor{contract: contract}, nil
}

// NewStakingPoolMangerFilterer creates a new log filterer instance of StakingPoolManger, bound to a specific deployed contract.
func NewStakingPoolMangerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolMangerFilterer, error) {
	contract, err := bindStakingPoolManger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolMangerFilterer{contract: contract}, nil
}

// bindStakingPoolManger binds a generic wrapper to an already deployed contract.
func bindStakingPoolManger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolMangerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolManger *StakingPoolMangerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolManger.Contract.StakingPoolMangerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolManger *StakingPoolMangerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.StakingPoolMangerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolManger *StakingPoolMangerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.StakingPoolMangerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolManger *StakingPoolMangerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolManger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolManger *StakingPoolMangerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolManger *StakingPoolMangerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.contract.Transact(opts, method, params...)
}

// AddressSetStorage is a free data retrieval call binding the contract method 0x528b04b5.
//
// Solidity: function AddressSetStorage() view returns(address)
func (_StakingPoolManger *StakingPoolMangerCaller) AddressSetStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "AddressSetStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressSetStorage is a free data retrieval call binding the contract method 0x528b04b5.
//
// Solidity: function AddressSetStorage() view returns(address)
func (_StakingPoolManger *StakingPoolMangerSession) AddressSetStorage() (common.Address, error) {
	return _StakingPoolManger.Contract.AddressSetStorage(&_StakingPoolManger.CallOpts)
}

// AddressSetStorage is a free data retrieval call binding the contract method 0x528b04b5.
//
// Solidity: function AddressSetStorage() view returns(address)
func (_StakingPoolManger *StakingPoolMangerCallerSession) AddressSetStorage() (common.Address, error) {
	return _StakingPoolManger.Contract.AddressSetStorage(&_StakingPoolManger.CallOpts)
}

// GetNodeStakingPoolAt is a free data retrieval call binding the contract method 0xc0f27c38.
//
// Solidity: function getNodeStakingPoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCaller) GetNodeStakingPoolAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getNodeStakingPoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeStakingPoolAt is a free data retrieval call binding the contract method 0xc0f27c38.
//
// Solidity: function getNodeStakingPoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerSession) GetNodeStakingPoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _StakingPoolManger.Contract.GetNodeStakingPoolAt(&_StakingPoolManger.CallOpts, _nodeAddress, _index)
}

// GetNodeStakingPoolAt is a free data retrieval call binding the contract method 0xc0f27c38.
//
// Solidity: function getNodeStakingPoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetNodeStakingPoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _StakingPoolManger.Contract.GetNodeStakingPoolAt(&_StakingPoolManger.CallOpts, _nodeAddress, _index)
}

// GetNodeStakingPoolCount is a free data retrieval call binding the contract method 0xc3d8073b.
//
// Solidity: function getNodeStakingPoolCount(address _nodeAddress) view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerCaller) GetNodeStakingPoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getNodeStakingPoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeStakingPoolCount is a free data retrieval call binding the contract method 0xc3d8073b.
//
// Solidity: function getNodeStakingPoolCount(address _nodeAddress) view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerSession) GetNodeStakingPoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _StakingPoolManger.Contract.GetNodeStakingPoolCount(&_StakingPoolManger.CallOpts, _nodeAddress)
}

// GetNodeStakingPoolCount is a free data retrieval call binding the contract method 0xc3d8073b.
//
// Solidity: function getNodeStakingPoolCount(address _nodeAddress) view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetNodeStakingPoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _StakingPoolManger.Contract.GetNodeStakingPoolCount(&_StakingPoolManger.CallOpts, _nodeAddress)
}

// GetNodeValidatingStakingPoolAt is a free data retrieval call binding the contract method 0x0fed8b63.
//
// Solidity: function getNodeValidatingStakingPoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCaller) GetNodeValidatingStakingPoolAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getNodeValidatingStakingPoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeValidatingStakingPoolAt is a free data retrieval call binding the contract method 0x0fed8b63.
//
// Solidity: function getNodeValidatingStakingPoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerSession) GetNodeValidatingStakingPoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _StakingPoolManger.Contract.GetNodeValidatingStakingPoolAt(&_StakingPoolManger.CallOpts, _nodeAddress, _index)
}

// GetNodeValidatingStakingPoolAt is a free data retrieval call binding the contract method 0x0fed8b63.
//
// Solidity: function getNodeValidatingStakingPoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetNodeValidatingStakingPoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _StakingPoolManger.Contract.GetNodeValidatingStakingPoolAt(&_StakingPoolManger.CallOpts, _nodeAddress, _index)
}

// GetNodeValidatingStakingPoolCount is a free data retrieval call binding the contract method 0x7be535ad.
//
// Solidity: function getNodeValidatingStakingPoolCount(address _nodeAddress) view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerCaller) GetNodeValidatingStakingPoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getNodeValidatingStakingPoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeValidatingStakingPoolCount is a free data retrieval call binding the contract method 0x7be535ad.
//
// Solidity: function getNodeValidatingStakingPoolCount(address _nodeAddress) view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerSession) GetNodeValidatingStakingPoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _StakingPoolManger.Contract.GetNodeValidatingStakingPoolCount(&_StakingPoolManger.CallOpts, _nodeAddress)
}

// GetNodeValidatingStakingPoolCount is a free data retrieval call binding the contract method 0x7be535ad.
//
// Solidity: function getNodeValidatingStakingPoolCount(address _nodeAddress) view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetNodeValidatingStakingPoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _StakingPoolManger.Contract.GetNodeValidatingStakingPoolCount(&_StakingPoolManger.CallOpts, _nodeAddress)
}

// GetPrelaunchStakingpools is a free data retrieval call binding the contract method 0xcc677617.
//
// Solidity: function getPrelaunchStakingpools(uint256 offset, uint256 limit) view returns(address[])
func (_StakingPoolManger *StakingPoolMangerCaller) GetPrelaunchStakingpools(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getPrelaunchStakingpools", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPrelaunchStakingpools is a free data retrieval call binding the contract method 0xcc677617.
//
// Solidity: function getPrelaunchStakingpools(uint256 offset, uint256 limit) view returns(address[])
func (_StakingPoolManger *StakingPoolMangerSession) GetPrelaunchStakingpools(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _StakingPoolManger.Contract.GetPrelaunchStakingpools(&_StakingPoolManger.CallOpts, offset, limit)
}

// GetPrelaunchStakingpools is a free data retrieval call binding the contract method 0xcc677617.
//
// Solidity: function getPrelaunchStakingpools(uint256 offset, uint256 limit) view returns(address[])
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetPrelaunchStakingpools(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _StakingPoolManger.Contract.GetPrelaunchStakingpools(&_StakingPoolManger.CallOpts, offset, limit)
}

// GetStakingPoolAt is a free data retrieval call binding the contract method 0x26a71b13.
//
// Solidity: function getStakingPoolAt(uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCaller) GetStakingPoolAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getStakingPoolAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingPoolAt is a free data retrieval call binding the contract method 0x26a71b13.
//
// Solidity: function getStakingPoolAt(uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerSession) GetStakingPoolAt(_index *big.Int) (common.Address, error) {
	return _StakingPoolManger.Contract.GetStakingPoolAt(&_StakingPoolManger.CallOpts, _index)
}

// GetStakingPoolAt is a free data retrieval call binding the contract method 0x26a71b13.
//
// Solidity: function getStakingPoolAt(uint256 _index) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetStakingPoolAt(_index *big.Int) (common.Address, error) {
	return _StakingPoolManger.Contract.GetStakingPoolAt(&_StakingPoolManger.CallOpts, _index)
}

// GetStakingPoolByPubkey is a free data retrieval call binding the contract method 0x02ec7481.
//
// Solidity: function getStakingPoolByPubkey(bytes _pubkey) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCaller) GetStakingPoolByPubkey(opts *bind.CallOpts, _pubkey []byte) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getStakingPoolByPubkey", _pubkey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingPoolByPubkey is a free data retrieval call binding the contract method 0x02ec7481.
//
// Solidity: function getStakingPoolByPubkey(bytes _pubkey) view returns(address)
func (_StakingPoolManger *StakingPoolMangerSession) GetStakingPoolByPubkey(_pubkey []byte) (common.Address, error) {
	return _StakingPoolManger.Contract.GetStakingPoolByPubkey(&_StakingPoolManger.CallOpts, _pubkey)
}

// GetStakingPoolByPubkey is a free data retrieval call binding the contract method 0x02ec7481.
//
// Solidity: function getStakingPoolByPubkey(bytes _pubkey) view returns(address)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetStakingPoolByPubkey(_pubkey []byte) (common.Address, error) {
	return _StakingPoolManger.Contract.GetStakingPoolByPubkey(&_StakingPoolManger.CallOpts, _pubkey)
}

// GetStakingPoolCount is a free data retrieval call binding the contract method 0x9bf00f18.
//
// Solidity: function getStakingPoolCount() view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerCaller) GetStakingPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getStakingPoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingPoolCount is a free data retrieval call binding the contract method 0x9bf00f18.
//
// Solidity: function getStakingPoolCount() view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerSession) GetStakingPoolCount() (*big.Int, error) {
	return _StakingPoolManger.Contract.GetStakingPoolCount(&_StakingPoolManger.CallOpts)
}

// GetStakingPoolCount is a free data retrieval call binding the contract method 0x9bf00f18.
//
// Solidity: function getStakingPoolCount() view returns(uint256)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetStakingPoolCount() (*big.Int, error) {
	return _StakingPoolManger.Contract.GetStakingPoolCount(&_StakingPoolManger.CallOpts)
}

// GetStakingPoolExists is a free data retrieval call binding the contract method 0x506419ca.
//
// Solidity: function getStakingPoolExists(address _stakingPoolAddress) view returns(bool)
func (_StakingPoolManger *StakingPoolMangerCaller) GetStakingPoolExists(opts *bind.CallOpts, _stakingPoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getStakingPoolExists", _stakingPoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingPoolExists is a free data retrieval call binding the contract method 0x506419ca.
//
// Solidity: function getStakingPoolExists(address _stakingPoolAddress) view returns(bool)
func (_StakingPoolManger *StakingPoolMangerSession) GetStakingPoolExists(_stakingPoolAddress common.Address) (bool, error) {
	return _StakingPoolManger.Contract.GetStakingPoolExists(&_StakingPoolManger.CallOpts, _stakingPoolAddress)
}

// GetStakingPoolExists is a free data retrieval call binding the contract method 0x506419ca.
//
// Solidity: function getStakingPoolExists(address _stakingPoolAddress) view returns(bool)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetStakingPoolExists(_stakingPoolAddress common.Address) (bool, error) {
	return _StakingPoolManger.Contract.GetStakingPoolExists(&_StakingPoolManger.CallOpts, _stakingPoolAddress)
}

// GetStakingPoolPubkey is a free data retrieval call binding the contract method 0x0d4354c6.
//
// Solidity: function getStakingPoolPubkey(address _stakingPoolAddress) view returns(bytes)
func (_StakingPoolManger *StakingPoolMangerCaller) GetStakingPoolPubkey(opts *bind.CallOpts, _stakingPoolAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getStakingPoolPubkey", _stakingPoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStakingPoolPubkey is a free data retrieval call binding the contract method 0x0d4354c6.
//
// Solidity: function getStakingPoolPubkey(address _stakingPoolAddress) view returns(bytes)
func (_StakingPoolManger *StakingPoolMangerSession) GetStakingPoolPubkey(_stakingPoolAddress common.Address) ([]byte, error) {
	return _StakingPoolManger.Contract.GetStakingPoolPubkey(&_StakingPoolManger.CallOpts, _stakingPoolAddress)
}

// GetStakingPoolPubkey is a free data retrieval call binding the contract method 0x0d4354c6.
//
// Solidity: function getStakingPoolPubkey(address _stakingPoolAddress) view returns(bytes)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetStakingPoolPubkey(_stakingPoolAddress common.Address) ([]byte, error) {
	return _StakingPoolManger.Contract.GetStakingPoolPubkey(&_StakingPoolManger.CallOpts, _stakingPoolAddress)
}

// GetStakingPoolWithdrawalProcessed is a free data retrieval call binding the contract method 0xf3e0ce85.
//
// Solidity: function getStakingPoolWithdrawalProcessed(address _stakingPoolAddress) view returns(bool)
func (_StakingPoolManger *StakingPoolMangerCaller) GetStakingPoolWithdrawalProcessed(opts *bind.CallOpts, _stakingPoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "getStakingPoolWithdrawalProcessed", _stakingPoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingPoolWithdrawalProcessed is a free data retrieval call binding the contract method 0xf3e0ce85.
//
// Solidity: function getStakingPoolWithdrawalProcessed(address _stakingPoolAddress) view returns(bool)
func (_StakingPoolManger *StakingPoolMangerSession) GetStakingPoolWithdrawalProcessed(_stakingPoolAddress common.Address) (bool, error) {
	return _StakingPoolManger.Contract.GetStakingPoolWithdrawalProcessed(&_StakingPoolManger.CallOpts, _stakingPoolAddress)
}

// GetStakingPoolWithdrawalProcessed is a free data retrieval call binding the contract method 0xf3e0ce85.
//
// Solidity: function getStakingPoolWithdrawalProcessed(address _stakingPoolAddress) view returns(bool)
func (_StakingPoolManger *StakingPoolMangerCallerSession) GetStakingPoolWithdrawalProcessed(_stakingPoolAddress common.Address) (bool, error) {
	return _StakingPoolManger.Contract.GetStakingPoolWithdrawalProcessed(&_StakingPoolManger.CallOpts, _stakingPoolAddress)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_StakingPoolManger *StakingPoolMangerCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakingPoolManger.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_StakingPoolManger *StakingPoolMangerSession) Version() (uint8, error) {
	return _StakingPoolManger.Contract.Version(&_StakingPoolManger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_StakingPoolManger *StakingPoolMangerCallerSession) Version() (uint8, error) {
	return _StakingPoolManger.Contract.Version(&_StakingPoolManger.CallOpts)
}

// CreateStakingPool is a paid mutator transaction binding the contract method 0x6e5da826.
//
// Solidity: function createStakingPool(address _nodeAddress, uint8 _depositType) returns(address)
func (_StakingPoolManger *StakingPoolMangerTransactor) CreateStakingPool(opts *bind.TransactOpts, _nodeAddress common.Address, _depositType uint8) (*types.Transaction, error) {
	return _StakingPoolManger.contract.Transact(opts, "createStakingPool", _nodeAddress, _depositType)
}

// CreateStakingPool is a paid mutator transaction binding the contract method 0x6e5da826.
//
// Solidity: function createStakingPool(address _nodeAddress, uint8 _depositType) returns(address)
func (_StakingPoolManger *StakingPoolMangerSession) CreateStakingPool(_nodeAddress common.Address, _depositType uint8) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.CreateStakingPool(&_StakingPoolManger.TransactOpts, _nodeAddress, _depositType)
}

// CreateStakingPool is a paid mutator transaction binding the contract method 0x6e5da826.
//
// Solidity: function createStakingPool(address _nodeAddress, uint8 _depositType) returns(address)
func (_StakingPoolManger *StakingPoolMangerTransactorSession) CreateStakingPool(_nodeAddress common.Address, _depositType uint8) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.CreateStakingPool(&_StakingPoolManger.TransactOpts, _nodeAddress, _depositType)
}

// DestroyStakingPool is a paid mutator transaction binding the contract method 0x6b16c49d.
//
// Solidity: function destroyStakingPool() returns()
func (_StakingPoolManger *StakingPoolMangerTransactor) DestroyStakingPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolManger.contract.Transact(opts, "destroyStakingPool")
}

// DestroyStakingPool is a paid mutator transaction binding the contract method 0x6b16c49d.
//
// Solidity: function destroyStakingPool() returns()
func (_StakingPoolManger *StakingPoolMangerSession) DestroyStakingPool() (*types.Transaction, error) {
	return _StakingPoolManger.Contract.DestroyStakingPool(&_StakingPoolManger.TransactOpts)
}

// DestroyStakingPool is a paid mutator transaction binding the contract method 0x6b16c49d.
//
// Solidity: function destroyStakingPool() returns()
func (_StakingPoolManger *StakingPoolMangerTransactorSession) DestroyStakingPool() (*types.Transaction, error) {
	return _StakingPoolManger.Contract.DestroyStakingPool(&_StakingPoolManger.TransactOpts)
}

// SetStakingPoolPubkey is a paid mutator transaction binding the contract method 0xaf40d55d.
//
// Solidity: function setStakingPoolPubkey(bytes _pubkey) returns()
func (_StakingPoolManger *StakingPoolMangerTransactor) SetStakingPoolPubkey(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _StakingPoolManger.contract.Transact(opts, "setStakingPoolPubkey", _pubkey)
}

// SetStakingPoolPubkey is a paid mutator transaction binding the contract method 0xaf40d55d.
//
// Solidity: function setStakingPoolPubkey(bytes _pubkey) returns()
func (_StakingPoolManger *StakingPoolMangerSession) SetStakingPoolPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.SetStakingPoolPubkey(&_StakingPoolManger.TransactOpts, _pubkey)
}

// SetStakingPoolPubkey is a paid mutator transaction binding the contract method 0xaf40d55d.
//
// Solidity: function setStakingPoolPubkey(bytes _pubkey) returns()
func (_StakingPoolManger *StakingPoolMangerTransactorSession) SetStakingPoolPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.SetStakingPoolPubkey(&_StakingPoolManger.TransactOpts, _pubkey)
}

// SetStakingPoolWithdrawalProcessed is a paid mutator transaction binding the contract method 0x1189504e.
//
// Solidity: function setStakingPoolWithdrawalProcessed(address _stakingPoolAddress, bool _processed) returns()
func (_StakingPoolManger *StakingPoolMangerTransactor) SetStakingPoolWithdrawalProcessed(opts *bind.TransactOpts, _stakingPoolAddress common.Address, _processed bool) (*types.Transaction, error) {
	return _StakingPoolManger.contract.Transact(opts, "setStakingPoolWithdrawalProcessed", _stakingPoolAddress, _processed)
}

// SetStakingPoolWithdrawalProcessed is a paid mutator transaction binding the contract method 0x1189504e.
//
// Solidity: function setStakingPoolWithdrawalProcessed(address _stakingPoolAddress, bool _processed) returns()
func (_StakingPoolManger *StakingPoolMangerSession) SetStakingPoolWithdrawalProcessed(_stakingPoolAddress common.Address, _processed bool) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.SetStakingPoolWithdrawalProcessed(&_StakingPoolManger.TransactOpts, _stakingPoolAddress, _processed)
}

// SetStakingPoolWithdrawalProcessed is a paid mutator transaction binding the contract method 0x1189504e.
//
// Solidity: function setStakingPoolWithdrawalProcessed(address _stakingPoolAddress, bool _processed) returns()
func (_StakingPoolManger *StakingPoolMangerTransactorSession) SetStakingPoolWithdrawalProcessed(_stakingPoolAddress common.Address, _processed bool) (*types.Transaction, error) {
	return _StakingPoolManger.Contract.SetStakingPoolWithdrawalProcessed(&_StakingPoolManger.TransactOpts, _stakingPoolAddress, _processed)
}

// StakingPoolMangerStakingPoolCreatedIterator is returned from FilterStakingPoolCreated and is used to iterate over the raw logs and unpacked data for StakingPoolCreated events raised by the StakingPoolManger contract.
type StakingPoolMangerStakingPoolCreatedIterator struct {
	Event *StakingPoolMangerStakingPoolCreated // Event containing the contract specifics and raw log

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
func (it *StakingPoolMangerStakingPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolMangerStakingPoolCreated)
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
		it.Event = new(StakingPoolMangerStakingPoolCreated)
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
func (it *StakingPoolMangerStakingPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolMangerStakingPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolMangerStakingPoolCreated represents a StakingPoolCreated event raised by the StakingPoolManger contract.
type StakingPoolMangerStakingPoolCreated struct {
	StakingPool common.Address
	Node        common.Address
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakingPoolCreated is a free log retrieval operation binding the contract event 0xdfba889e07ceb4b33e759a1015aa4ae7b45c3881d9cf96dbf09b8971a27fc4b7.
//
// Solidity: event StakingPoolCreated(address indexed stakingPool, address indexed node, uint256 time)
func (_StakingPoolManger *StakingPoolMangerFilterer) FilterStakingPoolCreated(opts *bind.FilterOpts, stakingPool []common.Address, node []common.Address) (*StakingPoolMangerStakingPoolCreatedIterator, error) {

	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StakingPoolManger.contract.FilterLogs(opts, "StakingPoolCreated", stakingPoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolMangerStakingPoolCreatedIterator{contract: _StakingPoolManger.contract, event: "StakingPoolCreated", logs: logs, sub: sub}, nil
}

// WatchStakingPoolCreated is a free log subscription operation binding the contract event 0xdfba889e07ceb4b33e759a1015aa4ae7b45c3881d9cf96dbf09b8971a27fc4b7.
//
// Solidity: event StakingPoolCreated(address indexed stakingPool, address indexed node, uint256 time)
func (_StakingPoolManger *StakingPoolMangerFilterer) WatchStakingPoolCreated(opts *bind.WatchOpts, sink chan<- *StakingPoolMangerStakingPoolCreated, stakingPool []common.Address, node []common.Address) (event.Subscription, error) {

	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StakingPoolManger.contract.WatchLogs(opts, "StakingPoolCreated", stakingPoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolMangerStakingPoolCreated)
				if err := _StakingPoolManger.contract.UnpackLog(event, "StakingPoolCreated", log); err != nil {
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

// ParseStakingPoolCreated is a log parse operation binding the contract event 0xdfba889e07ceb4b33e759a1015aa4ae7b45c3881d9cf96dbf09b8971a27fc4b7.
//
// Solidity: event StakingPoolCreated(address indexed stakingPool, address indexed node, uint256 time)
func (_StakingPoolManger *StakingPoolMangerFilterer) ParseStakingPoolCreated(log types.Log) (*StakingPoolMangerStakingPoolCreated, error) {
	event := new(StakingPoolMangerStakingPoolCreated)
	if err := _StakingPoolManger.contract.UnpackLog(event, "StakingPoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolMangerStakingPoolDestroyedIterator is returned from FilterStakingPoolDestroyed and is used to iterate over the raw logs and unpacked data for StakingPoolDestroyed events raised by the StakingPoolManger contract.
type StakingPoolMangerStakingPoolDestroyedIterator struct {
	Event *StakingPoolMangerStakingPoolDestroyed // Event containing the contract specifics and raw log

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
func (it *StakingPoolMangerStakingPoolDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolMangerStakingPoolDestroyed)
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
		it.Event = new(StakingPoolMangerStakingPoolDestroyed)
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
func (it *StakingPoolMangerStakingPoolDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolMangerStakingPoolDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolMangerStakingPoolDestroyed represents a StakingPoolDestroyed event raised by the StakingPoolManger contract.
type StakingPoolMangerStakingPoolDestroyed struct {
	StakingPool common.Address
	Node        common.Address
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakingPoolDestroyed is a free log retrieval operation binding the contract event 0x458a7280951da092252cc4009cfcf959d019fef8696665c814ebff7a8d4750ea.
//
// Solidity: event StakingPoolDestroyed(address indexed stakingPool, address indexed node, uint256 time)
func (_StakingPoolManger *StakingPoolMangerFilterer) FilterStakingPoolDestroyed(opts *bind.FilterOpts, stakingPool []common.Address, node []common.Address) (*StakingPoolMangerStakingPoolDestroyedIterator, error) {

	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StakingPoolManger.contract.FilterLogs(opts, "StakingPoolDestroyed", stakingPoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolMangerStakingPoolDestroyedIterator{contract: _StakingPoolManger.contract, event: "StakingPoolDestroyed", logs: logs, sub: sub}, nil
}

// WatchStakingPoolDestroyed is a free log subscription operation binding the contract event 0x458a7280951da092252cc4009cfcf959d019fef8696665c814ebff7a8d4750ea.
//
// Solidity: event StakingPoolDestroyed(address indexed stakingPool, address indexed node, uint256 time)
func (_StakingPoolManger *StakingPoolMangerFilterer) WatchStakingPoolDestroyed(opts *bind.WatchOpts, sink chan<- *StakingPoolMangerStakingPoolDestroyed, stakingPool []common.Address, node []common.Address) (event.Subscription, error) {

	var stakingPoolRule []interface{}
	for _, stakingPoolItem := range stakingPool {
		stakingPoolRule = append(stakingPoolRule, stakingPoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StakingPoolManger.contract.WatchLogs(opts, "StakingPoolDestroyed", stakingPoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolMangerStakingPoolDestroyed)
				if err := _StakingPoolManger.contract.UnpackLog(event, "StakingPoolDestroyed", log); err != nil {
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

// ParseStakingPoolDestroyed is a log parse operation binding the contract event 0x458a7280951da092252cc4009cfcf959d019fef8696665c814ebff7a8d4750ea.
//
// Solidity: event StakingPoolDestroyed(address indexed stakingPool, address indexed node, uint256 time)
func (_StakingPoolManger *StakingPoolMangerFilterer) ParseStakingPoolDestroyed(log types.Log) (*StakingPoolMangerStakingPoolDestroyed, error) {
	event := new(StakingPoolMangerStakingPoolDestroyed)
	if err := _StakingPoolManger.contract.UnpackLog(event, "StakingPoolDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
