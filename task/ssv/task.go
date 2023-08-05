package task_ssv

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prysmaticlabs/prysm/v3/encoding/bytesutil"

	// "github.com/prysmaticlabs/prysm/v3/config/params"
	types "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	erc20 "github.com/stafiprotocol/eth-ssv-client/bindings/Erc20"
	ssv_network "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-ssv-client/bindings/SsvNetworkViews"
	storage "github.com/stafiprotocol/eth-ssv-client/bindings/Storage"
	super_node "github.com/stafiprotocol/eth-ssv-client/bindings/SuperNode"
	user_deposit "github.com/stafiprotocol/eth-ssv-client/bindings/UserDeposit"
	withdraw "github.com/stafiprotocol/eth-ssv-client/bindings/Withdraw"
	"github.com/stafiprotocol/eth-ssv-client/pkg/config"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/constants"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/bls"
	"github.com/stafiprotocol/eth-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

var (
	minAmountNeedStake   = decimal.NewFromBigInt(big.NewInt(31), 18)
	minAmountNeedDeposit = decimal.NewFromBigInt(big.NewInt(32), 18)

	superNodeDepositAmount = decimal.NewFromBigInt(big.NewInt(1), 18)
	superNodeStakeAmount   = decimal.NewFromBigInt(big.NewInt(31), 18)

	blocksOfOneYear  = decimal.NewFromInt(2629800)
	blocksOfHalfYear = decimal.NewFromInt(1314900)
)

var (
	devPostUptimeUrl     = "https://test-drop-api.stafi.io/reth/v1/uploadEjectorUptime"
	mainnetPostUptimeUrl = "https://drop-api.stafi.io/reth/v1/uploadEjectorUptime"
)

var (
	domainVoluntaryExit  = bytesutil.Uint32ToBytes4(0x04000000)
	shardCommitteePeriod = types.Epoch(256) // ShardCommitteePeriod is the minimum amount of epochs a validator must participate before exiting.
)

// only support stafi super node account now !!!
// 0. find next key index and cache validator status on start
// 1. update validator status(on execution/ssv/beacon) periodically
// 2. check stakepool balance periodically, call stake/deposit if match
// 3. register validator on ssv, if status is staked on stafi contract
// 4. remove validator on ssv, if status is exited on beacon
type Task struct {
	stop            chan struct{}
	eth1StartHeight uint64
	eth1Endpoint    string
	eth2Endpoint    string

	superNodeKeyPair *secp256k1.Keypair
	ssvKeyPair       *secp256k1.Keypair

	gasLimit                       *big.Int
	maxGasPrice                    *big.Int
	storageContractAddress         common.Address
	ssvNetworkContractAddress      common.Address
	ssvNetworkViewsContractAddress common.Address
	ssvTokenContractAddress        common.Address
	seed                           []byte
	postUptimeUrl                  string

	// --- need init on start
	dev           bool
	ssvApiNetwork string
	chain         constants.Chain

	connectionOfSuperNodeAccount *connection.Connection
	connectionOfSsvAccount       *connection.Connection

	eth1WithdrawalAdress       common.Address
	feeRecipientAddressOnStafi common.Address

	eth1Client *ethclient.Client

	superNodeContract       *super_node.SuperNode
	userDepositContract     *user_deposit.UserDeposit
	ssvNetworkContract      *ssv_network.SsvNetwork
	ssvNetworkViewsContract *ssv_network_views.SsvNetworkViews
	ssvTokenContract        *erc20.Erc20
	withdrawContract        *withdraw.Withdraw

	ssvNetworkAbi abi.ABI

	nextKeyIndex               int
	dealedEth1Block            uint64 // for offchain state
	validatorsPerOperatorLimit uint64

	validatorsByKeyIndex      map[int]*Validator    // key index => validator, cache all validators(pending/active/exist) by keyIndex
	validatorsByPubkey        map[string]*Validator // pubkey => validator, cache all validators(pending/active/exist) by pubkey
	validatorsByValIndex      map[uint64]*Validator // val index => validator
	validatorsByValIndexMutex sync.RWMutex

	eth2Config beacon.Eth2Config

	// ssv offchain state
	clusters                 map[string]*Cluster // cluster key => cluster
	feeRecipientAddressOnSsv common.Address
}

type Cluster struct {
	operators               []*keyshare.Operator
	operatorIds             []uint64
	latestCluster           *ssv_network.ISSVNetworkCoreCluster
	latestRegistrationNonce uint64

	balance decimal.Decimal

	managingValidators map[int]struct{} // key index

	latestUpdateClusterBlockNumber    uint64
	latestValidatorAddedBlockNumber   uint64
	latestValidatorRemovedBlockNumber uint64
}

type Validator struct {
	privateKey *bls.PrivateKey
	keyIndex   int

	statusOnStafi  uint8
	statusOnSsv    uint8
	statusOnBeacon uint8

	validatorIndex uint64
	exitEpoch      uint64

	clusterKey string
}

const (
	valStatusUnInitiated = uint8(0)

	// on stafi
	valStatusDeposited = uint8(1)
	valStatusMatch     = uint8(2)
	valStatusUnmatch   = uint8(3)
	valStatusStaked    = uint8(4)

	// on ssv
	valStatusRegistedOnSsv = uint8(1)
	valStatusRemovedOnSsv  = uint8(2)

	// on beacon
	valStatusActiveOnBeacon = uint8(1)
	valStatusExitedOnBeacon = uint8(2)
)

func NewTask(cfg *config.Config, seed []byte, superNodeKeyPair, ssvKeyPair *secp256k1.Keypair) (*Task, error) {
	if !common.IsHexAddress(cfg.Contracts.StorageContractAddress) {
		return nil, fmt.Errorf("storage contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.SsvNetworkAddress) {
		return nil, fmt.Errorf("ssvnetwork contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.SsvNetworkViewsAddress) {
		return nil, fmt.Errorf("ssvnetworkviews contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.SsvTokenAddress) {
		return nil, fmt.Errorf("SsvTokenAddress contract address fmt err")
	}

	gasLimitDeci, err := decimal.NewFromString(cfg.GasLimit)
	if err != nil {
		return nil, err
	}

	if gasLimitDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("gas limit is zero")
	}
	maxGasPriceDeci, err := decimal.NewFromString(cfg.MaxGasPrice)
	if err != nil {
		return nil, err
	}
	if maxGasPriceDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("max gas price is zero")
	}

	eth1client, err := ethclient.Dial(cfg.Eth1Endpoint)
	if err != nil {
		return nil, err
	}
	s := &Task{
		stop:             make(chan struct{}),
		eth1Endpoint:     cfg.Eth1Endpoint,
		eth2Endpoint:     cfg.Eth2Endpoint,
		eth1Client:       eth1client,
		superNodeKeyPair: superNodeKeyPair,
		ssvKeyPair:       ssvKeyPair,
		seed:             seed,
		gasLimit:         gasLimitDeci.BigInt(),
		maxGasPrice:      maxGasPriceDeci.BigInt(),

		eth1StartHeight:                utils.TheMergeBlockNumber,
		storageContractAddress:         common.HexToAddress(cfg.Contracts.StorageContractAddress),
		ssvNetworkContractAddress:      common.HexToAddress(cfg.Contracts.SsvNetworkAddress),
		ssvNetworkViewsContractAddress: common.HexToAddress(cfg.Contracts.SsvNetworkViewsAddress),
		ssvTokenContractAddress:        common.HexToAddress(cfg.Contracts.SsvTokenAddress),

		validatorsByKeyIndex: make(map[int]*Validator),
		validatorsByPubkey:   make(map[string]*Validator),
		validatorsByValIndex: make(map[uint64]*Validator),

		clusters: make(map[string]*Cluster),
	}

	return s, nil
}

func (task *Task) Start() error {
	var err error
	task.connectionOfSuperNodeAccount, err = connection.NewConnection(task.eth1Endpoint, task.eth2Endpoint, task.superNodeKeyPair,
		task.gasLimit, task.maxGasPrice)
	if err != nil {
		return err
	}
	task.connectionOfSsvAccount, err = connection.NewConnection(task.eth1Endpoint, task.eth2Endpoint, task.ssvKeyPair,
		task.gasLimit, task.maxGasPrice)
	if err != nil {
		return err
	}
	chainId, err := task.eth1Client.ChainID(context.Background())
	if err != nil {
		return err
	}

	// task.eth2Config, err = task.connectionOfSuperNodeAccount.Eth2Client().GetEth2Config()
	// if err != nil {
	// 	return err
	// }

	switch chainId.Uint64() {
	case 1: //mainnet
		task.dev = false
		task.chain = constants.GetChain(constants.ChainMAINNET)
		// if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.MainnetConfig().GenesisForkVersion) {
		// 	return fmt.Errorf("endpoint network not match")
		// }
		task.dealedEth1Block = 17705353
		task.ssvApiNetwork = "mainnet"
		task.postUptimeUrl = mainnetPostUptimeUrl

	case 11155111: // sepolia
		task.dev = true
		task.chain = constants.GetChain(constants.ChainSEPOLIA)
		// if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.SepoliaConfig().GenesisForkVersion) {
		// 	return fmt.Errorf("endpoint network not match")
		// }
		task.dealedEth1Block = 9354882
		task.ssvApiNetwork = "prater"
		task.postUptimeUrl = devPostUptimeUrl
	case 5: // goerli
		task.dev = true
		task.chain = constants.GetChain(constants.ChainGOERLI)
		// if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.PraterConfig().GenesisForkVersion) {
		// 	return fmt.Errorf("endpoint network not match")
		// }
		task.dealedEth1Block = 9403883
		task.ssvApiNetwork = "prater"
		task.postUptimeUrl = devPostUptimeUrl

	default:
		return fmt.Errorf("unsupport chainId: %d", chainId.Int64())
	}
	if err != nil {
		return err
	}

	task.ssvNetworkAbi, err = abi.JSON(strings.NewReader(ssv_network.SsvNetworkABI))
	if err != nil {
		return err
	}

	err = task.initContract()
	if err != nil {
		return err
	}

	err = task.initValNextKeyIndex()
	if err != nil {
		return err
	}
	logrus.Infof("nextKeyIndex: %d", task.nextKeyIndex)

	utils.SafeGo(task.ssvHandler)
	utils.SafeGo(task.monitorHandler)
	utils.SafeGo(task.uptimeHandler)

	return nil
}

func (task *Task) Stop() {
	close(task.stop)
}

func (task *Task) copySeed() []byte {
	copyBts := make([]byte, len(task.seed))
	copy(copyBts, task.seed)
	return copyBts
}

func (task *Task) calClusterNeedDepositAmount(cluster *Cluster) (min, max *big.Int, err error) {
	networkFee, err := task.ssvNetworkViewsContract.GetNetworkFee(nil)
	if err != nil {
		return nil, nil, err
	}
	networkFeeDeci := decimal.NewFromBigInt(networkFee, 0)

	ltp, err := task.ssvNetworkViewsContract.GetLiquidationThresholdPeriod(nil)
	if err != nil {
		return nil, nil, err
	}
	ltpDeci := decimal.NewFromInt(int64(ltp))

	balance, err := task.ssvNetworkViewsContract.GetBalance(nil, task.ssvKeyPair.CommonAddress(), cluster.operatorIds,
		ssv_network_views.ISSVNetworkCoreCluster(*cluster.latestCluster))
	if err != nil {
		return nil, nil, err
	}
	balanceDeci := decimal.NewFromBigInt(balance, 0)

	totalOpFee := decimal.Zero
	for _, op := range cluster.operators {
		totalOpFee = totalOpFee.Add(op.Fee)
	}
	totalOpFee = totalOpFee.Add(networkFeeDeci)

	valAmount := decimal.NewFromInt(int64(len(cluster.managingValidators) + 1))

	maxExpected := valAmount.Mul(totalOpFee).Mul(blocksOfOneYear.Add(ltpDeci))
	minExpected := valAmount.Mul(totalOpFee).Mul(ltpDeci.Mul(decimal.NewFromInt(2)))

	if maxExpected.LessThan(minExpected) {
		maxExpected = minExpected
	}

	switch {
	case balanceDeci.GreaterThanOrEqual(maxExpected):
		return big.NewInt(0), big.NewInt(0), nil
	case balanceDeci.GreaterThanOrEqual(minExpected) && balanceDeci.LessThan(maxExpected):
		return big.NewInt(0), maxExpected.Sub(balanceDeci).BigInt(), nil
	case balanceDeci.LessThan(minExpected):
		return maxExpected.Sub(balanceDeci).BigInt(), maxExpected.Sub(balanceDeci).BigInt(), nil
	default:
		return nil, nil, fmt.Errorf("unreached balance")
	}

}

func (task *Task) initContract() error {
	storageContract, err := storage.NewStorage(task.storageContractAddress, task.eth1Client)
	if err != nil {
		return err
	}
	stafiWithdrawAddress, err := utils.GetContractAddress(storageContract, "stafiWithdraw")
	if err != nil {
		return err
	}
	task.eth1WithdrawalAdress = stafiWithdrawAddress
	logrus.Debugf("stafiWithdraw address: %s", task.eth1WithdrawalAdress.String())

	task.withdrawContract, err = withdraw.NewWithdraw(stafiWithdrawAddress, task.eth1Client)
	if err != nil {
		return err
	}

	superNodeAddress, err := utils.GetContractAddress(storageContract, "stafiSuperNode")
	if err != nil {
		return err
	}
	task.superNodeContract, err = super_node.NewSuperNode(superNodeAddress, task.eth1Client)
	if err != nil {
		return err
	}

	userDepositAddress, err := utils.GetContractAddress(storageContract, "stafiUserDeposit")
	if err != nil {
		return err
	}
	task.userDepositContract, err = user_deposit.NewUserDeposit(userDepositAddress, task.eth1Client)
	if err != nil {
		return err
	}

	superNodeFeePoolAddress, err := utils.GetContractAddress(storageContract, "stafiSuperNodeFeePool")
	if err != nil {
		return err
	}
	if (common.Address{}) == superNodeFeePoolAddress {
		return fmt.Errorf("superNodeFeePoolAddress is zero")
	}
	task.feeRecipientAddressOnStafi = superNodeFeePoolAddress

	task.ssvNetworkContract, err = ssv_network.NewSsvNetwork(task.ssvNetworkContractAddress, task.eth1Client)
	if err != nil {
		return err
	}
	task.ssvNetworkViewsContract, err = ssv_network_views.NewSsvNetworkViews(task.ssvNetworkViewsContractAddress, task.eth1Client)
	if err != nil {
		return err
	}

	task.ssvTokenContract, err = erc20.NewErc20(task.ssvTokenContractAddress, task.eth1Client)
	if err != nil {
		return err
	}
	return nil
}

func (task *Task) mustGetSuperNodePubkeyStatus(pubkey []byte) (uint8, error) {
	retry := 0
	var pubkeyStatus *big.Int
	var err error
	for {
		if retry > utils.RetryLimit {
			return 0, fmt.Errorf("updateValStatus reach retry limit")
		}
		pubkeyStatus, err = task.superNodeContract.GetSuperNodePubkeyStatus(nil, pubkey)
		if err != nil {
			logrus.Warnf("GetSuperNodePubkeyStatus err: %s", err.Error())
			time.Sleep(utils.RetryInterval)
			retry++
			continue
		}
		break
	}

	return uint8(pubkeyStatus.Uint64()), nil
}

func (task *Task) ssvHandler() {
	logrus.Info("start handler")
	retry := 0
	for {
		if retry > utils.RetryLimit {
			utils.ShutdownRequestChannel <- struct{}{}
			return
		}

		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		default:

			err := task.checkAndRepairValNexKeyIndex()
			if err != nil {
				logrus.Warnf("checkAndRepairValNe err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			err = task.updateSsvOffchainState()
			if err != nil {
				logrus.Warnf("updateOffchainState err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			err = task.updateValStatus()
			if err != nil {
				logrus.Warnf("updateValStatus err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}
			err = task.updateOperatorStatus()
			if err != nil {
				logrus.Warnf("updateOperatorStatus err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			// -------- stafi

			err = task.checkAndStake()
			if err != nil {
				logrus.Warnf("checkAndStake err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			err = task.checkAndDeposit()
			if err != nil {
				logrus.Warnf("checkAndDeposit err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			// -------- ssv

			err = task.checkAndSetFeeRecipient()
			if err != nil {
				logrus.Warnf("checkAndSetFeeRecipient err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			err = task.checkAndReactiveOnSSV()
			if err != nil {
				logrus.Warnf("checkAndReactiveOnSSV err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			err = task.checkAndOnboardOnSSV()
			if err != nil {
				logrus.Warnf("checkAndOnboardOnSSV err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			err = task.checkAndOffboardOnSSV()
			if err != nil {
				logrus.Warnf("checkAndOffboardOnSSV err %s", err)
				time.Sleep(utils.RetryInterval)
				retry++
				continue
			}

			retry = 0
		}

		time.Sleep(24 * time.Second) // 2 blocks
	}
}
