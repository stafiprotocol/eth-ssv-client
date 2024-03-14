package task

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/go-multicall"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v3/config/params"
	types "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v3/encoding/bytesutil"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	erc20 "github.com/stafiprotocol/eth-ssv-client/bindings/Erc20"
	"github.com/stafiprotocol/eth-ssv-client/bindings/Settings"
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

	blocksOfOneYear = decimal.NewFromInt(2629800)
)

const minNumberOfTargetOperators = 10

var (
	devPostUptimeUrl     = "https://test-drop-api.stafi.io/reth/v1/uploadEjectorUptime"
	mainnetPostUptimeUrl = "https://drop-api.stafi.io/reth/v1/uploadEjectorUptime"
)

var (
	domainVoluntaryExit  = bytesutil.Uint32ToBytes4(0x04000000)
	shardCommitteePeriod = types.Epoch(256) // ShardCommitteePeriod is the minimum amount of epochs a validator must participate before exiting.
)

const (
	valStatusUnInitiated = uint8(0)

	// on stafi
	valStatusDeposited = uint8(1)
	valStatusMatch     = uint8(2)
	valStatusUnmatch   = uint8(3)
	valStatusStaked    = uint8(4)

	// on ssv
	valStatusRegistedOnSsvValid   = uint8(1)
	valStatusRegistedOnSsvInvalid = uint8(2)
	valStatusRemovedOnSsv         = uint8(3)

	// on beacon
	valStatusActiveOnBeacon = uint8(1)
	valStatusExitedOnBeacon = uint8(2)
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

	gasLimit            *big.Int
	maxGasPrice         *big.Int
	poolReservedBalance *big.Int
	seed                []byte
	postUptimeUrl       string
	isViewMode          bool
	targetOperatorIds   []uint64
	operatorPubkeys     map[uint64]string

	storageContractAddress         common.Address
	ssvNetworkContractAddress      common.Address
	ssvNetworkViewsContractAddress common.Address
	ssvTokenContractAddress        common.Address

	// --- need init on start
	dev           bool
	ssvApiNetwork string
	chain         constants.Chain

	connectionOfSuperNodeAccount *connection.Connection
	connectionOfSsvAccount       *connection.Connection

	eth1WithdrawalAdress       common.Address
	feeRecipientAddressOnStafi common.Address
	latestRegistrationNonce    uint64
	latestTxBlock              uint64

	eth1Client *ethclient.Client
	eth2Config beacon.Eth2Config

	superNodeContract       *super_node.SuperNode
	userDepositContract     *user_deposit.UserDeposit
	ssvNetworkContract      *ssv_network.SsvNetwork
	ssvNetworkViewsContract *ssv_network_views.SsvNetworkViews
	ssvTokenContract        *erc20.Erc20
	withdrawContract        *withdraw.Withdraw
	networkSettingsContract *network_settings.NetworkSettings

	multicaler *multicall.Caller

	ssvNetworkAbi abi.ABI

	nextKeyIndex                int
	dealedEth1Block             uint64 // for offchain state
	validatorsPerOperatorLimit  uint64
	ValidatorsPerSuperNodeLimit uint64
	ValidatorsLimitByGas        uint64 // gas = 162917*n+268921

	validatorsByKeyIndex      map[int]*Validator    // key index => validator, cache all validators(pending/active/exist) by keyIndex
	validatorsByPubkey        map[string]*Validator // pubkey => validator, cache all validators(pending/active/exist) by pubkey
	validatorsByValIndex      map[uint64]*Validator // val index => validator
	validatorsByValIndexMutex sync.RWMutex

	targetOperators map[uint64]*keyshare.Operator

	// ssv offchain state
	clusters                 map[string]*Cluster // cluster key => cluster
	feeRecipientAddressOnSsv common.Address

	handlers     []func() error
	handlersName []string
}

type Cluster struct {
	operatorIds   []uint64
	latestCluster *ssv_network.ISSVNetworkCoreCluster

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

	removedFromSsvOnBlock uint64
}

func NewTask(cfg *config.Config, seed []byte, isViewMode bool, superNodeKeyPair, ssvKeyPair *secp256k1.Keypair) (*Task, error) {
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
	if len(cfg.TargetOperators) < minNumberOfTargetOperators {
		return nil, fmt.Errorf("target operators number %d less than %d", len(cfg.TargetOperators), minNumberOfTargetOperators)
	}

	gasLimitDeci, err := decimal.NewFromString(cfg.GasLimit)
	if err != nil {
		return nil, err
	}

	if gasLimitDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("gas limit is zero")
	}

	// gas = 162917*n+268921
	n := gasLimitDeci.Sub(decimal.NewFromInt(268921)).Div(decimal.NewFromInt(162917))
	validatorsLimitByGas := n.BigInt().Uint64()
	if validatorsLimitByGas == 0 {
		return nil, fmt.Errorf("gasLimit %s too small", cfg.GasLimit)
	}
	logrus.Debug("validatorsLimitByGas ", validatorsLimitByGas)

	maxGasPriceDeci, err := decimal.NewFromString(cfg.MaxGasPrice)
	if err != nil {
		return nil, err
	}
	if maxGasPriceDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("max gas price is zero")
	}

	poolReservedBalance := big.NewInt(0)
	if len(cfg.PoolReservedBalance) > 0 {
		reservedBalance, err := decimal.NewFromString(cfg.PoolReservedBalance)
		if err != nil {
			return nil, err
		}
		if maxGasPriceDeci.IsNegative() {
			return nil, fmt.Errorf("PoolReservedBalance is negative")
		}
		poolReservedBalance = reservedBalance.BigInt()
	}

	eth1client, err := ethclient.Dial(cfg.Eth1Endpoint)
	if err != nil {
		return nil, err
	}

	opPubkes := make(map[uint64]string)
	for _, op := range cfg.Operators {
		_, err := base64.StdEncoding.DecodeString(op.Pubkey)
		if err != nil {
			logrus.Warnf("operator: %d pubkey decode err, will skip this operator. pubkey: %s", op.Id, op.Pubkey)
			continue
		}
		opPubkes[op.Id] = op.Pubkey
	}

	s := &Task{
		stop:                 make(chan struct{}),
		eth1Endpoint:         cfg.Eth1Endpoint,
		eth2Endpoint:         cfg.Eth2Endpoint,
		eth1Client:           eth1client,
		superNodeKeyPair:     superNodeKeyPair,
		ssvKeyPair:           ssvKeyPair,
		seed:                 seed,
		isViewMode:           isViewMode,
		gasLimit:             gasLimitDeci.BigInt(),
		maxGasPrice:          maxGasPriceDeci.BigInt(),
		poolReservedBalance:  poolReservedBalance,
		eth1StartHeight:      utils.TheMergeBlockNumber,
		targetOperatorIds:    cfg.TargetOperators,
		ValidatorsLimitByGas: validatorsLimitByGas,
		operatorPubkeys:      opPubkes,

		storageContractAddress:         common.HexToAddress(cfg.Contracts.StorageContractAddress),
		ssvNetworkContractAddress:      common.HexToAddress(cfg.Contracts.SsvNetworkAddress),
		ssvNetworkViewsContractAddress: common.HexToAddress(cfg.Contracts.SsvNetworkViewsAddress),
		ssvTokenContractAddress:        common.HexToAddress(cfg.Contracts.SsvTokenAddress),

		validatorsByKeyIndex: make(map[int]*Validator),
		validatorsByPubkey:   make(map[string]*Validator),
		validatorsByValIndex: make(map[uint64]*Validator),

		targetOperators: make(map[uint64]*keyshare.Operator),

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

	task.eth2Config, err = task.connectionOfSuperNodeAccount.Eth2Client().GetEth2Config()
	if err != nil {
		return err
	}

	switch chainId.Uint64() {
	case 1: //mainnet
		task.dev = false
		task.chain = constants.GetChain(constants.ChainMAINNET)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.MainnetConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealedEth1Block = 17705353
		task.ssvApiNetwork = "mainnet"
		task.postUptimeUrl = mainnetPostUptimeUrl

	case 11155111: // sepolia
		task.dev = true
		task.chain = constants.GetChain(constants.ChainSEPOLIA)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.SepoliaConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealedEth1Block = 9354882
		task.ssvApiNetwork = "prater"
		task.postUptimeUrl = devPostUptimeUrl
	case 5: // goerli
		task.dev = true
		task.chain = constants.GetChain(constants.ChainGOERLI)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.PraterConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealedEth1Block = 9403883
		task.ssvApiNetwork = "prater"
		task.postUptimeUrl = devPostUptimeUrl

	default:
		return fmt.Errorf("unsupport chainId: %d", chainId.Int64())
	}
	if err != nil {
		return err
	}
	logrus.Infof("ssv network: %s", task.ssvApiNetwork)

	task.ssvNetworkAbi, err = abi.JSON(strings.NewReader(ssv_network.SsvNetworkABI))
	if err != nil {
		return err
	}

	err = task.initContract()
	if err != nil {
		return err
	}

	caller, err := multicall.Dial(context.Background(), task.eth1Endpoint)
	if err != nil {
		return err
	}
	task.multicaler = caller

	// check target operator id
	for _, opId := range task.targetOperatorIds {
		if _, exist := task.targetOperators[opId]; exist {
			return fmt.Errorf("duplicate operator id: %d", opId)
		}

		pubkey, exist := task.operatorPubkeys[opId]
		if !exist {
			return fmt.Errorf("operator %d pubkey not exist", opId)
		}

		_, fee, validatorCount, _, isPrivate, _, err := task.ssvNetworkViewsContract.GetOperatorById(nil, opId)
		if err != nil {
			return errors.Wrap(err, "ssvNetworkViewsContract.GetOperatorById failed")
		}
		if isPrivate {
			return fmt.Errorf("target operator %d is private", opId)
		}

		// fetch acitve status from api
		operatorFromApi, err := task.mustGetOperatorDetail(task.ssvApiNetwork, opId)
		if err != nil {
			return err
		}

		if operatorFromApi.IsActive != 1 {
			return fmt.Errorf("target operator %d is not active", opId)
		}

		task.targetOperators[opId] = &keyshare.Operator{
			Id:             opId,
			PublicKey:      pubkey,
			Fee:            decimal.NewFromBigInt(fee, 0),
			Active:         true,
			ValidatorCount: uint64(validatorCount),
		}
	}

	err = task.initValNextKeyIndex()
	if err != nil {
		return err
	}
	logrus.Infof("nextKeyIndex: %d", task.nextKeyIndex)
	logrus.Infof("init state...")

	retry := 0
	for {
		if retry > utils.RetryLimit {
			return fmt.Errorf("init state reach retry limit, err: %s", err.Error())
		}
		err = task.updateSsvOffchainState()
		if err != nil {
			retry++
			continue
		}
		err = task.updateValStatus()
		if err != nil {
			retry++
			continue
		}
		err = task.updateOperatorStatus()
		if err != nil {
			retry++
			continue
		}

		break
	}

	task.appendHandlers(
		task.checkAndRepairValNexKeyIndex,
		task.updateSsvOffchainState,
		task.updateValStatus,
		task.updateOperatorStatus,
	)
	if !task.isViewMode {
		task.appendHandlers(
			task.checkAndStake, //stafi
			task.checkAndDeposit,
			task.checkAndSetFeeRecipient, // ssv
			task.checkAndReactiveOnSSV,
			task.checkAndOnboardOnSSV,
			task.checkAndOffboardOnSSV,
			task.checkAndWithdrawOnSSV,
		)

		utils.SafeGo(task.ejectorService)
		utils.SafeGo(task.uptimeService)
	}

	utils.SafeGo(task.ssvService)

	return nil
}

func (task *Task) Stop() {
	close(task.stop)
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

	networkSettingsAddress, err := utils.GetContractAddress(storageContract, "stafiNetworkSettings")
	if err != nil {
		return err
	}
	task.networkSettingsContract, err = network_settings.NewNetworkSettings(networkSettingsAddress, task.eth1Client)
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

func (task *Task) appendHandlers(handlers ...func() error) {
	for _, handler := range handlers {

		funcNameRaw := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

		splits := strings.Split(funcNameRaw, "/")
		funcName := splits[len(splits)-1]

		task.handlersName = append(task.handlersName, funcName)
		task.handlers = append(task.handlers, handler)
	}
}

func (task *Task) waitTxOk(txHash common.Hash) error {
	blockNumber, err := utils.WaitTxOkCommon(task.eth1Client, txHash)
	if err != nil {
		return err
	}
	task.latestTxBlock = blockNumber
	return nil
}

func (task *Task) offchainStateIsLatest() bool {
	return task.dealedEth1Block > task.latestTxBlock
}
