package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	storage "github.com/stafiprotocol/eth-ssv-client/bindings/Storage"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
)

// 1 deposited { 2 withdrawl match 3 staked 4 withdrawl unmatch } { 5 offboard 6 OffBoard can withdraw 7 OffBoard withdrawed } 8 waiting 9 active 10 exited 11 withdrawable 12 withdrawdone { 13 distributed }
// 51 active+slash 52 exit+slash 53 withdrawable+slash 54 withdrawdone+slash 55 distributed+slash
const (
	ValidatorStatusUnInitial = uint8(0)
	ValidatorStatusDeposited = uint8(1)

	// lightnode + super node related
	ValidatorStatusWithdrawMatch   = uint8(2)
	ValidatorStatusStaked          = uint8(3)
	ValidatorStatusWithdrawUnmatch = uint8(4)

	// lightnode related
	ValidatorStatusOffBoard            = uint8(5)
	ValidatorStatusOffBoardCanWithdraw = uint8(6)
	ValidatorStatusOffBoardWithdrawed  = uint8(7)

	// status on beacon chain
	ValidatorStatusWaiting      = uint8(8)
	ValidatorStatusActive       = uint8(9)
	ValidatorStatusExited       = uint8(10)
	ValidatorStatusWithdrawable = uint8(11)
	ValidatorStatusWithdrawDone = uint8(12)

	// after distribute reward
	ValidatorStatusDistributed = uint8(13) // distribute full withdrawal

	// after slash
	ValidatorStatusActiveSlash       = uint8(51)
	ValidatorStatusExitedSlash       = uint8(52)
	ValidatorStatusWithdrawableSlash = uint8(53)
	ValidatorStatusWithdrawDoneSlash = uint8(54)

	ValidatorStatusDistributedSlash = uint8(55) // distribute full withdrawal
)

// 1 common node 2 trust node 3 light node 4 super node
const (
	NodeTypeCommon = uint8(1)
	NodeTypeTrust  = uint8(2)
	NodeTypeLight  = uint8(3)
	NodeTypeSuper  = uint8(4)
)

const (
	ValidatorEverSlashedFalse = uint8(0)
	ValidatorEverSlashedTrue  = uint8(1)
)

const (
	FeePool          = uint8(1)
	SuperNodeFeePool = uint8(2)
)

const (
	MetaTypeEth1BlockSyncer            = uint8(1) // dealed block height
	MetaTypeEth2ValidatorInfoSyncer    = uint8(2) // dealed epoch
	MetaTypeEth2ValidatorBalanceSyncer = uint8(3) // dealed epoch
	MetaTypeV1ValidatorSyncer          = uint8(4) // dealed block height
	MetaTypeEth2BlockSyncer            = uint8(5) // dealed epoch
	MetaTypeEth2NodeBalanceCollector   = uint8(6) // dealed epoch
)
const (
	SlashTypeFeeRecipient  = uint8(1)
	SlashTypeProposerSlash = uint8(2)
	SlashTypeAttesterSlash = uint8(3)

	SlashTypeNotExitSlash = uint8(7)

	// not show in front end
	SlashTypeSyncMiss     = uint8(4)
	SlashTypeAttesterMiss = uint8(5)
	SlashTypeProposerMiss = uint8(6)
)

const (
	TheMergeEpoch       = uint64(148000)
	TheMergeSlot        = uint64(4736000)
	TheMergeBlockNumber = uint64(15572967)

	RewardV1EndEpoch    = uint64(195975) //todo mainnet
	RewardEpochInterval = uint64(75)
	SlashStartEpoch     = uint64(180000)
	//dev
	DevTheMergeEpoch       = uint64(4400)
	DevTheMergeBlockNumber = uint64(133654)
	DevRewardV1EndEpoch    = uint64(75)

	StandardEffectiveBalance            = uint64(32e9) //gwei
	StandardSuperNodeFakeDepositBalance = uint64(1e9)  //gwei
	OfficialSlashAmount                 = uint64(1e9)  //gwei
	MaxPartialWithdrawalAmount          = uint64(8e9)  //gwei

	// node deposit amount(gwei)
	NodeDepositAmount0  = uint64(0)    //gwei super
	NodeDepositAmount4  = uint64(4e9)  //gwei solo 4
	NodeDepositAmount8  = uint64(8e9)  //gwei solo 8
	NodeDepositAmount12 = uint64(12e9) //gwei solo 12
)

//	enum ClaimType {
//	    None,
//	    CLAIMREWARD,
//	    CLAIMDEPOSIT,
//	    CLAIMTOTAL
//	}
const (
	NodeClaimTypeNone         = uint8(0)
	NodeClaimTypeClaimReward  = uint8(1)
	NodeClaimTypeClaimDeposit = uint8(2)
	NodeClaimTypeClaimTotal   = uint8(3)
)

var (
	// dev use
	OldRethSupply, _ = new(big.Int).SetString("25642334000000000000", 10)

	GweiDeci = decimal.NewFromInt(1e9)

	PlatformFeeV1Deci = decimal.NewFromFloat(0.1)
	NodeFeeV1Deci     = decimal.NewFromFloat(0.1)

	Percent5Deci  = decimal.NewFromFloat(0.05)
	Percent90Deci = decimal.NewFromFloat(0.9)
)

const (
	StakerWithdrawalClaimableTimestamp = uint64(1)
	MinValidatorWithdrawabilityDelay   = uint64(256 + 5)
	MaxDistributeWaitSeconds           = uint64(8 * 60 * 60)
	MaxDistributeWaitEpoch             = uint64(75)

	EjectorUptimeInterval = uint64(10 * 60)
)

// Get an eth2 epoch number by time
func EpochAtTimestamp(config beacon.Eth2Config, time uint64) uint64 {
	return config.GenesisEpoch + (time-config.GenesisTime)/config.SecondsPerEpoch
}

func StartTimestampOfEpoch(config beacon.Eth2Config, epoch uint64) uint64 {
	return (epoch-config.GenesisEpoch)*config.SecondsPerEpoch + config.GenesisTime
}

func TimestampOfSlot(config beacon.Eth2Config, slot uint64) uint64 {
	return slot*config.SecondsPerSlot + config.GenesisTime
}

// Get an eth2 first slot number by epoch
func StartSlotOfEpoch(config beacon.Eth2Config, epoch uint64) uint64 {
	return config.SlotsPerEpoch * epoch
}
func EndSlotOfEpoch(config beacon.Eth2Config, epoch uint64) uint64 {
	return config.SlotsPerEpoch*(epoch+1) - 1
}

func GetNodeManagedEth(nodeDeposit, balance uint64, status uint8) uint64 {
	switch status {
	case ValidatorStatusDeposited, ValidatorStatusWithdrawMatch, ValidatorStatusWithdrawUnmatch:
		return nodeDeposit

	case ValidatorStatusStaked, ValidatorStatusWaiting:
		return StandardEffectiveBalance

	default:
		return balance
	}
}

func GetGaspriceFromEthgasstation() (base, priority uint64, err error) {
	rsp, err := http.Get("https://api.ethgasstation.info/api/fee-estimate")
	if err != nil {
		return 0, 0, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("status err %d", rsp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return 0, 0, err
	}
	if len(bodyBytes) == 0 {
		return 0, 0, fmt.Errorf("bodyBytes zero err")
	}
	resGasPrice := ResGasPrice{}
	err = json.Unmarshal(bodyBytes, &resGasPrice)
	if err != nil {
		return 0, 0, err
	}
	return uint64(resGasPrice.BaseFee), uint64(resGasPrice.PriorityFee.Fast), nil
}

type ResGasPrice struct {
	BaseFee     int     `json:"baseFee"`
	BlockNumber int     `json:"blockNumber"`
	BlockTime   float64 `json:"blockTime"`
	GasPrice    struct {
		Fast     int `json:"fast"`
		Instant  int `json:"instant"`
		Standard int `json:"standard"`
	} `json:"gasPrice"`
	NextBaseFee int `json:"nextBaseFee"`
	PriorityFee struct {
		Fast     int `json:"fast"`
		Instant  int `json:"instant"`
		Standard int `json:"standard"`
	} `json:"priorityFee"`
}

func GetGaspriceFromBeacon() (base uint64, err error) {
	rsp, err := http.Get("https://beaconcha.in/api/v1/execution/gasnow")
	if err != nil {
		return 0, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status err %d", rsp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}
	if len(bodyBytes) == 0 {
		return 0, fmt.Errorf("bodyBytes zero err")
	}
	resGasPrice := ResGasPriceFromBeacon{}
	err = json.Unmarshal(bodyBytes, &resGasPrice)
	if err != nil {
		return 0, err
	}

	return uint64(resGasPrice.Data.Standard), nil
}

type ResGasPriceFromBeacon struct {
	Code int `json:"code"`
	Data struct {
		Rapid     int64   `json:"rapid"`
		Fast      int64   `json:"fast"`
		Standard  int64   `json:"standard"`
		Slow      int64   `json:"slow"`
		Timestamp int64   `json:"timestamp"`
		PriceUSD  float64 `json:"priceUSD"`
	} `json:"data"`
}

// statistic use
// return (user deposit, user reward, val deposit, val reward, paltform fee)
func GetUserValPlatformDepositAndRewardV1(validatorBalance, nodeDepositAmount uint64, platformFee, nodeFee decimal.Decimal) (uint64, uint64, uint64, uint64, uint64) {
	userDepositBalance := StandardEffectiveBalance - nodeDepositAmount

	switch {
	case validatorBalance == StandardEffectiveBalance:
		return userDepositBalance, 0, nodeDepositAmount, 0, 0
	case validatorBalance < StandardEffectiveBalance:
		loss := StandardEffectiveBalance - validatorBalance
		if loss < nodeDepositAmount {
			return userDepositBalance, 0, nodeDepositAmount - loss, 0, 0
		} else {
			return validatorBalance, 0, 0, 0, 0
		}
	case validatorBalance > StandardEffectiveBalance:
		// total staking reward
		reward := validatorBalance - StandardEffectiveBalance
		rewardDeci := decimal.NewFromInt(int64(reward))

		// platform Fee
		platformFeeDeci := rewardDeci.Mul(platformFee)

		// node+user stake reward
		nodeAndUserStakeRewardDeci := rewardDeci.Sub(platformFeeDeci)

		// user stake reward
		userStakeRewardDeci := nodeAndUserStakeRewardDeci.Mul(decimal.NewFromInt(int64(userDepositBalance))).Div(decimal.NewFromInt(int64(StandardEffectiveBalance)))
		// node stake reward
		nodeStakeRewardDeci := nodeAndUserStakeRewardDeci.Sub(userStakeRewardDeci)

		// node commisson reward from user
		nodeCommissionRewardFromUserDeci := userStakeRewardDeci.Mul(nodeFee)

		// user reward
		userRewardDeci := userStakeRewardDeci.Sub(nodeCommissionRewardFromUserDeci)
		// node reward
		nodeRewardDeci := nodeStakeRewardDeci.Add(nodeCommissionRewardFromUserDeci)

		return userDepositBalance, userRewardDeci.BigInt().Uint64(), nodeDepositAmount, nodeRewardDeci.BigInt().Uint64(), platformFeeDeci.BigInt().Uint64()
	default:
		// should not happen here
		panic("GetUserValPlatformDepositAndReward ")
	}
}

// v1:
// user = 90%*(1- nodedeposit/32)*90%
// node = 90%*(nodedeposit/32)+90%*(1- nodedeposit/32)*10%
// platform = 10%
// rewardDeci decimals maybe 9 or 18, also the returns
// return (user reward, node reward, paltform fee)
func GetUserNodePlatformRewardV1(nodeDepositAmount uint64, rewardDeci decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	if !rewardDeci.IsPositive() || nodeDepositAmount > StandardEffectiveBalance {
		return decimal.Zero, decimal.Zero, decimal.Zero
	}
	userDepositBalance := StandardEffectiveBalance - nodeDepositAmount

	// platform Fee
	platformFeeDeci := rewardDeci.Mul(PlatformFeeV1Deci) // 10%

	// node+user stake reward
	nodeAndUserStakeRewardDeci := rewardDeci.Sub(platformFeeDeci) // 90%

	// user stake reward
	userStakeRewardDeci := nodeAndUserStakeRewardDeci.Mul(decimal.NewFromInt(int64(userDepositBalance))).Div(decimal.NewFromInt(int64(StandardEffectiveBalance))) // 90%*(1 - nodedeposit/32)
	// node stake reward
	nodeStakeRewardDeci := nodeAndUserStakeRewardDeci.Sub(userStakeRewardDeci) // 90%*(nodedeposit/32)

	// node commisson reward from user
	nodeCommissionRewardFromUserDeci := userStakeRewardDeci.Mul(NodeFeeV1Deci) // 90%*(1 - nodedeposit/32)*10%

	// user reward
	userRewardDeci := userStakeRewardDeci.Sub(nodeCommissionRewardFromUserDeci) // 90%*(1 - nodedeposit/32)*90%
	// node reward
	nodeRewardDeci := nodeStakeRewardDeci.Add(nodeCommissionRewardFromUserDeci) // // 90%*(nodedeposit/32) + 90%*(1 - nodedeposit/32)*10%

	return userRewardDeci, nodeRewardDeci, platformFeeDeci
}

// v2:
// user = 90%*(1-nodedeposit/32)
// node = 5% + (90% * nodedeposit/32)
// platform = 5%
// rewardDeci decimals maybe 9 or 18, also the returns
// return (user reward, node reward, paltform fee)
func GetUserNodePlatformRewardV2(nodeDepositAmount uint64, rewardDeci decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	if !rewardDeci.IsPositive() || nodeDepositAmount > StandardEffectiveBalance {
		return decimal.Zero, decimal.Zero, decimal.Zero
	}
	nodeDepositAmountDeci := decimal.NewFromInt(int64(nodeDepositAmount))
	standEffectiveBalanceDeci := decimal.NewFromInt(int64(StandardEffectiveBalance))

	// platform Fee
	platformFeeDeci := rewardDeci.Mul(Percent5Deci)
	nodeRewardDeci := platformFeeDeci.Add(rewardDeci.Mul(Percent90Deci).Mul(nodeDepositAmountDeci).Div(standEffectiveBalanceDeci))

	userRewardDeci := rewardDeci.Sub(platformFeeDeci).Sub(nodeRewardDeci)
	if userRewardDeci.IsNegative() {
		userRewardDeci = decimal.Zero
	}

	return userRewardDeci, nodeRewardDeci, platformFeeDeci
}

func GetValidatorTotalReward(balance, totalWithdrawal, totalFee uint64) uint64 {
	totalBalance := balance + totalWithdrawal + totalFee
	if totalBalance > StandardEffectiveBalance {
		return totalBalance - StandardEffectiveBalance
	}
	return 0
}

// keccak256(abi.encodePacked("contract.address", _contractName))
func ContractStorageKey(name string) [32]byte {
	return crypto.Keccak256Hash([]byte("contract.address"), []byte(name))
}

func NodeSubmissionKey(sender common.Address, _block *big.Int, _totalEth *big.Int, _stakingEth *big.Int, _rethSupply *big.Int) [32]byte {
	// keccak256(abi.encodePacked("network.balances.submitted.node", sender, _block, _totalEth, _stakingEth, _rethSupply))
	return crypto.Keccak256Hash([]byte("network.balances.submitted.node"), sender.Bytes(), common.LeftPadBytes(_block.Bytes(), 32),
		common.LeftPadBytes(_totalEth.Bytes(), 32), common.LeftPadBytes(_stakingEth.Bytes(), 32), common.LeftPadBytes(_rethSupply.Bytes(), 32))
}

func StafiWithdrawProposalNodeKey(sender common.Address, proposalId [32]byte) [32]byte {
	return crypto.Keccak256Hash([]byte("stafiWithdraw.proposal.node.key"), proposalId[:], sender.Bytes())
}

func DistributeWithdrawalsProposalNodeKey(sender common.Address, _dealedHeight, _userAmount, _nodeAmount, _platformAmount, _maxClaimableWithdrawIndex *big.Int) [32]byte {
	proposalId := crypto.Keccak256Hash([]byte("distributeWithdrawals"), common.LeftPadBytes(_dealedHeight.Bytes(), 32), common.LeftPadBytes(_userAmount.Bytes(), 32),
		common.LeftPadBytes(_nodeAmount.Bytes(), 32), common.LeftPadBytes(_platformAmount.Bytes(), 32), common.LeftPadBytes(_maxClaimableWithdrawIndex.Bytes(), 32))
	return StafiWithdrawProposalNodeKey(sender, proposalId)
}

func StafiDistributorProposalNodeKey(sender common.Address, proposalId [32]byte) [32]byte {
	return crypto.Keccak256Hash([]byte("stafiDistributor.proposal.node.key"), proposalId[:], sender.Bytes())
}

func ReserveEthForWithdrawProposalId(cycle *big.Int) [32]byte {
	return crypto.Keccak256Hash([]byte("reserveEthForWithdraw"), common.LeftPadBytes(cycle.Bytes(), 32))
}

func DistributeFeeProposalNodeKey(sender common.Address, _dealedHeight, _userAmount, _nodeAmount, _platformAmount *big.Int) [32]byte {
	proposalId := crypto.Keccak256Hash([]byte("distributeFee"), common.LeftPadBytes(_dealedHeight.Bytes(), 32), common.LeftPadBytes(_userAmount.Bytes(), 32),
		common.LeftPadBytes(_nodeAmount.Bytes(), 32), common.LeftPadBytes(_platformAmount.Bytes(), 32))
	return StafiDistributorProposalNodeKey(sender, proposalId)
}

func DistributeSuperNodeFeeProposalNodeKey(sender common.Address, _dealedHeight, _userAmount, _nodeAmount, _platformAmount *big.Int) [32]byte {
	proposalId := crypto.Keccak256Hash([]byte("distributeSuperNodeFee"), common.LeftPadBytes(_dealedHeight.Bytes(), 32), common.LeftPadBytes(_userAmount.Bytes(), 32),
		common.LeftPadBytes(_nodeAmount.Bytes(), 32), common.LeftPadBytes(_platformAmount.Bytes(), 32))
	return StafiDistributorProposalNodeKey(sender, proposalId)
}

func GetContractAddress(storage *storage.Storage, name string) (common.Address, error) {
	address, err := storage.GetAddress(nil, ContractStorageKey(name))
	if err != nil {
		return common.Address{}, err
	}
	if bytes.Equal(address.Bytes(), common.Address{}.Bytes()) {
		return common.Address{}, fmt.Errorf("address empty")
	}
	return address, nil
}

func WaitTxOkCommon(client *ethclient.Client, txHash common.Hash) (err error) {
	defer func() {
		if err != nil {
			ShutdownRequestChannel <- struct{}{}
		}
	}()

	retry := 0
	for {
		if retry > RetryLimit {
			return fmt.Errorf("waitTx %s reach retry limit", txHash.String())
		}
		_, pending, err := client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"hash": txHash.String(),
				"err":  err.Error(),
			}).Warn("TransactionByHash")

			time.Sleep(RetryInterval)
			retry++
			continue
		} else {
			if pending {
				logrus.WithFields(logrus.Fields{
					"hash":    txHash.String(),
					"pending": pending,
				}).Warn("TransactionByHash")

				time.Sleep(RetryInterval)
				retry++
				continue
			} else {
				// check status
				var receipt *types.Receipt
				subRetry := 0
				for {
					if subRetry > RetryLimit {
						return fmt.Errorf("TransactionReceipt %s reach retry limit", txHash.String())
					}

					receipt, err = client.TransactionReceipt(context.Background(), txHash)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"hash": txHash.String(),
							"err":  err.Error(),
						}).Warn("tx TransactionReceipt")

						time.Sleep(RetryInterval)
						subRetry++
						continue
					}
					break
				}

				if receipt.Status == 1 { //success
					break
				} else { //failed
					return fmt.Errorf("tx %s failed", txHash.String())
				}
			}
		}
	}

	logrus.WithFields(logrus.Fields{
		"tx": txHash.String(),
	}).Info("tx send ok")

	return nil
}
