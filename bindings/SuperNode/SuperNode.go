// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package super_node

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

// SuperNodeMetaData contains all meta data concerning the SuperNode contract.
var SuperNodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"SetPubkeyStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_INITIAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_MATCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_STAKING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_UNINITIAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_UNMATCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PubkeySetStorage\",\"outputs\":[{\"internalType\":\"contractIPubkeySetStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPubkeyVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSuperNodeDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSuperNodePubkeyAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getSuperNodePubkeyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"}],\"name\":\"getSuperNodePubkeyStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setSuperNodeDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"setSuperNodePubkeyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"_matchs\",\"type\":\"bool[]\"}],\"name\":\"voteWithdrawCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SuperNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use SuperNodeMetaData.ABI instead.
var SuperNodeABI = SuperNodeMetaData.ABI

// SuperNode is an auto generated Go binding around an Ethereum contract.
type SuperNode struct {
	SuperNodeCaller     // Read-only binding to the contract
	SuperNodeTransactor // Write-only binding to the contract
	SuperNodeFilterer   // Log filterer for contract events
}

// SuperNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperNodeSession struct {
	Contract     *SuperNode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperNodeCallerSession struct {
	Contract *SuperNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SuperNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperNodeTransactorSession struct {
	Contract     *SuperNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SuperNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperNodeRaw struct {
	Contract *SuperNode // Generic contract binding to access the raw methods on
}

// SuperNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperNodeCallerRaw struct {
	Contract *SuperNodeCaller // Generic read-only contract binding to access the raw methods on
}

// SuperNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperNodeTransactorRaw struct {
	Contract *SuperNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperNode creates a new instance of SuperNode, bound to a specific deployed contract.
func NewSuperNode(address common.Address, backend bind.ContractBackend) (*SuperNode, error) {
	contract, err := bindSuperNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SuperNode{SuperNodeCaller: SuperNodeCaller{contract: contract}, SuperNodeTransactor: SuperNodeTransactor{contract: contract}, SuperNodeFilterer: SuperNodeFilterer{contract: contract}}, nil
}

// NewSuperNodeCaller creates a new read-only instance of SuperNode, bound to a specific deployed contract.
func NewSuperNodeCaller(address common.Address, caller bind.ContractCaller) (*SuperNodeCaller, error) {
	contract, err := bindSuperNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperNodeCaller{contract: contract}, nil
}

// NewSuperNodeTransactor creates a new write-only instance of SuperNode, bound to a specific deployed contract.
func NewSuperNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperNodeTransactor, error) {
	contract, err := bindSuperNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperNodeTransactor{contract: contract}, nil
}

// NewSuperNodeFilterer creates a new log filterer instance of SuperNode, bound to a specific deployed contract.
func NewSuperNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperNodeFilterer, error) {
	contract, err := bindSuperNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperNodeFilterer{contract: contract}, nil
}

// bindSuperNode binds a generic wrapper to an already deployed contract.
func bindSuperNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperNode *SuperNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperNode.Contract.SuperNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperNode *SuperNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperNode.Contract.SuperNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperNode *SuperNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperNode.Contract.SuperNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperNode *SuperNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperNode *SuperNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperNode *SuperNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperNode.Contract.contract.Transact(opts, method, params...)
}

// PUBKEYSTATUSINITIAL is a free data retrieval call binding the contract method 0xae3e885a.
//
// Solidity: function PUBKEY_STATUS_INITIAL() view returns(uint256)
func (_SuperNode *SuperNodeCaller) PUBKEYSTATUSINITIAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "PUBKEY_STATUS_INITIAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSINITIAL is a free data retrieval call binding the contract method 0xae3e885a.
//
// Solidity: function PUBKEY_STATUS_INITIAL() view returns(uint256)
func (_SuperNode *SuperNodeSession) PUBKEYSTATUSINITIAL() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSINITIAL(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSINITIAL is a free data retrieval call binding the contract method 0xae3e885a.
//
// Solidity: function PUBKEY_STATUS_INITIAL() view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) PUBKEYSTATUSINITIAL() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSINITIAL(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSMATCH is a free data retrieval call binding the contract method 0x64f14dba.
//
// Solidity: function PUBKEY_STATUS_MATCH() view returns(uint256)
func (_SuperNode *SuperNodeCaller) PUBKEYSTATUSMATCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "PUBKEY_STATUS_MATCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSMATCH is a free data retrieval call binding the contract method 0x64f14dba.
//
// Solidity: function PUBKEY_STATUS_MATCH() view returns(uint256)
func (_SuperNode *SuperNodeSession) PUBKEYSTATUSMATCH() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSMATCH(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSMATCH is a free data retrieval call binding the contract method 0x64f14dba.
//
// Solidity: function PUBKEY_STATUS_MATCH() view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) PUBKEYSTATUSMATCH() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSMATCH(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSSTAKING is a free data retrieval call binding the contract method 0x40ce18a0.
//
// Solidity: function PUBKEY_STATUS_STAKING() view returns(uint256)
func (_SuperNode *SuperNodeCaller) PUBKEYSTATUSSTAKING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "PUBKEY_STATUS_STAKING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSSTAKING is a free data retrieval call binding the contract method 0x40ce18a0.
//
// Solidity: function PUBKEY_STATUS_STAKING() view returns(uint256)
func (_SuperNode *SuperNodeSession) PUBKEYSTATUSSTAKING() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSSTAKING(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSSTAKING is a free data retrieval call binding the contract method 0x40ce18a0.
//
// Solidity: function PUBKEY_STATUS_STAKING() view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) PUBKEYSTATUSSTAKING() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSSTAKING(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSUNINITIAL is a free data retrieval call binding the contract method 0x3f25522d.
//
// Solidity: function PUBKEY_STATUS_UNINITIAL() view returns(uint256)
func (_SuperNode *SuperNodeCaller) PUBKEYSTATUSUNINITIAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "PUBKEY_STATUS_UNINITIAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSUNINITIAL is a free data retrieval call binding the contract method 0x3f25522d.
//
// Solidity: function PUBKEY_STATUS_UNINITIAL() view returns(uint256)
func (_SuperNode *SuperNodeSession) PUBKEYSTATUSUNINITIAL() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSUNINITIAL(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSUNINITIAL is a free data retrieval call binding the contract method 0x3f25522d.
//
// Solidity: function PUBKEY_STATUS_UNINITIAL() view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) PUBKEYSTATUSUNINITIAL() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSUNINITIAL(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSUNMATCH is a free data retrieval call binding the contract method 0x09ba5933.
//
// Solidity: function PUBKEY_STATUS_UNMATCH() view returns(uint256)
func (_SuperNode *SuperNodeCaller) PUBKEYSTATUSUNMATCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "PUBKEY_STATUS_UNMATCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSUNMATCH is a free data retrieval call binding the contract method 0x09ba5933.
//
// Solidity: function PUBKEY_STATUS_UNMATCH() view returns(uint256)
func (_SuperNode *SuperNodeSession) PUBKEYSTATUSUNMATCH() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSUNMATCH(&_SuperNode.CallOpts)
}

// PUBKEYSTATUSUNMATCH is a free data retrieval call binding the contract method 0x09ba5933.
//
// Solidity: function PUBKEY_STATUS_UNMATCH() view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) PUBKEYSTATUSUNMATCH() (*big.Int, error) {
	return _SuperNode.Contract.PUBKEYSTATUSUNMATCH(&_SuperNode.CallOpts)
}

// PubkeySetStorage is a free data retrieval call binding the contract method 0x90d76417.
//
// Solidity: function PubkeySetStorage() view returns(address)
func (_SuperNode *SuperNodeCaller) PubkeySetStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "PubkeySetStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeySetStorage is a free data retrieval call binding the contract method 0x90d76417.
//
// Solidity: function PubkeySetStorage() view returns(address)
func (_SuperNode *SuperNodeSession) PubkeySetStorage() (common.Address, error) {
	return _SuperNode.Contract.PubkeySetStorage(&_SuperNode.CallOpts)
}

// PubkeySetStorage is a free data retrieval call binding the contract method 0x90d76417.
//
// Solidity: function PubkeySetStorage() view returns(address)
func (_SuperNode *SuperNodeCallerSession) PubkeySetStorage() (common.Address, error) {
	return _SuperNode.Contract.PubkeySetStorage(&_SuperNode.CallOpts)
}

// GetPubkeyVoted is a free data retrieval call binding the contract method 0x18582e28.
//
// Solidity: function getPubkeyVoted(bytes _validatorPubkey, address user) view returns(bool)
func (_SuperNode *SuperNodeCaller) GetPubkeyVoted(opts *bind.CallOpts, _validatorPubkey []byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "getPubkeyVoted", _validatorPubkey, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPubkeyVoted is a free data retrieval call binding the contract method 0x18582e28.
//
// Solidity: function getPubkeyVoted(bytes _validatorPubkey, address user) view returns(bool)
func (_SuperNode *SuperNodeSession) GetPubkeyVoted(_validatorPubkey []byte, user common.Address) (bool, error) {
	return _SuperNode.Contract.GetPubkeyVoted(&_SuperNode.CallOpts, _validatorPubkey, user)
}

// GetPubkeyVoted is a free data retrieval call binding the contract method 0x18582e28.
//
// Solidity: function getPubkeyVoted(bytes _validatorPubkey, address user) view returns(bool)
func (_SuperNode *SuperNodeCallerSession) GetPubkeyVoted(_validatorPubkey []byte, user common.Address) (bool, error) {
	return _SuperNode.Contract.GetPubkeyVoted(&_SuperNode.CallOpts, _validatorPubkey, user)
}

// GetSuperNodeDepositEnabled is a free data retrieval call binding the contract method 0x722ad663.
//
// Solidity: function getSuperNodeDepositEnabled() view returns(bool)
func (_SuperNode *SuperNodeCaller) GetSuperNodeDepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "getSuperNodeDepositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSuperNodeDepositEnabled is a free data retrieval call binding the contract method 0x722ad663.
//
// Solidity: function getSuperNodeDepositEnabled() view returns(bool)
func (_SuperNode *SuperNodeSession) GetSuperNodeDepositEnabled() (bool, error) {
	return _SuperNode.Contract.GetSuperNodeDepositEnabled(&_SuperNode.CallOpts)
}

// GetSuperNodeDepositEnabled is a free data retrieval call binding the contract method 0x722ad663.
//
// Solidity: function getSuperNodeDepositEnabled() view returns(bool)
func (_SuperNode *SuperNodeCallerSession) GetSuperNodeDepositEnabled() (bool, error) {
	return _SuperNode.Contract.GetSuperNodeDepositEnabled(&_SuperNode.CallOpts)
}

// GetSuperNodePubkeyAt is a free data retrieval call binding the contract method 0x323c662b.
//
// Solidity: function getSuperNodePubkeyAt(address _nodeAddress, uint256 _index) view returns(bytes)
func (_SuperNode *SuperNodeCaller) GetSuperNodePubkeyAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "getSuperNodePubkeyAt", _nodeAddress, _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSuperNodePubkeyAt is a free data retrieval call binding the contract method 0x323c662b.
//
// Solidity: function getSuperNodePubkeyAt(address _nodeAddress, uint256 _index) view returns(bytes)
func (_SuperNode *SuperNodeSession) GetSuperNodePubkeyAt(_nodeAddress common.Address, _index *big.Int) ([]byte, error) {
	return _SuperNode.Contract.GetSuperNodePubkeyAt(&_SuperNode.CallOpts, _nodeAddress, _index)
}

// GetSuperNodePubkeyAt is a free data retrieval call binding the contract method 0x323c662b.
//
// Solidity: function getSuperNodePubkeyAt(address _nodeAddress, uint256 _index) view returns(bytes)
func (_SuperNode *SuperNodeCallerSession) GetSuperNodePubkeyAt(_nodeAddress common.Address, _index *big.Int) ([]byte, error) {
	return _SuperNode.Contract.GetSuperNodePubkeyAt(&_SuperNode.CallOpts, _nodeAddress, _index)
}

// GetSuperNodePubkeyCount is a free data retrieval call binding the contract method 0x6f4c0a82.
//
// Solidity: function getSuperNodePubkeyCount(address _nodeAddress) view returns(uint256)
func (_SuperNode *SuperNodeCaller) GetSuperNodePubkeyCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "getSuperNodePubkeyCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuperNodePubkeyCount is a free data retrieval call binding the contract method 0x6f4c0a82.
//
// Solidity: function getSuperNodePubkeyCount(address _nodeAddress) view returns(uint256)
func (_SuperNode *SuperNodeSession) GetSuperNodePubkeyCount(_nodeAddress common.Address) (*big.Int, error) {
	return _SuperNode.Contract.GetSuperNodePubkeyCount(&_SuperNode.CallOpts, _nodeAddress)
}

// GetSuperNodePubkeyCount is a free data retrieval call binding the contract method 0x6f4c0a82.
//
// Solidity: function getSuperNodePubkeyCount(address _nodeAddress) view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) GetSuperNodePubkeyCount(_nodeAddress common.Address) (*big.Int, error) {
	return _SuperNode.Contract.GetSuperNodePubkeyCount(&_SuperNode.CallOpts, _nodeAddress)
}

// GetSuperNodePubkeyStatus is a free data retrieval call binding the contract method 0x996a8ddb.
//
// Solidity: function getSuperNodePubkeyStatus(bytes _validatorPubkey) view returns(uint256)
func (_SuperNode *SuperNodeCaller) GetSuperNodePubkeyStatus(opts *bind.CallOpts, _validatorPubkey []byte) (*big.Int, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "getSuperNodePubkeyStatus", _validatorPubkey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuperNodePubkeyStatus is a free data retrieval call binding the contract method 0x996a8ddb.
//
// Solidity: function getSuperNodePubkeyStatus(bytes _validatorPubkey) view returns(uint256)
func (_SuperNode *SuperNodeSession) GetSuperNodePubkeyStatus(_validatorPubkey []byte) (*big.Int, error) {
	return _SuperNode.Contract.GetSuperNodePubkeyStatus(&_SuperNode.CallOpts, _validatorPubkey)
}

// GetSuperNodePubkeyStatus is a free data retrieval call binding the contract method 0x996a8ddb.
//
// Solidity: function getSuperNodePubkeyStatus(bytes _validatorPubkey) view returns(uint256)
func (_SuperNode *SuperNodeCallerSession) GetSuperNodePubkeyStatus(_validatorPubkey []byte) (*big.Int, error) {
	return _SuperNode.Contract.GetSuperNodePubkeyStatus(&_SuperNode.CallOpts, _validatorPubkey)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SuperNode *SuperNodeCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SuperNode.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SuperNode *SuperNodeSession) Version() (uint8, error) {
	return _SuperNode.Contract.Version(&_SuperNode.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SuperNode *SuperNodeCallerSession) Version() (uint8, error) {
	return _SuperNode.Contract.Version(&_SuperNode.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_SuperNode *SuperNodeTransactor) Deposit(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _SuperNode.contract.Transact(opts, "deposit", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_SuperNode *SuperNodeSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _SuperNode.Contract.Deposit(&_SuperNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_SuperNode *SuperNodeTransactorSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _SuperNode.Contract.Deposit(&_SuperNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_SuperNode *SuperNodeTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperNode.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_SuperNode *SuperNodeSession) DepositEth() (*types.Transaction, error) {
	return _SuperNode.Contract.DepositEth(&_SuperNode.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_SuperNode *SuperNodeTransactorSession) DepositEth() (*types.Transaction, error) {
	return _SuperNode.Contract.DepositEth(&_SuperNode.TransactOpts)
}

// SetSuperNodeDepositEnabled is a paid mutator transaction binding the contract method 0x24101a4b.
//
// Solidity: function setSuperNodeDepositEnabled(bool _value) returns()
func (_SuperNode *SuperNodeTransactor) SetSuperNodeDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _SuperNode.contract.Transact(opts, "setSuperNodeDepositEnabled", _value)
}

// SetSuperNodeDepositEnabled is a paid mutator transaction binding the contract method 0x24101a4b.
//
// Solidity: function setSuperNodeDepositEnabled(bool _value) returns()
func (_SuperNode *SuperNodeSession) SetSuperNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _SuperNode.Contract.SetSuperNodeDepositEnabled(&_SuperNode.TransactOpts, _value)
}

// SetSuperNodeDepositEnabled is a paid mutator transaction binding the contract method 0x24101a4b.
//
// Solidity: function setSuperNodeDepositEnabled(bool _value) returns()
func (_SuperNode *SuperNodeTransactorSession) SetSuperNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _SuperNode.Contract.SetSuperNodeDepositEnabled(&_SuperNode.TransactOpts, _value)
}

// SetSuperNodePubkeyStatus is a paid mutator transaction binding the contract method 0xac93a5bd.
//
// Solidity: function setSuperNodePubkeyStatus(bytes _validatorPubkey, uint256 _status) returns()
func (_SuperNode *SuperNodeTransactor) SetSuperNodePubkeyStatus(opts *bind.TransactOpts, _validatorPubkey []byte, _status *big.Int) (*types.Transaction, error) {
	return _SuperNode.contract.Transact(opts, "setSuperNodePubkeyStatus", _validatorPubkey, _status)
}

// SetSuperNodePubkeyStatus is a paid mutator transaction binding the contract method 0xac93a5bd.
//
// Solidity: function setSuperNodePubkeyStatus(bytes _validatorPubkey, uint256 _status) returns()
func (_SuperNode *SuperNodeSession) SetSuperNodePubkeyStatus(_validatorPubkey []byte, _status *big.Int) (*types.Transaction, error) {
	return _SuperNode.Contract.SetSuperNodePubkeyStatus(&_SuperNode.TransactOpts, _validatorPubkey, _status)
}

// SetSuperNodePubkeyStatus is a paid mutator transaction binding the contract method 0xac93a5bd.
//
// Solidity: function setSuperNodePubkeyStatus(bytes _validatorPubkey, uint256 _status) returns()
func (_SuperNode *SuperNodeTransactorSession) SetSuperNodePubkeyStatus(_validatorPubkey []byte, _status *big.Int) (*types.Transaction, error) {
	return _SuperNode.Contract.SetSuperNodePubkeyStatus(&_SuperNode.TransactOpts, _validatorPubkey, _status)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_SuperNode *SuperNodeTransactor) Stake(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _SuperNode.contract.Transact(opts, "stake", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_SuperNode *SuperNodeSession) Stake(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _SuperNode.Contract.Stake(&_SuperNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_SuperNode *SuperNodeTransactorSession) Stake(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _SuperNode.Contract.Stake(&_SuperNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_SuperNode *SuperNodeTransactor) VoteWithdrawCredentials(opts *bind.TransactOpts, _pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _SuperNode.contract.Transact(opts, "voteWithdrawCredentials", _pubkeys, _matchs)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_SuperNode *SuperNodeSession) VoteWithdrawCredentials(_pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _SuperNode.Contract.VoteWithdrawCredentials(&_SuperNode.TransactOpts, _pubkeys, _matchs)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_SuperNode *SuperNodeTransactorSession) VoteWithdrawCredentials(_pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _SuperNode.Contract.VoteWithdrawCredentials(&_SuperNode.TransactOpts, _pubkeys, _matchs)
}

// SuperNodeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the SuperNode contract.
type SuperNodeDepositedIterator struct {
	Event *SuperNodeDeposited // Event containing the contract specifics and raw log

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
func (it *SuperNodeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperNodeDeposited)
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
		it.Event = new(SuperNodeDeposited)
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
func (it *SuperNodeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperNodeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperNodeDeposited represents a Deposited event raised by the SuperNode contract.
type SuperNodeDeposited struct {
	Node               common.Address
	Pubkey             []byte
	ValidatorSignature []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xe1ad1f084336d5e44bd6540296e05704680f98d0e6770895a3c796c79ca6d138.
//
// Solidity: event Deposited(address node, bytes pubkey, bytes validatorSignature)
func (_SuperNode *SuperNodeFilterer) FilterDeposited(opts *bind.FilterOpts) (*SuperNodeDepositedIterator, error) {

	logs, sub, err := _SuperNode.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &SuperNodeDepositedIterator{contract: _SuperNode.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xe1ad1f084336d5e44bd6540296e05704680f98d0e6770895a3c796c79ca6d138.
//
// Solidity: event Deposited(address node, bytes pubkey, bytes validatorSignature)
func (_SuperNode *SuperNodeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *SuperNodeDeposited) (event.Subscription, error) {

	logs, sub, err := _SuperNode.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperNodeDeposited)
				if err := _SuperNode.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xe1ad1f084336d5e44bd6540296e05704680f98d0e6770895a3c796c79ca6d138.
//
// Solidity: event Deposited(address node, bytes pubkey, bytes validatorSignature)
func (_SuperNode *SuperNodeFilterer) ParseDeposited(log types.Log) (*SuperNodeDeposited, error) {
	event := new(SuperNodeDeposited)
	if err := _SuperNode.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperNodeEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the SuperNode contract.
type SuperNodeEtherDepositedIterator struct {
	Event *SuperNodeEtherDeposited // Event containing the contract specifics and raw log

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
func (it *SuperNodeEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperNodeEtherDeposited)
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
		it.Event = new(SuperNodeEtherDeposited)
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
func (it *SuperNodeEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperNodeEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperNodeEtherDeposited represents a EtherDeposited event raised by the SuperNode contract.
type SuperNodeEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_SuperNode *SuperNodeFilterer) FilterEtherDeposited(opts *bind.FilterOpts, from []common.Address) (*SuperNodeEtherDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SuperNode.contract.FilterLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &SuperNodeEtherDepositedIterator{contract: _SuperNode.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_SuperNode *SuperNodeFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *SuperNodeEtherDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SuperNode.contract.WatchLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperNodeEtherDeposited)
				if err := _SuperNode.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_SuperNode *SuperNodeFilterer) ParseEtherDeposited(log types.Log) (*SuperNodeEtherDeposited, error) {
	event := new(SuperNodeEtherDeposited)
	if err := _SuperNode.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperNodeSetPubkeyStatusIterator is returned from FilterSetPubkeyStatus and is used to iterate over the raw logs and unpacked data for SetPubkeyStatus events raised by the SuperNode contract.
type SuperNodeSetPubkeyStatusIterator struct {
	Event *SuperNodeSetPubkeyStatus // Event containing the contract specifics and raw log

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
func (it *SuperNodeSetPubkeyStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperNodeSetPubkeyStatus)
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
		it.Event = new(SuperNodeSetPubkeyStatus)
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
func (it *SuperNodeSetPubkeyStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperNodeSetPubkeyStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperNodeSetPubkeyStatus represents a SetPubkeyStatus event raised by the SuperNode contract.
type SuperNodeSetPubkeyStatus struct {
	Pubkey []byte
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPubkeyStatus is a free log retrieval operation binding the contract event 0xc581330a67ddadcf62d06e14db41c85bb45910369cedad649390e6dd2a9373c8.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint256 status)
func (_SuperNode *SuperNodeFilterer) FilterSetPubkeyStatus(opts *bind.FilterOpts) (*SuperNodeSetPubkeyStatusIterator, error) {

	logs, sub, err := _SuperNode.contract.FilterLogs(opts, "SetPubkeyStatus")
	if err != nil {
		return nil, err
	}
	return &SuperNodeSetPubkeyStatusIterator{contract: _SuperNode.contract, event: "SetPubkeyStatus", logs: logs, sub: sub}, nil
}

// WatchSetPubkeyStatus is a free log subscription operation binding the contract event 0xc581330a67ddadcf62d06e14db41c85bb45910369cedad649390e6dd2a9373c8.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint256 status)
func (_SuperNode *SuperNodeFilterer) WatchSetPubkeyStatus(opts *bind.WatchOpts, sink chan<- *SuperNodeSetPubkeyStatus) (event.Subscription, error) {

	logs, sub, err := _SuperNode.contract.WatchLogs(opts, "SetPubkeyStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperNodeSetPubkeyStatus)
				if err := _SuperNode.contract.UnpackLog(event, "SetPubkeyStatus", log); err != nil {
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

// ParseSetPubkeyStatus is a log parse operation binding the contract event 0xc581330a67ddadcf62d06e14db41c85bb45910369cedad649390e6dd2a9373c8.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint256 status)
func (_SuperNode *SuperNodeFilterer) ParseSetPubkeyStatus(log types.Log) (*SuperNodeSetPubkeyStatus, error) {
	event := new(SuperNodeSetPubkeyStatus)
	if err := _SuperNode.contract.UnpackLog(event, "SetPubkeyStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperNodeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the SuperNode contract.
type SuperNodeStakedIterator struct {
	Event *SuperNodeStaked // Event containing the contract specifics and raw log

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
func (it *SuperNodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperNodeStaked)
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
		it.Event = new(SuperNodeStaked)
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
func (it *SuperNodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperNodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperNodeStaked represents a Staked event raised by the SuperNode contract.
type SuperNodeStaked struct {
	Node   common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_SuperNode *SuperNodeFilterer) FilterStaked(opts *bind.FilterOpts) (*SuperNodeStakedIterator, error) {

	logs, sub, err := _SuperNode.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &SuperNodeStakedIterator{contract: _SuperNode.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_SuperNode *SuperNodeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *SuperNodeStaked) (event.Subscription, error) {

	logs, sub, err := _SuperNode.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperNodeStaked)
				if err := _SuperNode.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_SuperNode *SuperNodeFilterer) ParseStaked(log types.Log) (*SuperNodeStaked, error) {
	event := new(SuperNodeStaked)
	if err := _SuperNode.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
