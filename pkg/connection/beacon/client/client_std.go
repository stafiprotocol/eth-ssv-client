package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"golang.org/x/sync/errgroup"

	ethpb "github.com/prysmaticlabs/prysm/v3/proto/eth/v1"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

// Config
const (
	RequestUrlFormat   = "%s%s"
	RequestContentType = "application/json"

	RequestSyncStatusPath            = "/eth/v1/node/syncing"
	RequestEth2ConfigPath            = "/eth/v1/config/spec"
	RequestEth2DepositContractMethod = "/eth/v1/config/deposit_contract"
	RequestGenesisPath               = "/eth/v1/beacon/genesis"
	RequestCommitteePath             = "/eth/v1/beacon/states/%s/committees"
	RequestFinalityCheckpointsPath   = "/eth/v1/beacon/states/%s/finality_checkpoints"
	RequestForkPath                  = "/eth/v1/beacon/states/%s/fork"
	RequestValidatorsPath            = "/eth/v1/beacon/states/%s/validators"
	RequestVoluntaryExitPath         = "/eth/v1/beacon/pool/voluntary_exits"
	RequestAttestationsPath          = "/eth/v1/beacon/blocks/%s/attestations"
	RequestBeaconBlockPath           = "/eth/v2/beacon/blocks/%d"
	RequestValidatorSyncDuties       = "/eth/v1/validator/duties/sync/%s"
	RequestValidatorProposerDuties   = "/eth/v1/validator/duties/proposer/%s"
	RequestSyncCommittees            = "/eth/v1/beacon/states/%s/sync_committees"

	MaxRequestValidatorsCount = 50
)

var (
	retryLimit = 10
)

// Beacon client using the standard Beacon HTTP REST API (https://ethereum.github.io/beacon-APIs/)
type StandardHttpClient struct {
	providerAddress string
	eth2Config      beacon.Eth2Config
}

// Create a new client instance
func NewStandardHttpClient(providerAddress string) (*StandardHttpClient, error) {

	client := &StandardHttpClient{
		providerAddress: providerAddress,
	}
	config, err := client.GetEth2Config()
	if err != nil {
		return nil, err
	}
	client.eth2Config = config
	return client, nil
}

// Close the client connection
func (c *StandardHttpClient) Close() error {
	return nil
}

// Get the client's process configuration type
func (c *StandardHttpClient) GetClientType() (beacon.BeaconClientType, error) {
	return beacon.SplitProcess, nil
}

// Get the node's sync status
func (c *StandardHttpClient) GetSyncStatus() (beacon.SyncStatus, error) {

	// Get sync status
	syncStatus, err := c.getSyncStatus()
	if err != nil {
		return beacon.SyncStatus{}, err
	}

	// Calculate the progress
	progress := float64(syncStatus.Data.HeadSlot) / float64(syncStatus.Data.HeadSlot+syncStatus.Data.SyncDistance)

	// Return response
	return beacon.SyncStatus{
		Syncing:  syncStatus.Data.IsSyncing,
		Progress: progress,
	}, nil

}

// Get the eth2 config
func (c *StandardHttpClient) GetEth2Config() (beacon.Eth2Config, error) {

	// Data
	var wg errgroup.Group
	var eth2Config Eth2ConfigResponse
	var genesis GenesisResponse

	// Get eth2 config
	wg.Go(func() error {
		var err error
		eth2Config, err = c.getEth2Config()
		return err
	})

	// Get genesis
	wg.Go(func() error {
		var err error
		genesis, err = c.getGenesis()
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return beacon.Eth2Config{}, err
	}

	// Return response
	return beacon.Eth2Config{
		GenesisForkVersion:           genesis.Data.GenesisForkVersion,
		GenesisValidatorsRoot:        genesis.Data.GenesisValidatorsRoot,
		GenesisEpoch:                 0,
		GenesisTime:                  uint64(genesis.Data.GenesisTime),
		SecondsPerSlot:               uint64(eth2Config.Data.SecondsPerSlot),
		SlotsPerEpoch:                uint64(eth2Config.Data.SlotsPerEpoch),
		SecondsPerEpoch:              uint64(eth2Config.Data.SecondsPerSlot * eth2Config.Data.SlotsPerEpoch),
		EpochsPerSyncCommitteePeriod: uint64(eth2Config.Data.EpochsPerSyncCommitteePeriod),
	}, nil

}

// Get the eth2 deposit contract info
func (c *StandardHttpClient) GetEth2DepositContract() (beacon.Eth2DepositContract, error) {

	// Get the deposit contract
	depositContract, err := c.getEth2DepositContract()
	if err != nil {
		return beacon.Eth2DepositContract{}, err
	}

	// Return response
	return beacon.Eth2DepositContract{
		ChainID: uint64(depositContract.Data.ChainID),
		Address: depositContract.Data.Address,
	}, nil
}

// Get the beacon head
func (c *StandardHttpClient) GetBeaconHead() (beacon.BeaconHead, error) {

	var finalityCheckpoints FinalityCheckpointsResponse
	finalityCheckpoints, err := c.getFinalityCheckpoints("head")

	if err != nil {
		return beacon.BeaconHead{}, err
	}

	epoch := utils.EpochAtTimestamp(c.eth2Config, uint64(time.Now().Unix()))
	// Return response
	return beacon.BeaconHead{
		Epoch:                  epoch,
		Slot:                   utils.StartSlotOfEpoch(c.eth2Config, epoch),
		FinalizedEpoch:         uint64(finalityCheckpoints.Data.Finalized.Epoch),
		FinalizedSlot:          utils.StartSlotOfEpoch(c.eth2Config, uint64(finalityCheckpoints.Data.Finalized.Epoch)),
		JustifiedEpoch:         uint64(finalityCheckpoints.Data.CurrentJustified.Epoch),
		PreviousJustifiedEpoch: uint64(finalityCheckpoints.Data.PreviousJustified.Epoch),
	}, nil

}

// Get a validator's status
func (c *StandardHttpClient) GetValidatorStatus(pubkey types.ValidatorPubkey, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {

	return c.getValidatorStatus(utils.AddPrefix(pubkey.Hex()), opts)

}
func (c *StandardHttpClient) GetValidatorStatusByIndex(index string, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {

	return c.getValidatorStatus(index, opts)

}

func (c *StandardHttpClient) getValidatorStatus(pubkeyOrIndex string, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {

	// Return zero status for null pubkeyOrIndex
	if pubkeyOrIndex == "" {
		return beacon.ValidatorStatus{}, nil
	}

	// Get validator
	validators, err := c.getValidatorsByOpts([]string{pubkeyOrIndex}, opts)
	if err != nil {
		return beacon.ValidatorStatus{}, err
	}
	if len(validators.Data) == 0 {
		return beacon.ValidatorStatus{}, nil
	}
	validator := validators.Data[0]

	// Return response
	return beacon.ValidatorStatus{
		Pubkey:                     types.BytesToValidatorPubkey(validator.Validator.Pubkey),
		Index:                      uint64(validator.Index),
		WithdrawalCredentials:      common.BytesToHash(validator.Validator.WithdrawalCredentials),
		Balance:                    uint64(validator.Balance),
		EffectiveBalance:           uint64(validator.Validator.EffectiveBalance),
		Slashed:                    validator.Validator.Slashed,
		ActivationEligibilityEpoch: uint64(validator.Validator.ActivationEligibilityEpoch),
		ActivationEpoch:            uint64(validator.Validator.ActivationEpoch),
		ExitEpoch:                  uint64(validator.Validator.ExitEpoch),
		WithdrawableEpoch:          uint64(validator.Validator.WithdrawableEpoch),
		Exists:                     true,
		Status:                     ethpb.ValidatorStatus(ethpb.ValidatorStatus_value[strings.ToUpper(validator.Status)]),
	}, nil

}

// Get multiple validators' statuses
// epoch in opts == the first slot of epoch
func (c *StandardHttpClient) GetValidatorStatuses(pubkeys []types.ValidatorPubkey, opts *beacon.ValidatorStatusOptions) (map[types.ValidatorPubkey]beacon.ValidatorStatus, error) {

	// The null validator pubkey
	nullPubkey := types.ValidatorPubkey{}

	// Filter out null pubkeys
	nullPubkeyExists := false
	for _, pubkey := range pubkeys {
		if bytes.Equal(pubkey.Bytes(), nullPubkey.Bytes()) {
			nullPubkeyExists = true
		}
	}

	// Convert pubkeys into hex strings
	pubkeysHex := make([]string, len(pubkeys))
	for vi := 0; vi < len(pubkeys); vi++ {
		pubkeysHex[vi] = utils.AddPrefix(pubkeys[vi].Hex())
	}

	// Get validators
	validators, err := c.getValidatorsByOpts(pubkeysHex, opts)
	if err != nil {
		return nil, err
	}

	// Build validator status map
	statuses := make(map[types.ValidatorPubkey]beacon.ValidatorStatus)
	for _, validator := range validators.Data {

		// Get validator pubkey
		pubkey := types.BytesToValidatorPubkey(validator.Validator.Pubkey)

		// Add status
		statuses[pubkey] = beacon.ValidatorStatus{
			Pubkey:                     types.BytesToValidatorPubkey(validator.Validator.Pubkey),
			Index:                      uint64(validator.Index),
			WithdrawalCredentials:      common.BytesToHash(validator.Validator.WithdrawalCredentials),
			Balance:                    uint64(validator.Balance),
			EffectiveBalance:           uint64(validator.Validator.EffectiveBalance),
			Slashed:                    validator.Validator.Slashed,
			ActivationEligibilityEpoch: uint64(validator.Validator.ActivationEligibilityEpoch),
			ActivationEpoch:            uint64(validator.Validator.ActivationEpoch),
			ExitEpoch:                  uint64(validator.Validator.ExitEpoch),
			WithdrawableEpoch:          uint64(validator.Validator.WithdrawableEpoch),
			Exists:                     true,
			Status:                     ethpb.ValidatorStatus(ethpb.ValidatorStatus_value[strings.ToUpper(validator.Status)]),
		}

	}

	// Add zero status for null pubkey if requested
	if nullPubkeyExists {
		statuses[nullPubkey] = beacon.ValidatorStatus{}
	}

	// Return
	return statuses, nil

}

// Get whether validators have sync duties to perform at given epoch
func (c *StandardHttpClient) GetValidatorSyncDuties(indices []uint64, epoch uint64) (map[uint64]bool, error) {

	// Convert incoming uint64 validator indices into an array of string for the request
	indicesStrings := make([]string, len(indices))

	for i, index := range indices {
		indicesStrings[i] = strconv.FormatUint(index, 10)
	}

	// Perform the post request
	responseBody, status, err := c.postRequest(fmt.Sprintf(RequestValidatorSyncDuties, strconv.FormatUint(epoch, 10)), indicesStrings)

	if err != nil {
		return nil, fmt.Errorf("could not get validator sync duties: %w", err)
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("could not get validator sync duties: HTTP status %d; response body: '%s'", status, string(responseBody))
	}

	var response SyncDutiesResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("could not decode validator sync duties data: %w", err)
	}

	// Map the results
	validatorMap := make(map[uint64]bool)

	for _, index := range indices {
		validatorMap[index] = false
		for _, duty := range response.Data {
			if uint64(duty.ValidatorIndex) == index {
				validatorMap[index] = true
				break
			}
		}
	}

	return validatorMap, nil
}

// proposer duties for a given epoch, return [slot][valIndex]
func (c *StandardHttpClient) GetValidatorProposerDuties(epoch uint64) (map[uint64]uint64, error) {

	// Perform the post request
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestValidatorProposerDuties, strconv.FormatUint(epoch, 10)))

	if err != nil {
		return nil, fmt.Errorf("could not get validator proposer duties: %w", err)
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("could not get validator proposer duties: HTTP status %d; response body: '%s'", status, string(responseBody))
	}

	var response ProposerDutiesResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("could not decode validator proposer duties data: %w", err)
	}

	// Map the results
	proposerMap := make(map[uint64]uint64)

	for _, duty := range response.Data {
		proposerMap[uint64(duty.Slot)] = uint64(duty.ValidatorIndex)
	}

	return proposerMap, nil
}

// Get a validator's index
func (c *StandardHttpClient) GetValidatorIndex(pubkey types.ValidatorPubkey) (uint64, error) {

	// Get validator
	pubkeyString := utils.AddPrefix(pubkey.Hex())
	validators, err := c.getValidatorsByOpts([]string{pubkeyString}, nil)
	if err != nil {
		return 0, err
	}
	if len(validators.Data) == 0 {
		return 0, fmt.Errorf("validator %s index not found", pubkeyString)
	}
	validator := validators.Data[0]

	// Return validator index
	return uint64(validator.Index), nil

}

// Get domain data for a domain type at a given epoch
func (c *StandardHttpClient) GetDomainData(domainType []byte, epoch uint64) ([]byte, error) {

	// Data
	var wg errgroup.Group
	var genesis GenesisResponse
	var fork ForkResponse

	// Get genesis
	wg.Go(func() error {
		var err error
		genesis, err = c.getGenesis()
		return err
	})

	// Get fork
	wg.Go(func() error {
		var err error
		fork, err = c.getFork("head")
		return err
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		return []byte{}, err
	}

	// Get fork version
	var forkVersion []byte
	if epoch < uint64(fork.Data.Epoch) {
		forkVersion = fork.Data.PreviousVersion
	} else {
		forkVersion = fork.Data.CurrentVersion
	}

	// Compute & return domain
	var dt [4]byte
	copy(dt[:], domainType[:])
	return eth2types.Domain(dt, forkVersion, genesis.Data.GenesisValidatorsRoot), nil

}

// Perform a voluntary exit on a validator
func (c *StandardHttpClient) ExitValidator(validatorIndex, epoch uint64, signature types.ValidatorSignature) error {
	return c.postVoluntaryExit(VoluntaryExitRequest{
		Message: VoluntaryExitMessage{
			Epoch:          uinteger(epoch),
			ValidatorIndex: uinteger(validatorIndex),
		},
		Signature: signature.Bytes(),
	})
}

// Get the ETH1 data for the target beacon block
func (c *StandardHttpClient) GetEth1DataForEth2Block(blockId uint64) (beacon.Eth1Data, bool, error) {

	// Get the Beacon block
	block, exists, err := c.getBeaconBlock(blockId)
	if err != nil {
		return beacon.Eth1Data{}, false, err
	}
	if !exists {
		return beacon.Eth1Data{}, false, nil
	}

	// Convert the response to the eth1 data struct
	return beacon.Eth1Data{
		DepositRoot:  common.HexToHash(block.Data.Message.Body.Eth1Data.DepositRoot),
		DepositCount: uint64(block.Data.Message.Body.Eth1Data.DepositCount),
		BlockHash:    common.HexToHash(block.Data.Message.Body.Eth1Data.BlockHash),
	}, true, nil

}

func (c *StandardHttpClient) GetAttestations(blockId string) ([]beacon.AttestationInfo, bool, error) {
	attestations, exists, err := c.getAttestations(blockId)
	if err != nil {
		return nil, false, err
	}
	if !exists {
		return nil, false, nil
	}

	// Add attestation info
	attestationInfo := make([]beacon.AttestationInfo, len(attestations.Data))
	for i, attestation := range attestations.Data {
		bitString := utils.RemovePrefix(attestation.AggregationBits)
		attestationInfo[i].SlotIndex = uint64(attestation.Data.Slot)
		attestationInfo[i].CommitteeIndex = uint64(attestation.Data.Index)
		attestationInfo[i].AggregationBits, err = hex.DecodeString(bitString)
		if err != nil {
			return nil, false, fmt.Errorf("decoding aggregation bits for attestation %d of block %s err: %w", i, blockId, err)
		}
	}

	return attestationInfo, true, nil
}

func (c *StandardHttpClient) GetBeaconBlock(blockId uint64) (beacon.BeaconBlock, bool, error) {
	block, exists, err := c.getBeaconBlock(blockId)
	if err != nil {
		return beacon.BeaconBlock{}, false, err
	}
	if !exists {
		return beacon.BeaconBlock{}, false, nil
	}

	beaconBlock := beacon.BeaconBlock{
		Slot:          uint64(block.Data.Message.Slot),
		ProposerIndex: uint64(block.Data.Message.ProposerIndex),
	}

	// Add attestation info
	for i, attestation := range block.Data.Message.Body.Attestations {
		bitString := utils.RemovePrefix(attestation.AggregationBits)
		info := beacon.AttestationInfo{
			SlotIndex:      uint64(attestation.Data.Slot),
			CommitteeIndex: uint64(attestation.Data.Index),
		}
		info.AggregationBits, err = hex.DecodeString(bitString)
		if err != nil {
			return beacon.BeaconBlock{}, false, fmt.Errorf("decoding aggregation bits for attestation %d of block %d err: %w", i, blockId, err)
		}
		beaconBlock.Attestations = append(beaconBlock.Attestations, info)
	}

	//add syncAggregate
	if len(block.Data.Message.Body.SyncAggregate.SyncCommitteeBits) > 0 {
		syncAggregate := beacon.SyncAggregate{}
		bitString := utils.RemovePrefix(block.Data.Message.Body.SyncAggregate.SyncCommitteeBits)
		syncAggregate.SyncCommitteeBits, err = hex.DecodeString(bitString)
		if err != nil {
			return beacon.BeaconBlock{}, false, fmt.Errorf("decoding aggregation bits for SyncCommitteeBits of block %d err: %w", blockId, err)
		}
		syncAggregate.SyncCommitteeSignature = block.Data.Message.Body.SyncAggregate.SyncCommitteeSignature

		beaconBlock.SyncAggregate = syncAggregate
	}

	// Add proposer slash
	for _, proposerSlash := range block.Data.Message.Body.ProposerSlashings {
		newProposerSlash := beacon.ProposerSlashing{
			SignedHeader1: beacon.SignedHeader{
				Slot:          uint64(proposerSlash.SignedHeader1.Message.Slot),
				ProposerIndex: uint64(proposerSlash.SignedHeader1.Message.ProposerIndex),
				ParentRoot:    proposerSlash.SignedHeader1.Message.ParentRoot,
				StateRoot:     proposerSlash.SignedHeader1.Message.StateRoot,
				BodyRoot:      proposerSlash.SignedHeader1.Message.BodyRoot,
				Signature:     proposerSlash.SignedHeader1.Signature,
			},
			SignedHeader2: beacon.SignedHeader{
				Slot:          uint64(proposerSlash.SignedHeader2.Message.Slot),
				ProposerIndex: uint64(proposerSlash.SignedHeader2.Message.ProposerIndex),
				ParentRoot:    proposerSlash.SignedHeader2.Message.ParentRoot,
				StateRoot:     proposerSlash.SignedHeader2.Message.StateRoot,
				BodyRoot:      proposerSlash.SignedHeader2.Message.BodyRoot,
				Signature:     proposerSlash.SignedHeader2.Signature,
			},
		}
		beaconBlock.ProposerSlashings = append(beaconBlock.ProposerSlashings, newProposerSlash)
	}

	// Add attester slash
	for _, attesterSlash := range block.Data.Message.Body.AttesterSlashings {
		newAttestingIndeces1 := make([]uint64, len(attesterSlash.Attestation1.AttestingIndices))
		for i, indice := range attesterSlash.Attestation1.AttestingIndices {
			newAttestingIndeces1[i] = uint64(indice)
		}
		newAttestingIndeces2 := make([]uint64, len(attesterSlash.Attestation2.AttestingIndices))
		for i, indice := range attesterSlash.Attestation2.AttestingIndices {
			newAttestingIndeces2[i] = uint64(indice)
		}
		newAttesterSalsh := beacon.AttesterSlashing{
			Attestation1: beacon.Attestation{
				AttestingIndices: newAttestingIndeces1,
				Signature:        attesterSlash.Attestation1.Signature,
				Slot:             uint64(attesterSlash.Attestation1.Data.Slot),
				Index:            uint64(attesterSlash.Attestation1.Data.Index),
				BeaconBlockRoot:  attesterSlash.Attestation1.Data.BeaconBlockRoot,
				SourceEpoch:      uint64(attesterSlash.Attestation1.Data.Source.Epoch),
				SourceRoot:       attesterSlash.Attestation1.Data.Source.Root,
				TargetEpoch:      uint64(attesterSlash.Attestation1.Data.Target.Epoch),
				TargetRoot:       attesterSlash.Attestation1.Data.Target.Root,
			},
			Attestation2: beacon.Attestation{
				AttestingIndices: newAttestingIndeces2,
				Signature:        attesterSlash.Attestation2.Signature,
				Slot:             uint64(attesterSlash.Attestation2.Data.Slot),
				Index:            uint64(attesterSlash.Attestation2.Data.Index),
				BeaconBlockRoot:  attesterSlash.Attestation2.Data.BeaconBlockRoot,
				SourceEpoch:      uint64(attesterSlash.Attestation2.Data.Source.Epoch),
				SourceRoot:       attesterSlash.Attestation2.Data.Source.Root,
				TargetEpoch:      uint64(attesterSlash.Attestation2.Data.Target.Epoch),
				TargetRoot:       attesterSlash.Attestation2.Data.Target.Root,
			},
		}
		beaconBlock.AttesterSlashing = append(beaconBlock.AttesterSlashing, newAttesterSalsh)
	}

	for _, withdrawal := range block.Data.Message.Body.ExecutionPayload.Withdrawals {
		beaconBlock.Withdrawals = append(beaconBlock.Withdrawals, beacon.Withdrawal{
			WithdrawIndex:  uint64(withdrawal.Index),
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Address:        common.HexToAddress(withdrawal.Address),
			Amount:         uint64(withdrawal.Amount),
		})
	}

	for _, exitMsg := range block.Data.Message.Body.VoluntaryExits {
		beaconBlock.VoluntaryExits = append(beaconBlock.VoluntaryExits, beacon.VoluntaryExit{
			ValidatorIndex: uint64(exitMsg.Message.ValidatorIndex),
			Epoch:          uint64(exitMsg.Message.Epoch),
		})
	}

	// Execution payload only exists after the merge, so check for its existence
	if block.Data.Message.Body.ExecutionPayload == nil {
		beaconBlock.HasExecutionPayload = false
	} else {
		beaconBlock.HasExecutionPayload = true

		beaconBlock.FeeRecipient = common.HexToAddress(block.Data.Message.Body.ExecutionPayload.FeeRecipient)
		beaconBlock.ExecutionBlockNumber = uint64(block.Data.Message.Body.ExecutionPayload.BlockNumber)
	}

	return beaconBlock, true, nil
}

// Get the attestation committees for the given epoch, or the current epoch if nil
func (c *StandardHttpClient) GetCommitteesForEpoch(epoch uint64) ([]beacon.Committee, error) {
	response, err := c.getCommittees("head", &epoch)
	if err != nil {
		return nil, err
	}

	committees := []beacon.Committee{}
	for _, committee := range response.Data {
		validators := []uint64{}
		for _, validator := range committee.Validators {
			validators = append(validators, uint64(validator))
		}
		committees = append(committees, beacon.Committee{
			Index:      uint64(committee.Index),
			Slot:       uint64(committee.Slot),
			Validators: validators,
		})
	}

	return committees, nil
}

func (c *StandardHttpClient) GetSyncCommitteesForEpoch(epoch uint64) ([]beacon.SyncCommittee, error) {
	response, err := c.getSyncCommittees("head", &epoch)
	if err != nil {
		return nil, err
	}

	committees := []beacon.SyncCommittee{}
	for _, valIndexStr := range response.Data.Validators {
		valIndexU64, err := strconv.ParseUint(valIndexStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("in sync_committee for epoch %d validator has bad validator index: %q", epoch, valIndexStr)
		}
		committees = append(committees, beacon.SyncCommittee{
			ValIndex: valIndexU64,
		})
	}

	return committees, nil
}

// Get sync status
func (c *StandardHttpClient) getSyncStatus() (SyncStatusResponse, error) {
	responseBody, status, err := c.getRequest(RequestSyncStatusPath)
	if err != nil {
		return SyncStatusResponse{}, fmt.Errorf("could not get node sync status: %w", err)
	}
	if status != http.StatusOK {
		return SyncStatusResponse{}, fmt.Errorf("could not get node sync status: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var syncStatus SyncStatusResponse
	if err := json.Unmarshal(responseBody, &syncStatus); err != nil {
		return SyncStatusResponse{}, fmt.Errorf("could not decode node sync status: %w", err)
	}
	return syncStatus, nil
}

// Get the eth2 config
func (c *StandardHttpClient) getEth2Config() (Eth2ConfigResponse, error) {
	responseBody, status, err := c.getRequest(RequestEth2ConfigPath)
	if err != nil {
		return Eth2ConfigResponse{}, fmt.Errorf("could not get eth2 config: %w", err)
	}
	if status != http.StatusOK {
		return Eth2ConfigResponse{}, fmt.Errorf("could not get eth2 config: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var eth2Config Eth2ConfigResponse
	if err := json.Unmarshal(responseBody, &eth2Config); err != nil {
		return Eth2ConfigResponse{}, fmt.Errorf("could not decode eth2 config: %w", err)
	}
	return eth2Config, nil
}

// Get the eth2 deposit contract info
func (c *StandardHttpClient) getEth2DepositContract() (Eth2DepositContractResponse, error) {
	responseBody, status, err := c.getRequest(RequestEth2DepositContractMethod)
	if err != nil {
		return Eth2DepositContractResponse{}, fmt.Errorf("could not get eth2 deposit contract: %w", err)
	}
	if status != http.StatusOK {
		return Eth2DepositContractResponse{}, fmt.Errorf("could not get eth2 deposit contract: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var eth2DepositContract Eth2DepositContractResponse
	if err := json.Unmarshal(responseBody, &eth2DepositContract); err != nil {
		return Eth2DepositContractResponse{}, fmt.Errorf("could not decode eth2 deposit contract: %w", err)
	}
	return eth2DepositContract, nil
}

// Get genesis information
func (c *StandardHttpClient) getGenesis() (GenesisResponse, error) {
	responseBody, status, err := c.getRequest(RequestGenesisPath)
	if err != nil {
		return GenesisResponse{}, fmt.Errorf("could not get genesis data: %w", err)
	}
	if status != http.StatusOK {
		return GenesisResponse{}, fmt.Errorf("could not get genesis data: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var genesis GenesisResponse
	if err := json.Unmarshal(responseBody, &genesis); err != nil {
		return GenesisResponse{}, fmt.Errorf("could not decode genesis: %w", err)
	}
	return genesis, nil
}

// Get finality checkpoints
func (c *StandardHttpClient) getFinalityCheckpoints(stateId string) (FinalityCheckpointsResponse, error) {
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestFinalityCheckpointsPath, stateId))
	if err != nil {
		return FinalityCheckpointsResponse{}, fmt.Errorf("could not get finality checkpoints: %w", err)
	}
	if status != http.StatusOK {
		return FinalityCheckpointsResponse{}, fmt.Errorf("could not get finality checkpoints: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var finalityCheckpoints FinalityCheckpointsResponse
	if err := json.Unmarshal(responseBody, &finalityCheckpoints); err != nil {
		return FinalityCheckpointsResponse{}, fmt.Errorf("could not decode finality checkpoints: %w", err)
	}
	return finalityCheckpoints, nil
}

// Get fork
func (c *StandardHttpClient) getFork(stateId string) (ForkResponse, error) {
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestForkPath, stateId))
	if err != nil {
		return ForkResponse{}, fmt.Errorf("could not get fork data: %w", err)
	}
	if status != http.StatusOK {
		return ForkResponse{}, fmt.Errorf("could not get fork data: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var fork ForkResponse
	if err := json.Unmarshal(responseBody, &fork); err != nil {
		return ForkResponse{}, fmt.Errorf("could not decode fork data: %w", err)
	}
	return fork, nil
}

// Get validators
func (c *StandardHttpClient) getValidators(stateId string, pubkeys []string) (ValidatorsResponse, error) {
	var query string
	if len(pubkeys) > 0 {
		query = fmt.Sprintf("?id=%s", strings.Join(pubkeys, ","))
	}
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestValidatorsPath, stateId) + query)
	if err != nil {
		return ValidatorsResponse{}, fmt.Errorf("could not get validators: %w", err)
	}
	if status != http.StatusOK {
		return ValidatorsResponse{}, fmt.Errorf("could not get validators: HTTP status %d; response body: '%s'", status, string(responseBody))
	}

	var validators ValidatorsResponse
	if err := json.Unmarshal(responseBody, &validators); err != nil {
		return ValidatorsResponse{}, fmt.Errorf("could not decode validators: %w", err)
	}
	return validators, nil
}

// Get validators by pubkeys and status options
func (c *StandardHttpClient) getValidatorsByOpts(pubkeysOrIndices []string, opts *beacon.ValidatorStatusOptions) (ValidatorsResponse, error) {

	// Get state ID
	var stateId string
	if opts == nil {
		stateId = "head"
	} else if opts.Slot != nil {
		stateId = strconv.FormatInt(int64(*opts.Slot), 10)
	} else if opts.Epoch != nil {
		// Get slot nuumber
		slot := *opts.Epoch * uint64(c.eth2Config.SlotsPerEpoch)
		stateId = strconv.FormatInt(int64(slot), 10)

	} else {
		return ValidatorsResponse{}, fmt.Errorf("must specify a slot or epoch when calling getValidatorsByOpts")
	}

	// Load validator data in batches & return
	data := make([]Validator, 0, len(pubkeysOrIndices))
	for bsi := 0; bsi < len(pubkeysOrIndices); bsi += MaxRequestValidatorsCount {

		// Get batch start & end index
		vsi := bsi
		vei := bsi + MaxRequestValidatorsCount
		if vei > len(pubkeysOrIndices) {
			vei = len(pubkeysOrIndices)
		}

		// Get validator pubkeysOrIndices for batch request
		batch := make([]string, vei-vsi)
		for vi := vsi; vi < vei; vi++ {
			batch[vi-vsi] = pubkeysOrIndices[vi]
		}

		// Get & add validators
		var validators ValidatorsResponse
		var err error
		retry := 0
		for {
			if retry > retryLimit {
				return ValidatorsResponse{}, err
			}
			validators, err = c.getValidators(stateId, batch)
			if err != nil {
				time.Sleep(time.Millisecond * 200)
				retry++
				continue
			}
			break
		}
		data = append(data, validators.Data...)

	}
	return ValidatorsResponse{Data: data}, nil

}

// Send voluntary exit request
func (c *StandardHttpClient) postVoluntaryExit(request VoluntaryExitRequest) error {
	responseBody, status, err := c.postRequest(RequestVoluntaryExitPath, request)
	if err != nil {
		return fmt.Errorf("could not broadcast exit for validator at index %d: %w", request.Message.ValidatorIndex, err)
	}
	if status != http.StatusOK {
		return fmt.Errorf("could not broadcast exit for validator at index %d: HTTP status %d; response body: '%s'", request.Message.ValidatorIndex, status, string(responseBody))
	}
	return nil
}

// Get the target beacon block
func (c *StandardHttpClient) getAttestations(blockId string) (AttestationsResponse, bool, error) {
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestAttestationsPath, blockId))
	if err != nil {
		return AttestationsResponse{}, false, fmt.Errorf("could not get attestations data for slot %s: %w", blockId, err)
	}
	if status == http.StatusNotFound {
		return AttestationsResponse{}, false, nil
	}
	if status != http.StatusOK {
		return AttestationsResponse{}, false, fmt.Errorf("could not get attestations data for slot %s: HTTP status %d; response body: '%s'", blockId, status, string(responseBody))
	}
	var attestations AttestationsResponse
	if err := json.Unmarshal(responseBody, &attestations); err != nil {
		return AttestationsResponse{}, false, fmt.Errorf("could not decode attestations data for slot %s: %w", blockId, err)
	}
	return attestations, true, nil
}

// Get the target beacon block
func (c *StandardHttpClient) getBeaconBlock(blockId uint64) (BeaconBlockResponse, bool, error) {
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestBeaconBlockPath, blockId))
	if err != nil {
		return BeaconBlockResponse{}, false, fmt.Errorf("could not get beacon block data: %w", err)
	}
	if status == http.StatusNotFound {
		return BeaconBlockResponse{}, false, nil
	}
	if status != http.StatusOK {
		return BeaconBlockResponse{}, false, fmt.Errorf("could not get beacon block data: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var beaconBlock BeaconBlockResponse
	if err := json.Unmarshal(responseBody, &beaconBlock); err != nil {
		return BeaconBlockResponse{}, false, fmt.Errorf("could not decode beacon block data: %w", err)
	}
	return beaconBlock, true, nil
}

// Get the committees for the epoch
func (c *StandardHttpClient) getCommittees(stateId string, epoch *uint64) (CommitteesResponse, error) {
	query := ""
	if epoch != nil {
		query = fmt.Sprintf("?epoch=%d", *epoch)
	}
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestCommitteePath, stateId) + query)
	if err != nil {
		return CommitteesResponse{}, fmt.Errorf("could not get committees: %w", err)
	}
	if status != http.StatusOK {
		return CommitteesResponse{}, fmt.Errorf("could not get committees: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var committees CommitteesResponse
	if err := json.Unmarshal(responseBody, &committees); err != nil {
		return CommitteesResponse{}, fmt.Errorf("could not decode committees: %w", err)
	}
	return committees, nil
}

// Get the sync committees for the epoch
func (c *StandardHttpClient) getSyncCommittees(stateId string, epoch *uint64) (SyncCommitteesResponse, error) {
	query := ""
	if epoch != nil {
		query = fmt.Sprintf("?epoch=%d", *epoch)
	}
	responseBody, status, err := c.getRequest(fmt.Sprintf(RequestSyncCommittees, stateId) + query)
	if err != nil {
		return SyncCommitteesResponse{}, fmt.Errorf("could not get sync committees: %w", err)
	}
	if status != http.StatusOK {
		return SyncCommitteesResponse{}, fmt.Errorf("could not get sync committees: HTTP status %d; response body: '%s'", status, string(responseBody))
	}
	var committees SyncCommitteesResponse
	if err := json.Unmarshal(responseBody, &committees); err != nil {
		return SyncCommitteesResponse{}, fmt.Errorf("could not decode sync committees: %w", err)
	}
	return committees, nil
}

// Make a GET request to the beacon node
func (c *StandardHttpClient) getRequest(requestPath string) ([]byte, int, error) {

	// Send request
	client := http.Client{Timeout: 120 * time.Second}

	response, err := client.Get(fmt.Sprintf(RequestUrlFormat, c.providerAddress, requestPath))
	if err != nil {
		return []byte{}, 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Get response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	// Return
	return body, response.StatusCode, nil

}

// Make a POST request to the beacon node
func (c *StandardHttpClient) postRequest(requestPath string, requestBody interface{}) ([]byte, int, error) {

	// Get request body
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return []byte{}, 0, err
	}
	requestBodyReader := bytes.NewReader(requestBodyBytes)

	// Send request
	client := http.Client{Timeout: 60 * time.Second}
	response, err := client.Post(fmt.Sprintf(RequestUrlFormat, c.providerAddress, requestPath), RequestContentType, requestBodyReader)
	if err != nil {
		return []byte{}, 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Get response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	// Return
	return body, response.StatusCode, nil

}
