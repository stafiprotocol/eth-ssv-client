// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stafi_ether

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

// StafiEtherMetaData contains all meta data concerning the StafiEther contract.
var StafiEtherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"by\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"by\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StafiEtherABI is the input ABI used to generate the binding from.
// Deprecated: Use StafiEtherMetaData.ABI instead.
var StafiEtherABI = StafiEtherMetaData.ABI

// StafiEther is an auto generated Go binding around an Ethereum contract.
type StafiEther struct {
	StafiEtherCaller     // Read-only binding to the contract
	StafiEtherTransactor // Write-only binding to the contract
	StafiEtherFilterer   // Log filterer for contract events
}

// StafiEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type StafiEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StafiEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StafiEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StafiEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StafiEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StafiEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StafiEtherSession struct {
	Contract     *StafiEther       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StafiEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StafiEtherCallerSession struct {
	Contract *StafiEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StafiEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StafiEtherTransactorSession struct {
	Contract     *StafiEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StafiEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type StafiEtherRaw struct {
	Contract *StafiEther // Generic contract binding to access the raw methods on
}

// StafiEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StafiEtherCallerRaw struct {
	Contract *StafiEtherCaller // Generic read-only contract binding to access the raw methods on
}

// StafiEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StafiEtherTransactorRaw struct {
	Contract *StafiEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStafiEther creates a new instance of StafiEther, bound to a specific deployed contract.
func NewStafiEther(address common.Address, backend bind.ContractBackend) (*StafiEther, error) {
	contract, err := bindStafiEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StafiEther{StafiEtherCaller: StafiEtherCaller{contract: contract}, StafiEtherTransactor: StafiEtherTransactor{contract: contract}, StafiEtherFilterer: StafiEtherFilterer{contract: contract}}, nil
}

// NewStafiEtherCaller creates a new read-only instance of StafiEther, bound to a specific deployed contract.
func NewStafiEtherCaller(address common.Address, caller bind.ContractCaller) (*StafiEtherCaller, error) {
	contract, err := bindStafiEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StafiEtherCaller{contract: contract}, nil
}

// NewStafiEtherTransactor creates a new write-only instance of StafiEther, bound to a specific deployed contract.
func NewStafiEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*StafiEtherTransactor, error) {
	contract, err := bindStafiEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StafiEtherTransactor{contract: contract}, nil
}

// NewStafiEtherFilterer creates a new log filterer instance of StafiEther, bound to a specific deployed contract.
func NewStafiEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*StafiEtherFilterer, error) {
	contract, err := bindStafiEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StafiEtherFilterer{contract: contract}, nil
}

// bindStafiEther binds a generic wrapper to an already deployed contract.
func bindStafiEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StafiEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StafiEther *StafiEtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StafiEther.Contract.StafiEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StafiEther *StafiEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StafiEther.Contract.StafiEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StafiEther *StafiEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StafiEther.Contract.StafiEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StafiEther *StafiEtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StafiEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StafiEther *StafiEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StafiEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StafiEther *StafiEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StafiEther.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _contractAddress) view returns(uint256)
func (_StafiEther *StafiEtherCaller) BalanceOf(opts *bind.CallOpts, _contractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StafiEther.contract.Call(opts, &out, "balanceOf", _contractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _contractAddress) view returns(uint256)
func (_StafiEther *StafiEtherSession) BalanceOf(_contractAddress common.Address) (*big.Int, error) {
	return _StafiEther.Contract.BalanceOf(&_StafiEther.CallOpts, _contractAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _contractAddress) view returns(uint256)
func (_StafiEther *StafiEtherCallerSession) BalanceOf(_contractAddress common.Address) (*big.Int, error) {
	return _StafiEther.Contract.BalanceOf(&_StafiEther.CallOpts, _contractAddress)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_StafiEther *StafiEtherCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StafiEther.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_StafiEther *StafiEtherSession) Version() (uint8, error) {
	return _StafiEther.Contract.Version(&_StafiEther.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_StafiEther *StafiEtherCallerSession) Version() (uint8, error) {
	return _StafiEther.Contract.Version(&_StafiEther.CallOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_StafiEther *StafiEtherTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StafiEther.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_StafiEther *StafiEtherSession) DepositEther() (*types.Transaction, error) {
	return _StafiEther.Contract.DepositEther(&_StafiEther.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_StafiEther *StafiEtherTransactorSession) DepositEther() (*types.Transaction, error) {
	return _StafiEther.Contract.DepositEther(&_StafiEther.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_StafiEther *StafiEtherTransactor) WithdrawEther(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StafiEther.contract.Transact(opts, "withdrawEther", _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_StafiEther *StafiEtherSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _StafiEther.Contract.WithdrawEther(&_StafiEther.TransactOpts, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_StafiEther *StafiEtherTransactorSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _StafiEther.Contract.WithdrawEther(&_StafiEther.TransactOpts, _amount)
}

// StafiEtherEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the StafiEther contract.
type StafiEtherEtherDepositedIterator struct {
	Event *StafiEtherEtherDeposited // Event containing the contract specifics and raw log

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
func (it *StafiEtherEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StafiEtherEtherDeposited)
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
		it.Event = new(StafiEtherEtherDeposited)
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
func (it *StafiEtherEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StafiEtherEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StafiEtherEtherDeposited represents a EtherDeposited event raised by the StafiEther contract.
type StafiEtherEtherDeposited struct {
	By     [32]byte
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0x2c7d80ba9bc6395644b4ff4a878353ac20adeed6e23cead48c8cec7a58b6e719.
//
// Solidity: event EtherDeposited(bytes32 indexed by, uint256 amount, uint256 time)
func (_StafiEther *StafiEtherFilterer) FilterEtherDeposited(opts *bind.FilterOpts, by [][32]byte) (*StafiEtherEtherDepositedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _StafiEther.contract.FilterLogs(opts, "EtherDeposited", byRule)
	if err != nil {
		return nil, err
	}
	return &StafiEtherEtherDepositedIterator{contract: _StafiEther.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0x2c7d80ba9bc6395644b4ff4a878353ac20adeed6e23cead48c8cec7a58b6e719.
//
// Solidity: event EtherDeposited(bytes32 indexed by, uint256 amount, uint256 time)
func (_StafiEther *StafiEtherFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *StafiEtherEtherDeposited, by [][32]byte) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _StafiEther.contract.WatchLogs(opts, "EtherDeposited", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StafiEtherEtherDeposited)
				if err := _StafiEther.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0x2c7d80ba9bc6395644b4ff4a878353ac20adeed6e23cead48c8cec7a58b6e719.
//
// Solidity: event EtherDeposited(bytes32 indexed by, uint256 amount, uint256 time)
func (_StafiEther *StafiEtherFilterer) ParseEtherDeposited(log types.Log) (*StafiEtherEtherDeposited, error) {
	event := new(StafiEtherEtherDeposited)
	if err := _StafiEther.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StafiEtherEtherWithdrawnIterator is returned from FilterEtherWithdrawn and is used to iterate over the raw logs and unpacked data for EtherWithdrawn events raised by the StafiEther contract.
type StafiEtherEtherWithdrawnIterator struct {
	Event *StafiEtherEtherWithdrawn // Event containing the contract specifics and raw log

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
func (it *StafiEtherEtherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StafiEtherEtherWithdrawn)
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
		it.Event = new(StafiEtherEtherWithdrawn)
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
func (it *StafiEtherEtherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StafiEtherEtherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StafiEtherEtherWithdrawn represents a EtherWithdrawn event raised by the StafiEther contract.
type StafiEtherEtherWithdrawn struct {
	By     [32]byte
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawn is a free log retrieval operation binding the contract event 0x73bee9f217b293842a2fbe73ad32a4c1175e05bb940c97d6e7b8da25b6865828.
//
// Solidity: event EtherWithdrawn(bytes32 indexed by, uint256 amount, uint256 time)
func (_StafiEther *StafiEtherFilterer) FilterEtherWithdrawn(opts *bind.FilterOpts, by [][32]byte) (*StafiEtherEtherWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _StafiEther.contract.FilterLogs(opts, "EtherWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return &StafiEtherEtherWithdrawnIterator{contract: _StafiEther.contract, event: "EtherWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawn is a free log subscription operation binding the contract event 0x73bee9f217b293842a2fbe73ad32a4c1175e05bb940c97d6e7b8da25b6865828.
//
// Solidity: event EtherWithdrawn(bytes32 indexed by, uint256 amount, uint256 time)
func (_StafiEther *StafiEtherFilterer) WatchEtherWithdrawn(opts *bind.WatchOpts, sink chan<- *StafiEtherEtherWithdrawn, by [][32]byte) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _StafiEther.contract.WatchLogs(opts, "EtherWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StafiEtherEtherWithdrawn)
				if err := _StafiEther.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
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

// ParseEtherWithdrawn is a log parse operation binding the contract event 0x73bee9f217b293842a2fbe73ad32a4c1175e05bb940c97d6e7b8da25b6865828.
//
// Solidity: event EtherWithdrawn(bytes32 indexed by, uint256 amount, uint256 time)
func (_StafiEther *StafiEtherFilterer) ParseEtherWithdrawn(log types.Log) (*StafiEtherEtherWithdrawn, error) {
	event := new(StafiEtherEtherWithdrawn)
	if err := _StafiEther.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
