// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node_manager

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

// NodeManagerMetaData contains all meta data concerning the NodeManager contract.
var NodeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"trusted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeSuperSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"trusted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeTrustedSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeTrusted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSuperNodeAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSuperNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getSuperNodeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTrustedNodeAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_super\",\"type\":\"bool\"}],\"name\":\"setNodeSuper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_trusted\",\"type\":\"bool\"}],\"name\":\"setNodeTrusted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NodeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeManagerMetaData.ABI instead.
var NodeManagerABI = NodeManagerMetaData.ABI

// NodeManager is an auto generated Go binding around an Ethereum contract.
type NodeManager struct {
	NodeManagerCaller     // Read-only binding to the contract
	NodeManagerTransactor // Write-only binding to the contract
	NodeManagerFilterer   // Log filterer for contract events
}

// NodeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerSession struct {
	Contract     *NodeManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerCallerSession struct {
	Contract *NodeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerTransactorSession struct {
	Contract     *NodeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerRaw struct {
	Contract *NodeManager // Generic contract binding to access the raw methods on
}

// NodeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerCallerRaw struct {
	Contract *NodeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerTransactorRaw struct {
	Contract *NodeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManager creates a new instance of NodeManager, bound to a specific deployed contract.
func NewNodeManager(address common.Address, backend bind.ContractBackend) (*NodeManager, error) {
	contract, err := bindNodeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManager{NodeManagerCaller: NodeManagerCaller{contract: contract}, NodeManagerTransactor: NodeManagerTransactor{contract: contract}, NodeManagerFilterer: NodeManagerFilterer{contract: contract}}, nil
}

// NewNodeManagerCaller creates a new read-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerCaller(address common.Address, caller bind.ContractCaller) (*NodeManagerCaller, error) {
	contract, err := bindNodeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerCaller{contract: contract}, nil
}

// NewNodeManagerTransactor creates a new write-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerTransactor, error) {
	contract, err := bindNodeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerTransactor{contract: contract}, nil
}

// NewNodeManagerFilterer creates a new log filterer instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerFilterer, error) {
	contract, err := bindNodeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerFilterer{contract: contract}, nil
}

// bindNodeManager binds a generic wrapper to an already deployed contract.
func bindNodeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.NodeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transact(opts, method, params...)
}

// GetNodeAt is a free data retrieval call binding the contract method 0xba75d806.
//
// Solidity: function getNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerCaller) GetNodeAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getNodeAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeAt is a free data retrieval call binding the contract method 0xba75d806.
//
// Solidity: function getNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerSession) GetNodeAt(_index *big.Int) (common.Address, error) {
	return _NodeManager.Contract.GetNodeAt(&_NodeManager.CallOpts, _index)
}

// GetNodeAt is a free data retrieval call binding the contract method 0xba75d806.
//
// Solidity: function getNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerCallerSession) GetNodeAt(_index *big.Int) (common.Address, error) {
	return _NodeManager.Contract.GetNodeAt(&_NodeManager.CallOpts, _index)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerSession) GetNodeCount() (*big.Int, error) {
	return _NodeManager.Contract.GetNodeCount(&_NodeManager.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetNodeCount() (*big.Int, error) {
	return _NodeManager.Contract.GetNodeCount(&_NodeManager.CallOpts)
}

// GetNodeExists is a free data retrieval call binding the contract method 0x65d4176f.
//
// Solidity: function getNodeExists(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerCaller) GetNodeExists(opts *bind.CallOpts, _nodeAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getNodeExists", _nodeAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeExists is a free data retrieval call binding the contract method 0x65d4176f.
//
// Solidity: function getNodeExists(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerSession) GetNodeExists(_nodeAddress common.Address) (bool, error) {
	return _NodeManager.Contract.GetNodeExists(&_NodeManager.CallOpts, _nodeAddress)
}

// GetNodeExists is a free data retrieval call binding the contract method 0x65d4176f.
//
// Solidity: function getNodeExists(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerCallerSession) GetNodeExists(_nodeAddress common.Address) (bool, error) {
	return _NodeManager.Contract.GetNodeExists(&_NodeManager.CallOpts, _nodeAddress)
}

// GetNodeTrusted is a free data retrieval call binding the contract method 0xb13e9d29.
//
// Solidity: function getNodeTrusted(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerCaller) GetNodeTrusted(opts *bind.CallOpts, _nodeAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getNodeTrusted", _nodeAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeTrusted is a free data retrieval call binding the contract method 0xb13e9d29.
//
// Solidity: function getNodeTrusted(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerSession) GetNodeTrusted(_nodeAddress common.Address) (bool, error) {
	return _NodeManager.Contract.GetNodeTrusted(&_NodeManager.CallOpts, _nodeAddress)
}

// GetNodeTrusted is a free data retrieval call binding the contract method 0xb13e9d29.
//
// Solidity: function getNodeTrusted(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerCallerSession) GetNodeTrusted(_nodeAddress common.Address) (bool, error) {
	return _NodeManager.Contract.GetNodeTrusted(&_NodeManager.CallOpts, _nodeAddress)
}

// GetSuperNodeAt is a free data retrieval call binding the contract method 0xc8974907.
//
// Solidity: function getSuperNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerCaller) GetSuperNodeAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getSuperNodeAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSuperNodeAt is a free data retrieval call binding the contract method 0xc8974907.
//
// Solidity: function getSuperNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerSession) GetSuperNodeAt(_index *big.Int) (common.Address, error) {
	return _NodeManager.Contract.GetSuperNodeAt(&_NodeManager.CallOpts, _index)
}

// GetSuperNodeAt is a free data retrieval call binding the contract method 0xc8974907.
//
// Solidity: function getSuperNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerCallerSession) GetSuperNodeAt(_index *big.Int) (common.Address, error) {
	return _NodeManager.Contract.GetSuperNodeAt(&_NodeManager.CallOpts, _index)
}

// GetSuperNodeCount is a free data retrieval call binding the contract method 0x56071cf0.
//
// Solidity: function getSuperNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerCaller) GetSuperNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getSuperNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuperNodeCount is a free data retrieval call binding the contract method 0x56071cf0.
//
// Solidity: function getSuperNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerSession) GetSuperNodeCount() (*big.Int, error) {
	return _NodeManager.Contract.GetSuperNodeCount(&_NodeManager.CallOpts)
}

// GetSuperNodeCount is a free data retrieval call binding the contract method 0x56071cf0.
//
// Solidity: function getSuperNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetSuperNodeCount() (*big.Int, error) {
	return _NodeManager.Contract.GetSuperNodeCount(&_NodeManager.CallOpts)
}

// GetSuperNodeExists is a free data retrieval call binding the contract method 0xca66d1d1.
//
// Solidity: function getSuperNodeExists(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerCaller) GetSuperNodeExists(opts *bind.CallOpts, _nodeAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getSuperNodeExists", _nodeAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSuperNodeExists is a free data retrieval call binding the contract method 0xca66d1d1.
//
// Solidity: function getSuperNodeExists(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerSession) GetSuperNodeExists(_nodeAddress common.Address) (bool, error) {
	return _NodeManager.Contract.GetSuperNodeExists(&_NodeManager.CallOpts, _nodeAddress)
}

// GetSuperNodeExists is a free data retrieval call binding the contract method 0xca66d1d1.
//
// Solidity: function getSuperNodeExists(address _nodeAddress) view returns(bool)
func (_NodeManager *NodeManagerCallerSession) GetSuperNodeExists(_nodeAddress common.Address) (bool, error) {
	return _NodeManager.Contract.GetSuperNodeExists(&_NodeManager.CallOpts, _nodeAddress)
}

// GetTrustedNodeAt is a free data retrieval call binding the contract method 0xfe26a5c8.
//
// Solidity: function getTrustedNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerCaller) GetTrustedNodeAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getTrustedNodeAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTrustedNodeAt is a free data retrieval call binding the contract method 0xfe26a5c8.
//
// Solidity: function getTrustedNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerSession) GetTrustedNodeAt(_index *big.Int) (common.Address, error) {
	return _NodeManager.Contract.GetTrustedNodeAt(&_NodeManager.CallOpts, _index)
}

// GetTrustedNodeAt is a free data retrieval call binding the contract method 0xfe26a5c8.
//
// Solidity: function getTrustedNodeAt(uint256 _index) view returns(address)
func (_NodeManager *NodeManagerCallerSession) GetTrustedNodeAt(_index *big.Int) (common.Address, error) {
	return _NodeManager.Contract.GetTrustedNodeAt(&_NodeManager.CallOpts, _index)
}

// GetTrustedNodeCount is a free data retrieval call binding the contract method 0x66e3c026.
//
// Solidity: function getTrustedNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerCaller) GetTrustedNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getTrustedNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTrustedNodeCount is a free data retrieval call binding the contract method 0x66e3c026.
//
// Solidity: function getTrustedNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerSession) GetTrustedNodeCount() (*big.Int, error) {
	return _NodeManager.Contract.GetTrustedNodeCount(&_NodeManager.CallOpts)
}

// GetTrustedNodeCount is a free data retrieval call binding the contract method 0x66e3c026.
//
// Solidity: function getTrustedNodeCount() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetTrustedNodeCount() (*big.Int, error) {
	return _NodeManager.Contract.GetTrustedNodeCount(&_NodeManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeManager *NodeManagerCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeManager *NodeManagerSession) Version() (uint8, error) {
	return _NodeManager.Contract.Version(&_NodeManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeManager *NodeManagerCallerSession) Version() (uint8, error) {
	return _NodeManager.Contract.Version(&_NodeManager.CallOpts)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x672d7a0d.
//
// Solidity: function registerNode(address _nodeAddress) returns()
func (_NodeManager *NodeManagerTransactor) RegisterNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "registerNode", _nodeAddress)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x672d7a0d.
//
// Solidity: function registerNode(address _nodeAddress) returns()
func (_NodeManager *NodeManagerSession) RegisterNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _NodeManager.Contract.RegisterNode(&_NodeManager.TransactOpts, _nodeAddress)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x672d7a0d.
//
// Solidity: function registerNode(address _nodeAddress) returns()
func (_NodeManager *NodeManagerTransactorSession) RegisterNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _NodeManager.Contract.RegisterNode(&_NodeManager.TransactOpts, _nodeAddress)
}

// SetNodeSuper is a paid mutator transaction binding the contract method 0xaabcb8aa.
//
// Solidity: function setNodeSuper(address _nodeAddress, bool _super) returns()
func (_NodeManager *NodeManagerTransactor) SetNodeSuper(opts *bind.TransactOpts, _nodeAddress common.Address, _super bool) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "setNodeSuper", _nodeAddress, _super)
}

// SetNodeSuper is a paid mutator transaction binding the contract method 0xaabcb8aa.
//
// Solidity: function setNodeSuper(address _nodeAddress, bool _super) returns()
func (_NodeManager *NodeManagerSession) SetNodeSuper(_nodeAddress common.Address, _super bool) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNodeSuper(&_NodeManager.TransactOpts, _nodeAddress, _super)
}

// SetNodeSuper is a paid mutator transaction binding the contract method 0xaabcb8aa.
//
// Solidity: function setNodeSuper(address _nodeAddress, bool _super) returns()
func (_NodeManager *NodeManagerTransactorSession) SetNodeSuper(_nodeAddress common.Address, _super bool) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNodeSuper(&_NodeManager.TransactOpts, _nodeAddress, _super)
}

// SetNodeTrusted is a paid mutator transaction binding the contract method 0x57d63792.
//
// Solidity: function setNodeTrusted(address _nodeAddress, bool _trusted) returns()
func (_NodeManager *NodeManagerTransactor) SetNodeTrusted(opts *bind.TransactOpts, _nodeAddress common.Address, _trusted bool) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "setNodeTrusted", _nodeAddress, _trusted)
}

// SetNodeTrusted is a paid mutator transaction binding the contract method 0x57d63792.
//
// Solidity: function setNodeTrusted(address _nodeAddress, bool _trusted) returns()
func (_NodeManager *NodeManagerSession) SetNodeTrusted(_nodeAddress common.Address, _trusted bool) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNodeTrusted(&_NodeManager.TransactOpts, _nodeAddress, _trusted)
}

// SetNodeTrusted is a paid mutator transaction binding the contract method 0x57d63792.
//
// Solidity: function setNodeTrusted(address _nodeAddress, bool _trusted) returns()
func (_NodeManager *NodeManagerTransactorSession) SetNodeTrusted(_nodeAddress common.Address, _trusted bool) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNodeTrusted(&_NodeManager.TransactOpts, _nodeAddress, _trusted)
}

// NodeManagerNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the NodeManager contract.
type NodeManagerNodeRegisteredIterator struct {
	Event *NodeManagerNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeRegistered)
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
		it.Event = new(NodeManagerNodeRegistered)
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
func (it *NodeManagerNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeRegistered represents a NodeRegistered event raised by the NodeManager contract.
type NodeManagerNodeRegistered struct {
	Node common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0xf773bca07d020a1bc1fdd45ea3db573da547dd27180143afaf075c158a847594.
//
// Solidity: event NodeRegistered(address indexed node, uint256 time)
func (_NodeManager *NodeManagerFilterer) FilterNodeRegistered(opts *bind.FilterOpts, node []common.Address) (*NodeManagerNodeRegisteredIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeRegistered", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeRegisteredIterator{contract: _NodeManager.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0xf773bca07d020a1bc1fdd45ea3db573da547dd27180143afaf075c158a847594.
//
// Solidity: event NodeRegistered(address indexed node, uint256 time)
func (_NodeManager *NodeManagerFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeRegistered, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeRegistered", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeRegistered)
				if err := _NodeManager.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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

// ParseNodeRegistered is a log parse operation binding the contract event 0xf773bca07d020a1bc1fdd45ea3db573da547dd27180143afaf075c158a847594.
//
// Solidity: event NodeRegistered(address indexed node, uint256 time)
func (_NodeManager *NodeManagerFilterer) ParseNodeRegistered(log types.Log) (*NodeManagerNodeRegistered, error) {
	event := new(NodeManagerNodeRegistered)
	if err := _NodeManager.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerNodeSuperSetIterator is returned from FilterNodeSuperSet and is used to iterate over the raw logs and unpacked data for NodeSuperSet events raised by the NodeManager contract.
type NodeManagerNodeSuperSetIterator struct {
	Event *NodeManagerNodeSuperSet // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeSuperSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeSuperSet)
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
		it.Event = new(NodeManagerNodeSuperSet)
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
func (it *NodeManagerNodeSuperSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeSuperSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeSuperSet represents a NodeSuperSet event raised by the NodeManager contract.
type NodeManagerNodeSuperSet struct {
	Node    common.Address
	Trusted bool
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeSuperSet is a free log retrieval operation binding the contract event 0x4e254d8fe04424f1b37e26b86bded3b00bb5d712b323242e413c84b25115b4f8.
//
// Solidity: event NodeSuperSet(address indexed node, bool trusted, uint256 time)
func (_NodeManager *NodeManagerFilterer) FilterNodeSuperSet(opts *bind.FilterOpts, node []common.Address) (*NodeManagerNodeSuperSetIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeSuperSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeSuperSetIterator{contract: _NodeManager.contract, event: "NodeSuperSet", logs: logs, sub: sub}, nil
}

// WatchNodeSuperSet is a free log subscription operation binding the contract event 0x4e254d8fe04424f1b37e26b86bded3b00bb5d712b323242e413c84b25115b4f8.
//
// Solidity: event NodeSuperSet(address indexed node, bool trusted, uint256 time)
func (_NodeManager *NodeManagerFilterer) WatchNodeSuperSet(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeSuperSet, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeSuperSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeSuperSet)
				if err := _NodeManager.contract.UnpackLog(event, "NodeSuperSet", log); err != nil {
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

// ParseNodeSuperSet is a log parse operation binding the contract event 0x4e254d8fe04424f1b37e26b86bded3b00bb5d712b323242e413c84b25115b4f8.
//
// Solidity: event NodeSuperSet(address indexed node, bool trusted, uint256 time)
func (_NodeManager *NodeManagerFilterer) ParseNodeSuperSet(log types.Log) (*NodeManagerNodeSuperSet, error) {
	event := new(NodeManagerNodeSuperSet)
	if err := _NodeManager.contract.UnpackLog(event, "NodeSuperSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerNodeTrustedSetIterator is returned from FilterNodeTrustedSet and is used to iterate over the raw logs and unpacked data for NodeTrustedSet events raised by the NodeManager contract.
type NodeManagerNodeTrustedSetIterator struct {
	Event *NodeManagerNodeTrustedSet // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeTrustedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeTrustedSet)
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
		it.Event = new(NodeManagerNodeTrustedSet)
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
func (it *NodeManagerNodeTrustedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeTrustedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeTrustedSet represents a NodeTrustedSet event raised by the NodeManager contract.
type NodeManagerNodeTrustedSet struct {
	Node    common.Address
	Trusted bool
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeTrustedSet is a free log retrieval operation binding the contract event 0xc1308995890d10efade1f87f0bdb6020a8e7b6d3c3cfb6a1d9813d4ecf200aca.
//
// Solidity: event NodeTrustedSet(address indexed node, bool trusted, uint256 time)
func (_NodeManager *NodeManagerFilterer) FilterNodeTrustedSet(opts *bind.FilterOpts, node []common.Address) (*NodeManagerNodeTrustedSetIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeTrustedSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeTrustedSetIterator{contract: _NodeManager.contract, event: "NodeTrustedSet", logs: logs, sub: sub}, nil
}

// WatchNodeTrustedSet is a free log subscription operation binding the contract event 0xc1308995890d10efade1f87f0bdb6020a8e7b6d3c3cfb6a1d9813d4ecf200aca.
//
// Solidity: event NodeTrustedSet(address indexed node, bool trusted, uint256 time)
func (_NodeManager *NodeManagerFilterer) WatchNodeTrustedSet(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeTrustedSet, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeTrustedSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeTrustedSet)
				if err := _NodeManager.contract.UnpackLog(event, "NodeTrustedSet", log); err != nil {
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

// ParseNodeTrustedSet is a log parse operation binding the contract event 0xc1308995890d10efade1f87f0bdb6020a8e7b6d3c3cfb6a1d9813d4ecf200aca.
//
// Solidity: event NodeTrustedSet(address indexed node, bool trusted, uint256 time)
func (_NodeManager *NodeManagerFilterer) ParseNodeTrustedSet(log types.Log) (*NodeManagerNodeTrustedSet, error) {
	event := new(NodeManagerNodeTrustedSet)
	if err := _NodeManager.contract.UnpackLog(event, "NodeTrustedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
