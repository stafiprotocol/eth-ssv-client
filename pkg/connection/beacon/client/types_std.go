package client

import (
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	hexutil "github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

// Request types
type VoluntaryExitRequest struct {
	Message   VoluntaryExitMessage `json:"message"`
	Signature byteArray            `json:"signature"`
}
type VoluntaryExitMessage struct {
	Epoch          uinteger `json:"epoch"`
	ValidatorIndex uinteger `json:"validator_index"`
}

// Response types
type SyncStatusResponse struct {
	Data struct {
		IsSyncing    bool     `json:"is_syncing"`
		HeadSlot     uinteger `json:"head_slot"`
		SyncDistance uinteger `json:"sync_distance"`
	} `json:"data"`
}
type Eth2ConfigResponse struct {
	Data struct {
		SecondsPerSlot               uinteger `json:"SECONDS_PER_SLOT"`
		SlotsPerEpoch                uinteger `json:"SLOTS_PER_EPOCH"`
		EpochsPerSyncCommitteePeriod uinteger `json:"EPOCHS_PER_SYNC_COMMITTEE_PERIOD"`
	} `json:"data"`
}
type Eth2DepositContractResponse struct {
	Data struct {
		ChainID uinteger       `json:"chain_id"`
		Address common.Address `json:"address"`
	} `json:"data"`
}
type GenesisResponse struct {
	Data struct {
		GenesisTime           uinteger  `json:"genesis_time"`
		GenesisForkVersion    byteArray `json:"genesis_fork_version"`
		GenesisValidatorsRoot byteArray `json:"genesis_validators_root"`
	} `json:"data"`
}
type FinalityCheckpointsResponse struct {
	Data struct {
		PreviousJustified struct {
			Epoch uinteger `json:"epoch"`
		} `json:"previous_justified"`
		CurrentJustified struct {
			Epoch uinteger `json:"epoch"`
		} `json:"current_justified"`
		Finalized struct {
			Epoch uinteger `json:"epoch"`
		} `json:"finalized"`
	} `json:"data"`
}
type ForkResponse struct {
	Data struct {
		PreviousVersion byteArray `json:"previous_version"`
		CurrentVersion  byteArray `json:"current_version"`
		Epoch           uinteger  `json:"epoch"`
	} `json:"data"`
}
type AttestationsResponse struct {
	Data []Attestation `json:"data"`
}

type BeaconBlockResponse struct {
	Version             string `json:"version"`
	ExecutionOptimistic bool   `json:"execution_optimistic"`
	Data                struct {
		Message struct {
			Slot          uinteger `json:"slot"`
			ProposerIndex uinteger `json:"proposer_index"`
			ParentRoot    string   `json:"parent_root"`
			StateRoot     string   `json:"state_root"`
			Body          struct {
				RandaoReveal string `json:"randao_reveal"`
				Eth1Data     struct {
					DepositRoot  string   `json:"deposit_root"`
					DepositCount uinteger `json:"deposit_count"`
					BlockHash    string   `json:"block_hash"`
				} `json:"eth1_data"`
				Graffiti          string `json:"graffiti"`
				ProposerSlashings []struct {
					SignedHeader1 struct {
						Message struct {
							Slot          uinteger `json:"slot"`
							ProposerIndex uinteger `json:"proposer_index"`
							ParentRoot    string   `json:"parent_root"`
							StateRoot     string   `json:"state_root"`
							BodyRoot      string   `json:"body_root"`
						} `json:"message"`
						Signature string `json:"signature"`
					} `json:"signed_header_1"`
					SignedHeader2 struct {
						Message struct {
							Slot          uinteger `json:"slot"`
							ProposerIndex uinteger `json:"proposer_index"`
							ParentRoot    string   `json:"parent_root"`
							StateRoot     string   `json:"state_root"`
							BodyRoot      string   `json:"body_root"`
						} `json:"message"`
						Signature string `json:"signature"`
					} `json:"signed_header_2"`
				} `json:"proposer_slashings"`
				AttesterSlashings []struct {
					Attestation1 struct {
						AttestingIndices []uinteger `json:"attesting_indices"`
						Signature        string     `json:"signature"`
						Data             struct {
							Slot            uinteger `json:"slot"`
							Index           uinteger `json:"index"`
							BeaconBlockRoot string   `json:"beacon_block_root"`
							Source          struct {
								Epoch uinteger `json:"epoch"`
								Root  string   `json:"root"`
							} `json:"source"`
							Target struct {
								Epoch uinteger `json:"epoch"`
								Root  string   `json:"root"`
							} `json:"target"`
						} `json:"data"`
					} `json:"attestation_1"`
					Attestation2 struct {
						AttestingIndices []uinteger `json:"attesting_indices"`
						Signature        string     `json:"signature"`
						Data             struct {
							Slot            uinteger `json:"slot"`
							Index           uinteger `json:"index"`
							BeaconBlockRoot string   `json:"beacon_block_root"`
							Source          struct {
								Epoch uinteger `json:"epoch"`
								Root  string   `json:"root"`
							} `json:"source"`
							Target struct {
								Epoch uinteger `json:"epoch"`
								Root  string   `json:"root"`
							} `json:"target"`
						} `json:"data"`
					} `json:"attestation_2"`
				} `json:"attester_slashings"`
				Attestations []struct {
					AggregationBits string `json:"aggregation_bits"`
					Signature       string `json:"signature"`
					Data            struct {
						Slot            uinteger `json:"slot"`
						Index           uinteger `json:"index"`
						BeaconBlockRoot string   `json:"beacon_block_root"`
						Source          struct {
							Epoch uinteger `json:"epoch"`
							Root  string   `json:"root"`
						} `json:"source"`
						Target struct {
							Epoch uinteger `json:"epoch"`
							Root  string   `json:"root"`
						} `json:"target"`
					} `json:"data"`
				} `json:"attestations"`
				Deposits []struct {
					Proof []string `json:"proof"`
					Data  struct {
						Pubkey                string   `json:"pubkey"`
						WithdrawalCredentials string   `json:"withdrawal_credentials"`
						Amount                uinteger `json:"amount"`
						Signature             string   `json:"signature"`
					} `json:"data"`
				} `json:"deposits"`
				VoluntaryExits []struct {
					Message struct {
						Epoch          uinteger `json:"epoch"`
						ValidatorIndex uinteger `json:"validator_index"`
					} `json:"message"`
					Signature string `json:"signature"`
				} `json:"voluntary_exits"`
				SyncAggregate struct {
					SyncCommitteeBits      string `json:"sync_committee_bits"`
					SyncCommitteeSignature string `json:"sync_committee_signature"`
				} `json:"sync_aggregate"`
				ExecutionPayload *struct {
					ParentHash    string   `json:"parent_hash"`
					FeeRecipient  string   `json:"fee_recipient"`
					StateRoot     string   `json:"state_root"`
					ReceiptsRoot  string   `json:"receipts_root"`
					LogsBloom     string   `json:"logs_bloom"`
					PrevRandao    string   `json:"prev_randao"`
					BlockNumber   uinteger `json:"block_number"`
					GasLimit      uinteger `json:"gas_limit"`
					GasUsed       uinteger `json:"gas_used"`
					Timestamp     uinteger `json:"timestamp"`
					ExtraData     string   `json:"extra_data"`
					BaseFeePerGas uinteger `json:"base_fee_per_gas"`
					BlockHash     string   `json:"block_hash"`
					Transactions  []string `json:"transactions"`
					// present only after capella
					Withdrawals []WithdrawalPayload `json:"withdrawals"`
				} `json:"execution_payload"`
			} `json:"body"`
		} `json:"message"`
		Signature string `json:"signature"`
	} `json:"data"`
}

type WithdrawalPayload struct {
	Index          uinteger `json:"index"`
	ValidatorIndex uinteger `json:"validator_index"`
	Address        string   `json:"address"`
	Amount         uinteger `json:"amount"`
}

type ValidatorsResponse struct {
	Data []Validator `json:"data"`
}
type Validator struct {
	Index     uinteger `json:"index"`
	Balance   uinteger `json:"balance"`
	Status    string   `json:"status"`
	Validator struct {
		Pubkey                     byteArray `json:"pubkey"`
		WithdrawalCredentials      byteArray `json:"withdrawal_credentials"`
		EffectiveBalance           uinteger  `json:"effective_balance"`
		Slashed                    bool      `json:"slashed"`
		ActivationEligibilityEpoch uinteger  `json:"activation_eligibility_epoch"`
		ActivationEpoch            uinteger  `json:"activation_epoch"`
		ExitEpoch                  uinteger  `json:"exit_epoch"`
		WithdrawableEpoch          uinteger  `json:"withdrawable_epoch"`
	} `json:"validator"`
}
type SyncDutiesResponse struct {
	Data []SyncDuty `json:"data"`
}
type SyncDuty struct {
	Pubkey               byteArray  `json:"pubkey"`
	ValidatorIndex       uinteger   `json:"validator_index"`
	SyncCommitteeIndices []uinteger `json:"validator_sync_committee_indices"`
}
type ProposerDutiesResponse struct {
	Data []ProposerDuty `json:"data"`
}
type ProposerDuty struct {
	ValidatorIndex uinteger `json:"validator_index"`
	Pubkey         string   `json:"pubkey"`
	Slot           uinteger `json:"slot"`
}

type SyncCommittee struct {
	Validators          []string   `json:"validators"`
	ValidatorAggregates [][]string `json:"validator_aggregates"`
}

type SyncCommitteesResponse struct {
	Data SyncCommittee `json:"data"`
}

type CommitteesResponse struct {
	Data []Committee `json:"data"`
}

type Committee struct {
	Index      uinteger   `json:"index"`
	Slot       uinteger   `json:"slot"`
	Validators []uinteger `json:"validators"`
}

type Attestation struct {
	AggregationBits string `json:"aggregation_bits"`
	Data            struct {
		Slot  uinteger `json:"slot"`
		Index uinteger `json:"index"`
	} `json:"data"`
}

// Unsigned integer type
type uinteger uint64

func (i uinteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(i)))
}
func (i *uinteger) UnmarshalJSON(data []byte) error {

	// Unmarshal string
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	// Parse integer value
	value, err := strconv.ParseUint(dataStr, 10, 64)
	if err != nil {
		return err
	}

	// Set value and return
	*i = uinteger(value)
	return nil

}

// Byte array type
type byteArray []byte

func (b byteArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(hexutil.AddPrefix(hex.EncodeToString(b)))
}
func (b *byteArray) UnmarshalJSON(data []byte) error {

	// Unmarshal string
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	// Decode hex
	value, err := hex.DecodeString(hexutil.RemovePrefix(dataStr))
	if err != nil {
		return err
	}

	// Set value and return
	*b = value
	return nil

}
