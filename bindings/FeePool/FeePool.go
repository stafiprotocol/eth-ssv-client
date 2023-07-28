// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_pool

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

// FeePoolMetaData contains all meta data concerning the FeePool contract.
var FeePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"by\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FeePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use FeePoolMetaData.ABI instead.
var FeePoolABI = FeePoolMetaData.ABI

// FeePool is an auto generated Go binding around an Ethereum contract.
type FeePool struct {
	FeePoolCaller     // Read-only binding to the contract
	FeePoolTransactor // Write-only binding to the contract
	FeePoolFilterer   // Log filterer for contract events
}

// FeePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeePoolSession struct {
	Contract     *FeePool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeePoolCallerSession struct {
	Contract *FeePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FeePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeePoolTransactorSession struct {
	Contract     *FeePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FeePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeePoolRaw struct {
	Contract *FeePool // Generic contract binding to access the raw methods on
}

// FeePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeePoolCallerRaw struct {
	Contract *FeePoolCaller // Generic read-only contract binding to access the raw methods on
}

// FeePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeePoolTransactorRaw struct {
	Contract *FeePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeePool creates a new instance of FeePool, bound to a specific deployed contract.
func NewFeePool(address common.Address, backend bind.ContractBackend) (*FeePool, error) {
	contract, err := bindFeePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeePool{FeePoolCaller: FeePoolCaller{contract: contract}, FeePoolTransactor: FeePoolTransactor{contract: contract}, FeePoolFilterer: FeePoolFilterer{contract: contract}}, nil
}

// NewFeePoolCaller creates a new read-only instance of FeePool, bound to a specific deployed contract.
func NewFeePoolCaller(address common.Address, caller bind.ContractCaller) (*FeePoolCaller, error) {
	contract, err := bindFeePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeePoolCaller{contract: contract}, nil
}

// NewFeePoolTransactor creates a new write-only instance of FeePool, bound to a specific deployed contract.
func NewFeePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*FeePoolTransactor, error) {
	contract, err := bindFeePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeePoolTransactor{contract: contract}, nil
}

// NewFeePoolFilterer creates a new log filterer instance of FeePool, bound to a specific deployed contract.
func NewFeePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*FeePoolFilterer, error) {
	contract, err := bindFeePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeePoolFilterer{contract: contract}, nil
}

// bindFeePool binds a generic wrapper to an already deployed contract.
func bindFeePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeePool *FeePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeePool.Contract.FeePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeePool *FeePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeePool.Contract.FeePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeePool *FeePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeePool.Contract.FeePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeePool *FeePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeePool *FeePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeePool *FeePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeePool.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_FeePool *FeePoolCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FeePool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_FeePool *FeePoolSession) Version() (uint8, error) {
	return _FeePool.Contract.Version(&_FeePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_FeePool *FeePoolCallerSession) Version() (uint8, error) {
	return _FeePool.Contract.Version(&_FeePool.CallOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address _to, uint256 _amount) returns()
func (_FeePool *FeePoolTransactor) WithdrawEther(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeePool.contract.Transact(opts, "withdrawEther", _to, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address _to, uint256 _amount) returns()
func (_FeePool *FeePoolSession) WithdrawEther(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeePool.Contract.WithdrawEther(&_FeePool.TransactOpts, _to, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address _to, uint256 _amount) returns()
func (_FeePool *FeePoolTransactorSession) WithdrawEther(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FeePool.Contract.WithdrawEther(&_FeePool.TransactOpts, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeePool *FeePoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeePool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeePool *FeePoolSession) Receive() (*types.Transaction, error) {
	return _FeePool.Contract.Receive(&_FeePool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeePool *FeePoolTransactorSession) Receive() (*types.Transaction, error) {
	return _FeePool.Contract.Receive(&_FeePool.TransactOpts)
}

// FeePoolEtherWithdrawnIterator is returned from FilterEtherWithdrawn and is used to iterate over the raw logs and unpacked data for EtherWithdrawn events raised by the FeePool contract.
type FeePoolEtherWithdrawnIterator struct {
	Event *FeePoolEtherWithdrawn // Event containing the contract specifics and raw log

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
func (it *FeePoolEtherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeePoolEtherWithdrawn)
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
		it.Event = new(FeePoolEtherWithdrawn)
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
func (it *FeePoolEtherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeePoolEtherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeePoolEtherWithdrawn represents a EtherWithdrawn event raised by the FeePool contract.
type FeePoolEtherWithdrawn struct {
	By     common.Hash
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawn is a free log retrieval operation binding the contract event 0xff381a086696de98df05e634263390296dd1d097ef34a1be9a91daead07fb01d.
//
// Solidity: event EtherWithdrawn(string indexed by, address indexed to, uint256 amount, uint256 time)
func (_FeePool *FeePoolFilterer) FilterEtherWithdrawn(opts *bind.FilterOpts, by []string, to []common.Address) (*FeePoolEtherWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeePool.contract.FilterLogs(opts, "EtherWithdrawn", byRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeePoolEtherWithdrawnIterator{contract: _FeePool.contract, event: "EtherWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawn is a free log subscription operation binding the contract event 0xff381a086696de98df05e634263390296dd1d097ef34a1be9a91daead07fb01d.
//
// Solidity: event EtherWithdrawn(string indexed by, address indexed to, uint256 amount, uint256 time)
func (_FeePool *FeePoolFilterer) WatchEtherWithdrawn(opts *bind.WatchOpts, sink chan<- *FeePoolEtherWithdrawn, by []string, to []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeePool.contract.WatchLogs(opts, "EtherWithdrawn", byRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeePoolEtherWithdrawn)
				if err := _FeePool.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
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

// ParseEtherWithdrawn is a log parse operation binding the contract event 0xff381a086696de98df05e634263390296dd1d097ef34a1be9a91daead07fb01d.
//
// Solidity: event EtherWithdrawn(string indexed by, address indexed to, uint256 amount, uint256 time)
func (_FeePool *FeePoolFilterer) ParseEtherWithdrawn(log types.Log) (*FeePoolEtherWithdrawn, error) {
	event := new(FeePoolEtherWithdrawn)
	if err := _FeePool.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
