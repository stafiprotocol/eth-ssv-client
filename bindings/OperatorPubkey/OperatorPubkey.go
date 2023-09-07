// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operator_pubkey

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

// OperatorPubkeyMetaData contains all meta data concerning the OperatorPubkey contract.
var OperatorPubkeyMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"method\",\"type\":\"function\",\"outputs\":[{\"type\":\"bytes\"}]}]",
}

// OperatorPubkeyABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorPubkeyMetaData.ABI instead.
var OperatorPubkeyABI = OperatorPubkeyMetaData.ABI

// OperatorPubkey is an auto generated Go binding around an Ethereum contract.
type OperatorPubkey struct {
	OperatorPubkeyCaller     // Read-only binding to the contract
	OperatorPubkeyTransactor // Write-only binding to the contract
	OperatorPubkeyFilterer   // Log filterer for contract events
}

// OperatorPubkeyCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorPubkeyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorPubkeyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorPubkeyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorPubkeyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorPubkeyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorPubkeySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorPubkeySession struct {
	Contract     *OperatorPubkey   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorPubkeyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorPubkeyCallerSession struct {
	Contract *OperatorPubkeyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OperatorPubkeyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorPubkeyTransactorSession struct {
	Contract     *OperatorPubkeyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OperatorPubkeyRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorPubkeyRaw struct {
	Contract *OperatorPubkey // Generic contract binding to access the raw methods on
}

// OperatorPubkeyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorPubkeyCallerRaw struct {
	Contract *OperatorPubkeyCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorPubkeyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorPubkeyTransactorRaw struct {
	Contract *OperatorPubkeyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorPubkey creates a new instance of OperatorPubkey, bound to a specific deployed contract.
func NewOperatorPubkey(address common.Address, backend bind.ContractBackend) (*OperatorPubkey, error) {
	contract, err := bindOperatorPubkey(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorPubkey{OperatorPubkeyCaller: OperatorPubkeyCaller{contract: contract}, OperatorPubkeyTransactor: OperatorPubkeyTransactor{contract: contract}, OperatorPubkeyFilterer: OperatorPubkeyFilterer{contract: contract}}, nil
}

// NewOperatorPubkeyCaller creates a new read-only instance of OperatorPubkey, bound to a specific deployed contract.
func NewOperatorPubkeyCaller(address common.Address, caller bind.ContractCaller) (*OperatorPubkeyCaller, error) {
	contract, err := bindOperatorPubkey(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorPubkeyCaller{contract: contract}, nil
}

// NewOperatorPubkeyTransactor creates a new write-only instance of OperatorPubkey, bound to a specific deployed contract.
func NewOperatorPubkeyTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorPubkeyTransactor, error) {
	contract, err := bindOperatorPubkey(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorPubkeyTransactor{contract: contract}, nil
}

// NewOperatorPubkeyFilterer creates a new log filterer instance of OperatorPubkey, bound to a specific deployed contract.
func NewOperatorPubkeyFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorPubkeyFilterer, error) {
	contract, err := bindOperatorPubkey(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorPubkeyFilterer{contract: contract}, nil
}

// bindOperatorPubkey binds a generic wrapper to an already deployed contract.
func bindOperatorPubkey(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperatorPubkeyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorPubkey *OperatorPubkeyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorPubkey.Contract.OperatorPubkeyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorPubkey *OperatorPubkeyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorPubkey.Contract.OperatorPubkeyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorPubkey *OperatorPubkeyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorPubkey.Contract.OperatorPubkeyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorPubkey *OperatorPubkeyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorPubkey.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorPubkey *OperatorPubkeyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorPubkey.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorPubkey *OperatorPubkeyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorPubkey.Contract.contract.Transact(opts, method, params...)
}

// Method is a paid mutator transaction binding the contract method 0x2c383a9f.
//
// Solidity: function method() returns(bytes)
func (_OperatorPubkey *OperatorPubkeyTransactor) Method(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorPubkey.contract.Transact(opts, "method")
}

// Method is a paid mutator transaction binding the contract method 0x2c383a9f.
//
// Solidity: function method() returns(bytes)
func (_OperatorPubkey *OperatorPubkeySession) Method() (*types.Transaction, error) {
	return _OperatorPubkey.Contract.Method(&_OperatorPubkey.TransactOpts)
}

// Method is a paid mutator transaction binding the contract method 0x2c383a9f.
//
// Solidity: function method() returns(bytes)
func (_OperatorPubkey *OperatorPubkeyTransactorSession) Method() (*types.Transaction, error) {
	return _OperatorPubkey.Contract.Method(&_OperatorPubkey.TransactOpts)
}
