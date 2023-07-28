// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network_settings

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

// NetworkSettingsMetaData contains all meta data concerning the NetworkSettings contract.
var NetworkSettingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stafiStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getNodeConsensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeTrustedRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessWithdrawalsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubmitBalancesEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSuperNodePubkeyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNodeConsensusThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNodeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNodeRefundRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNodeTrustedRefundRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setProcessWithdrawalsEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setSubmitBalancesEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setSuperNodePubkeyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NetworkSettingsABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkSettingsMetaData.ABI instead.
var NetworkSettingsABI = NetworkSettingsMetaData.ABI

// NetworkSettings is an auto generated Go binding around an Ethereum contract.
type NetworkSettings struct {
	NetworkSettingsCaller     // Read-only binding to the contract
	NetworkSettingsTransactor // Write-only binding to the contract
	NetworkSettingsFilterer   // Log filterer for contract events
}

// NetworkSettingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkSettingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkSettingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkSettingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkSettingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkSettingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkSettingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkSettingsSession struct {
	Contract     *NetworkSettings  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkSettingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkSettingsCallerSession struct {
	Contract *NetworkSettingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkSettingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkSettingsTransactorSession struct {
	Contract     *NetworkSettingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkSettingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkSettingsRaw struct {
	Contract *NetworkSettings // Generic contract binding to access the raw methods on
}

// NetworkSettingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkSettingsCallerRaw struct {
	Contract *NetworkSettingsCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkSettingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkSettingsTransactorRaw struct {
	Contract *NetworkSettingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkSettings creates a new instance of NetworkSettings, bound to a specific deployed contract.
func NewNetworkSettings(address common.Address, backend bind.ContractBackend) (*NetworkSettings, error) {
	contract, err := bindNetworkSettings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkSettings{NetworkSettingsCaller: NetworkSettingsCaller{contract: contract}, NetworkSettingsTransactor: NetworkSettingsTransactor{contract: contract}, NetworkSettingsFilterer: NetworkSettingsFilterer{contract: contract}}, nil
}

// NewNetworkSettingsCaller creates a new read-only instance of NetworkSettings, bound to a specific deployed contract.
func NewNetworkSettingsCaller(address common.Address, caller bind.ContractCaller) (*NetworkSettingsCaller, error) {
	contract, err := bindNetworkSettings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkSettingsCaller{contract: contract}, nil
}

// NewNetworkSettingsTransactor creates a new write-only instance of NetworkSettings, bound to a specific deployed contract.
func NewNetworkSettingsTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkSettingsTransactor, error) {
	contract, err := bindNetworkSettings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkSettingsTransactor{contract: contract}, nil
}

// NewNetworkSettingsFilterer creates a new log filterer instance of NetworkSettings, bound to a specific deployed contract.
func NewNetworkSettingsFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkSettingsFilterer, error) {
	contract, err := bindNetworkSettings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkSettingsFilterer{contract: contract}, nil
}

// bindNetworkSettings binds a generic wrapper to an already deployed contract.
func bindNetworkSettings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkSettingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkSettings *NetworkSettingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkSettings.Contract.NetworkSettingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkSettings *NetworkSettingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkSettings.Contract.NetworkSettingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkSettings *NetworkSettingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkSettings.Contract.NetworkSettingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkSettings *NetworkSettingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkSettings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkSettings *NetworkSettingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkSettings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkSettings *NetworkSettingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkSettings.Contract.contract.Transact(opts, method, params...)
}

// GetNodeConsensusThreshold is a free data retrieval call binding the contract method 0x1f66e8ed.
//
// Solidity: function getNodeConsensusThreshold() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCaller) GetNodeConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getNodeConsensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeConsensusThreshold is a free data retrieval call binding the contract method 0x1f66e8ed.
//
// Solidity: function getNodeConsensusThreshold() view returns(uint256)
func (_NetworkSettings *NetworkSettingsSession) GetNodeConsensusThreshold() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeConsensusThreshold(&_NetworkSettings.CallOpts)
}

// GetNodeConsensusThreshold is a free data retrieval call binding the contract method 0x1f66e8ed.
//
// Solidity: function getNodeConsensusThreshold() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCallerSession) GetNodeConsensusThreshold() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeConsensusThreshold(&_NetworkSettings.CallOpts)
}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCaller) GetNodeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getNodeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_NetworkSettings *NetworkSettingsSession) GetNodeFee() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeFee(&_NetworkSettings.CallOpts)
}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCallerSession) GetNodeFee() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeFee(&_NetworkSettings.CallOpts)
}

// GetNodeRefundRatio is a free data retrieval call binding the contract method 0xeb364204.
//
// Solidity: function getNodeRefundRatio() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCaller) GetNodeRefundRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getNodeRefundRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRefundRatio is a free data retrieval call binding the contract method 0xeb364204.
//
// Solidity: function getNodeRefundRatio() view returns(uint256)
func (_NetworkSettings *NetworkSettingsSession) GetNodeRefundRatio() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeRefundRatio(&_NetworkSettings.CallOpts)
}

// GetNodeRefundRatio is a free data retrieval call binding the contract method 0xeb364204.
//
// Solidity: function getNodeRefundRatio() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCallerSession) GetNodeRefundRatio() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeRefundRatio(&_NetworkSettings.CallOpts)
}

// GetNodeTrustedRefundRatio is a free data retrieval call binding the contract method 0xf4cf0830.
//
// Solidity: function getNodeTrustedRefundRatio() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCaller) GetNodeTrustedRefundRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getNodeTrustedRefundRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeTrustedRefundRatio is a free data retrieval call binding the contract method 0xf4cf0830.
//
// Solidity: function getNodeTrustedRefundRatio() view returns(uint256)
func (_NetworkSettings *NetworkSettingsSession) GetNodeTrustedRefundRatio() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeTrustedRefundRatio(&_NetworkSettings.CallOpts)
}

// GetNodeTrustedRefundRatio is a free data retrieval call binding the contract method 0xf4cf0830.
//
// Solidity: function getNodeTrustedRefundRatio() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCallerSession) GetNodeTrustedRefundRatio() (*big.Int, error) {
	return _NetworkSettings.Contract.GetNodeTrustedRefundRatio(&_NetworkSettings.CallOpts)
}

// GetPlatformFee is a free data retrieval call binding the contract method 0x6ea8bc10.
//
// Solidity: function getPlatformFee() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCaller) GetPlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getPlatformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlatformFee is a free data retrieval call binding the contract method 0x6ea8bc10.
//
// Solidity: function getPlatformFee() view returns(uint256)
func (_NetworkSettings *NetworkSettingsSession) GetPlatformFee() (*big.Int, error) {
	return _NetworkSettings.Contract.GetPlatformFee(&_NetworkSettings.CallOpts)
}

// GetPlatformFee is a free data retrieval call binding the contract method 0x6ea8bc10.
//
// Solidity: function getPlatformFee() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCallerSession) GetPlatformFee() (*big.Int, error) {
	return _NetworkSettings.Contract.GetPlatformFee(&_NetworkSettings.CallOpts)
}

// GetProcessWithdrawalsEnabled is a free data retrieval call binding the contract method 0x28ea8d06.
//
// Solidity: function getProcessWithdrawalsEnabled() view returns(bool)
func (_NetworkSettings *NetworkSettingsCaller) GetProcessWithdrawalsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getProcessWithdrawalsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetProcessWithdrawalsEnabled is a free data retrieval call binding the contract method 0x28ea8d06.
//
// Solidity: function getProcessWithdrawalsEnabled() view returns(bool)
func (_NetworkSettings *NetworkSettingsSession) GetProcessWithdrawalsEnabled() (bool, error) {
	return _NetworkSettings.Contract.GetProcessWithdrawalsEnabled(&_NetworkSettings.CallOpts)
}

// GetProcessWithdrawalsEnabled is a free data retrieval call binding the contract method 0x28ea8d06.
//
// Solidity: function getProcessWithdrawalsEnabled() view returns(bool)
func (_NetworkSettings *NetworkSettingsCallerSession) GetProcessWithdrawalsEnabled() (bool, error) {
	return _NetworkSettings.Contract.GetProcessWithdrawalsEnabled(&_NetworkSettings.CallOpts)
}

// GetSubmitBalancesEnabled is a free data retrieval call binding the contract method 0xfcdb60db.
//
// Solidity: function getSubmitBalancesEnabled() view returns(bool)
func (_NetworkSettings *NetworkSettingsCaller) GetSubmitBalancesEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getSubmitBalancesEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSubmitBalancesEnabled is a free data retrieval call binding the contract method 0xfcdb60db.
//
// Solidity: function getSubmitBalancesEnabled() view returns(bool)
func (_NetworkSettings *NetworkSettingsSession) GetSubmitBalancesEnabled() (bool, error) {
	return _NetworkSettings.Contract.GetSubmitBalancesEnabled(&_NetworkSettings.CallOpts)
}

// GetSubmitBalancesEnabled is a free data retrieval call binding the contract method 0xfcdb60db.
//
// Solidity: function getSubmitBalancesEnabled() view returns(bool)
func (_NetworkSettings *NetworkSettingsCallerSession) GetSubmitBalancesEnabled() (bool, error) {
	return _NetworkSettings.Contract.GetSubmitBalancesEnabled(&_NetworkSettings.CallOpts)
}

// GetSuperNodePubkeyLimit is a free data retrieval call binding the contract method 0xb195384f.
//
// Solidity: function getSuperNodePubkeyLimit() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCaller) GetSuperNodePubkeyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getSuperNodePubkeyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuperNodePubkeyLimit is a free data retrieval call binding the contract method 0xb195384f.
//
// Solidity: function getSuperNodePubkeyLimit() view returns(uint256)
func (_NetworkSettings *NetworkSettingsSession) GetSuperNodePubkeyLimit() (*big.Int, error) {
	return _NetworkSettings.Contract.GetSuperNodePubkeyLimit(&_NetworkSettings.CallOpts)
}

// GetSuperNodePubkeyLimit is a free data retrieval call binding the contract method 0xb195384f.
//
// Solidity: function getSuperNodePubkeyLimit() view returns(uint256)
func (_NetworkSettings *NetworkSettingsCallerSession) GetSuperNodePubkeyLimit() (*big.Int, error) {
	return _NetworkSettings.Contract.GetSuperNodePubkeyLimit(&_NetworkSettings.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes)
func (_NetworkSettings *NetworkSettingsCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes)
func (_NetworkSettings *NetworkSettingsSession) GetWithdrawalCredentials() ([]byte, error) {
	return _NetworkSettings.Contract.GetWithdrawalCredentials(&_NetworkSettings.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes)
func (_NetworkSettings *NetworkSettingsCallerSession) GetWithdrawalCredentials() ([]byte, error) {
	return _NetworkSettings.Contract.GetWithdrawalCredentials(&_NetworkSettings.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkSettings *NetworkSettingsCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NetworkSettings.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkSettings *NetworkSettingsSession) Version() (uint8, error) {
	return _NetworkSettings.Contract.Version(&_NetworkSettings.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkSettings *NetworkSettingsCallerSession) Version() (uint8, error) {
	return _NetworkSettings.Contract.Version(&_NetworkSettings.CallOpts)
}

// SetNodeConsensusThreshold is a paid mutator transaction binding the contract method 0x977886ea.
//
// Solidity: function setNodeConsensusThreshold(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetNodeConsensusThreshold(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setNodeConsensusThreshold", _value)
}

// SetNodeConsensusThreshold is a paid mutator transaction binding the contract method 0x977886ea.
//
// Solidity: function setNodeConsensusThreshold(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetNodeConsensusThreshold(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeConsensusThreshold(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeConsensusThreshold is a paid mutator transaction binding the contract method 0x977886ea.
//
// Solidity: function setNodeConsensusThreshold(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetNodeConsensusThreshold(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeConsensusThreshold(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeFee is a paid mutator transaction binding the contract method 0x89093310.
//
// Solidity: function setNodeFee(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetNodeFee(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setNodeFee", _value)
}

// SetNodeFee is a paid mutator transaction binding the contract method 0x89093310.
//
// Solidity: function setNodeFee(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetNodeFee(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeFee(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeFee is a paid mutator transaction binding the contract method 0x89093310.
//
// Solidity: function setNodeFee(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetNodeFee(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeFee(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeRefundRatio is a paid mutator transaction binding the contract method 0xb67b446a.
//
// Solidity: function setNodeRefundRatio(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetNodeRefundRatio(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setNodeRefundRatio", _value)
}

// SetNodeRefundRatio is a paid mutator transaction binding the contract method 0xb67b446a.
//
// Solidity: function setNodeRefundRatio(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetNodeRefundRatio(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeRefundRatio(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeRefundRatio is a paid mutator transaction binding the contract method 0xb67b446a.
//
// Solidity: function setNodeRefundRatio(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetNodeRefundRatio(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeRefundRatio(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeTrustedRefundRatio is a paid mutator transaction binding the contract method 0x89799aa1.
//
// Solidity: function setNodeTrustedRefundRatio(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetNodeTrustedRefundRatio(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setNodeTrustedRefundRatio", _value)
}

// SetNodeTrustedRefundRatio is a paid mutator transaction binding the contract method 0x89799aa1.
//
// Solidity: function setNodeTrustedRefundRatio(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetNodeTrustedRefundRatio(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeTrustedRefundRatio(&_NetworkSettings.TransactOpts, _value)
}

// SetNodeTrustedRefundRatio is a paid mutator transaction binding the contract method 0x89799aa1.
//
// Solidity: function setNodeTrustedRefundRatio(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetNodeTrustedRefundRatio(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetNodeTrustedRefundRatio(&_NetworkSettings.TransactOpts, _value)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetPlatformFee(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setPlatformFee", _value)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetPlatformFee(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetPlatformFee(&_NetworkSettings.TransactOpts, _value)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetPlatformFee(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetPlatformFee(&_NetworkSettings.TransactOpts, _value)
}

// SetProcessWithdrawalsEnabled is a paid mutator transaction binding the contract method 0xc6fbae5e.
//
// Solidity: function setProcessWithdrawalsEnabled(bool _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetProcessWithdrawalsEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setProcessWithdrawalsEnabled", _value)
}

// SetProcessWithdrawalsEnabled is a paid mutator transaction binding the contract method 0xc6fbae5e.
//
// Solidity: function setProcessWithdrawalsEnabled(bool _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetProcessWithdrawalsEnabled(_value bool) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetProcessWithdrawalsEnabled(&_NetworkSettings.TransactOpts, _value)
}

// SetProcessWithdrawalsEnabled is a paid mutator transaction binding the contract method 0xc6fbae5e.
//
// Solidity: function setProcessWithdrawalsEnabled(bool _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetProcessWithdrawalsEnabled(_value bool) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetProcessWithdrawalsEnabled(&_NetworkSettings.TransactOpts, _value)
}

// SetSubmitBalancesEnabled is a paid mutator transaction binding the contract method 0xa37a9e61.
//
// Solidity: function setSubmitBalancesEnabled(bool _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetSubmitBalancesEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setSubmitBalancesEnabled", _value)
}

// SetSubmitBalancesEnabled is a paid mutator transaction binding the contract method 0xa37a9e61.
//
// Solidity: function setSubmitBalancesEnabled(bool _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetSubmitBalancesEnabled(_value bool) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetSubmitBalancesEnabled(&_NetworkSettings.TransactOpts, _value)
}

// SetSubmitBalancesEnabled is a paid mutator transaction binding the contract method 0xa37a9e61.
//
// Solidity: function setSubmitBalancesEnabled(bool _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetSubmitBalancesEnabled(_value bool) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetSubmitBalancesEnabled(&_NetworkSettings.TransactOpts, _value)
}

// SetSuperNodePubkeyLimit is a paid mutator transaction binding the contract method 0xa8520bec.
//
// Solidity: function setSuperNodePubkeyLimit(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetSuperNodePubkeyLimit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setSuperNodePubkeyLimit", _value)
}

// SetSuperNodePubkeyLimit is a paid mutator transaction binding the contract method 0xa8520bec.
//
// Solidity: function setSuperNodePubkeyLimit(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetSuperNodePubkeyLimit(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetSuperNodePubkeyLimit(&_NetworkSettings.TransactOpts, _value)
}

// SetSuperNodePubkeyLimit is a paid mutator transaction binding the contract method 0xa8520bec.
//
// Solidity: function setSuperNodePubkeyLimit(uint256 _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetSuperNodePubkeyLimit(_value *big.Int) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetSuperNodePubkeyLimit(&_NetworkSettings.TransactOpts, _value)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc4cf71c9.
//
// Solidity: function setWithdrawalCredentials(bytes _value) returns()
func (_NetworkSettings *NetworkSettingsTransactor) SetWithdrawalCredentials(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _NetworkSettings.contract.Transact(opts, "setWithdrawalCredentials", _value)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc4cf71c9.
//
// Solidity: function setWithdrawalCredentials(bytes _value) returns()
func (_NetworkSettings *NetworkSettingsSession) SetWithdrawalCredentials(_value []byte) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetWithdrawalCredentials(&_NetworkSettings.TransactOpts, _value)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xc4cf71c9.
//
// Solidity: function setWithdrawalCredentials(bytes _value) returns()
func (_NetworkSettings *NetworkSettingsTransactorSession) SetWithdrawalCredentials(_value []byte) (*types.Transaction, error) {
	return _NetworkSettings.Contract.SetWithdrawalCredentials(&_NetworkSettings.TransactOpts, _value)
}
