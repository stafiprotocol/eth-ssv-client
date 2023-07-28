// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network_balances

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

// NetworkBalancesMetaData contains all meta data concerning the NetworkBalances contract.
var NetworkBalancesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rethSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BalancesSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rethSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BalancesUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBalancesBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHStakingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRETHSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rethSupply\",\"type\":\"uint256\"}],\"name\":\"nodeVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rethSupply\",\"type\":\"uint256\"}],\"name\":\"submitBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NetworkBalancesABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkBalancesMetaData.ABI instead.
var NetworkBalancesABI = NetworkBalancesMetaData.ABI

// NetworkBalances is an auto generated Go binding around an Ethereum contract.
type NetworkBalances struct {
	NetworkBalancesCaller     // Read-only binding to the contract
	NetworkBalancesTransactor // Write-only binding to the contract
	NetworkBalancesFilterer   // Log filterer for contract events
}

// NetworkBalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkBalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkBalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkBalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkBalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkBalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkBalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkBalancesSession struct {
	Contract     *NetworkBalances  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkBalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkBalancesCallerSession struct {
	Contract *NetworkBalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkBalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkBalancesTransactorSession struct {
	Contract     *NetworkBalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkBalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkBalancesRaw struct {
	Contract *NetworkBalances // Generic contract binding to access the raw methods on
}

// NetworkBalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkBalancesCallerRaw struct {
	Contract *NetworkBalancesCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkBalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkBalancesTransactorRaw struct {
	Contract *NetworkBalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkBalances creates a new instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalances(address common.Address, backend bind.ContractBackend) (*NetworkBalances, error) {
	contract, err := bindNetworkBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkBalances{NetworkBalancesCaller: NetworkBalancesCaller{contract: contract}, NetworkBalancesTransactor: NetworkBalancesTransactor{contract: contract}, NetworkBalancesFilterer: NetworkBalancesFilterer{contract: contract}}, nil
}

// NewNetworkBalancesCaller creates a new read-only instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalancesCaller(address common.Address, caller bind.ContractCaller) (*NetworkBalancesCaller, error) {
	contract, err := bindNetworkBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesCaller{contract: contract}, nil
}

// NewNetworkBalancesTransactor creates a new write-only instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkBalancesTransactor, error) {
	contract, err := bindNetworkBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesTransactor{contract: contract}, nil
}

// NewNetworkBalancesFilterer creates a new log filterer instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkBalancesFilterer, error) {
	contract, err := bindNetworkBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesFilterer{contract: contract}, nil
}

// bindNetworkBalances binds a generic wrapper to an already deployed contract.
func bindNetworkBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkBalancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkBalances *NetworkBalancesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkBalances.Contract.NetworkBalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkBalances *NetworkBalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkBalances.Contract.NetworkBalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkBalances *NetworkBalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkBalances.Contract.NetworkBalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkBalances *NetworkBalancesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkBalances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkBalances *NetworkBalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkBalances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkBalances *NetworkBalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkBalances.Contract.contract.Transact(opts, method, params...)
}

// GetBalancesBlock is a free data retrieval call binding the contract method 0x9100c13d.
//
// Solidity: function getBalancesBlock() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetBalancesBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getBalancesBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalancesBlock is a free data retrieval call binding the contract method 0x9100c13d.
//
// Solidity: function getBalancesBlock() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetBalancesBlock() (*big.Int, error) {
	return _NetworkBalances.Contract.GetBalancesBlock(&_NetworkBalances.CallOpts)
}

// GetBalancesBlock is a free data retrieval call binding the contract method 0x9100c13d.
//
// Solidity: function getBalancesBlock() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetBalancesBlock() (*big.Int, error) {
	return _NetworkBalances.Contract.GetBalancesBlock(&_NetworkBalances.CallOpts)
}

// GetETHStakingRate is a free data retrieval call binding the contract method 0xcd482567.
//
// Solidity: function getETHStakingRate() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetETHStakingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getETHStakingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETHStakingRate is a free data retrieval call binding the contract method 0xcd482567.
//
// Solidity: function getETHStakingRate() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetETHStakingRate() (*big.Int, error) {
	return _NetworkBalances.Contract.GetETHStakingRate(&_NetworkBalances.CallOpts)
}

// GetETHStakingRate is a free data retrieval call binding the contract method 0xcd482567.
//
// Solidity: function getETHStakingRate() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetETHStakingRate() (*big.Int, error) {
	return _NetworkBalances.Contract.GetETHStakingRate(&_NetworkBalances.CallOpts)
}

// GetStakingETHBalance is a free data retrieval call binding the contract method 0xf1eda634.
//
// Solidity: function getStakingETHBalance() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetStakingETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getStakingETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingETHBalance is a free data retrieval call binding the contract method 0xf1eda634.
//
// Solidity: function getStakingETHBalance() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetStakingETHBalance() (*big.Int, error) {
	return _NetworkBalances.Contract.GetStakingETHBalance(&_NetworkBalances.CallOpts)
}

// GetStakingETHBalance is a free data retrieval call binding the contract method 0xf1eda634.
//
// Solidity: function getStakingETHBalance() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetStakingETHBalance() (*big.Int, error) {
	return _NetworkBalances.Contract.GetStakingETHBalance(&_NetworkBalances.CallOpts)
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetTotalETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getTotalETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetTotalETHBalance() (*big.Int, error) {
	return _NetworkBalances.Contract.GetTotalETHBalance(&_NetworkBalances.CallOpts)
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetTotalETHBalance() (*big.Int, error) {
	return _NetworkBalances.Contract.GetTotalETHBalance(&_NetworkBalances.CallOpts)
}

// GetTotalRETHSupply is a free data retrieval call binding the contract method 0xc4c8d0ad.
//
// Solidity: function getTotalRETHSupply() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetTotalRETHSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getTotalRETHSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRETHSupply is a free data retrieval call binding the contract method 0xc4c8d0ad.
//
// Solidity: function getTotalRETHSupply() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetTotalRETHSupply() (*big.Int, error) {
	return _NetworkBalances.Contract.GetTotalRETHSupply(&_NetworkBalances.CallOpts)
}

// GetTotalRETHSupply is a free data retrieval call binding the contract method 0xc4c8d0ad.
//
// Solidity: function getTotalRETHSupply() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetTotalRETHSupply() (*big.Int, error) {
	return _NetworkBalances.Contract.GetTotalRETHSupply(&_NetworkBalances.CallOpts)
}

// NodeVoted is a free data retrieval call binding the contract method 0xcad46175.
//
// Solidity: function nodeVoted(address sender, uint256 _block, uint256 _totalEth, uint256 _stakingEth, uint256 _rethSupply) view returns(bool)
func (_NetworkBalances *NetworkBalancesCaller) NodeVoted(opts *bind.CallOpts, sender common.Address, _block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) (bool, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "nodeVoted", sender, _block, _totalEth, _stakingEth, _rethSupply)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NodeVoted is a free data retrieval call binding the contract method 0xcad46175.
//
// Solidity: function nodeVoted(address sender, uint256 _block, uint256 _totalEth, uint256 _stakingEth, uint256 _rethSupply) view returns(bool)
func (_NetworkBalances *NetworkBalancesSession) NodeVoted(sender common.Address, _block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) (bool, error) {
	return _NetworkBalances.Contract.NodeVoted(&_NetworkBalances.CallOpts, sender, _block, _totalEth, _stakingEth, _rethSupply)
}

// NodeVoted is a free data retrieval call binding the contract method 0xcad46175.
//
// Solidity: function nodeVoted(address sender, uint256 _block, uint256 _totalEth, uint256 _stakingEth, uint256 _rethSupply) view returns(bool)
func (_NetworkBalances *NetworkBalancesCallerSession) NodeVoted(sender common.Address, _block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) (bool, error) {
	return _NetworkBalances.Contract.NodeVoted(&_NetworkBalances.CallOpts, sender, _block, _totalEth, _stakingEth, _rethSupply)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkBalances *NetworkBalancesCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkBalances *NetworkBalancesSession) Version() (uint8, error) {
	return _NetworkBalances.Contract.Version(&_NetworkBalances.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkBalances *NetworkBalancesCallerSession) Version() (uint8, error) {
	return _NetworkBalances.Contract.Version(&_NetworkBalances.CallOpts)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0x32db5470.
//
// Solidity: function submitBalances(uint256 _block, uint256 _totalEth, uint256 _stakingEth, uint256 _rethSupply) returns()
func (_NetworkBalances *NetworkBalancesTransactor) SubmitBalances(opts *bind.TransactOpts, _block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "submitBalances", _block, _totalEth, _stakingEth, _rethSupply)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0x32db5470.
//
// Solidity: function submitBalances(uint256 _block, uint256 _totalEth, uint256 _stakingEth, uint256 _rethSupply) returns()
func (_NetworkBalances *NetworkBalancesSession) SubmitBalances(_block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SubmitBalances(&_NetworkBalances.TransactOpts, _block, _totalEth, _stakingEth, _rethSupply)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0x32db5470.
//
// Solidity: function submitBalances(uint256 _block, uint256 _totalEth, uint256 _stakingEth, uint256 _rethSupply) returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) SubmitBalances(_block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SubmitBalances(&_NetworkBalances.TransactOpts, _block, _totalEth, _stakingEth, _rethSupply)
}

// NetworkBalancesBalancesSubmittedIterator is returned from FilterBalancesSubmitted and is used to iterate over the raw logs and unpacked data for BalancesSubmitted events raised by the NetworkBalances contract.
type NetworkBalancesBalancesSubmittedIterator struct {
	Event *NetworkBalancesBalancesSubmitted // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesBalancesSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesBalancesSubmitted)
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
		it.Event = new(NetworkBalancesBalancesSubmitted)
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
func (it *NetworkBalancesBalancesSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesBalancesSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesBalancesSubmitted represents a BalancesSubmitted event raised by the NetworkBalances contract.
type NetworkBalancesBalancesSubmitted struct {
	From       common.Address
	Block      *big.Int
	TotalEth   *big.Int
	StakingEth *big.Int
	RethSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBalancesSubmitted is a free log retrieval operation binding the contract event 0xe657a6d6957f4fabb37b86d4d6571e82df061bd2d8a3ede5d197b0b98a5a1bdf.
//
// Solidity: event BalancesSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 stakingEth, uint256 rethSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) FilterBalancesSubmitted(opts *bind.FilterOpts, from []common.Address) (*NetworkBalancesBalancesSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "BalancesSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesBalancesSubmittedIterator{contract: _NetworkBalances.contract, event: "BalancesSubmitted", logs: logs, sub: sub}, nil
}

// WatchBalancesSubmitted is a free log subscription operation binding the contract event 0xe657a6d6957f4fabb37b86d4d6571e82df061bd2d8a3ede5d197b0b98a5a1bdf.
//
// Solidity: event BalancesSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 stakingEth, uint256 rethSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) WatchBalancesSubmitted(opts *bind.WatchOpts, sink chan<- *NetworkBalancesBalancesSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "BalancesSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesBalancesSubmitted)
				if err := _NetworkBalances.contract.UnpackLog(event, "BalancesSubmitted", log); err != nil {
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

// ParseBalancesSubmitted is a log parse operation binding the contract event 0xe657a6d6957f4fabb37b86d4d6571e82df061bd2d8a3ede5d197b0b98a5a1bdf.
//
// Solidity: event BalancesSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 stakingEth, uint256 rethSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) ParseBalancesSubmitted(log types.Log) (*NetworkBalancesBalancesSubmitted, error) {
	event := new(NetworkBalancesBalancesSubmitted)
	if err := _NetworkBalances.contract.UnpackLog(event, "BalancesSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkBalancesBalancesUpdatedIterator is returned from FilterBalancesUpdated and is used to iterate over the raw logs and unpacked data for BalancesUpdated events raised by the NetworkBalances contract.
type NetworkBalancesBalancesUpdatedIterator struct {
	Event *NetworkBalancesBalancesUpdated // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesBalancesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesBalancesUpdated)
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
		it.Event = new(NetworkBalancesBalancesUpdated)
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
func (it *NetworkBalancesBalancesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesBalancesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesBalancesUpdated represents a BalancesUpdated event raised by the NetworkBalances contract.
type NetworkBalancesBalancesUpdated struct {
	Block      *big.Int
	TotalEth   *big.Int
	StakingEth *big.Int
	RethSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBalancesUpdated is a free log retrieval operation binding the contract event 0x7bbbb137fdad433d6168b1c75c714c72b8abe8d07460f0c0b433063e7bf1f394.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 stakingEth, uint256 rethSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) FilterBalancesUpdated(opts *bind.FilterOpts) (*NetworkBalancesBalancesUpdatedIterator, error) {

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "BalancesUpdated")
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesBalancesUpdatedIterator{contract: _NetworkBalances.contract, event: "BalancesUpdated", logs: logs, sub: sub}, nil
}

// WatchBalancesUpdated is a free log subscription operation binding the contract event 0x7bbbb137fdad433d6168b1c75c714c72b8abe8d07460f0c0b433063e7bf1f394.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 stakingEth, uint256 rethSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) WatchBalancesUpdated(opts *bind.WatchOpts, sink chan<- *NetworkBalancesBalancesUpdated) (event.Subscription, error) {

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "BalancesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesBalancesUpdated)
				if err := _NetworkBalances.contract.UnpackLog(event, "BalancesUpdated", log); err != nil {
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

// ParseBalancesUpdated is a log parse operation binding the contract event 0x7bbbb137fdad433d6168b1c75c714c72b8abe8d07460f0c0b433063e7bf1f394.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 stakingEth, uint256 rethSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) ParseBalancesUpdated(log types.Log) (*NetworkBalancesBalancesUpdated, error) {
	event := new(NetworkBalancesBalancesUpdated)
	if err := _NetworkBalances.contract.UnpackLog(event, "BalancesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
