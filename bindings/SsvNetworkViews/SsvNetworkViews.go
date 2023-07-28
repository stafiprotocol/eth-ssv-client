// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssv_network_views

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

// ISSVNetworkCoreCluster is an auto generated low-level Go binding around an user-defined struct.
type ISSVNetworkCoreCluster struct {
	ValidatorCount  uint32
	NetworkFeeIndex uint64
	Index           uint64
	Active          bool
	Balance         *big.Int
}

// SsvNetworkViewsMetaData contains all meta data concerning the SsvNetworkViews contract.
var SsvNetworkViewsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalNotWithinTimeframe\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterDoesNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterIsLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterNotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedValidatorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsIncreaseLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeIncreaseNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectClusterState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectValidatorState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIdsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxValueExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewBlockPeriodIsBelowMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeDeclared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorsListNotUnique\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameFeeChangeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetModuleDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsortedOperatorsList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"getBurnRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationThresholdPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumLiquidationCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetworkEarnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetworkFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"getOperatorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"getOperatorDeclaredFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"getOperatorEarnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"getOperatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorFeeIncreaseLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorMaxFeeIncrease\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorFeePeriods\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"declareOperatorFeePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executeOperatorFeePeriod\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsPerOperatorLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISSVViews\",\"name\":\"ssvNetwork_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"isLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"isLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ssvNetwork\",\"outputs\":[{\"internalType\":\"contractISSVViews\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SsvNetworkViewsABI is the input ABI used to generate the binding from.
// Deprecated: Use SsvNetworkViewsMetaData.ABI instead.
var SsvNetworkViewsABI = SsvNetworkViewsMetaData.ABI

// SsvNetworkViews is an auto generated Go binding around an Ethereum contract.
type SsvNetworkViews struct {
	SsvNetworkViewsCaller     // Read-only binding to the contract
	SsvNetworkViewsTransactor // Write-only binding to the contract
	SsvNetworkViewsFilterer   // Log filterer for contract events
}

// SsvNetworkViewsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsvNetworkViewsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvNetworkViewsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsvNetworkViewsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvNetworkViewsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsvNetworkViewsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvNetworkViewsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsvNetworkViewsSession struct {
	Contract     *SsvNetworkViews  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsvNetworkViewsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsvNetworkViewsCallerSession struct {
	Contract *SsvNetworkViewsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SsvNetworkViewsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsvNetworkViewsTransactorSession struct {
	Contract     *SsvNetworkViewsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SsvNetworkViewsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsvNetworkViewsRaw struct {
	Contract *SsvNetworkViews // Generic contract binding to access the raw methods on
}

// SsvNetworkViewsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsvNetworkViewsCallerRaw struct {
	Contract *SsvNetworkViewsCaller // Generic read-only contract binding to access the raw methods on
}

// SsvNetworkViewsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsvNetworkViewsTransactorRaw struct {
	Contract *SsvNetworkViewsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsvNetworkViews creates a new instance of SsvNetworkViews, bound to a specific deployed contract.
func NewSsvNetworkViews(address common.Address, backend bind.ContractBackend) (*SsvNetworkViews, error) {
	contract, err := bindSsvNetworkViews(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViews{SsvNetworkViewsCaller: SsvNetworkViewsCaller{contract: contract}, SsvNetworkViewsTransactor: SsvNetworkViewsTransactor{contract: contract}, SsvNetworkViewsFilterer: SsvNetworkViewsFilterer{contract: contract}}, nil
}

// NewSsvNetworkViewsCaller creates a new read-only instance of SsvNetworkViews, bound to a specific deployed contract.
func NewSsvNetworkViewsCaller(address common.Address, caller bind.ContractCaller) (*SsvNetworkViewsCaller, error) {
	contract, err := bindSsvNetworkViews(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsCaller{contract: contract}, nil
}

// NewSsvNetworkViewsTransactor creates a new write-only instance of SsvNetworkViews, bound to a specific deployed contract.
func NewSsvNetworkViewsTransactor(address common.Address, transactor bind.ContractTransactor) (*SsvNetworkViewsTransactor, error) {
	contract, err := bindSsvNetworkViews(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsTransactor{contract: contract}, nil
}

// NewSsvNetworkViewsFilterer creates a new log filterer instance of SsvNetworkViews, bound to a specific deployed contract.
func NewSsvNetworkViewsFilterer(address common.Address, filterer bind.ContractFilterer) (*SsvNetworkViewsFilterer, error) {
	contract, err := bindSsvNetworkViews(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsFilterer{contract: contract}, nil
}

// bindSsvNetworkViews binds a generic wrapper to an already deployed contract.
func bindSsvNetworkViews(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SsvNetworkViewsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsvNetworkViews *SsvNetworkViewsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsvNetworkViews.Contract.SsvNetworkViewsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsvNetworkViews *SsvNetworkViewsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.SsvNetworkViewsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsvNetworkViews *SsvNetworkViewsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.SsvNetworkViewsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsvNetworkViews *SsvNetworkViewsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsvNetworkViews.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsvNetworkViews *SsvNetworkViewsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsvNetworkViews *SsvNetworkViewsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xeb8ecfa7.
//
// Solidity: function getBalance(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetBalance(opts *bind.CallOpts, clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getBalance", clusterOwner, operatorIds, cluster)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xeb8ecfa7.
//
// Solidity: function getBalance(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetBalance(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetBalance(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// GetBalance is a free data retrieval call binding the contract method 0xeb8ecfa7.
//
// Solidity: function getBalance(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetBalance(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetBalance(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// GetBurnRate is a free data retrieval call binding the contract method 0xca162e5e.
//
// Solidity: function getBurnRate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetBurnRate(opts *bind.CallOpts, clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getBurnRate", clusterOwner, operatorIds, cluster)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnRate is a free data retrieval call binding the contract method 0xca162e5e.
//
// Solidity: function getBurnRate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetBurnRate(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetBurnRate(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// GetBurnRate is a free data retrieval call binding the contract method 0xca162e5e.
//
// Solidity: function getBurnRate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetBurnRate(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetBurnRate(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// GetLiquidationThresholdPeriod is a free data retrieval call binding the contract method 0x9040f7c3.
//
// Solidity: function getLiquidationThresholdPeriod() view returns(uint64)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetLiquidationThresholdPeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getLiquidationThresholdPeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLiquidationThresholdPeriod is a free data retrieval call binding the contract method 0x9040f7c3.
//
// Solidity: function getLiquidationThresholdPeriod() view returns(uint64)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetLiquidationThresholdPeriod() (uint64, error) {
	return _SsvNetworkViews.Contract.GetLiquidationThresholdPeriod(&_SsvNetworkViews.CallOpts)
}

// GetLiquidationThresholdPeriod is a free data retrieval call binding the contract method 0x9040f7c3.
//
// Solidity: function getLiquidationThresholdPeriod() view returns(uint64)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetLiquidationThresholdPeriod() (uint64, error) {
	return _SsvNetworkViews.Contract.GetLiquidationThresholdPeriod(&_SsvNetworkViews.CallOpts)
}

// GetMinimumLiquidationCollateral is a free data retrieval call binding the contract method 0x5ba3d62a.
//
// Solidity: function getMinimumLiquidationCollateral() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetMinimumLiquidationCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getMinimumLiquidationCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumLiquidationCollateral is a free data retrieval call binding the contract method 0x5ba3d62a.
//
// Solidity: function getMinimumLiquidationCollateral() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetMinimumLiquidationCollateral() (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetMinimumLiquidationCollateral(&_SsvNetworkViews.CallOpts)
}

// GetMinimumLiquidationCollateral is a free data retrieval call binding the contract method 0x5ba3d62a.
//
// Solidity: function getMinimumLiquidationCollateral() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetMinimumLiquidationCollateral() (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetMinimumLiquidationCollateral(&_SsvNetworkViews.CallOpts)
}

// GetNetworkEarnings is a free data retrieval call binding the contract method 0x777915cb.
//
// Solidity: function getNetworkEarnings() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetNetworkEarnings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getNetworkEarnings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkEarnings is a free data retrieval call binding the contract method 0x777915cb.
//
// Solidity: function getNetworkEarnings() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetNetworkEarnings() (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetNetworkEarnings(&_SsvNetworkViews.CallOpts)
}

// GetNetworkEarnings is a free data retrieval call binding the contract method 0x777915cb.
//
// Solidity: function getNetworkEarnings() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetNetworkEarnings() (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetNetworkEarnings(&_SsvNetworkViews.CallOpts)
}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetNetworkFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getNetworkFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetNetworkFee() (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetNetworkFee(&_SsvNetworkViews.CallOpts)
}

// GetNetworkFee is a free data retrieval call binding the contract method 0xfc043830.
//
// Solidity: function getNetworkFee() view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetNetworkFee() (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetNetworkFee(&_SsvNetworkViews.CallOpts)
}

// GetOperatorById is a free data retrieval call binding the contract method 0xbe3f058e.
//
// Solidity: function getOperatorById(uint64 operatorId) view returns(address, uint256, uint32, address, bool, bool)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetOperatorById(opts *bind.CallOpts, operatorId uint64) (common.Address, *big.Int, uint32, common.Address, bool, bool, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getOperatorById", operatorId)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(uint32), *new(common.Address), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetOperatorById is a free data retrieval call binding the contract method 0xbe3f058e.
//
// Solidity: function getOperatorById(uint64 operatorId) view returns(address, uint256, uint32, address, bool, bool)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetOperatorById(operatorId uint64) (common.Address, *big.Int, uint32, common.Address, bool, bool, error) {
	return _SsvNetworkViews.Contract.GetOperatorById(&_SsvNetworkViews.CallOpts, operatorId)
}

// GetOperatorById is a free data retrieval call binding the contract method 0xbe3f058e.
//
// Solidity: function getOperatorById(uint64 operatorId) view returns(address, uint256, uint32, address, bool, bool)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetOperatorById(operatorId uint64) (common.Address, *big.Int, uint32, common.Address, bool, bool, error) {
	return _SsvNetworkViews.Contract.GetOperatorById(&_SsvNetworkViews.CallOpts, operatorId)
}

// GetOperatorDeclaredFee is a free data retrieval call binding the contract method 0x03b3d436.
//
// Solidity: function getOperatorDeclaredFee(uint64 operatorId) view returns(bool, uint256, uint64, uint64)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetOperatorDeclaredFee(opts *bind.CallOpts, operatorId uint64) (bool, *big.Int, uint64, uint64, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getOperatorDeclaredFee", operatorId)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return out0, out1, out2, out3, err

}

// GetOperatorDeclaredFee is a free data retrieval call binding the contract method 0x03b3d436.
//
// Solidity: function getOperatorDeclaredFee(uint64 operatorId) view returns(bool, uint256, uint64, uint64)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetOperatorDeclaredFee(operatorId uint64) (bool, *big.Int, uint64, uint64, error) {
	return _SsvNetworkViews.Contract.GetOperatorDeclaredFee(&_SsvNetworkViews.CallOpts, operatorId)
}

// GetOperatorDeclaredFee is a free data retrieval call binding the contract method 0x03b3d436.
//
// Solidity: function getOperatorDeclaredFee(uint64 operatorId) view returns(bool, uint256, uint64, uint64)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetOperatorDeclaredFee(operatorId uint64) (bool, *big.Int, uint64, uint64, error) {
	return _SsvNetworkViews.Contract.GetOperatorDeclaredFee(&_SsvNetworkViews.CallOpts, operatorId)
}

// GetOperatorEarnings is a free data retrieval call binding the contract method 0x6d0db0e4.
//
// Solidity: function getOperatorEarnings(uint64 id) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetOperatorEarnings(opts *bind.CallOpts, id uint64) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getOperatorEarnings", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorEarnings is a free data retrieval call binding the contract method 0x6d0db0e4.
//
// Solidity: function getOperatorEarnings(uint64 id) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetOperatorEarnings(id uint64) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetOperatorEarnings(&_SsvNetworkViews.CallOpts, id)
}

// GetOperatorEarnings is a free data retrieval call binding the contract method 0x6d0db0e4.
//
// Solidity: function getOperatorEarnings(uint64 id) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetOperatorEarnings(id uint64) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetOperatorEarnings(&_SsvNetworkViews.CallOpts, id)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0x9ad3c745.
//
// Solidity: function getOperatorFee(uint64 operatorId) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetOperatorFee(opts *bind.CallOpts, operatorId uint64) (*big.Int, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getOperatorFee", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorFee is a free data retrieval call binding the contract method 0x9ad3c745.
//
// Solidity: function getOperatorFee(uint64 operatorId) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetOperatorFee(operatorId uint64) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetOperatorFee(&_SsvNetworkViews.CallOpts, operatorId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0x9ad3c745.
//
// Solidity: function getOperatorFee(uint64 operatorId) view returns(uint256)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetOperatorFee(operatorId uint64) (*big.Int, error) {
	return _SsvNetworkViews.Contract.GetOperatorFee(&_SsvNetworkViews.CallOpts, operatorId)
}

// GetOperatorFeeIncreaseLimit is a free data retrieval call binding the contract method 0x68465f7d.
//
// Solidity: function getOperatorFeeIncreaseLimit() view returns(uint64 operatorMaxFeeIncrease)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetOperatorFeeIncreaseLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getOperatorFeeIncreaseLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOperatorFeeIncreaseLimit is a free data retrieval call binding the contract method 0x68465f7d.
//
// Solidity: function getOperatorFeeIncreaseLimit() view returns(uint64 operatorMaxFeeIncrease)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetOperatorFeeIncreaseLimit() (uint64, error) {
	return _SsvNetworkViews.Contract.GetOperatorFeeIncreaseLimit(&_SsvNetworkViews.CallOpts)
}

// GetOperatorFeeIncreaseLimit is a free data retrieval call binding the contract method 0x68465f7d.
//
// Solidity: function getOperatorFeeIncreaseLimit() view returns(uint64 operatorMaxFeeIncrease)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetOperatorFeeIncreaseLimit() (uint64, error) {
	return _SsvNetworkViews.Contract.GetOperatorFeeIncreaseLimit(&_SsvNetworkViews.CallOpts)
}

// GetOperatorFeePeriods is a free data retrieval call binding the contract method 0xe6d2834d.
//
// Solidity: function getOperatorFeePeriods() view returns(uint64 declareOperatorFeePeriod, uint64 executeOperatorFeePeriod)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetOperatorFeePeriods(opts *bind.CallOpts) (struct {
	DeclareOperatorFeePeriod uint64
	ExecuteOperatorFeePeriod uint64
}, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getOperatorFeePeriods")

	outstruct := new(struct {
		DeclareOperatorFeePeriod uint64
		ExecuteOperatorFeePeriod uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DeclareOperatorFeePeriod = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ExecuteOperatorFeePeriod = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetOperatorFeePeriods is a free data retrieval call binding the contract method 0xe6d2834d.
//
// Solidity: function getOperatorFeePeriods() view returns(uint64 declareOperatorFeePeriod, uint64 executeOperatorFeePeriod)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetOperatorFeePeriods() (struct {
	DeclareOperatorFeePeriod uint64
	ExecuteOperatorFeePeriod uint64
}, error) {
	return _SsvNetworkViews.Contract.GetOperatorFeePeriods(&_SsvNetworkViews.CallOpts)
}

// GetOperatorFeePeriods is a free data retrieval call binding the contract method 0xe6d2834d.
//
// Solidity: function getOperatorFeePeriods() view returns(uint64 declareOperatorFeePeriod, uint64 executeOperatorFeePeriod)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetOperatorFeePeriods() (struct {
	DeclareOperatorFeePeriod uint64
	ExecuteOperatorFeePeriod uint64
}, error) {
	return _SsvNetworkViews.Contract.GetOperatorFeePeriods(&_SsvNetworkViews.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0x3e2ec160.
//
// Solidity: function getValidator(address clusterOwner, bytes publicKey) view returns(bool active)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetValidator(opts *bind.CallOpts, clusterOwner common.Address, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getValidator", clusterOwner, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x3e2ec160.
//
// Solidity: function getValidator(address clusterOwner, bytes publicKey) view returns(bool active)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetValidator(clusterOwner common.Address, publicKey []byte) (bool, error) {
	return _SsvNetworkViews.Contract.GetValidator(&_SsvNetworkViews.CallOpts, clusterOwner, publicKey)
}

// GetValidator is a free data retrieval call binding the contract method 0x3e2ec160.
//
// Solidity: function getValidator(address clusterOwner, bytes publicKey) view returns(bool active)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetValidator(clusterOwner common.Address, publicKey []byte) (bool, error) {
	return _SsvNetworkViews.Contract.GetValidator(&_SsvNetworkViews.CallOpts, clusterOwner, publicKey)
}

// GetValidatorsPerOperatorLimit is a free data retrieval call binding the contract method 0x14cb9d7b.
//
// Solidity: function getValidatorsPerOperatorLimit() view returns(uint32)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetValidatorsPerOperatorLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getValidatorsPerOperatorLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetValidatorsPerOperatorLimit is a free data retrieval call binding the contract method 0x14cb9d7b.
//
// Solidity: function getValidatorsPerOperatorLimit() view returns(uint32)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetValidatorsPerOperatorLimit() (uint32, error) {
	return _SsvNetworkViews.Contract.GetValidatorsPerOperatorLimit(&_SsvNetworkViews.CallOpts)
}

// GetValidatorsPerOperatorLimit is a free data retrieval call binding the contract method 0x14cb9d7b.
//
// Solidity: function getValidatorsPerOperatorLimit() view returns(uint32)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetValidatorsPerOperatorLimit() (uint32, error) {
	return _SsvNetworkViews.Contract.GetValidatorsPerOperatorLimit(&_SsvNetworkViews.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string version)
func (_SsvNetworkViews *SsvNetworkViewsCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string version)
func (_SsvNetworkViews *SsvNetworkViewsSession) GetVersion() (string, error) {
	return _SsvNetworkViews.Contract.GetVersion(&_SsvNetworkViews.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string version)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) GetVersion() (string, error) {
	return _SsvNetworkViews.Contract.GetVersion(&_SsvNetworkViews.CallOpts)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x16cff008.
//
// Solidity: function isLiquidatable(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(bool)
func (_SsvNetworkViews *SsvNetworkViewsCaller) IsLiquidatable(opts *bind.CallOpts, clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (bool, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "isLiquidatable", clusterOwner, operatorIds, cluster)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x16cff008.
//
// Solidity: function isLiquidatable(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(bool)
func (_SsvNetworkViews *SsvNetworkViewsSession) IsLiquidatable(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (bool, error) {
	return _SsvNetworkViews.Contract.IsLiquidatable(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x16cff008.
//
// Solidity: function isLiquidatable(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(bool)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) IsLiquidatable(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (bool, error) {
	return _SsvNetworkViews.Contract.IsLiquidatable(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// IsLiquidated is a free data retrieval call binding the contract method 0xa694695b.
//
// Solidity: function isLiquidated(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(bool)
func (_SsvNetworkViews *SsvNetworkViewsCaller) IsLiquidated(opts *bind.CallOpts, clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (bool, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "isLiquidated", clusterOwner, operatorIds, cluster)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidated is a free data retrieval call binding the contract method 0xa694695b.
//
// Solidity: function isLiquidated(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(bool)
func (_SsvNetworkViews *SsvNetworkViewsSession) IsLiquidated(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (bool, error) {
	return _SsvNetworkViews.Contract.IsLiquidated(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// IsLiquidated is a free data retrieval call binding the contract method 0xa694695b.
//
// Solidity: function isLiquidated(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) view returns(bool)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) IsLiquidated(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (bool, error) {
	return _SsvNetworkViews.Contract.IsLiquidated(&_SsvNetworkViews.CallOpts, clusterOwner, operatorIds, cluster)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsSession) Owner() (common.Address, error) {
	return _SsvNetworkViews.Contract.Owner(&_SsvNetworkViews.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) Owner() (common.Address, error) {
	return _SsvNetworkViews.Contract.Owner(&_SsvNetworkViews.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsSession) PendingOwner() (common.Address, error) {
	return _SsvNetworkViews.Contract.PendingOwner(&_SsvNetworkViews.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) PendingOwner() (common.Address, error) {
	return _SsvNetworkViews.Contract.PendingOwner(&_SsvNetworkViews.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SsvNetworkViews *SsvNetworkViewsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SsvNetworkViews *SsvNetworkViewsSession) ProxiableUUID() ([32]byte, error) {
	return _SsvNetworkViews.Contract.ProxiableUUID(&_SsvNetworkViews.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SsvNetworkViews.Contract.ProxiableUUID(&_SsvNetworkViews.CallOpts)
}

// SsvNetwork is a free data retrieval call binding the contract method 0x10d04858.
//
// Solidity: function ssvNetwork() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsCaller) SsvNetwork(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsvNetworkViews.contract.Call(opts, &out, "ssvNetwork")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SsvNetwork is a free data retrieval call binding the contract method 0x10d04858.
//
// Solidity: function ssvNetwork() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsSession) SsvNetwork() (common.Address, error) {
	return _SsvNetworkViews.Contract.SsvNetwork(&_SsvNetworkViews.CallOpts)
}

// SsvNetwork is a free data retrieval call binding the contract method 0x10d04858.
//
// Solidity: function ssvNetwork() view returns(address)
func (_SsvNetworkViews *SsvNetworkViewsCallerSession) SsvNetwork() (common.Address, error) {
	return _SsvNetworkViews.Contract.SsvNetwork(&_SsvNetworkViews.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetworkViews.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SsvNetworkViews *SsvNetworkViewsSession) AcceptOwnership() (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.AcceptOwnership(&_SsvNetworkViews.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.AcceptOwnership(&_SsvNetworkViews.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ssvNetwork_) returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactor) Initialize(opts *bind.TransactOpts, ssvNetwork_ common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.contract.Transact(opts, "initialize", ssvNetwork_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ssvNetwork_) returns()
func (_SsvNetworkViews *SsvNetworkViewsSession) Initialize(ssvNetwork_ common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.Initialize(&_SsvNetworkViews.TransactOpts, ssvNetwork_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ssvNetwork_) returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactorSession) Initialize(ssvNetwork_ common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.Initialize(&_SsvNetworkViews.TransactOpts, ssvNetwork_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetworkViews.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsvNetworkViews *SsvNetworkViewsSession) RenounceOwnership() (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.RenounceOwnership(&_SsvNetworkViews.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.RenounceOwnership(&_SsvNetworkViews.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsvNetworkViews *SsvNetworkViewsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.TransferOwnership(&_SsvNetworkViews.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.TransferOwnership(&_SsvNetworkViews.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsvNetworkViews *SsvNetworkViewsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.UpgradeTo(&_SsvNetworkViews.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.UpgradeTo(&_SsvNetworkViews.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SsvNetworkViews.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SsvNetworkViews *SsvNetworkViewsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.UpgradeToAndCall(&_SsvNetworkViews.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SsvNetworkViews *SsvNetworkViewsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SsvNetworkViews.Contract.UpgradeToAndCall(&_SsvNetworkViews.TransactOpts, newImplementation, data)
}

// SsvNetworkViewsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SsvNetworkViews contract.
type SsvNetworkViewsAdminChangedIterator struct {
	Event *SsvNetworkViewsAdminChanged // Event containing the contract specifics and raw log

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
func (it *SsvNetworkViewsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkViewsAdminChanged)
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
		it.Event = new(SsvNetworkViewsAdminChanged)
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
func (it *SsvNetworkViewsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkViewsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkViewsAdminChanged represents a AdminChanged event raised by the SsvNetworkViews contract.
type SsvNetworkViewsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SsvNetworkViewsAdminChangedIterator, error) {

	logs, sub, err := _SsvNetworkViews.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsAdminChangedIterator{contract: _SsvNetworkViews.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SsvNetworkViewsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _SsvNetworkViews.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkViewsAdminChanged)
				if err := _SsvNetworkViews.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) ParseAdminChanged(log types.Log) (*SsvNetworkViewsAdminChanged, error) {
	event := new(SsvNetworkViewsAdminChanged)
	if err := _SsvNetworkViews.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkViewsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SsvNetworkViews contract.
type SsvNetworkViewsBeaconUpgradedIterator struct {
	Event *SsvNetworkViewsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SsvNetworkViewsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkViewsBeaconUpgraded)
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
		it.Event = new(SsvNetworkViewsBeaconUpgraded)
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
func (it *SsvNetworkViewsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkViewsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkViewsBeaconUpgraded represents a BeaconUpgraded event raised by the SsvNetworkViews contract.
type SsvNetworkViewsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SsvNetworkViewsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsBeaconUpgradedIterator{contract: _SsvNetworkViews.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SsvNetworkViewsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkViewsBeaconUpgraded)
				if err := _SsvNetworkViews.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) ParseBeaconUpgraded(log types.Log) (*SsvNetworkViewsBeaconUpgraded, error) {
	event := new(SsvNetworkViewsBeaconUpgraded)
	if err := _SsvNetworkViews.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkViewsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SsvNetworkViews contract.
type SsvNetworkViewsInitializedIterator struct {
	Event *SsvNetworkViewsInitialized // Event containing the contract specifics and raw log

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
func (it *SsvNetworkViewsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkViewsInitialized)
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
		it.Event = new(SsvNetworkViewsInitialized)
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
func (it *SsvNetworkViewsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkViewsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkViewsInitialized represents a Initialized event raised by the SsvNetworkViews contract.
type SsvNetworkViewsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) FilterInitialized(opts *bind.FilterOpts) (*SsvNetworkViewsInitializedIterator, error) {

	logs, sub, err := _SsvNetworkViews.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsInitializedIterator{contract: _SsvNetworkViews.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SsvNetworkViewsInitialized) (event.Subscription, error) {

	logs, sub, err := _SsvNetworkViews.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkViewsInitialized)
				if err := _SsvNetworkViews.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) ParseInitialized(log types.Log) (*SsvNetworkViewsInitialized, error) {
	event := new(SsvNetworkViewsInitialized)
	if err := _SsvNetworkViews.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkViewsOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SsvNetworkViews contract.
type SsvNetworkViewsOwnershipTransferStartedIterator struct {
	Event *SsvNetworkViewsOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SsvNetworkViewsOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkViewsOwnershipTransferStarted)
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
		it.Event = new(SsvNetworkViewsOwnershipTransferStarted)
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
func (it *SsvNetworkViewsOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkViewsOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkViewsOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SsvNetworkViews contract.
type SsvNetworkViewsOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvNetworkViewsOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsOwnershipTransferStartedIterator{contract: _SsvNetworkViews.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SsvNetworkViewsOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkViewsOwnershipTransferStarted)
				if err := _SsvNetworkViews.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) ParseOwnershipTransferStarted(log types.Log) (*SsvNetworkViewsOwnershipTransferStarted, error) {
	event := new(SsvNetworkViewsOwnershipTransferStarted)
	if err := _SsvNetworkViews.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkViewsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SsvNetworkViews contract.
type SsvNetworkViewsOwnershipTransferredIterator struct {
	Event *SsvNetworkViewsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SsvNetworkViewsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkViewsOwnershipTransferred)
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
		it.Event = new(SsvNetworkViewsOwnershipTransferred)
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
func (it *SsvNetworkViewsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkViewsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkViewsOwnershipTransferred represents a OwnershipTransferred event raised by the SsvNetworkViews contract.
type SsvNetworkViewsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvNetworkViewsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsOwnershipTransferredIterator{contract: _SsvNetworkViews.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SsvNetworkViewsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkViewsOwnershipTransferred)
				if err := _SsvNetworkViews.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) ParseOwnershipTransferred(log types.Log) (*SsvNetworkViewsOwnershipTransferred, error) {
	event := new(SsvNetworkViewsOwnershipTransferred)
	if err := _SsvNetworkViews.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkViewsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SsvNetworkViews contract.
type SsvNetworkViewsUpgradedIterator struct {
	Event *SsvNetworkViewsUpgraded // Event containing the contract specifics and raw log

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
func (it *SsvNetworkViewsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkViewsUpgraded)
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
		it.Event = new(SsvNetworkViewsUpgraded)
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
func (it *SsvNetworkViewsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkViewsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkViewsUpgraded represents a Upgraded event raised by the SsvNetworkViews contract.
type SsvNetworkViewsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SsvNetworkViewsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkViewsUpgradedIterator{contract: _SsvNetworkViews.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SsvNetworkViewsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SsvNetworkViews.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkViewsUpgraded)
				if err := _SsvNetworkViews.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsvNetworkViews *SsvNetworkViewsFilterer) ParseUpgraded(log types.Log) (*SsvNetworkViewsUpgraded, error) {
	event := new(SsvNetworkViewsUpgraded)
	if err := _SsvNetworkViews.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
