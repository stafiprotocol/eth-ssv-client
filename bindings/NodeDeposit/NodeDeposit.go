// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node_deposit

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

// NodeDepositMetaData contains all meta data concerning the NodeDeposit contract.
var NodeDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"DepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentNodeDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setCurrentNodeDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakingPools\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NodeDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeDepositMetaData.ABI instead.
var NodeDepositABI = NodeDepositMetaData.ABI

// NodeDeposit is an auto generated Go binding around an Ethereum contract.
type NodeDeposit struct {
	NodeDepositCaller     // Read-only binding to the contract
	NodeDepositTransactor // Write-only binding to the contract
	NodeDepositFilterer   // Log filterer for contract events
}

// NodeDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeDepositSession struct {
	Contract     *NodeDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeDepositCallerSession struct {
	Contract *NodeDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeDepositTransactorSession struct {
	Contract     *NodeDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeDepositRaw struct {
	Contract *NodeDeposit // Generic contract binding to access the raw methods on
}

// NodeDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeDepositCallerRaw struct {
	Contract *NodeDepositCaller // Generic read-only contract binding to access the raw methods on
}

// NodeDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeDepositTransactorRaw struct {
	Contract *NodeDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeDeposit creates a new instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDeposit(address common.Address, backend bind.ContractBackend) (*NodeDeposit, error) {
	contract, err := bindNodeDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeDeposit{NodeDepositCaller: NodeDepositCaller{contract: contract}, NodeDepositTransactor: NodeDepositTransactor{contract: contract}, NodeDepositFilterer: NodeDepositFilterer{contract: contract}}, nil
}

// NewNodeDepositCaller creates a new read-only instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDepositCaller(address common.Address, caller bind.ContractCaller) (*NodeDepositCaller, error) {
	contract, err := bindNodeDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeDepositCaller{contract: contract}, nil
}

// NewNodeDepositTransactor creates a new write-only instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeDepositTransactor, error) {
	contract, err := bindNodeDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeDepositTransactor{contract: contract}, nil
}

// NewNodeDepositFilterer creates a new log filterer instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeDepositFilterer, error) {
	contract, err := bindNodeDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeDepositFilterer{contract: contract}, nil
}

// bindNodeDeposit binds a generic wrapper to an already deployed contract.
func bindNodeDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeDeposit *NodeDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeDeposit.Contract.NodeDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeDeposit *NodeDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeDeposit.Contract.NodeDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeDeposit *NodeDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeDeposit.Contract.NodeDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeDeposit *NodeDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeDeposit *NodeDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeDeposit *NodeDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeDeposit.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositCaller) GetCurrentNodeDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getCurrentNodeDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositSession) GetCurrentNodeDepositAmount() (*big.Int, error) {
	return _NodeDeposit.Contract.GetCurrentNodeDepositAmount(&_NodeDeposit.CallOpts)
}

// GetCurrentNodeDepositAmount is a free data retrieval call binding the contract method 0x83f3f086.
//
// Solidity: function getCurrentNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositCallerSession) GetCurrentNodeDepositAmount() (*big.Int, error) {
	return _NodeDeposit.Contract.GetCurrentNodeDepositAmount(&_NodeDeposit.CallOpts)
}

// GetDepositEnabled is a free data retrieval call binding the contract method 0x6ada7847.
//
// Solidity: function getDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCaller) GetDepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getDepositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetDepositEnabled is a free data retrieval call binding the contract method 0x6ada7847.
//
// Solidity: function getDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositSession) GetDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.GetDepositEnabled(&_NodeDeposit.CallOpts)
}

// GetDepositEnabled is a free data retrieval call binding the contract method 0x6ada7847.
//
// Solidity: function getDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCallerSession) GetDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.GetDepositEnabled(&_NodeDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeDeposit *NodeDepositCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeDeposit *NodeDepositSession) Version() (uint8, error) {
	return _NodeDeposit.Contract.Version(&_NodeDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeDeposit *NodeDepositCallerSession) Version() (uint8, error) {
	return _NodeDeposit.Contract.Version(&_NodeDeposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_NodeDeposit *NodeDepositTransactor) Deposit(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "deposit", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_NodeDeposit *NodeDepositSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Deposit(&_NodeDeposit.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_NodeDeposit *NodeDepositTransactorSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Deposit(&_NodeDeposit.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// SetCurrentNodeDepositAmount is a paid mutator transaction binding the contract method 0x33f30a5e.
//
// Solidity: function setCurrentNodeDepositAmount(uint256 _value) returns()
func (_NodeDeposit *NodeDepositTransactor) SetCurrentNodeDepositAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setCurrentNodeDepositAmount", _value)
}

// SetCurrentNodeDepositAmount is a paid mutator transaction binding the contract method 0x33f30a5e.
//
// Solidity: function setCurrentNodeDepositAmount(uint256 _value) returns()
func (_NodeDeposit *NodeDepositSession) SetCurrentNodeDepositAmount(_value *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetCurrentNodeDepositAmount(&_NodeDeposit.TransactOpts, _value)
}

// SetCurrentNodeDepositAmount is a paid mutator transaction binding the contract method 0x33f30a5e.
//
// Solidity: function setCurrentNodeDepositAmount(uint256 _value) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetCurrentNodeDepositAmount(_value *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetCurrentNodeDepositAmount(&_NodeDeposit.TransactOpts, _value)
}

// SetDepositEnabled is a paid mutator transaction binding the contract method 0x5b17d04b.
//
// Solidity: function setDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactor) SetDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setDepositEnabled", _value)
}

// SetDepositEnabled is a paid mutator transaction binding the contract method 0x5b17d04b.
//
// Solidity: function setDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositSession) SetDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// SetDepositEnabled is a paid mutator transaction binding the contract method 0x5b17d04b.
//
// Solidity: function setDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// Stake is a paid mutator transaction binding the contract method 0xa481513a.
//
// Solidity: function stake(address[] _stakingPools, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_NodeDeposit *NodeDepositTransactor) Stake(opts *bind.TransactOpts, _stakingPools []common.Address, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "stake", _stakingPools, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xa481513a.
//
// Solidity: function stake(address[] _stakingPools, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_NodeDeposit *NodeDepositSession) Stake(_stakingPools []common.Address, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Stake(&_NodeDeposit.TransactOpts, _stakingPools, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xa481513a.
//
// Solidity: function stake(address[] _stakingPools, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_NodeDeposit *NodeDepositTransactorSession) Stake(_stakingPools []common.Address, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Stake(&_NodeDeposit.TransactOpts, _stakingPools, _validatorSignatures, _depositDataRoots)
}

// NodeDepositDepositReceivedIterator is returned from FilterDepositReceived and is used to iterate over the raw logs and unpacked data for DepositReceived events raised by the NodeDeposit contract.
type NodeDepositDepositReceivedIterator struct {
	Event *NodeDepositDepositReceived // Event containing the contract specifics and raw log

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
func (it *NodeDepositDepositReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositDepositReceived)
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
		it.Event = new(NodeDepositDepositReceived)
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
func (it *NodeDepositDepositReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositDepositReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositDepositReceived represents a DepositReceived event raised by the NodeDeposit contract.
type NodeDepositDepositReceived struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositReceived is a free log retrieval operation binding the contract event 0x7aa1a8eb998c779420645fc14513bf058edb347d95c2fc2e6845bdc22f888631.
//
// Solidity: event DepositReceived(address indexed from, uint256 amount, uint256 time)
func (_NodeDeposit *NodeDepositFilterer) FilterDepositReceived(opts *bind.FilterOpts, from []common.Address) (*NodeDepositDepositReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "DepositReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &NodeDepositDepositReceivedIterator{contract: _NodeDeposit.contract, event: "DepositReceived", logs: logs, sub: sub}, nil
}

// WatchDepositReceived is a free log subscription operation binding the contract event 0x7aa1a8eb998c779420645fc14513bf058edb347d95c2fc2e6845bdc22f888631.
//
// Solidity: event DepositReceived(address indexed from, uint256 amount, uint256 time)
func (_NodeDeposit *NodeDepositFilterer) WatchDepositReceived(opts *bind.WatchOpts, sink chan<- *NodeDepositDepositReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "DepositReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositDepositReceived)
				if err := _NodeDeposit.contract.UnpackLog(event, "DepositReceived", log); err != nil {
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

// ParseDepositReceived is a log parse operation binding the contract event 0x7aa1a8eb998c779420645fc14513bf058edb347d95c2fc2e6845bdc22f888631.
//
// Solidity: event DepositReceived(address indexed from, uint256 amount, uint256 time)
func (_NodeDeposit *NodeDepositFilterer) ParseDepositReceived(log types.Log) (*NodeDepositDepositReceived, error) {
	event := new(NodeDepositDepositReceived)
	if err := _NodeDeposit.contract.UnpackLog(event, "DepositReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the NodeDeposit contract.
type NodeDepositDepositedIterator struct {
	Event *NodeDepositDeposited // Event containing the contract specifics and raw log

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
func (it *NodeDepositDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositDeposited)
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
		it.Event = new(NodeDepositDeposited)
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
func (it *NodeDepositDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositDeposited represents a Deposited event raised by the NodeDeposit contract.
type NodeDepositDeposited struct {
	Node               common.Address
	Pool               common.Address
	ValidatorPubkey    []byte
	ValidatorSignature []byte
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2c4863664bc90d75b80d943cf7e8cc0974d20e97b949374d963637af4b80dce6.
//
// Solidity: event Deposited(address node, address pool, bytes validatorPubkey, bytes validatorSignature, uint256 amount)
func (_NodeDeposit *NodeDepositFilterer) FilterDeposited(opts *bind.FilterOpts) (*NodeDepositDepositedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &NodeDepositDepositedIterator{contract: _NodeDeposit.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2c4863664bc90d75b80d943cf7e8cc0974d20e97b949374d963637af4b80dce6.
//
// Solidity: event Deposited(address node, address pool, bytes validatorPubkey, bytes validatorSignature, uint256 amount)
func (_NodeDeposit *NodeDepositFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *NodeDepositDeposited) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositDeposited)
				if err := _NodeDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2c4863664bc90d75b80d943cf7e8cc0974d20e97b949374d963637af4b80dce6.
//
// Solidity: event Deposited(address node, address pool, bytes validatorPubkey, bytes validatorSignature, uint256 amount)
func (_NodeDeposit *NodeDepositFilterer) ParseDeposited(log types.Log) (*NodeDepositDeposited, error) {
	event := new(NodeDepositDeposited)
	if err := _NodeDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the NodeDeposit contract.
type NodeDepositStakedIterator struct {
	Event *NodeDepositStaked // Event containing the contract specifics and raw log

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
func (it *NodeDepositStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositStaked)
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
		it.Event = new(NodeDepositStaked)
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
func (it *NodeDepositStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositStaked represents a Staked event raised by the NodeDeposit contract.
type NodeDepositStaked struct {
	Node            common.Address
	ValidatorPubkey []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes validatorPubkey)
func (_NodeDeposit *NodeDepositFilterer) FilterStaked(opts *bind.FilterOpts) (*NodeDepositStakedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &NodeDepositStakedIterator{contract: _NodeDeposit.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes validatorPubkey)
func (_NodeDeposit *NodeDepositFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *NodeDepositStaked) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositStaked)
				if err := _NodeDeposit.contract.UnpackLog(event, "Staked", log); err != nil {
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
// Solidity: event Staked(address node, bytes validatorPubkey)
func (_NodeDeposit *NodeDepositFilterer) ParseStaked(log types.Log) (*NodeDepositStaked, error) {
	event := new(NodeDepositStaked)
	if err := _NodeDeposit.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
