// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package light_node

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

// LightNodeMetaData contains all meta data concerning the LightNode contract.
var LightNodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"OffBoarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"SetPubkeyStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_CANWITHDRAW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_INITIAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_MATCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_OFFBOARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_STAKING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_UNINITIAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_UNMATCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_STATUS_WITHDRAWED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PubkeySetStorage\",\"outputs\":[{\"internalType\":\"contractIPubkeySetStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentNodeDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLightNodeDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLightNodePubkeyAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getLightNodePubkeyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"}],\"name\":\"getLightNodePubkeyStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPubkeyVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"}],\"name\":\"offBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"}],\"name\":\"provideNodeDepositToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveEtherWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setLightNodeDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"setLightNodePubkeyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"_matchs\",\"type\":\"bool[]\"}],\"name\":\"voteWithdrawCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"}],\"name\":\"withdrawNodeDepositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LightNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use LightNodeMetaData.ABI instead.
var LightNodeABI = LightNodeMetaData.ABI

// LightNode is an auto generated Go binding around an Ethereum contract.
type LightNode struct {
	LightNodeCaller     // Read-only binding to the contract
	LightNodeTransactor // Write-only binding to the contract
	LightNodeFilterer   // Log filterer for contract events
}

// LightNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightNodeSession struct {
	Contract     *LightNode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightNodeCallerSession struct {
	Contract *LightNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LightNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightNodeTransactorSession struct {
	Contract     *LightNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LightNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightNodeRaw struct {
	Contract *LightNode // Generic contract binding to access the raw methods on
}

// LightNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightNodeCallerRaw struct {
	Contract *LightNodeCaller // Generic read-only contract binding to access the raw methods on
}

// LightNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightNodeTransactorRaw struct {
	Contract *LightNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightNode creates a new instance of LightNode, bound to a specific deployed contract.
func NewLightNode(address common.Address, backend bind.ContractBackend) (*LightNode, error) {
	contract, err := bindLightNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightNode{LightNodeCaller: LightNodeCaller{contract: contract}, LightNodeTransactor: LightNodeTransactor{contract: contract}, LightNodeFilterer: LightNodeFilterer{contract: contract}}, nil
}

// NewLightNodeCaller creates a new read-only instance of LightNode, bound to a specific deployed contract.
func NewLightNodeCaller(address common.Address, caller bind.ContractCaller) (*LightNodeCaller, error) {
	contract, err := bindLightNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightNodeCaller{contract: contract}, nil
}

// NewLightNodeTransactor creates a new write-only instance of LightNode, bound to a specific deployed contract.
func NewLightNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*LightNodeTransactor, error) {
	contract, err := bindLightNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightNodeTransactor{contract: contract}, nil
}

// NewLightNodeFilterer creates a new log filterer instance of LightNode, bound to a specific deployed contract.
func NewLightNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*LightNodeFilterer, error) {
	contract, err := bindLightNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightNodeFilterer{contract: contract}, nil
}

// bindLightNode binds a generic wrapper to an already deployed contract.
func bindLightNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LightNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightNode *LightNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightNode.Contract.LightNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightNode *LightNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightNode.Contract.LightNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightNode *LightNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightNode.Contract.LightNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightNode *LightNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightNode *LightNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightNode *LightNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightNode.Contract.contract.Transact(opts, method, params...)
}

// PUBKEYSTATUSCANWITHDRAW is a free data retrieval call binding the contract method 0x1e8bc826.
//
// Solidity: function PUBKEY_STATUS_CANWITHDRAW() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSCANWITHDRAW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_CANWITHDRAW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSCANWITHDRAW is a free data retrieval call binding the contract method 0x1e8bc826.
//
// Solidity: function PUBKEY_STATUS_CANWITHDRAW() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSCANWITHDRAW() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSCANWITHDRAW(&_LightNode.CallOpts)
}

// PUBKEYSTATUSCANWITHDRAW is a free data retrieval call binding the contract method 0x1e8bc826.
//
// Solidity: function PUBKEY_STATUS_CANWITHDRAW() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSCANWITHDRAW() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSCANWITHDRAW(&_LightNode.CallOpts)
}

// PUBKEYSTATUSINITIAL is a free data retrieval call binding the contract method 0xae3e885a.
//
// Solidity: function PUBKEY_STATUS_INITIAL() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSINITIAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_INITIAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSINITIAL is a free data retrieval call binding the contract method 0xae3e885a.
//
// Solidity: function PUBKEY_STATUS_INITIAL() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSINITIAL() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSINITIAL(&_LightNode.CallOpts)
}

// PUBKEYSTATUSINITIAL is a free data retrieval call binding the contract method 0xae3e885a.
//
// Solidity: function PUBKEY_STATUS_INITIAL() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSINITIAL() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSINITIAL(&_LightNode.CallOpts)
}

// PUBKEYSTATUSMATCH is a free data retrieval call binding the contract method 0x64f14dba.
//
// Solidity: function PUBKEY_STATUS_MATCH() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSMATCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_MATCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSMATCH is a free data retrieval call binding the contract method 0x64f14dba.
//
// Solidity: function PUBKEY_STATUS_MATCH() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSMATCH() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSMATCH(&_LightNode.CallOpts)
}

// PUBKEYSTATUSMATCH is a free data retrieval call binding the contract method 0x64f14dba.
//
// Solidity: function PUBKEY_STATUS_MATCH() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSMATCH() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSMATCH(&_LightNode.CallOpts)
}

// PUBKEYSTATUSOFFBOARD is a free data retrieval call binding the contract method 0x9532e08a.
//
// Solidity: function PUBKEY_STATUS_OFFBOARD() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSOFFBOARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_OFFBOARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSOFFBOARD is a free data retrieval call binding the contract method 0x9532e08a.
//
// Solidity: function PUBKEY_STATUS_OFFBOARD() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSOFFBOARD() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSOFFBOARD(&_LightNode.CallOpts)
}

// PUBKEYSTATUSOFFBOARD is a free data retrieval call binding the contract method 0x9532e08a.
//
// Solidity: function PUBKEY_STATUS_OFFBOARD() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSOFFBOARD() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSOFFBOARD(&_LightNode.CallOpts)
}

// PUBKEYSTATUSSTAKING is a free data retrieval call binding the contract method 0x40ce18a0.
//
// Solidity: function PUBKEY_STATUS_STAKING() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSSTAKING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_STAKING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSSTAKING is a free data retrieval call binding the contract method 0x40ce18a0.
//
// Solidity: function PUBKEY_STATUS_STAKING() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSSTAKING() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSSTAKING(&_LightNode.CallOpts)
}

// PUBKEYSTATUSSTAKING is a free data retrieval call binding the contract method 0x40ce18a0.
//
// Solidity: function PUBKEY_STATUS_STAKING() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSSTAKING() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSSTAKING(&_LightNode.CallOpts)
}

// PUBKEYSTATUSUNINITIAL is a free data retrieval call binding the contract method 0x3f25522d.
//
// Solidity: function PUBKEY_STATUS_UNINITIAL() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSUNINITIAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_UNINITIAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSUNINITIAL is a free data retrieval call binding the contract method 0x3f25522d.
//
// Solidity: function PUBKEY_STATUS_UNINITIAL() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSUNINITIAL() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSUNINITIAL(&_LightNode.CallOpts)
}

// PUBKEYSTATUSUNINITIAL is a free data retrieval call binding the contract method 0x3f25522d.
//
// Solidity: function PUBKEY_STATUS_UNINITIAL() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSUNINITIAL() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSUNINITIAL(&_LightNode.CallOpts)
}

// PUBKEYSTATUSUNMATCH is a free data retrieval call binding the contract method 0x09ba5933.
//
// Solidity: function PUBKEY_STATUS_UNMATCH() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSUNMATCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_UNMATCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSUNMATCH is a free data retrieval call binding the contract method 0x09ba5933.
//
// Solidity: function PUBKEY_STATUS_UNMATCH() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSUNMATCH() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSUNMATCH(&_LightNode.CallOpts)
}

// PUBKEYSTATUSUNMATCH is a free data retrieval call binding the contract method 0x09ba5933.
//
// Solidity: function PUBKEY_STATUS_UNMATCH() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSUNMATCH() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSUNMATCH(&_LightNode.CallOpts)
}

// PUBKEYSTATUSWITHDRAWED is a free data retrieval call binding the contract method 0x68d25fe8.
//
// Solidity: function PUBKEY_STATUS_WITHDRAWED() view returns(uint256)
func (_LightNode *LightNodeCaller) PUBKEYSTATUSWITHDRAWED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PUBKEY_STATUS_WITHDRAWED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYSTATUSWITHDRAWED is a free data retrieval call binding the contract method 0x68d25fe8.
//
// Solidity: function PUBKEY_STATUS_WITHDRAWED() view returns(uint256)
func (_LightNode *LightNodeSession) PUBKEYSTATUSWITHDRAWED() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSWITHDRAWED(&_LightNode.CallOpts)
}

// PUBKEYSTATUSWITHDRAWED is a free data retrieval call binding the contract method 0x68d25fe8.
//
// Solidity: function PUBKEY_STATUS_WITHDRAWED() view returns(uint256)
func (_LightNode *LightNodeCallerSession) PUBKEYSTATUSWITHDRAWED() (*big.Int, error) {
	return _LightNode.Contract.PUBKEYSTATUSWITHDRAWED(&_LightNode.CallOpts)
}

// PubkeySetStorage is a free data retrieval call binding the contract method 0x90d76417.
//
// Solidity: function PubkeySetStorage() view returns(address)
func (_LightNode *LightNodeCaller) PubkeySetStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "PubkeySetStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeySetStorage is a free data retrieval call binding the contract method 0x90d76417.
//
// Solidity: function PubkeySetStorage() view returns(address)
func (_LightNode *LightNodeSession) PubkeySetStorage() (common.Address, error) {
	return _LightNode.Contract.PubkeySetStorage(&_LightNode.CallOpts)
}

// PubkeySetStorage is a free data retrieval call binding the contract method 0x90d76417.
//
// Solidity: function PubkeySetStorage() view returns(address)
func (_LightNode *LightNodeCallerSession) PubkeySetStorage() (common.Address, error) {
	return _LightNode.Contract.PubkeySetStorage(&_LightNode.CallOpts)
}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_LightNode *LightNodeCaller) GetCurrentNodeDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getCurrentNodeDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_LightNode *LightNodeSession) GetCurrentNodeDepositAmount() (*big.Int, error) {
	return _LightNode.Contract.GetCurrentNodeDepositAmount(&_LightNode.CallOpts)
}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_LightNode *LightNodeCallerSession) GetCurrentNodeDepositAmount() (*big.Int, error) {
	return _LightNode.Contract.GetCurrentNodeDepositAmount(&_LightNode.CallOpts)
}

// GetLightNodeDepositEnabled is a free data retrieval call binding the contract method 0xa32cb341.
//
// Solidity: function getLightNodeDepositEnabled() view returns(bool)
func (_LightNode *LightNodeCaller) GetLightNodeDepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getLightNodeDepositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetLightNodeDepositEnabled is a free data retrieval call binding the contract method 0xa32cb341.
//
// Solidity: function getLightNodeDepositEnabled() view returns(bool)
func (_LightNode *LightNodeSession) GetLightNodeDepositEnabled() (bool, error) {
	return _LightNode.Contract.GetLightNodeDepositEnabled(&_LightNode.CallOpts)
}

// GetLightNodeDepositEnabled is a free data retrieval call binding the contract method 0xa32cb341.
//
// Solidity: function getLightNodeDepositEnabled() view returns(bool)
func (_LightNode *LightNodeCallerSession) GetLightNodeDepositEnabled() (bool, error) {
	return _LightNode.Contract.GetLightNodeDepositEnabled(&_LightNode.CallOpts)
}

// GetLightNodePubkeyAt is a free data retrieval call binding the contract method 0xba7fe814.
//
// Solidity: function getLightNodePubkeyAt(address _nodeAddress, uint256 _index) view returns(bytes)
func (_LightNode *LightNodeCaller) GetLightNodePubkeyAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getLightNodePubkeyAt", _nodeAddress, _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLightNodePubkeyAt is a free data retrieval call binding the contract method 0xba7fe814.
//
// Solidity: function getLightNodePubkeyAt(address _nodeAddress, uint256 _index) view returns(bytes)
func (_LightNode *LightNodeSession) GetLightNodePubkeyAt(_nodeAddress common.Address, _index *big.Int) ([]byte, error) {
	return _LightNode.Contract.GetLightNodePubkeyAt(&_LightNode.CallOpts, _nodeAddress, _index)
}

// GetLightNodePubkeyAt is a free data retrieval call binding the contract method 0xba7fe814.
//
// Solidity: function getLightNodePubkeyAt(address _nodeAddress, uint256 _index) view returns(bytes)
func (_LightNode *LightNodeCallerSession) GetLightNodePubkeyAt(_nodeAddress common.Address, _index *big.Int) ([]byte, error) {
	return _LightNode.Contract.GetLightNodePubkeyAt(&_LightNode.CallOpts, _nodeAddress, _index)
}

// GetLightNodePubkeyCount is a free data retrieval call binding the contract method 0xccdf1d79.
//
// Solidity: function getLightNodePubkeyCount(address _nodeAddress) view returns(uint256)
func (_LightNode *LightNodeCaller) GetLightNodePubkeyCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getLightNodePubkeyCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLightNodePubkeyCount is a free data retrieval call binding the contract method 0xccdf1d79.
//
// Solidity: function getLightNodePubkeyCount(address _nodeAddress) view returns(uint256)
func (_LightNode *LightNodeSession) GetLightNodePubkeyCount(_nodeAddress common.Address) (*big.Int, error) {
	return _LightNode.Contract.GetLightNodePubkeyCount(&_LightNode.CallOpts, _nodeAddress)
}

// GetLightNodePubkeyCount is a free data retrieval call binding the contract method 0xccdf1d79.
//
// Solidity: function getLightNodePubkeyCount(address _nodeAddress) view returns(uint256)
func (_LightNode *LightNodeCallerSession) GetLightNodePubkeyCount(_nodeAddress common.Address) (*big.Int, error) {
	return _LightNode.Contract.GetLightNodePubkeyCount(&_LightNode.CallOpts, _nodeAddress)
}

// GetLightNodePubkeyStatus is a free data retrieval call binding the contract method 0xa37212cb.
//
// Solidity: function getLightNodePubkeyStatus(bytes _validatorPubkey) view returns(uint256)
func (_LightNode *LightNodeCaller) GetLightNodePubkeyStatus(opts *bind.CallOpts, _validatorPubkey []byte) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getLightNodePubkeyStatus", _validatorPubkey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLightNodePubkeyStatus is a free data retrieval call binding the contract method 0xa37212cb.
//
// Solidity: function getLightNodePubkeyStatus(bytes _validatorPubkey) view returns(uint256)
func (_LightNode *LightNodeSession) GetLightNodePubkeyStatus(_validatorPubkey []byte) (*big.Int, error) {
	return _LightNode.Contract.GetLightNodePubkeyStatus(&_LightNode.CallOpts, _validatorPubkey)
}

// GetLightNodePubkeyStatus is a free data retrieval call binding the contract method 0xa37212cb.
//
// Solidity: function getLightNodePubkeyStatus(bytes _validatorPubkey) view returns(uint256)
func (_LightNode *LightNodeCallerSession) GetLightNodePubkeyStatus(_validatorPubkey []byte) (*big.Int, error) {
	return _LightNode.Contract.GetLightNodePubkeyStatus(&_LightNode.CallOpts, _validatorPubkey)
}

// GetPubkeyVoted is a free data retrieval call binding the contract method 0x18582e28.
//
// Solidity: function getPubkeyVoted(bytes _validatorPubkey, address user) view returns(bool)
func (_LightNode *LightNodeCaller) GetPubkeyVoted(opts *bind.CallOpts, _validatorPubkey []byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getPubkeyVoted", _validatorPubkey, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPubkeyVoted is a free data retrieval call binding the contract method 0x18582e28.
//
// Solidity: function getPubkeyVoted(bytes _validatorPubkey, address user) view returns(bool)
func (_LightNode *LightNodeSession) GetPubkeyVoted(_validatorPubkey []byte, user common.Address) (bool, error) {
	return _LightNode.Contract.GetPubkeyVoted(&_LightNode.CallOpts, _validatorPubkey, user)
}

// GetPubkeyVoted is a free data retrieval call binding the contract method 0x18582e28.
//
// Solidity: function getPubkeyVoted(bytes _validatorPubkey, address user) view returns(bool)
func (_LightNode *LightNodeCallerSession) GetPubkeyVoted(_validatorPubkey []byte, user common.Address) (bool, error) {
	return _LightNode.Contract.GetPubkeyVoted(&_LightNode.CallOpts, _validatorPubkey, user)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_LightNode *LightNodeCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_LightNode *LightNodeSession) Version() (uint8, error) {
	return _LightNode.Contract.Version(&_LightNode.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_LightNode *LightNodeCallerSession) Version() (uint8, error) {
	return _LightNode.Contract.Version(&_LightNode.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_LightNode *LightNodeTransactor) Deposit(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "deposit", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_LightNode *LightNodeSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LightNode.Contract.Deposit(&_LightNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_LightNode *LightNodeTransactorSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LightNode.Contract.Deposit(&_LightNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_LightNode *LightNodeTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_LightNode *LightNodeSession) DepositEth() (*types.Transaction, error) {
	return _LightNode.Contract.DepositEth(&_LightNode.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_LightNode *LightNodeTransactorSession) DepositEth() (*types.Transaction, error) {
	return _LightNode.Contract.DepositEth(&_LightNode.TransactOpts)
}

// OffBoard is a paid mutator transaction binding the contract method 0x84d9eb7c.
//
// Solidity: function offBoard(bytes _validatorPubkey) returns()
func (_LightNode *LightNodeTransactor) OffBoard(opts *bind.TransactOpts, _validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "offBoard", _validatorPubkey)
}

// OffBoard is a paid mutator transaction binding the contract method 0x84d9eb7c.
//
// Solidity: function offBoard(bytes _validatorPubkey) returns()
func (_LightNode *LightNodeSession) OffBoard(_validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.Contract.OffBoard(&_LightNode.TransactOpts, _validatorPubkey)
}

// OffBoard is a paid mutator transaction binding the contract method 0x84d9eb7c.
//
// Solidity: function offBoard(bytes _validatorPubkey) returns()
func (_LightNode *LightNodeTransactorSession) OffBoard(_validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.Contract.OffBoard(&_LightNode.TransactOpts, _validatorPubkey)
}

// ProvideNodeDepositToken is a paid mutator transaction binding the contract method 0xeee1fa44.
//
// Solidity: function provideNodeDepositToken(bytes _validatorPubkey) payable returns()
func (_LightNode *LightNodeTransactor) ProvideNodeDepositToken(opts *bind.TransactOpts, _validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "provideNodeDepositToken", _validatorPubkey)
}

// ProvideNodeDepositToken is a paid mutator transaction binding the contract method 0xeee1fa44.
//
// Solidity: function provideNodeDepositToken(bytes _validatorPubkey) payable returns()
func (_LightNode *LightNodeSession) ProvideNodeDepositToken(_validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.Contract.ProvideNodeDepositToken(&_LightNode.TransactOpts, _validatorPubkey)
}

// ProvideNodeDepositToken is a paid mutator transaction binding the contract method 0xeee1fa44.
//
// Solidity: function provideNodeDepositToken(bytes _validatorPubkey) payable returns()
func (_LightNode *LightNodeTransactorSession) ProvideNodeDepositToken(_validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.Contract.ProvideNodeDepositToken(&_LightNode.TransactOpts, _validatorPubkey)
}

// ReceiveEtherWithdrawal is a paid mutator transaction binding the contract method 0x0a019eaf.
//
// Solidity: function receiveEtherWithdrawal() payable returns()
func (_LightNode *LightNodeTransactor) ReceiveEtherWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "receiveEtherWithdrawal")
}

// ReceiveEtherWithdrawal is a paid mutator transaction binding the contract method 0x0a019eaf.
//
// Solidity: function receiveEtherWithdrawal() payable returns()
func (_LightNode *LightNodeSession) ReceiveEtherWithdrawal() (*types.Transaction, error) {
	return _LightNode.Contract.ReceiveEtherWithdrawal(&_LightNode.TransactOpts)
}

// ReceiveEtherWithdrawal is a paid mutator transaction binding the contract method 0x0a019eaf.
//
// Solidity: function receiveEtherWithdrawal() payable returns()
func (_LightNode *LightNodeTransactorSession) ReceiveEtherWithdrawal() (*types.Transaction, error) {
	return _LightNode.Contract.ReceiveEtherWithdrawal(&_LightNode.TransactOpts)
}

// SetLightNodeDepositEnabled is a paid mutator transaction binding the contract method 0xadf1d8d6.
//
// Solidity: function setLightNodeDepositEnabled(bool _value) returns()
func (_LightNode *LightNodeTransactor) SetLightNodeDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "setLightNodeDepositEnabled", _value)
}

// SetLightNodeDepositEnabled is a paid mutator transaction binding the contract method 0xadf1d8d6.
//
// Solidity: function setLightNodeDepositEnabled(bool _value) returns()
func (_LightNode *LightNodeSession) SetLightNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _LightNode.Contract.SetLightNodeDepositEnabled(&_LightNode.TransactOpts, _value)
}

// SetLightNodeDepositEnabled is a paid mutator transaction binding the contract method 0xadf1d8d6.
//
// Solidity: function setLightNodeDepositEnabled(bool _value) returns()
func (_LightNode *LightNodeTransactorSession) SetLightNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _LightNode.Contract.SetLightNodeDepositEnabled(&_LightNode.TransactOpts, _value)
}

// SetLightNodePubkeyStatus is a paid mutator transaction binding the contract method 0x9630f63c.
//
// Solidity: function setLightNodePubkeyStatus(bytes _validatorPubkey, uint256 _status) returns()
func (_LightNode *LightNodeTransactor) SetLightNodePubkeyStatus(opts *bind.TransactOpts, _validatorPubkey []byte, _status *big.Int) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "setLightNodePubkeyStatus", _validatorPubkey, _status)
}

// SetLightNodePubkeyStatus is a paid mutator transaction binding the contract method 0x9630f63c.
//
// Solidity: function setLightNodePubkeyStatus(bytes _validatorPubkey, uint256 _status) returns()
func (_LightNode *LightNodeSession) SetLightNodePubkeyStatus(_validatorPubkey []byte, _status *big.Int) (*types.Transaction, error) {
	return _LightNode.Contract.SetLightNodePubkeyStatus(&_LightNode.TransactOpts, _validatorPubkey, _status)
}

// SetLightNodePubkeyStatus is a paid mutator transaction binding the contract method 0x9630f63c.
//
// Solidity: function setLightNodePubkeyStatus(bytes _validatorPubkey, uint256 _status) returns()
func (_LightNode *LightNodeTransactorSession) SetLightNodePubkeyStatus(_validatorPubkey []byte, _status *big.Int) (*types.Transaction, error) {
	return _LightNode.Contract.SetLightNodePubkeyStatus(&_LightNode.TransactOpts, _validatorPubkey, _status)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_LightNode *LightNodeTransactor) Stake(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "stake", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_LightNode *LightNodeSession) Stake(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LightNode.Contract.Stake(&_LightNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_LightNode *LightNodeTransactorSession) Stake(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _LightNode.Contract.Stake(&_LightNode.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_LightNode *LightNodeTransactor) VoteWithdrawCredentials(opts *bind.TransactOpts, _pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "voteWithdrawCredentials", _pubkeys, _matchs)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_LightNode *LightNodeSession) VoteWithdrawCredentials(_pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _LightNode.Contract.VoteWithdrawCredentials(&_LightNode.TransactOpts, _pubkeys, _matchs)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_LightNode *LightNodeTransactorSession) VoteWithdrawCredentials(_pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _LightNode.Contract.VoteWithdrawCredentials(&_LightNode.TransactOpts, _pubkeys, _matchs)
}

// WithdrawNodeDepositToken is a paid mutator transaction binding the contract method 0x3c4bc635.
//
// Solidity: function withdrawNodeDepositToken(bytes _validatorPubkey) returns()
func (_LightNode *LightNodeTransactor) WithdrawNodeDepositToken(opts *bind.TransactOpts, _validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "withdrawNodeDepositToken", _validatorPubkey)
}

// WithdrawNodeDepositToken is a paid mutator transaction binding the contract method 0x3c4bc635.
//
// Solidity: function withdrawNodeDepositToken(bytes _validatorPubkey) returns()
func (_LightNode *LightNodeSession) WithdrawNodeDepositToken(_validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.Contract.WithdrawNodeDepositToken(&_LightNode.TransactOpts, _validatorPubkey)
}

// WithdrawNodeDepositToken is a paid mutator transaction binding the contract method 0x3c4bc635.
//
// Solidity: function withdrawNodeDepositToken(bytes _validatorPubkey) returns()
func (_LightNode *LightNodeTransactorSession) WithdrawNodeDepositToken(_validatorPubkey []byte) (*types.Transaction, error) {
	return _LightNode.Contract.WithdrawNodeDepositToken(&_LightNode.TransactOpts, _validatorPubkey)
}

// LightNodeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the LightNode contract.
type LightNodeDepositedIterator struct {
	Event *LightNodeDeposited // Event containing the contract specifics and raw log

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
func (it *LightNodeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeDeposited)
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
		it.Event = new(LightNodeDeposited)
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
func (it *LightNodeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeDeposited represents a Deposited event raised by the LightNode contract.
type LightNodeDeposited struct {
	Node               common.Address
	Pubkey             []byte
	ValidatorSignature []byte
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1fc33daf7f37b2e02f49dbee7b9707bddd7c8cef1f08deff23bff3d105e5aa72.
//
// Solidity: event Deposited(address node, bytes pubkey, bytes validatorSignature, uint256 amount)
func (_LightNode *LightNodeFilterer) FilterDeposited(opts *bind.FilterOpts) (*LightNodeDepositedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &LightNodeDepositedIterator{contract: _LightNode.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1fc33daf7f37b2e02f49dbee7b9707bddd7c8cef1f08deff23bff3d105e5aa72.
//
// Solidity: event Deposited(address node, bytes pubkey, bytes validatorSignature, uint256 amount)
func (_LightNode *LightNodeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *LightNodeDeposited) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeDeposited)
				if err := _LightNode.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1fc33daf7f37b2e02f49dbee7b9707bddd7c8cef1f08deff23bff3d105e5aa72.
//
// Solidity: event Deposited(address node, bytes pubkey, bytes validatorSignature, uint256 amount)
func (_LightNode *LightNodeFilterer) ParseDeposited(log types.Log) (*LightNodeDeposited, error) {
	event := new(LightNodeDeposited)
	if err := _LightNode.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the LightNode contract.
type LightNodeEtherDepositedIterator struct {
	Event *LightNodeEtherDeposited // Event containing the contract specifics and raw log

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
func (it *LightNodeEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeEtherDeposited)
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
		it.Event = new(LightNodeEtherDeposited)
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
func (it *LightNodeEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeEtherDeposited represents a EtherDeposited event raised by the LightNode contract.
type LightNodeEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_LightNode *LightNodeFilterer) FilterEtherDeposited(opts *bind.FilterOpts, from []common.Address) (*LightNodeEtherDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &LightNodeEtherDepositedIterator{contract: _LightNode.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_LightNode *LightNodeFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *LightNodeEtherDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeEtherDeposited)
				if err := _LightNode.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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
func (_LightNode *LightNodeFilterer) ParseEtherDeposited(log types.Log) (*LightNodeEtherDeposited, error) {
	event := new(LightNodeEtherDeposited)
	if err := _LightNode.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeOffBoardedIterator is returned from FilterOffBoarded and is used to iterate over the raw logs and unpacked data for OffBoarded events raised by the LightNode contract.
type LightNodeOffBoardedIterator struct {
	Event *LightNodeOffBoarded // Event containing the contract specifics and raw log

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
func (it *LightNodeOffBoardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeOffBoarded)
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
		it.Event = new(LightNodeOffBoarded)
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
func (it *LightNodeOffBoardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeOffBoardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeOffBoarded represents a OffBoarded event raised by the LightNode contract.
type LightNodeOffBoarded struct {
	Node   common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOffBoarded is a free log retrieval operation binding the contract event 0x9d5023d85497e8c264e3b053f8da9415f4db76eb5af3ecef3fe953fe9f981470.
//
// Solidity: event OffBoarded(address node, bytes pubkey)
func (_LightNode *LightNodeFilterer) FilterOffBoarded(opts *bind.FilterOpts) (*LightNodeOffBoardedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "OffBoarded")
	if err != nil {
		return nil, err
	}
	return &LightNodeOffBoardedIterator{contract: _LightNode.contract, event: "OffBoarded", logs: logs, sub: sub}, nil
}

// WatchOffBoarded is a free log subscription operation binding the contract event 0x9d5023d85497e8c264e3b053f8da9415f4db76eb5af3ecef3fe953fe9f981470.
//
// Solidity: event OffBoarded(address node, bytes pubkey)
func (_LightNode *LightNodeFilterer) WatchOffBoarded(opts *bind.WatchOpts, sink chan<- *LightNodeOffBoarded) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "OffBoarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeOffBoarded)
				if err := _LightNode.contract.UnpackLog(event, "OffBoarded", log); err != nil {
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

// ParseOffBoarded is a log parse operation binding the contract event 0x9d5023d85497e8c264e3b053f8da9415f4db76eb5af3ecef3fe953fe9f981470.
//
// Solidity: event OffBoarded(address node, bytes pubkey)
func (_LightNode *LightNodeFilterer) ParseOffBoarded(log types.Log) (*LightNodeOffBoarded, error) {
	event := new(LightNodeOffBoarded)
	if err := _LightNode.contract.UnpackLog(event, "OffBoarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeSetPubkeyStatusIterator is returned from FilterSetPubkeyStatus and is used to iterate over the raw logs and unpacked data for SetPubkeyStatus events raised by the LightNode contract.
type LightNodeSetPubkeyStatusIterator struct {
	Event *LightNodeSetPubkeyStatus // Event containing the contract specifics and raw log

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
func (it *LightNodeSetPubkeyStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeSetPubkeyStatus)
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
		it.Event = new(LightNodeSetPubkeyStatus)
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
func (it *LightNodeSetPubkeyStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeSetPubkeyStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeSetPubkeyStatus represents a SetPubkeyStatus event raised by the LightNode contract.
type LightNodeSetPubkeyStatus struct {
	Pubkey []byte
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPubkeyStatus is a free log retrieval operation binding the contract event 0xc581330a67ddadcf62d06e14db41c85bb45910369cedad649390e6dd2a9373c8.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint256 status)
func (_LightNode *LightNodeFilterer) FilterSetPubkeyStatus(opts *bind.FilterOpts) (*LightNodeSetPubkeyStatusIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "SetPubkeyStatus")
	if err != nil {
		return nil, err
	}
	return &LightNodeSetPubkeyStatusIterator{contract: _LightNode.contract, event: "SetPubkeyStatus", logs: logs, sub: sub}, nil
}

// WatchSetPubkeyStatus is a free log subscription operation binding the contract event 0xc581330a67ddadcf62d06e14db41c85bb45910369cedad649390e6dd2a9373c8.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint256 status)
func (_LightNode *LightNodeFilterer) WatchSetPubkeyStatus(opts *bind.WatchOpts, sink chan<- *LightNodeSetPubkeyStatus) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "SetPubkeyStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeSetPubkeyStatus)
				if err := _LightNode.contract.UnpackLog(event, "SetPubkeyStatus", log); err != nil {
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
func (_LightNode *LightNodeFilterer) ParseSetPubkeyStatus(log types.Log) (*LightNodeSetPubkeyStatus, error) {
	event := new(LightNodeSetPubkeyStatus)
	if err := _LightNode.contract.UnpackLog(event, "SetPubkeyStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the LightNode contract.
type LightNodeStakedIterator struct {
	Event *LightNodeStaked // Event containing the contract specifics and raw log

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
func (it *LightNodeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeStaked)
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
		it.Event = new(LightNodeStaked)
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
func (it *LightNodeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeStaked represents a Staked event raised by the LightNode contract.
type LightNodeStaked struct {
	Node   common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_LightNode *LightNodeFilterer) FilterStaked(opts *bind.FilterOpts) (*LightNodeStakedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &LightNodeStakedIterator{contract: _LightNode.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_LightNode *LightNodeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LightNodeStaked) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeStaked)
				if err := _LightNode.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_LightNode *LightNodeFilterer) ParseStaked(log types.Log) (*LightNodeStaked, error) {
	event := new(LightNodeStaked)
	if err := _LightNode.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
