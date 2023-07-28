// Copyright 2020 Stafi Protocol
// SPDX-License-Identifier: LGPL-3.0-only

package connection

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon/client"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/types"
	"golang.org/x/sync/errgroup"
)

var Gwei5 = big.NewInt(5e9)
var Gwei10 = big.NewInt(10e9)
var Gwei20 = big.NewInt(20e9)

var retryLimit = 100
var waitInterval = 6 * time.Second

type Connection struct {
	eth1Endpoint string
	eth2Endpoint string
	kp           *secp256k1.Keypair
	gasLimit     *big.Int
	maxGasPrice  *big.Int
	eth1Client   *ethclient.Client
	eth1Rpc      *rpc.Client
	eth2Client   *client.StandardHttpClient
	txOpts       *bind.TransactOpts
	callOpts     bind.CallOpts
	optsLock     sync.Mutex
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(eth1Endpoint, eth2Endpoint string, kp *secp256k1.Keypair, gasLimit, maxGasPrice *big.Int) (*Connection, error) {
	if kp != nil {
		if maxGasPrice.Cmp(big.NewInt(0)) <= 0 {
			return nil, fmt.Errorf("max gas price empty")
		}
		if gasLimit.Cmp(big.NewInt(0)) <= 0 {
			return nil, fmt.Errorf("gas limit empty")
		}
	}
	c := &Connection{
		eth1Endpoint: eth1Endpoint,
		eth2Endpoint: eth2Endpoint,
		kp:           kp,
		gasLimit:     gasLimit,
		maxGasPrice:  maxGasPrice,
	}
	err := c.connect()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Connect starts the ethereum WS connection
func (c *Connection) connect() error {
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	u, err := url.Parse(c.eth1Endpoint)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case "http", "https":
		rpcClient, err = rpc.DialHTTP(c.eth1Endpoint)
	case "ws", "wss":
		rpcClient, err = rpc.DialWebsocket(context.Background(), c.eth1Endpoint, fmt.Sprintf("/%s", u.Scheme))
	default:
		err = fmt.Errorf("unsupport scheme: %s", u.Scheme)
	}
	if err != nil {
		return err
	}

	c.eth1Client = ethclient.NewClient(rpcClient)

	c.eth1Rpc = rpcClient

	// eth2 client
	if len(c.eth2Endpoint) != 0 {
		c.eth2Client, err = client.NewStandardHttpClient(c.eth2Endpoint)
		if err != nil {
			return err
		}
	}

	if c.kp != nil {
		// Construct tx opts, call opts, and nonce mechanism
		opts, err := c.newTransactOpts(big.NewInt(0), c.gasLimit)
		if err != nil {
			return err
		}
		c.txOpts = opts
		c.callOpts = bind.CallOpts{Pending: false, From: c.kp.CommonAddress(), BlockNumber: nil, Context: context.Background()}
	} else {
		c.callOpts = bind.CallOpts{Pending: false, From: common.Address{}, BlockNumber: nil, Context: context.Background()}
	}
	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit *big.Int) (*bind.TransactOpts, error) {
	privateKey := c.kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.eth1Client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	chainId, err := c.eth1Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(gasLimit.Int64())
	auth.Context = context.Background()

	return auth, nil
}

func (c *Connection) Keypair() *secp256k1.Keypair {
	return c.kp
}

func (c *Connection) Eth1Client() *ethclient.Client {
	return c.eth1Client
}

func (c *Connection) Eth2Client() *client.StandardHttpClient {
	return c.eth2Client
}

func (c *Connection) TxOpts() *bind.TransactOpts {
	return c.txOpts
}

func (c *Connection) CallOpts(blocknumber *big.Int) *bind.CallOpts {
	newCallOpts := c.callOpts
	newCallOpts.BlockNumber = blocknumber
	return &newCallOpts
}

// return suggest gastipcap gasfeecap
func (c *Connection) SafeEstimateFee(ctx context.Context) (*big.Int, *big.Int, error) {
	gasTipCap, err := c.eth1Client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, err
	}
	gasFeeCap, err := c.eth1Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, err
	}

	if gasFeeCap.Cmp(Gwei20) < 0 {
		gasFeeCap = new(big.Int).Add(gasFeeCap, Gwei5)
	} else {
		gasFeeCap = new(big.Int).Add(gasFeeCap, Gwei10)
	}

	if gasFeeCap.Cmp(c.maxGasPrice) > 0 {
		gasFeeCap = c.maxGasPrice
	}

	return gasTipCap, gasFeeCap, nil
}

// LockAndUpdateOpts acquires a lock on the opts before updating the nonce
// and gas price.
func (c *Connection) LockAndUpdateTxOpts() error {
	c.optsLock.Lock()

	gasTipCap, gasFeeCap, err := c.SafeEstimateFee(context.Background())
	if err != nil {
		c.optsLock.Unlock()
		return err
	}
	c.txOpts.GasTipCap = gasTipCap
	c.txOpts.GasFeeCap = gasFeeCap

	nonce, err := c.eth1Client.NonceAt(context.Background(), c.txOpts.From, nil)
	if err != nil {
		c.optsLock.Unlock()
		return err
	}
	c.txOpts.Nonce.SetUint64(nonce)
	return nil
}

func (c *Connection) UnlockTxOpts() {
	c.optsLock.Unlock()
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) Eth1LatestBlock() (uint64, error) {
	header, err := c.eth1Client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return header, nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Connection) EnsureHasBytecode(addr common.Address) error {
	code, err := c.eth1Client.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

func (c *Connection) Eth2BeaconHead() (beacon.BeaconHead, error) {
	return c.eth2Client.GetBeaconHead()
}

func (c *Connection) GetValidatorStatus(pubkey types.ValidatorPubkey, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {
	var retErr error
	for i := 0; i < retryLimit; i++ {
		status, err := c.eth2Client.GetValidatorStatus(pubkey, opts)
		if err != nil {
			retErr = err
			logrus.Warnf("eth2Client.GetValidatorStatus err: %s", err)
			time.Sleep(waitInterval)
			continue
		}
		return status, nil
	}
	return beacon.ValidatorStatus{}, fmt.Errorf("eth2Client.GetValidatorStatus reach RetryLimit, err: %s", retErr)

}

func (c *Connection) GetValidatorStatuses(pubkeys []types.ValidatorPubkey, opts *beacon.ValidatorStatusOptions) (map[types.ValidatorPubkey]beacon.ValidatorStatus, error) {
	var retErr error
	for i := 0; i < retryLimit; i++ {
		status, err := c.eth2Client.GetValidatorStatuses(pubkeys, opts)
		if err != nil {
			retErr = err
			logrus.Warnf("eth2Client.GetValidatorStatuses err: %s", err)
			time.Sleep(waitInterval)
			continue
		}
		return status, nil
	}
	return nil, fmt.Errorf("eth2Client.GetValidatorStatuses reach RetryLimit, err: %s", retErr)

}

func (c *Connection) GetBeaconBlock(blockId uint64) (beacon.BeaconBlock, bool, error) {
	var retErr error
	for i := 0; i < retryLimit; i++ {
		status, ok, err := c.eth2Client.GetBeaconBlock(blockId)
		if err != nil {
			retErr = err
			logrus.Warnf("eth2Client.GetBeaconBlock err: %s", err)
			time.Sleep(waitInterval)
			continue
		}
		return status, ok, nil
	}
	return beacon.BeaconBlock{}, false, fmt.Errorf("eth2Client.GetBeaconBlock reach RetryLimit, err: %s", retErr)

}

func (c *Connection) GetValidatorProposerDuties(epoch uint64) (map[uint64]uint64, error) {
	var retErr error
	for i := 0; i < retryLimit; i++ {
		status, err := c.eth2Client.GetValidatorProposerDuties(epoch)
		if err != nil {
			retErr = err
			logrus.Warnf("GetValidatorProposerDuties err: %s, epoch: %d", err, epoch)
			time.Sleep(waitInterval)
			continue
		}
		return status, nil
	}
	return nil, fmt.Errorf("GetValidatorProposerDuties reach RetryLimit, err: %s", retErr)

}

func (c *Connection) GetValidatorStatusByIndex(index string, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {
	var retErr error
	for i := 0; i < retryLimit; i++ {
		status, err := c.eth2Client.GetValidatorStatusByIndex(index, opts)
		if err != nil {
			retErr = err
			logrus.Warnf("eth2Client.GetValidatorStatusByIndex err: %s", err)
			time.Sleep(waitInterval)
			continue
		}
		return status, nil
	}
	return beacon.ValidatorStatus{}, fmt.Errorf("eth2Client.GetValidatorStatusByIndex reach RetryLimit, err: %s", retErr)
}

func (c *Connection) GetSyncCommitteesForEpoch(epoch uint64) ([]beacon.SyncCommittee, error) {
	var retErr error
	for i := 0; i < retryLimit; i++ {
		status, err := c.eth2Client.GetSyncCommitteesForEpoch(epoch)
		if err != nil {
			if strings.Contains(err.Error(), "has no sync committee") {
				return []beacon.SyncCommittee{}, nil
			}
			retErr = err
			logrus.Warnf("eth2Client.GetSyncCommitteesForEpoch err: %s", err)
			time.Sleep(waitInterval)
			continue
		}
		return status, nil
	}
	return nil, fmt.Errorf("eth2Client.GetSyncCommitteesForEpoch reach RetryLimit, err: %s", retErr)
}

func (c *Connection) GetELRewardForBlock(executionBlockNumber uint64) (*big.Int, error) {

	block, err := c.eth1Client.BlockByNumber(context.Background(), big.NewInt(int64(executionBlockNumber)))
	if err != nil {
		return nil, err
	}

	if len(block.Transactions()) == 0 {
		return big.NewInt(0), nil
	}

	txHashes := []common.Hash{}
	for _, tx := range block.Transactions() {
		txHashes = append(txHashes, tx.Hash())
	}

	var txReceipts []*client.TxReceipt
	for j := 1; j <= 16; j++ { // retry up to 16 times
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
		txReceipts, err = c.batchRequestReceipts(ctx, txHashes)
		if err == nil {
			cancel()
			break
		} else {
			logrus.Infof("error (%d) doing batchRequestReceipts for execution block %v: %v", j, executionBlockNumber, err)
			time.Sleep(time.Duration(j) * time.Second)
		}
		cancel()
	}
	if err != nil {
		return nil, fmt.Errorf("error doing batchRequestReceipts for execution block %v: %w", executionBlockNumber, err)
	}

	totalTxFee := big.NewInt(0)
	for _, r := range txReceipts {
		if r.EffectiveGasPrice == nil {
			return nil, fmt.Errorf("no EffectiveGasPrice for execution block %v: %v", executionBlockNumber, txHashes)
		}
		txFee := new(big.Int).Mul(r.EffectiveGasPrice.ToInt(), new(big.Int).SetUint64(uint64(r.GasUsed)))
		totalTxFee.Add(totalTxFee, txFee)
	}

	// base fee per gas is stored little-endian but we need it
	// big-endian for big.Int.

	burntFee := new(big.Int).Mul(block.BaseFee(), new(big.Int).SetUint64(block.GasUsed()))

	totalTxFee.Sub(totalTxFee, burntFee)

	return totalTxFee, nil
}

func (c *Connection) batchRequestReceipts(ctx context.Context, txHashes []common.Hash) ([]*client.TxReceipt, error) {
	elems := make([]rpc.BatchElem, 0, len(txHashes))
	errors := make([]error, 0, len(txHashes))
	txReceipts := make([]*client.TxReceipt, len(txHashes))
	for i, h := range txHashes {
		txReceipt := &client.TxReceipt{}
		err := error(nil)
		elems = append(elems, rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{h.Hex()},
			Result: txReceipt,
			Error:  err,
		})
		txReceipts[i] = txReceipt
		errors = append(errors, err)
	}

	ioErr := c.eth1Rpc.BatchCallContext(ctx, elems)
	if ioErr != nil {
		return nil, fmt.Errorf("io-error when fetching tx-receipts: %w", ioErr)
	}
	for _, e := range errors {
		if e != nil {
			return nil, fmt.Errorf("error when fetching tx-receipts: %w", e)
		}
	}
	return txReceipts, nil
}

// if validator not exist on beacon chain will return err
// if exit after epoch will return zero reward
func (c *Connection) GetRewardsForEpochWithValidators(epoch uint64, valIndexs []uint64) (map[uint64]*client.ValidatorEpochIncome, error) {
	valIndexStrs := make([]string, 0)
	valIndexMap := make(map[uint64]bool)
	for _, index := range valIndexs {
		if !valIndexMap[index] {
			valIndexStrs = append(valIndexStrs, fmt.Sprintf("%d", index))
			valIndexMap[index] = true
		}
	}

	logrus.Trace("GetRewardsForEpochWithValidators", valIndexStrs)

	proposerAssignments, err := c.eth2Client.ProposerAssignments(epoch)
	if err != nil {
		return nil, fmt.Errorf("client.ProposerAssignments %s", err)
	}

	slotsPerEpoch := uint64(len(proposerAssignments.Data))

	startSlot := epoch * slotsPerEpoch
	endSlot := startSlot + slotsPerEpoch - 1

	g := new(errgroup.Group)
	g.SetLimit(16)

	slotsToProposerIndex := make(map[uint64]uint64)
	for _, pa := range proposerAssignments.Data {
		slotsToProposerIndex[uint64(pa.Slot)] = uint64(pa.ValidatorIndex)
	}

	rewardsMux := &sync.Mutex{}

	rewards := make(map[uint64]*client.ValidatorEpochIncome)

	for i := startSlot + 1; i <= endSlot; i++ {
		slot := i
		g.Go(func() error {
			proposer, found := slotsToProposerIndex[slot]
			if !found {
				return fmt.Errorf("assigned proposer for slot %v not found", slot)
			}

			// get sync rewards
			syncRewards, err := c.eth2Client.SyncCommitteeRewards(slot)
			if err != nil {
				switch {
				case strings.Contains(err.Error(), client.ErrBlockNotFound.Error()):
					// block not exit, should return
					return nil
				case strings.Contains(err.Error(), client.ErrSlotPreSyncCommittees.Error()):
					// skip err
				default:
					return fmt.Errorf("client.SyncCommitteeRewards err %s, slot: %d", err, slot)
				}
			}

			rewardsMux.Lock()
			if syncRewards != nil {
				for _, sr := range syncRewards.Data {
					if !valIndexMap[sr.ValidatorIndex] {
						continue
					}

					if rewards[sr.ValidatorIndex] == nil {
						rewards[sr.ValidatorIndex] = &client.ValidatorEpochIncome{}
					}

					if sr.Reward > 0 {
						rewards[sr.ValidatorIndex].SyncCommitteeReward += uint64(sr.Reward)
					} else {
						rewards[sr.ValidatorIndex].SyncCommitteePenalty += uint64(sr.Reward * -1)
					}
				}
			}
			rewardsMux.Unlock()

			if !valIndexMap[proposer] {
				return nil
			}

			// get proposer fee reward
			// execBlockNumber, err := c.eth2Client.ExecutionBlockNumber(slot)
			// rewardsMux.Lock()
			// if rewards[proposer] == nil {
			// 	rewards[proposer] = &client.ValidatorEpochIncome{}
			// }
			// rewardsMux.Unlock()

			// if err != nil {
			// 	if err == client.ErrBlockNotFound {
			// 		rewardsMux.Lock()
			// 		rewards[proposer].ProposalsMissed += 1
			// 		rewardsMux.Unlock()
			// 		return nil
			// 	} else if err != client.ErrSlotPreMerge { // ignore
			// 		logrus.Errorf("error retrieving execution block number for slot %v: %v", slot, err)
			// 		return err
			// 	}
			// } else {
			// 	txFeeIncome, err := c.GetELRewardForBlock(execBlockNumber)
			// 	if err != nil {
			// 		return fmt.Errorf("elrewards.GetELRewardForBlock %s", err)
			// 	}

			// 	rewardsMux.Lock()
			// 	rewards[proposer].TxFeeRewardWei = txFeeIncome.Bytes()
			// 	rewardsMux.Unlock()
			// }

			// get proposer block rewards
			blockRewards, err := c.eth2Client.BlockRewards(slot)
			if err != nil {
				return fmt.Errorf("client.BlockRewards %s", err)
			}

			rewardsMux.Lock()
			if rewards[blockRewards.Data.ProposerIndex] == nil {
				rewards[blockRewards.Data.ProposerIndex] = &client.ValidatorEpochIncome{}
			}
			rewards[blockRewards.Data.ProposerIndex].ProposerAttestationInclusionReward += blockRewards.Data.Attestations
			rewards[blockRewards.Data.ProposerIndex].ProposerSlashingInclusionReward += blockRewards.Data.AttesterSlashings + blockRewards.Data.ProposerSlashings
			rewards[blockRewards.Data.ProposerIndex].ProposerSyncInclusionReward += blockRewards.Data.SyncAggregate
			rewardsMux.Unlock()
			return nil
		})
	}

	g.Go(func() error {
		// get attestion reward
		ar, err := c.eth2Client.AttestationRewardsWithVals(epoch, valIndexStrs)
		if err != nil {
			return errors.Wrapf(err, "eth2Client.AttestationRewardsWithVals failed, epoch: %d", epoch)
		}
		rewardsMux.Lock()
		defer rewardsMux.Unlock()
		for _, ar := range ar.Data.TotalRewards {
			if !valIndexMap[ar.ValidatorIndex] {
				continue
			}

			if rewards[ar.ValidatorIndex] == nil {
				rewards[ar.ValidatorIndex] = &client.ValidatorEpochIncome{}
			}

			if ar.Head >= 0 {
				rewards[ar.ValidatorIndex].AttestationHeadReward = uint64(ar.Head)
			} else {
				return fmt.Errorf("retrieved negative attestation head reward for validator %v: %v", ar.ValidatorIndex, ar.Head)
			}

			if ar.Source > 0 {
				rewards[ar.ValidatorIndex].AttestationSourceReward = uint64(ar.Source)
			} else {
				rewards[ar.ValidatorIndex].AttestationSourcePenalty = uint64(ar.Source * -1)
			}

			if ar.Target > 0 {
				rewards[ar.ValidatorIndex].AttestationTargetReward = uint64(ar.Target)
			} else {
				rewards[ar.ValidatorIndex].AttestationTargetPenalty = uint64(ar.Target * -1)
			}

			if ar.InclusionDelay <= 0 {
				rewards[ar.ValidatorIndex].FinalityDelayPenalty = uint64(ar.InclusionDelay * -1)
			} else {
				return fmt.Errorf("retrieved positive inclusion delay penalty for validator %v: %v", ar.ValidatorIndex, ar.InclusionDelay)
			}
		}

		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, err
	}

	return rewards, nil
}
