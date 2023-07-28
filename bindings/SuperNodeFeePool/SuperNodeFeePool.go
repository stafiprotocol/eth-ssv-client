// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package super_node_fee_pool

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

// SuperNodeFeePoolMetaData contains all meta data concerning the SuperNodeFeePool contract.
var SuperNodeFeePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"by\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SuperNodeFeePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SuperNodeFeePoolMetaData.ABI instead.
var SuperNodeFeePoolABI = SuperNodeFeePoolMetaData.ABI

// SuperNodeFeePool is an auto generated Go binding around an Ethereum contract.
type SuperNodeFeePool struct {
	SuperNodeFeePoolCaller     // Read-only binding to the contract
	SuperNodeFeePoolTransactor // Write-only binding to the contract
	SuperNodeFeePoolFilterer   // Log filterer for contract events
}

// SuperNodeFeePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperNodeFeePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperNodeFeePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperNodeFeePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperNodeFeePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperNodeFeePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperNodeFeePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperNodeFeePoolSession struct {
	Contract     *SuperNodeFeePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperNodeFeePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperNodeFeePoolCallerSession struct {
	Contract *SuperNodeFeePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SuperNodeFeePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperNodeFeePoolTransactorSession struct {
	Contract     *SuperNodeFeePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SuperNodeFeePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperNodeFeePoolRaw struct {
	Contract *SuperNodeFeePool // Generic contract binding to access the raw methods on
}

// SuperNodeFeePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperNodeFeePoolCallerRaw struct {
	Contract *SuperNodeFeePoolCaller // Generic read-only contract binding to access the raw methods on
}

// SuperNodeFeePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperNodeFeePoolTransactorRaw struct {
	Contract *SuperNodeFeePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperNodeFeePool creates a new instance of SuperNodeFeePool, bound to a specific deployed contract.
func NewSuperNodeFeePool(address common.Address, backend bind.ContractBackend) (*SuperNodeFeePool, error) {
	contract, err := bindSuperNodeFeePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SuperNodeFeePool{SuperNodeFeePoolCaller: SuperNodeFeePoolCaller{contract: contract}, SuperNodeFeePoolTransactor: SuperNodeFeePoolTransactor{contract: contract}, SuperNodeFeePoolFilterer: SuperNodeFeePoolFilterer{contract: contract}}, nil
}

// NewSuperNodeFeePoolCaller creates a new read-only instance of SuperNodeFeePool, bound to a specific deployed contract.
func NewSuperNodeFeePoolCaller(address common.Address, caller bind.ContractCaller) (*SuperNodeFeePoolCaller, error) {
	contract, err := bindSuperNodeFeePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperNodeFeePoolCaller{contract: contract}, nil
}

// NewSuperNodeFeePoolTransactor creates a new write-only instance of SuperNodeFeePool, bound to a specific deployed contract.
func NewSuperNodeFeePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperNodeFeePoolTransactor, error) {
	contract, err := bindSuperNodeFeePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperNodeFeePoolTransactor{contract: contract}, nil
}

// NewSuperNodeFeePoolFilterer creates a new log filterer instance of SuperNodeFeePool, bound to a specific deployed contract.
func NewSuperNodeFeePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperNodeFeePoolFilterer, error) {
	contract, err := bindSuperNodeFeePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperNodeFeePoolFilterer{contract: contract}, nil
}

// bindSuperNodeFeePool binds a generic wrapper to an already deployed contract.
func bindSuperNodeFeePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperNodeFeePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperNodeFeePool *SuperNodeFeePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperNodeFeePool.Contract.SuperNodeFeePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperNodeFeePool *SuperNodeFeePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.SuperNodeFeePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperNodeFeePool *SuperNodeFeePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.SuperNodeFeePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperNodeFeePool *SuperNodeFeePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SuperNodeFeePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperNodeFeePool *SuperNodeFeePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperNodeFeePool *SuperNodeFeePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SuperNodeFeePool *SuperNodeFeePoolCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SuperNodeFeePool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SuperNodeFeePool *SuperNodeFeePoolSession) Version() (uint8, error) {
	return _SuperNodeFeePool.Contract.Version(&_SuperNodeFeePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_SuperNodeFeePool *SuperNodeFeePoolCallerSession) Version() (uint8, error) {
	return _SuperNodeFeePool.Contract.Version(&_SuperNodeFeePool.CallOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address _to, uint256 _amount) returns()
func (_SuperNodeFeePool *SuperNodeFeePoolTransactor) WithdrawEther(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SuperNodeFeePool.contract.Transact(opts, "withdrawEther", _to, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address _to, uint256 _amount) returns()
func (_SuperNodeFeePool *SuperNodeFeePoolSession) WithdrawEther(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.WithdrawEther(&_SuperNodeFeePool.TransactOpts, _to, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address _to, uint256 _amount) returns()
func (_SuperNodeFeePool *SuperNodeFeePoolTransactorSession) WithdrawEther(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.WithdrawEther(&_SuperNodeFeePool.TransactOpts, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SuperNodeFeePool *SuperNodeFeePoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperNodeFeePool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SuperNodeFeePool *SuperNodeFeePoolSession) Receive() (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.Receive(&_SuperNodeFeePool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SuperNodeFeePool *SuperNodeFeePoolTransactorSession) Receive() (*types.Transaction, error) {
	return _SuperNodeFeePool.Contract.Receive(&_SuperNodeFeePool.TransactOpts)
}

// SuperNodeFeePoolEtherWithdrawnIterator is returned from FilterEtherWithdrawn and is used to iterate over the raw logs and unpacked data for EtherWithdrawn events raised by the SuperNodeFeePool contract.
type SuperNodeFeePoolEtherWithdrawnIterator struct {
	Event *SuperNodeFeePoolEtherWithdrawn // Event containing the contract specifics and raw log

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
func (it *SuperNodeFeePoolEtherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperNodeFeePoolEtherWithdrawn)
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
		it.Event = new(SuperNodeFeePoolEtherWithdrawn)
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
func (it *SuperNodeFeePoolEtherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperNodeFeePoolEtherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperNodeFeePoolEtherWithdrawn represents a EtherWithdrawn event raised by the SuperNodeFeePool contract.
type SuperNodeFeePoolEtherWithdrawn struct {
	By     common.Hash
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawn is a free log retrieval operation binding the contract event 0xff381a086696de98df05e634263390296dd1d097ef34a1be9a91daead07fb01d.
//
// Solidity: event EtherWithdrawn(string indexed by, address indexed to, uint256 amount, uint256 time)
func (_SuperNodeFeePool *SuperNodeFeePoolFilterer) FilterEtherWithdrawn(opts *bind.FilterOpts, by []string, to []common.Address) (*SuperNodeFeePoolEtherWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SuperNodeFeePool.contract.FilterLogs(opts, "EtherWithdrawn", byRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SuperNodeFeePoolEtherWithdrawnIterator{contract: _SuperNodeFeePool.contract, event: "EtherWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawn is a free log subscription operation binding the contract event 0xff381a086696de98df05e634263390296dd1d097ef34a1be9a91daead07fb01d.
//
// Solidity: event EtherWithdrawn(string indexed by, address indexed to, uint256 amount, uint256 time)
func (_SuperNodeFeePool *SuperNodeFeePoolFilterer) WatchEtherWithdrawn(opts *bind.WatchOpts, sink chan<- *SuperNodeFeePoolEtherWithdrawn, by []string, to []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SuperNodeFeePool.contract.WatchLogs(opts, "EtherWithdrawn", byRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperNodeFeePoolEtherWithdrawn)
				if err := _SuperNodeFeePool.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
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
func (_SuperNodeFeePool *SuperNodeFeePoolFilterer) ParseEtherWithdrawn(log types.Log) (*SuperNodeFeePoolEtherWithdrawn, error) {
	event := new(SuperNodeFeePoolEtherWithdrawn)
	if err := _SuperNodeFeePool.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
