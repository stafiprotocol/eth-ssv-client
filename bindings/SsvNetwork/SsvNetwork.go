// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssv_network

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

// SsvNetworkMetaData contains all meta data concerning the SsvNetwork contract.
var SsvNetworkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalNotWithinTimeframe\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterDoesNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterIsLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterNotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedValidatorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsIncreaseLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeIncreaseNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectClusterState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectValidatorState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIdsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxValueExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewBlockPeriodIsBelowMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeDeclared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorsListNotUnique\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameFeeChangeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetModuleDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsortedOperatorsList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterReactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"DeclareOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"ExecuteOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"FeeRecipientAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"LiquidationThresholdPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinimumLiquidationCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"NetworkEarningsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"NetworkFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeDeclarationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeIncreaseLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"OperatorWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"cancelDeclaredOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"declareOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"executeOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getRegisterAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"authOperators\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"authValidators\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"contractISSVOperators\",\"name\":\"ssvOperators_\",\"type\":\"address\"},{\"internalType\":\"contractISSVClusters\",\"name\":\"ssvClusters_\",\"type\":\"address\"},{\"internalType\":\"contractISSVDAO\",\"name\":\"ssvDAO_\",\"type\":\"address\"},{\"internalType\":\"contractISSVViews\",\"name\":\"ssvViews_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"minimumBlocksBeforeLiquidation_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidationCollateral_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validatorsPerOperatorLimit_\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"declareOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executeOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"operatorMaxFeeIncrease_\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"reactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"reduceOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes\",\"name\":\"sharesData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"setFeeRecipientAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"setOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authOperator\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"authValidator\",\"type\":\"bool\"}],\"name\":\"setRegisterAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timeInSeconds\",\"type\":\"uint64\"}],\"name\":\"updateDeclareOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timeInSeconds\",\"type\":\"uint64\"}],\"name\":\"updateExecuteOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blocks\",\"type\":\"uint64\"}],\"name\":\"updateLiquidationThresholdPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinimumLiquidationCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSSVModules\",\"name\":\"moduleId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"updateModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateNetworkFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"percentage\",\"type\":\"uint64\"}],\"name\":\"updateOperatorFeeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"withdrawAllOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNetworkEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SsvNetworkABI is the input ABI used to generate the binding from.
// Deprecated: Use SsvNetworkMetaData.ABI instead.
var SsvNetworkABI = SsvNetworkMetaData.ABI

// SsvNetwork is an auto generated Go binding around an Ethereum contract.
type SsvNetwork struct {
	SsvNetworkCaller     // Read-only binding to the contract
	SsvNetworkTransactor // Write-only binding to the contract
	SsvNetworkFilterer   // Log filterer for contract events
}

// SsvNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsvNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsvNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsvNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsvNetworkSession struct {
	Contract     *SsvNetwork       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsvNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsvNetworkCallerSession struct {
	Contract *SsvNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SsvNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsvNetworkTransactorSession struct {
	Contract     *SsvNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SsvNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsvNetworkRaw struct {
	Contract *SsvNetwork // Generic contract binding to access the raw methods on
}

// SsvNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsvNetworkCallerRaw struct {
	Contract *SsvNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// SsvNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsvNetworkTransactorRaw struct {
	Contract *SsvNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsvNetwork creates a new instance of SsvNetwork, bound to a specific deployed contract.
func NewSsvNetwork(address common.Address, backend bind.ContractBackend) (*SsvNetwork, error) {
	contract, err := bindSsvNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SsvNetwork{SsvNetworkCaller: SsvNetworkCaller{contract: contract}, SsvNetworkTransactor: SsvNetworkTransactor{contract: contract}, SsvNetworkFilterer: SsvNetworkFilterer{contract: contract}}, nil
}

// NewSsvNetworkCaller creates a new read-only instance of SsvNetwork, bound to a specific deployed contract.
func NewSsvNetworkCaller(address common.Address, caller bind.ContractCaller) (*SsvNetworkCaller, error) {
	contract, err := bindSsvNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkCaller{contract: contract}, nil
}

// NewSsvNetworkTransactor creates a new write-only instance of SsvNetwork, bound to a specific deployed contract.
func NewSsvNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*SsvNetworkTransactor, error) {
	contract, err := bindSsvNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkTransactor{contract: contract}, nil
}

// NewSsvNetworkFilterer creates a new log filterer instance of SsvNetwork, bound to a specific deployed contract.
func NewSsvNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*SsvNetworkFilterer, error) {
	contract, err := bindSsvNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkFilterer{contract: contract}, nil
}

// bindSsvNetwork binds a generic wrapper to an already deployed contract.
func bindSsvNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SsvNetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsvNetwork *SsvNetworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsvNetwork.Contract.SsvNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsvNetwork *SsvNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SsvNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsvNetwork *SsvNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SsvNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsvNetwork *SsvNetworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsvNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsvNetwork *SsvNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsvNetwork *SsvNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsvNetwork.Contract.contract.Transact(opts, method, params...)
}

// GetRegisterAuth is a free data retrieval call binding the contract method 0x7398ca6c.
//
// Solidity: function getRegisterAuth(address userAddress) view returns(bool authOperators, bool authValidators)
func (_SsvNetwork *SsvNetworkCaller) GetRegisterAuth(opts *bind.CallOpts, userAddress common.Address) (struct {
	AuthOperators  bool
	AuthValidators bool
}, error) {
	var out []interface{}
	err := _SsvNetwork.contract.Call(opts, &out, "getRegisterAuth", userAddress)

	outstruct := new(struct {
		AuthOperators  bool
		AuthValidators bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AuthOperators = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AuthValidators = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetRegisterAuth is a free data retrieval call binding the contract method 0x7398ca6c.
//
// Solidity: function getRegisterAuth(address userAddress) view returns(bool authOperators, bool authValidators)
func (_SsvNetwork *SsvNetworkSession) GetRegisterAuth(userAddress common.Address) (struct {
	AuthOperators  bool
	AuthValidators bool
}, error) {
	return _SsvNetwork.Contract.GetRegisterAuth(&_SsvNetwork.CallOpts, userAddress)
}

// GetRegisterAuth is a free data retrieval call binding the contract method 0x7398ca6c.
//
// Solidity: function getRegisterAuth(address userAddress) view returns(bool authOperators, bool authValidators)
func (_SsvNetwork *SsvNetworkCallerSession) GetRegisterAuth(userAddress common.Address) (struct {
	AuthOperators  bool
	AuthValidators bool
}, error) {
	return _SsvNetwork.Contract.GetRegisterAuth(&_SsvNetwork.CallOpts, userAddress)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_SsvNetwork *SsvNetworkCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SsvNetwork.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_SsvNetwork *SsvNetworkSession) GetVersion() (string, error) {
	return _SsvNetwork.Contract.GetVersion(&_SsvNetwork.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_SsvNetwork *SsvNetworkCallerSession) GetVersion() (string, error) {
	return _SsvNetwork.Contract.GetVersion(&_SsvNetwork.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsvNetwork *SsvNetworkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsvNetwork.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsvNetwork *SsvNetworkSession) Owner() (common.Address, error) {
	return _SsvNetwork.Contract.Owner(&_SsvNetwork.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsvNetwork *SsvNetworkCallerSession) Owner() (common.Address, error) {
	return _SsvNetwork.Contract.Owner(&_SsvNetwork.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SsvNetwork *SsvNetworkCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsvNetwork.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SsvNetwork *SsvNetworkSession) PendingOwner() (common.Address, error) {
	return _SsvNetwork.Contract.PendingOwner(&_SsvNetwork.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SsvNetwork *SsvNetworkCallerSession) PendingOwner() (common.Address, error) {
	return _SsvNetwork.Contract.PendingOwner(&_SsvNetwork.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SsvNetwork *SsvNetworkCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SsvNetwork.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SsvNetwork *SsvNetworkSession) ProxiableUUID() ([32]byte, error) {
	return _SsvNetwork.Contract.ProxiableUUID(&_SsvNetwork.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SsvNetwork *SsvNetworkCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SsvNetwork.Contract.ProxiableUUID(&_SsvNetwork.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SsvNetwork *SsvNetworkTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SsvNetwork *SsvNetworkSession) AcceptOwnership() (*types.Transaction, error) {
	return _SsvNetwork.Contract.AcceptOwnership(&_SsvNetwork.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SsvNetwork *SsvNetworkTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SsvNetwork.Contract.AcceptOwnership(&_SsvNetwork.TransactOpts)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactor) CancelDeclaredOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "cancelDeclaredOperatorFee", operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.CancelDeclaredOperatorFee(&_SsvNetwork.TransactOpts, operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.CancelDeclaredOperatorFee(&_SsvNetwork.TransactOpts, operatorId)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SsvNetwork *SsvNetworkTransactor) DeclareOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "declareOperatorFee", operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SsvNetwork *SsvNetworkSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.DeclareOperatorFee(&_SsvNetwork.TransactOpts, operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.DeclareOperatorFee(&_SsvNetwork.TransactOpts, operatorId, fee)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address clusterOwner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactor) Deposit(opts *bind.TransactOpts, clusterOwner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "deposit", clusterOwner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address clusterOwner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkSession) Deposit(clusterOwner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Deposit(&_SsvNetwork.TransactOpts, clusterOwner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address clusterOwner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) Deposit(clusterOwner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Deposit(&_SsvNetwork.TransactOpts, clusterOwner, operatorIds, amount, cluster)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactor) ExecuteOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "executeOperatorFee", operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.ExecuteOperatorFee(&_SsvNetwork.TransactOpts, operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.ExecuteOperatorFee(&_SsvNetwork.TransactOpts, operatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_SsvNetwork *SsvNetworkTransactor) Initialize(opts *bind.TransactOpts, token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "initialize", token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_SsvNetwork *SsvNetworkSession) Initialize(token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Initialize(&_SsvNetwork.TransactOpts, token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) Initialize(token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Initialize(&_SsvNetwork.TransactOpts, token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactor) Liquidate(opts *bind.TransactOpts, clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "liquidate", clusterOwner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkSession) Liquidate(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Liquidate(&_SsvNetwork.TransactOpts, clusterOwner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) Liquidate(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Liquidate(&_SsvNetwork.TransactOpts, clusterOwner, operatorIds, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactor) Reactivate(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "reactivate", operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Reactivate(&_SsvNetwork.TransactOpts, operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Reactivate(&_SsvNetwork.TransactOpts, operatorIds, amount, cluster)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SsvNetwork *SsvNetworkTransactor) ReduceOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "reduceOperatorFee", operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SsvNetwork *SsvNetworkSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.ReduceOperatorFee(&_SsvNetwork.TransactOpts, operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.ReduceOperatorFee(&_SsvNetwork.TransactOpts, operatorId, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_SsvNetwork *SsvNetworkTransactor) RegisterOperator(opts *bind.TransactOpts, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "registerOperator", publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_SsvNetwork *SsvNetworkSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RegisterOperator(&_SsvNetwork.TransactOpts, publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_SsvNetwork *SsvNetworkTransactorSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RegisterOperator(&_SsvNetwork.TransactOpts, publicKey, fee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactor) RegisterValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "registerValidator", publicKey, operatorIds, sharesData, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkSession) RegisterValidator(publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RegisterValidator(&_SsvNetwork.TransactOpts, publicKey, operatorIds, sharesData, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) RegisterValidator(publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RegisterValidator(&_SsvNetwork.TransactOpts, publicKey, operatorIds, sharesData, amount, cluster)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactor) RemoveOperator(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "removeOperator", operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RemoveOperator(&_SsvNetwork.TransactOpts, operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RemoveOperator(&_SsvNetwork.TransactOpts, operatorId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactor) RemoveValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "removeValidator", publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RemoveValidator(&_SsvNetwork.TransactOpts, publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.RemoveValidator(&_SsvNetwork.TransactOpts, publicKey, operatorIds, cluster)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsvNetwork *SsvNetworkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsvNetwork *SsvNetworkSession) RenounceOwnership() (*types.Transaction, error) {
	return _SsvNetwork.Contract.RenounceOwnership(&_SsvNetwork.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsvNetwork *SsvNetworkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SsvNetwork.Contract.RenounceOwnership(&_SsvNetwork.TransactOpts)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_SsvNetwork *SsvNetworkTransactor) SetFeeRecipientAddress(opts *bind.TransactOpts, recipientAddress common.Address) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "setFeeRecipientAddress", recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_SsvNetwork *SsvNetworkSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SetFeeRecipientAddress(&_SsvNetwork.TransactOpts, recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SetFeeRecipientAddress(&_SsvNetwork.TransactOpts, recipientAddress)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_SsvNetwork *SsvNetworkTransactor) SetOperatorWhitelist(opts *bind.TransactOpts, operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "setOperatorWhitelist", operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_SsvNetwork *SsvNetworkSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SetOperatorWhitelist(&_SsvNetwork.TransactOpts, operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SetOperatorWhitelist(&_SsvNetwork.TransactOpts, operatorId, whitelisted)
}

// SetRegisterAuth is a paid mutator transaction binding the contract method 0x3ed00469.
//
// Solidity: function setRegisterAuth(address userAddress, bool authOperator, bool authValidator) returns()
func (_SsvNetwork *SsvNetworkTransactor) SetRegisterAuth(opts *bind.TransactOpts, userAddress common.Address, authOperator bool, authValidator bool) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "setRegisterAuth", userAddress, authOperator, authValidator)
}

// SetRegisterAuth is a paid mutator transaction binding the contract method 0x3ed00469.
//
// Solidity: function setRegisterAuth(address userAddress, bool authOperator, bool authValidator) returns()
func (_SsvNetwork *SsvNetworkSession) SetRegisterAuth(userAddress common.Address, authOperator bool, authValidator bool) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SetRegisterAuth(&_SsvNetwork.TransactOpts, userAddress, authOperator, authValidator)
}

// SetRegisterAuth is a paid mutator transaction binding the contract method 0x3ed00469.
//
// Solidity: function setRegisterAuth(address userAddress, bool authOperator, bool authValidator) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) SetRegisterAuth(userAddress common.Address, authOperator bool, authValidator bool) (*types.Transaction, error) {
	return _SsvNetwork.Contract.SetRegisterAuth(&_SsvNetwork.TransactOpts, userAddress, authOperator, authValidator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsvNetwork *SsvNetworkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsvNetwork *SsvNetworkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.TransferOwnership(&_SsvNetwork.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.TransferOwnership(&_SsvNetwork.TransactOpts, newOwner)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateDeclareOperatorFeePeriod(opts *bind.TransactOpts, timeInSeconds uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateDeclareOperatorFeePeriod", timeInSeconds)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateDeclareOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateDeclareOperatorFeePeriod(&_SsvNetwork.TransactOpts, timeInSeconds)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateDeclareOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateDeclareOperatorFeePeriod(&_SsvNetwork.TransactOpts, timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateExecuteOperatorFeePeriod(opts *bind.TransactOpts, timeInSeconds uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateExecuteOperatorFeePeriod", timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateExecuteOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateExecuteOperatorFeePeriod(&_SsvNetwork.TransactOpts, timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateExecuteOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateExecuteOperatorFeePeriod(&_SsvNetwork.TransactOpts, timeInSeconds)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateLiquidationThresholdPeriod(opts *bind.TransactOpts, blocks uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateLiquidationThresholdPeriod", blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateLiquidationThresholdPeriod(&_SsvNetwork.TransactOpts, blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateLiquidationThresholdPeriod(&_SsvNetwork.TransactOpts, blocks)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateMinimumLiquidationCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateMinimumLiquidationCollateral", amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateMinimumLiquidationCollateral(&_SsvNetwork.TransactOpts, amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateMinimumLiquidationCollateral(&_SsvNetwork.TransactOpts, amount)
}

// UpdateModule is a paid mutator transaction binding the contract method 0xe3e324b0.
//
// Solidity: function updateModule(uint8 moduleId, address moduleAddress) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateModule(opts *bind.TransactOpts, moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateModule", moduleId, moduleAddress)
}

// UpdateModule is a paid mutator transaction binding the contract method 0xe3e324b0.
//
// Solidity: function updateModule(uint8 moduleId, address moduleAddress) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateModule(moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateModule(&_SsvNetwork.TransactOpts, moduleId, moduleAddress)
}

// UpdateModule is a paid mutator transaction binding the contract method 0xe3e324b0.
//
// Solidity: function updateModule(uint8 moduleId, address moduleAddress) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateModule(moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateModule(&_SsvNetwork.TransactOpts, moduleId, moduleAddress)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateNetworkFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateNetworkFee", fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateNetworkFee(&_SsvNetwork.TransactOpts, fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateNetworkFee(&_SsvNetwork.TransactOpts, fee)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpdateOperatorFeeIncreaseLimit(opts *bind.TransactOpts, percentage uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "updateOperatorFeeIncreaseLimit", percentage)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_SsvNetwork *SsvNetworkSession) UpdateOperatorFeeIncreaseLimit(percentage uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateOperatorFeeIncreaseLimit(&_SsvNetwork.TransactOpts, percentage)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpdateOperatorFeeIncreaseLimit(percentage uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpdateOperatorFeeIncreaseLimit(&_SsvNetwork.TransactOpts, percentage)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsvNetwork *SsvNetworkTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsvNetwork *SsvNetworkSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpgradeTo(&_SsvNetwork.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpgradeTo(&_SsvNetwork.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SsvNetwork *SsvNetworkTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SsvNetwork *SsvNetworkSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpgradeToAndCall(&_SsvNetwork.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SsvNetwork *SsvNetworkTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SsvNetwork.Contract.UpgradeToAndCall(&_SsvNetwork.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactor) Withdraw(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "withdraw", operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Withdraw(&_SsvNetwork.TransactOpts, operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Withdraw(&_SsvNetwork.TransactOpts, operatorIds, amount, cluster)
}

// WithdrawAllOperatorEarnings is a paid mutator transaction binding the contract method 0x4bc93b64.
//
// Solidity: function withdrawAllOperatorEarnings(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactor) WithdrawAllOperatorEarnings(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "withdrawAllOperatorEarnings", operatorId)
}

// WithdrawAllOperatorEarnings is a paid mutator transaction binding the contract method 0x4bc93b64.
//
// Solidity: function withdrawAllOperatorEarnings(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkSession) WithdrawAllOperatorEarnings(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.WithdrawAllOperatorEarnings(&_SsvNetwork.TransactOpts, operatorId)
}

// WithdrawAllOperatorEarnings is a paid mutator transaction binding the contract method 0x4bc93b64.
//
// Solidity: function withdrawAllOperatorEarnings(uint64 operatorId) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) WithdrawAllOperatorEarnings(operatorId uint64) (*types.Transaction, error) {
	return _SsvNetwork.Contract.WithdrawAllOperatorEarnings(&_SsvNetwork.TransactOpts, operatorId)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_SsvNetwork *SsvNetworkTransactor) WithdrawNetworkEarnings(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "withdrawNetworkEarnings", amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_SsvNetwork *SsvNetworkSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.WithdrawNetworkEarnings(&_SsvNetwork.TransactOpts, amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.WithdrawNetworkEarnings(&_SsvNetwork.TransactOpts, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_SsvNetwork *SsvNetworkTransactor) WithdrawOperatorEarnings(opts *bind.TransactOpts, operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.contract.Transact(opts, "withdrawOperatorEarnings", operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_SsvNetwork *SsvNetworkSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.WithdrawOperatorEarnings(&_SsvNetwork.TransactOpts, operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_SsvNetwork *SsvNetworkTransactorSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _SsvNetwork.Contract.WithdrawOperatorEarnings(&_SsvNetwork.TransactOpts, operatorId, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SsvNetwork *SsvNetworkTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SsvNetwork.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SsvNetwork *SsvNetworkSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Fallback(&_SsvNetwork.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SsvNetwork *SsvNetworkTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SsvNetwork.Contract.Fallback(&_SsvNetwork.TransactOpts, calldata)
}

// SsvNetworkAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SsvNetwork contract.
type SsvNetworkAdminChangedIterator struct {
	Event *SsvNetworkAdminChanged // Event containing the contract specifics and raw log

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
func (it *SsvNetworkAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkAdminChanged)
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
		it.Event = new(SsvNetworkAdminChanged)
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
func (it *SsvNetworkAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkAdminChanged represents a AdminChanged event raised by the SsvNetwork contract.
type SsvNetworkAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SsvNetwork *SsvNetworkFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SsvNetworkAdminChangedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkAdminChangedIterator{contract: _SsvNetwork.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SsvNetwork *SsvNetworkFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SsvNetworkAdminChanged) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkAdminChanged)
				if err := _SsvNetwork.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_SsvNetwork *SsvNetworkFilterer) ParseAdminChanged(log types.Log) (*SsvNetworkAdminChanged, error) {
	event := new(SsvNetworkAdminChanged)
	if err := _SsvNetwork.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SsvNetwork contract.
type SsvNetworkBeaconUpgradedIterator struct {
	Event *SsvNetworkBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SsvNetworkBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkBeaconUpgraded)
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
		it.Event = new(SsvNetworkBeaconUpgraded)
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
func (it *SsvNetworkBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkBeaconUpgraded represents a BeaconUpgraded event raised by the SsvNetwork contract.
type SsvNetworkBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SsvNetwork *SsvNetworkFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SsvNetworkBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkBeaconUpgradedIterator{contract: _SsvNetwork.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SsvNetwork *SsvNetworkFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SsvNetworkBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkBeaconUpgraded)
				if err := _SsvNetwork.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_SsvNetwork *SsvNetworkFilterer) ParseBeaconUpgraded(log types.Log) (*SsvNetworkBeaconUpgraded, error) {
	event := new(SsvNetworkBeaconUpgraded)
	if err := _SsvNetwork.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkClusterDepositedIterator is returned from FilterClusterDeposited and is used to iterate over the raw logs and unpacked data for ClusterDeposited events raised by the SsvNetwork contract.
type SsvNetworkClusterDepositedIterator struct {
	Event *SsvNetworkClusterDeposited // Event containing the contract specifics and raw log

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
func (it *SsvNetworkClusterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkClusterDeposited)
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
		it.Event = new(SsvNetworkClusterDeposited)
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
func (it *SsvNetworkClusterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkClusterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkClusterDeposited represents a ClusterDeposited event raised by the SsvNetwork contract.
type SsvNetworkClusterDeposited struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterDeposited is a free log retrieval operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) FilterClusterDeposited(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkClusterDepositedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkClusterDepositedIterator{contract: _SsvNetwork.contract, event: "ClusterDeposited", logs: logs, sub: sub}, nil
}

// WatchClusterDeposited is a free log subscription operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) WatchClusterDeposited(opts *bind.WatchOpts, sink chan<- *SsvNetworkClusterDeposited, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkClusterDeposited)
				if err := _SsvNetwork.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
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

// ParseClusterDeposited is a log parse operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) ParseClusterDeposited(log types.Log) (*SsvNetworkClusterDeposited, error) {
	event := new(SsvNetworkClusterDeposited)
	if err := _SsvNetwork.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkClusterLiquidatedIterator is returned from FilterClusterLiquidated and is used to iterate over the raw logs and unpacked data for ClusterLiquidated events raised by the SsvNetwork contract.
type SsvNetworkClusterLiquidatedIterator struct {
	Event *SsvNetworkClusterLiquidated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkClusterLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkClusterLiquidated)
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
		it.Event = new(SsvNetworkClusterLiquidated)
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
func (it *SsvNetworkClusterLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkClusterLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkClusterLiquidated represents a ClusterLiquidated event raised by the SsvNetwork contract.
type SsvNetworkClusterLiquidated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterLiquidated is a free log retrieval operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) FilterClusterLiquidated(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkClusterLiquidatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkClusterLiquidatedIterator{contract: _SsvNetwork.contract, event: "ClusterLiquidated", logs: logs, sub: sub}, nil
}

// WatchClusterLiquidated is a free log subscription operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) WatchClusterLiquidated(opts *bind.WatchOpts, sink chan<- *SsvNetworkClusterLiquidated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkClusterLiquidated)
				if err := _SsvNetwork.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
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

// ParseClusterLiquidated is a log parse operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) ParseClusterLiquidated(log types.Log) (*SsvNetworkClusterLiquidated, error) {
	event := new(SsvNetworkClusterLiquidated)
	if err := _SsvNetwork.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkClusterReactivatedIterator is returned from FilterClusterReactivated and is used to iterate over the raw logs and unpacked data for ClusterReactivated events raised by the SsvNetwork contract.
type SsvNetworkClusterReactivatedIterator struct {
	Event *SsvNetworkClusterReactivated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkClusterReactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkClusterReactivated)
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
		it.Event = new(SsvNetworkClusterReactivated)
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
func (it *SsvNetworkClusterReactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkClusterReactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkClusterReactivated represents a ClusterReactivated event raised by the SsvNetwork contract.
type SsvNetworkClusterReactivated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterReactivated is a free log retrieval operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) FilterClusterReactivated(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkClusterReactivatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkClusterReactivatedIterator{contract: _SsvNetwork.contract, event: "ClusterReactivated", logs: logs, sub: sub}, nil
}

// WatchClusterReactivated is a free log subscription operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) WatchClusterReactivated(opts *bind.WatchOpts, sink chan<- *SsvNetworkClusterReactivated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkClusterReactivated)
				if err := _SsvNetwork.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
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

// ParseClusterReactivated is a log parse operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) ParseClusterReactivated(log types.Log) (*SsvNetworkClusterReactivated, error) {
	event := new(SsvNetworkClusterReactivated)
	if err := _SsvNetwork.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkClusterWithdrawnIterator is returned from FilterClusterWithdrawn and is used to iterate over the raw logs and unpacked data for ClusterWithdrawn events raised by the SsvNetwork contract.
type SsvNetworkClusterWithdrawnIterator struct {
	Event *SsvNetworkClusterWithdrawn // Event containing the contract specifics and raw log

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
func (it *SsvNetworkClusterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkClusterWithdrawn)
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
		it.Event = new(SsvNetworkClusterWithdrawn)
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
func (it *SsvNetworkClusterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkClusterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkClusterWithdrawn represents a ClusterWithdrawn event raised by the SsvNetwork contract.
type SsvNetworkClusterWithdrawn struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterWithdrawn is a free log retrieval operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) FilterClusterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkClusterWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkClusterWithdrawnIterator{contract: _SsvNetwork.contract, event: "ClusterWithdrawn", logs: logs, sub: sub}, nil
}

// WatchClusterWithdrawn is a free log subscription operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) WatchClusterWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvNetworkClusterWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkClusterWithdrawn)
				if err := _SsvNetwork.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
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

// ParseClusterWithdrawn is a log parse operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) ParseClusterWithdrawn(log types.Log) (*SsvNetworkClusterWithdrawn, error) {
	event := new(SsvNetworkClusterWithdrawn)
	if err := _SsvNetwork.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkDeclareOperatorFeePeriodUpdatedIterator is returned from FilterDeclareOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for DeclareOperatorFeePeriodUpdated events raised by the SsvNetwork contract.
type SsvNetworkDeclareOperatorFeePeriodUpdatedIterator struct {
	Event *SsvNetworkDeclareOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkDeclareOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkDeclareOperatorFeePeriodUpdated)
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
		it.Event = new(SsvNetworkDeclareOperatorFeePeriodUpdated)
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
func (it *SsvNetworkDeclareOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkDeclareOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkDeclareOperatorFeePeriodUpdated represents a DeclareOperatorFeePeriodUpdated event raised by the SsvNetwork contract.
type SsvNetworkDeclareOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeclareOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) FilterDeclareOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SsvNetworkDeclareOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkDeclareOperatorFeePeriodUpdatedIterator{contract: _SsvNetwork.contract, event: "DeclareOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDeclareOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) WatchDeclareOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkDeclareOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkDeclareOperatorFeePeriodUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
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

// ParseDeclareOperatorFeePeriodUpdated is a log parse operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) ParseDeclareOperatorFeePeriodUpdated(log types.Log) (*SsvNetworkDeclareOperatorFeePeriodUpdated, error) {
	event := new(SsvNetworkDeclareOperatorFeePeriodUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkExecuteOperatorFeePeriodUpdatedIterator is returned from FilterExecuteOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for ExecuteOperatorFeePeriodUpdated events raised by the SsvNetwork contract.
type SsvNetworkExecuteOperatorFeePeriodUpdatedIterator struct {
	Event *SsvNetworkExecuteOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkExecuteOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkExecuteOperatorFeePeriodUpdated)
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
		it.Event = new(SsvNetworkExecuteOperatorFeePeriodUpdated)
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
func (it *SsvNetworkExecuteOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkExecuteOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkExecuteOperatorFeePeriodUpdated represents a ExecuteOperatorFeePeriodUpdated event raised by the SsvNetwork contract.
type SsvNetworkExecuteOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecuteOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) FilterExecuteOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SsvNetworkExecuteOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkExecuteOperatorFeePeriodUpdatedIterator{contract: _SsvNetwork.contract, event: "ExecuteOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchExecuteOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) WatchExecuteOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkExecuteOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkExecuteOperatorFeePeriodUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
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

// ParseExecuteOperatorFeePeriodUpdated is a log parse operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) ParseExecuteOperatorFeePeriodUpdated(log types.Log) (*SsvNetworkExecuteOperatorFeePeriodUpdated, error) {
	event := new(SsvNetworkExecuteOperatorFeePeriodUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkFeeRecipientAddressUpdatedIterator is returned from FilterFeeRecipientAddressUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientAddressUpdated events raised by the SsvNetwork contract.
type SsvNetworkFeeRecipientAddressUpdatedIterator struct {
	Event *SsvNetworkFeeRecipientAddressUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkFeeRecipientAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkFeeRecipientAddressUpdated)
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
		it.Event = new(SsvNetworkFeeRecipientAddressUpdated)
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
func (it *SsvNetworkFeeRecipientAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkFeeRecipientAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkFeeRecipientAddressUpdated represents a FeeRecipientAddressUpdated event raised by the SsvNetwork contract.
type SsvNetworkFeeRecipientAddressUpdated struct {
	Owner            common.Address
	RecipientAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientAddressUpdated is a free log retrieval operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_SsvNetwork *SsvNetworkFilterer) FilterFeeRecipientAddressUpdated(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkFeeRecipientAddressUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkFeeRecipientAddressUpdatedIterator{contract: _SsvNetwork.contract, event: "FeeRecipientAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientAddressUpdated is a free log subscription operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_SsvNetwork *SsvNetworkFilterer) WatchFeeRecipientAddressUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkFeeRecipientAddressUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkFeeRecipientAddressUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
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

// ParseFeeRecipientAddressUpdated is a log parse operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_SsvNetwork *SsvNetworkFilterer) ParseFeeRecipientAddressUpdated(log types.Log) (*SsvNetworkFeeRecipientAddressUpdated, error) {
	event := new(SsvNetworkFeeRecipientAddressUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SsvNetwork contract.
type SsvNetworkInitializedIterator struct {
	Event *SsvNetworkInitialized // Event containing the contract specifics and raw log

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
func (it *SsvNetworkInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkInitialized)
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
		it.Event = new(SsvNetworkInitialized)
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
func (it *SsvNetworkInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkInitialized represents a Initialized event raised by the SsvNetwork contract.
type SsvNetworkInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsvNetwork *SsvNetworkFilterer) FilterInitialized(opts *bind.FilterOpts) (*SsvNetworkInitializedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkInitializedIterator{contract: _SsvNetwork.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsvNetwork *SsvNetworkFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SsvNetworkInitialized) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkInitialized)
				if err := _SsvNetwork.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SsvNetwork *SsvNetworkFilterer) ParseInitialized(log types.Log) (*SsvNetworkInitialized, error) {
	event := new(SsvNetworkInitialized)
	if err := _SsvNetwork.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkLiquidationThresholdPeriodUpdatedIterator is returned from FilterLiquidationThresholdPeriodUpdated and is used to iterate over the raw logs and unpacked data for LiquidationThresholdPeriodUpdated events raised by the SsvNetwork contract.
type SsvNetworkLiquidationThresholdPeriodUpdatedIterator struct {
	Event *SsvNetworkLiquidationThresholdPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkLiquidationThresholdPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkLiquidationThresholdPeriodUpdated)
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
		it.Event = new(SsvNetworkLiquidationThresholdPeriodUpdated)
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
func (it *SsvNetworkLiquidationThresholdPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkLiquidationThresholdPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkLiquidationThresholdPeriodUpdated represents a LiquidationThresholdPeriodUpdated event raised by the SsvNetwork contract.
type SsvNetworkLiquidationThresholdPeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidationThresholdPeriodUpdated is a free log retrieval operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) FilterLiquidationThresholdPeriodUpdated(opts *bind.FilterOpts) (*SsvNetworkLiquidationThresholdPeriodUpdatedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkLiquidationThresholdPeriodUpdatedIterator{contract: _SsvNetwork.contract, event: "LiquidationThresholdPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidationThresholdPeriodUpdated is a free log subscription operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) WatchLiquidationThresholdPeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkLiquidationThresholdPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkLiquidationThresholdPeriodUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
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

// ParseLiquidationThresholdPeriodUpdated is a log parse operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) ParseLiquidationThresholdPeriodUpdated(log types.Log) (*SsvNetworkLiquidationThresholdPeriodUpdated, error) {
	event := new(SsvNetworkLiquidationThresholdPeriodUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkMinimumLiquidationCollateralUpdatedIterator is returned from FilterMinimumLiquidationCollateralUpdated and is used to iterate over the raw logs and unpacked data for MinimumLiquidationCollateralUpdated events raised by the SsvNetwork contract.
type SsvNetworkMinimumLiquidationCollateralUpdatedIterator struct {
	Event *SsvNetworkMinimumLiquidationCollateralUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkMinimumLiquidationCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkMinimumLiquidationCollateralUpdated)
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
		it.Event = new(SsvNetworkMinimumLiquidationCollateralUpdated)
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
func (it *SsvNetworkMinimumLiquidationCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkMinimumLiquidationCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkMinimumLiquidationCollateralUpdated represents a MinimumLiquidationCollateralUpdated event raised by the SsvNetwork contract.
type SsvNetworkMinimumLiquidationCollateralUpdated struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinimumLiquidationCollateralUpdated is a free log retrieval operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_SsvNetwork *SsvNetworkFilterer) FilterMinimumLiquidationCollateralUpdated(opts *bind.FilterOpts) (*SsvNetworkMinimumLiquidationCollateralUpdatedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkMinimumLiquidationCollateralUpdatedIterator{contract: _SsvNetwork.contract, event: "MinimumLiquidationCollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumLiquidationCollateralUpdated is a free log subscription operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_SsvNetwork *SsvNetworkFilterer) WatchMinimumLiquidationCollateralUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkMinimumLiquidationCollateralUpdated) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkMinimumLiquidationCollateralUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
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

// ParseMinimumLiquidationCollateralUpdated is a log parse operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_SsvNetwork *SsvNetworkFilterer) ParseMinimumLiquidationCollateralUpdated(log types.Log) (*SsvNetworkMinimumLiquidationCollateralUpdated, error) {
	event := new(SsvNetworkMinimumLiquidationCollateralUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkNetworkEarningsWithdrawnIterator is returned from FilterNetworkEarningsWithdrawn and is used to iterate over the raw logs and unpacked data for NetworkEarningsWithdrawn events raised by the SsvNetwork contract.
type SsvNetworkNetworkEarningsWithdrawnIterator struct {
	Event *SsvNetworkNetworkEarningsWithdrawn // Event containing the contract specifics and raw log

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
func (it *SsvNetworkNetworkEarningsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkNetworkEarningsWithdrawn)
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
		it.Event = new(SsvNetworkNetworkEarningsWithdrawn)
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
func (it *SsvNetworkNetworkEarningsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkNetworkEarningsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkNetworkEarningsWithdrawn represents a NetworkEarningsWithdrawn event raised by the SsvNetwork contract.
type SsvNetworkNetworkEarningsWithdrawn struct {
	Value     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNetworkEarningsWithdrawn is a free log retrieval operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_SsvNetwork *SsvNetworkFilterer) FilterNetworkEarningsWithdrawn(opts *bind.FilterOpts) (*SsvNetworkNetworkEarningsWithdrawnIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkNetworkEarningsWithdrawnIterator{contract: _SsvNetwork.contract, event: "NetworkEarningsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNetworkEarningsWithdrawn is a free log subscription operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_SsvNetwork *SsvNetworkFilterer) WatchNetworkEarningsWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvNetworkNetworkEarningsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkNetworkEarningsWithdrawn)
				if err := _SsvNetwork.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
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

// ParseNetworkEarningsWithdrawn is a log parse operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_SsvNetwork *SsvNetworkFilterer) ParseNetworkEarningsWithdrawn(log types.Log) (*SsvNetworkNetworkEarningsWithdrawn, error) {
	event := new(SsvNetworkNetworkEarningsWithdrawn)
	if err := _SsvNetwork.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkNetworkFeeUpdatedIterator is returned from FilterNetworkFeeUpdated and is used to iterate over the raw logs and unpacked data for NetworkFeeUpdated events raised by the SsvNetwork contract.
type SsvNetworkNetworkFeeUpdatedIterator struct {
	Event *SsvNetworkNetworkFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkNetworkFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkNetworkFeeUpdated)
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
		it.Event = new(SsvNetworkNetworkFeeUpdated)
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
func (it *SsvNetworkNetworkFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkNetworkFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkNetworkFeeUpdated represents a NetworkFeeUpdated event raised by the SsvNetwork contract.
type SsvNetworkNetworkFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNetworkFeeUpdated is a free log retrieval operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_SsvNetwork *SsvNetworkFilterer) FilterNetworkFeeUpdated(opts *bind.FilterOpts) (*SsvNetworkNetworkFeeUpdatedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkNetworkFeeUpdatedIterator{contract: _SsvNetwork.contract, event: "NetworkFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchNetworkFeeUpdated is a free log subscription operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_SsvNetwork *SsvNetworkFilterer) WatchNetworkFeeUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkNetworkFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkNetworkFeeUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
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

// ParseNetworkFeeUpdated is a log parse operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_SsvNetwork *SsvNetworkFilterer) ParseNetworkFeeUpdated(log types.Log) (*SsvNetworkNetworkFeeUpdated, error) {
	event := new(SsvNetworkNetworkFeeUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the SsvNetwork contract.
type SsvNetworkOperatorAddedIterator struct {
	Event *SsvNetworkOperatorAdded // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorAdded)
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
		it.Event = new(SsvNetworkOperatorAdded)
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
func (it *SsvNetworkOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorAdded represents a OperatorAdded event raised by the SsvNetwork contract.
type SsvNetworkOperatorAdded struct {
	OperatorId uint64
	Owner      common.Address
	PublicKey  []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operatorId []uint64, owner []common.Address) (*SsvNetworkOperatorAddedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorAddedIterator{contract: _SsvNetwork.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorAdded, operatorId []uint64, owner []common.Address) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorAdded)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorAdded(log types.Log) (*SsvNetworkOperatorAdded, error) {
	event := new(SsvNetworkOperatorAdded)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorFeeDeclarationCancelledIterator is returned from FilterOperatorFeeDeclarationCancelled and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclarationCancelled events raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeDeclarationCancelledIterator struct {
	Event *SsvNetworkOperatorFeeDeclarationCancelled // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorFeeDeclarationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorFeeDeclarationCancelled)
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
		it.Event = new(SsvNetworkOperatorFeeDeclarationCancelled)
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
func (it *SsvNetworkOperatorFeeDeclarationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorFeeDeclarationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorFeeDeclarationCancelled represents a OperatorFeeDeclarationCancelled event raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeDeclarationCancelled struct {
	Owner      common.Address
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclarationCancelled is a free log retrieval operation binding the contract event 0x5055fa347441172447637c015e80a3ee748b9382212ceb5dca5a3683298fd6f3.
//
// Solidity: event OperatorFeeDeclarationCancelled(address indexed owner, uint64 indexed operatorId)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorFeeDeclarationCancelled(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvNetworkOperatorFeeDeclarationCancelledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorFeeDeclarationCancelled", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorFeeDeclarationCancelledIterator{contract: _SsvNetwork.contract, event: "OperatorFeeDeclarationCancelled", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclarationCancelled is a free log subscription operation binding the contract event 0x5055fa347441172447637c015e80a3ee748b9382212ceb5dca5a3683298fd6f3.
//
// Solidity: event OperatorFeeDeclarationCancelled(address indexed owner, uint64 indexed operatorId)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorFeeDeclarationCancelled(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorFeeDeclarationCancelled, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorFeeDeclarationCancelled", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorFeeDeclarationCancelled)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeDeclarationCancelled", log); err != nil {
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

// ParseOperatorFeeDeclarationCancelled is a log parse operation binding the contract event 0x5055fa347441172447637c015e80a3ee748b9382212ceb5dca5a3683298fd6f3.
//
// Solidity: event OperatorFeeDeclarationCancelled(address indexed owner, uint64 indexed operatorId)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorFeeDeclarationCancelled(log types.Log) (*SsvNetworkOperatorFeeDeclarationCancelled, error) {
	event := new(SsvNetworkOperatorFeeDeclarationCancelled)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeDeclarationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorFeeDeclaredIterator is returned from FilterOperatorFeeDeclared and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclared events raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeDeclaredIterator struct {
	Event *SsvNetworkOperatorFeeDeclared // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorFeeDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorFeeDeclared)
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
		it.Event = new(SsvNetworkOperatorFeeDeclared)
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
func (it *SsvNetworkOperatorFeeDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorFeeDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorFeeDeclared represents a OperatorFeeDeclared event raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeDeclared struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclared is a free log retrieval operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorFeeDeclared(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvNetworkOperatorFeeDeclaredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorFeeDeclaredIterator{contract: _SsvNetwork.contract, event: "OperatorFeeDeclared", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclared is a free log subscription operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorFeeDeclared(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorFeeDeclared, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorFeeDeclared)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
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

// ParseOperatorFeeDeclared is a log parse operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorFeeDeclared(log types.Log) (*SsvNetworkOperatorFeeDeclared, error) {
	event := new(SsvNetworkOperatorFeeDeclared)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorFeeExecutedIterator is returned from FilterOperatorFeeExecuted and is used to iterate over the raw logs and unpacked data for OperatorFeeExecuted events raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeExecutedIterator struct {
	Event *SsvNetworkOperatorFeeExecuted // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorFeeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorFeeExecuted)
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
		it.Event = new(SsvNetworkOperatorFeeExecuted)
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
func (it *SsvNetworkOperatorFeeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorFeeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorFeeExecuted represents a OperatorFeeExecuted event raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeExecuted struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeExecuted is a free log retrieval operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorFeeExecuted(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvNetworkOperatorFeeExecutedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorFeeExecutedIterator{contract: _SsvNetwork.contract, event: "OperatorFeeExecuted", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeExecuted is a free log subscription operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorFeeExecuted(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorFeeExecuted, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorFeeExecuted)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
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

// ParseOperatorFeeExecuted is a log parse operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorFeeExecuted(log types.Log) (*SsvNetworkOperatorFeeExecuted, error) {
	event := new(SsvNetworkOperatorFeeExecuted)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator is returned from FilterOperatorFeeIncreaseLimitUpdated and is used to iterate over the raw logs and unpacked data for OperatorFeeIncreaseLimitUpdated events raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator struct {
	Event *SsvNetworkOperatorFeeIncreaseLimitUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorFeeIncreaseLimitUpdated)
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
		it.Event = new(SsvNetworkOperatorFeeIncreaseLimitUpdated)
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
func (it *SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorFeeIncreaseLimitUpdated represents a OperatorFeeIncreaseLimitUpdated event raised by the SsvNetwork contract.
type SsvNetworkOperatorFeeIncreaseLimitUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeIncreaseLimitUpdated is a free log retrieval operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorFeeIncreaseLimitUpdated(opts *bind.FilterOpts) (*SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator, error) {

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorFeeIncreaseLimitUpdatedIterator{contract: _SsvNetwork.contract, event: "OperatorFeeIncreaseLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeIncreaseLimitUpdated is a free log subscription operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorFeeIncreaseLimitUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorFeeIncreaseLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorFeeIncreaseLimitUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
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

// ParseOperatorFeeIncreaseLimitUpdated is a log parse operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorFeeIncreaseLimitUpdated(log types.Log) (*SsvNetworkOperatorFeeIncreaseLimitUpdated, error) {
	event := new(SsvNetworkOperatorFeeIncreaseLimitUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the SsvNetwork contract.
type SsvNetworkOperatorRemovedIterator struct {
	Event *SsvNetworkOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorRemoved)
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
		it.Event = new(SsvNetworkOperatorRemoved)
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
func (it *SsvNetworkOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorRemoved represents a OperatorRemoved event raised by the SsvNetwork contract.
type SsvNetworkOperatorRemoved struct {
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operatorId []uint64) (*SsvNetworkOperatorRemovedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorRemovedIterator{contract: _SsvNetwork.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorRemoved, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorRemoved)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorRemoved(log types.Log) (*SsvNetworkOperatorRemoved, error) {
	event := new(SsvNetworkOperatorRemoved)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorWhitelistUpdatedIterator is returned from FilterOperatorWhitelistUpdated and is used to iterate over the raw logs and unpacked data for OperatorWhitelistUpdated events raised by the SsvNetwork contract.
type SsvNetworkOperatorWhitelistUpdatedIterator struct {
	Event *SsvNetworkOperatorWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorWhitelistUpdated)
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
		it.Event = new(SsvNetworkOperatorWhitelistUpdated)
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
func (it *SsvNetworkOperatorWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorWhitelistUpdated represents a OperatorWhitelistUpdated event raised by the SsvNetwork contract.
type SsvNetworkOperatorWhitelistUpdated struct {
	OperatorId  uint64
	Whitelisted common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistUpdated is a free log retrieval operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorWhitelistUpdated(opts *bind.FilterOpts, operatorId []uint64) (*SsvNetworkOperatorWhitelistUpdatedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorWhitelistUpdatedIterator{contract: _SsvNetwork.contract, event: "OperatorWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistUpdated is a free log subscription operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorWhitelistUpdated, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorWhitelistUpdated)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
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

// ParseOperatorWhitelistUpdated is a log parse operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorWhitelistUpdated(log types.Log) (*SsvNetworkOperatorWhitelistUpdated, error) {
	event := new(SsvNetworkOperatorWhitelistUpdated)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOperatorWithdrawnIterator is returned from FilterOperatorWithdrawn and is used to iterate over the raw logs and unpacked data for OperatorWithdrawn events raised by the SsvNetwork contract.
type SsvNetworkOperatorWithdrawnIterator struct {
	Event *SsvNetworkOperatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOperatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOperatorWithdrawn)
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
		it.Event = new(SsvNetworkOperatorWithdrawn)
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
func (it *SsvNetworkOperatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOperatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOperatorWithdrawn represents a OperatorWithdrawn event raised by the SsvNetwork contract.
type SsvNetworkOperatorWithdrawn struct {
	Owner      common.Address
	OperatorId uint64
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrawn is a free log retrieval operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_SsvNetwork *SsvNetworkFilterer) FilterOperatorWithdrawn(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvNetworkOperatorWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOperatorWithdrawnIterator{contract: _SsvNetwork.contract, event: "OperatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrawn is a free log subscription operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_SsvNetwork *SsvNetworkFilterer) WatchOperatorWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvNetworkOperatorWithdrawn, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOperatorWithdrawn)
				if err := _SsvNetwork.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
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

// ParseOperatorWithdrawn is a log parse operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_SsvNetwork *SsvNetworkFilterer) ParseOperatorWithdrawn(log types.Log) (*SsvNetworkOperatorWithdrawn, error) {
	event := new(SsvNetworkOperatorWithdrawn)
	if err := _SsvNetwork.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SsvNetwork contract.
type SsvNetworkOwnershipTransferStartedIterator struct {
	Event *SsvNetworkOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOwnershipTransferStarted)
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
		it.Event = new(SsvNetworkOwnershipTransferStarted)
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
func (it *SsvNetworkOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SsvNetwork contract.
type SsvNetworkOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SsvNetwork *SsvNetworkFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvNetworkOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOwnershipTransferStartedIterator{contract: _SsvNetwork.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SsvNetwork *SsvNetworkFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SsvNetworkOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOwnershipTransferStarted)
				if err := _SsvNetwork.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_SsvNetwork *SsvNetworkFilterer) ParseOwnershipTransferStarted(log types.Log) (*SsvNetworkOwnershipTransferStarted, error) {
	event := new(SsvNetworkOwnershipTransferStarted)
	if err := _SsvNetwork.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SsvNetwork contract.
type SsvNetworkOwnershipTransferredIterator struct {
	Event *SsvNetworkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SsvNetworkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkOwnershipTransferred)
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
		it.Event = new(SsvNetworkOwnershipTransferred)
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
func (it *SsvNetworkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkOwnershipTransferred represents a OwnershipTransferred event raised by the SsvNetwork contract.
type SsvNetworkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsvNetwork *SsvNetworkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvNetworkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkOwnershipTransferredIterator{contract: _SsvNetwork.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsvNetwork *SsvNetworkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SsvNetworkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkOwnershipTransferred)
				if err := _SsvNetwork.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SsvNetwork *SsvNetworkFilterer) ParseOwnershipTransferred(log types.Log) (*SsvNetworkOwnershipTransferred, error) {
	event := new(SsvNetworkOwnershipTransferred)
	if err := _SsvNetwork.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SsvNetwork contract.
type SsvNetworkUpgradedIterator struct {
	Event *SsvNetworkUpgraded // Event containing the contract specifics and raw log

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
func (it *SsvNetworkUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkUpgraded)
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
		it.Event = new(SsvNetworkUpgraded)
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
func (it *SsvNetworkUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkUpgraded represents a Upgraded event raised by the SsvNetwork contract.
type SsvNetworkUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsvNetwork *SsvNetworkFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SsvNetworkUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkUpgradedIterator{contract: _SsvNetwork.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsvNetwork *SsvNetworkFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SsvNetworkUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkUpgraded)
				if err := _SsvNetwork.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SsvNetwork *SsvNetworkFilterer) ParseUpgraded(log types.Log) (*SsvNetworkUpgraded, error) {
	event := new(SsvNetworkUpgraded)
	if err := _SsvNetwork.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the SsvNetwork contract.
type SsvNetworkValidatorAddedIterator struct {
	Event *SsvNetworkValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SsvNetworkValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkValidatorAdded)
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
		it.Event = new(SsvNetworkValidatorAdded)
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
func (it *SsvNetworkValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkValidatorAdded represents a ValidatorAdded event raised by the SsvNetwork contract.
type SsvNetworkValidatorAdded struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Shares      []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) FilterValidatorAdded(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkValidatorAddedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkValidatorAddedIterator{contract: _SsvNetwork.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SsvNetworkValidatorAdded, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkValidatorAdded)
				if err := _SsvNetwork.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) ParseValidatorAdded(log types.Log) (*SsvNetworkValidatorAdded, error) {
	event := new(SsvNetworkValidatorAdded)
	if err := _SsvNetwork.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the SsvNetwork contract.
type SsvNetworkValidatorRemovedIterator struct {
	Event *SsvNetworkValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SsvNetworkValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkValidatorRemoved)
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
		it.Event = new(SsvNetworkValidatorRemoved)
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
func (it *SsvNetworkValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkValidatorRemoved represents a ValidatorRemoved event raised by the SsvNetwork contract.
type SsvNetworkValidatorRemoved struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, owner []common.Address) (*SsvNetworkValidatorRemovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.FilterLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvNetworkValidatorRemovedIterator{contract: _SsvNetwork.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SsvNetworkValidatorRemoved, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SsvNetwork.contract.WatchLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkValidatorRemoved)
				if err := _SsvNetwork.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SsvNetwork *SsvNetworkFilterer) ParseValidatorRemoved(log types.Log) (*SsvNetworkValidatorRemoved, error) {
	event := new(SsvNetworkValidatorRemoved)
	if err := _SsvNetwork.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
