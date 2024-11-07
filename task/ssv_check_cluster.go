package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/signing"
	"github.com/prysmaticlabs/prysm/v3/config/params"
	primTypes "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v3/crypto/bls"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/types"
	"math"
)

func (task *Task) checkClusterOnSSV() error {
	latestBlockNumber, err := task.connectionOfSuperNodeAccount.Eth1LatestBlock()
	if err != nil {
		return err
	}

	if latestBlockNumber > task.balanceUpdateEth1Block+100 {
		return nil
	}

	for _, cluster := range task.clusters {
		if cluster.balance.IsZero() {
			for valKeyIndex := range cluster.managingValidators {
				validator, exist := task.validatorsByKeyIndex[valKeyIndex]
				if !exist {
					return fmt.Errorf("validator at index %d not exist", valKeyIndex)
				}

				// check beacon sync status
				syncStatus, err := task.connectionOfSuperNodeAccount.Eth2Client().GetSyncStatus()
				if err != nil {
					return err
				}
				if syncStatus.Syncing {
					return errors.New("could not perform exit: beacon node is syncing.")
				}
				beaconHead, err := task.connectionOfSuperNodeAccount.Eth2Client().GetBeaconHead()
				if err != nil {
					return err
				}
				// check exited before
				pubkey := types.BytesToValidatorPubkey(validator.privateKey.PublicKey().Marshal())
				status, err := task.connectionOfSuperNodeAccount.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{Epoch: &beaconHead.Epoch})
				if err != nil {
					return err
				}
				// will skip if already sign exit
				if status.ExitEpoch != math.MaxUint64 {
					logrus.Infof("validator %d will exit at epoch %d", validator.validatorIndex, status.ExitEpoch)
					validator.exitEpoch = status.ExitEpoch
					continue
				}

				currentEpoch := primTypes.Epoch(beaconHead.Epoch)

				// not active
				if uint64(currentEpoch) < status.ActivationEpoch {
					logrus.Warnf("validator %d is not active and can't exit, will skip, active epoch: %d current epoch: %d", validator.validatorIndex, status.ActivationEpoch, currentEpoch)
					continue
				}
				if currentEpoch < primTypes.Epoch(status.ActivationEpoch)+shardCommitteePeriod {
					logrus.Warnf("validator %d is not active long enough and can't exit, will skip, active epoch: %d current epoch: %d", validator.validatorIndex, status.ActivationEpoch, currentEpoch)
					continue
				}

				// will sign and broadcast exit msg
				exit := &ethpb.VoluntaryExit{Epoch: currentEpoch, ValidatorIndex: primTypes.ValidatorIndex(validator.validatorIndex)}

				fork := &ethpb.Fork{
					PreviousVersion: params.BeaconConfig().CapellaForkVersion,
					CurrentVersion:  params.BeaconConfig().CapellaForkVersion,
					Epoch:           params.BeaconConfig().CapellaForkEpoch,
				}

				signatureDomain, err := signing.Domain(fork, exit.Epoch, domainVoluntaryExit,
					task.eth2Config.GenesisValidatorsRoot)
				if err != nil {
					return errors.Wrap(err, "failed to compute signature domain")
				}

				exitRoot, err := signing.ComputeSigningRoot(exit, signatureDomain)
				if err != nil {
					return errors.Wrap(err, "ComputeSigningRoot failed")
				}

				secretKey, err := bls.SecretKeyFromBytes(validator.privateKey.Marshal())
				if err != nil {
					return errors.Wrap(err, "failed to initialize keys caches from account keystore")
				}
				sig := secretKey.Sign(exitRoot[:])

				err = task.connectionOfSuperNodeAccount.Eth2Client().ExitValidator(validator.validatorIndex,
					uint64(currentEpoch), types.BytesToValidatorSignature(sig.Marshal()))
				if err != nil {
					return err
				}

				logrus.Infof("validator %d broadcast voluntary exit ok", validator.validatorIndex)

			}
		}

	}

	return nil
}
