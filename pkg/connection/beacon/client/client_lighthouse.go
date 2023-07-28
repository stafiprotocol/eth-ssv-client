package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (c *StandardHttpClient) Balance(slot uint64, validator uint64) (uint64, error) {

	url := fmt.Sprintf("/eth/v1/beacon/states/%d/validator_balances?id=%d", slot, validator)

	bts, status, err := c.getRequest(url)
	if err != nil {
		return 0, err
	}

	if status != 200 {
		return 0, fmt.Errorf("http request error: %d", status)
	}

	r := &BalanceApiResponse{}

	err = json.Unmarshal(bts, r)

	if err != nil {
		return 0, err
	}
	return r.Data[0].Balance, nil

}

func (c *StandardHttpClient) AttestationRewards(epoch uint64) (*AttestationRewardsApiResponse, error) {
	url := fmt.Sprintf("/eth/v1/beacon/rewards/attestations/%d", epoch)

	bts, status, err := c.postRequest(url, []string{})
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("http request error: %d", status)
	}

	r := &AttestationRewardsApiResponse{}

	err = json.Unmarshal(bts, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *StandardHttpClient) AttestationRewardsWithVals(epoch uint64, pubkeysOrIndexes []string) (*AttestationRewardsApiResponse, error) {
	url := fmt.Sprintf("/eth/v1/beacon/rewards/attestations/%d", epoch)

	bts, status, err := c.postRequest(url, pubkeysOrIndexes)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("http request AttestationRewardsWithVals error: %d", status)
	}

	r := &AttestationRewardsApiResponse{}

	err = json.Unmarshal(bts, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *StandardHttpClient) SyncCommitteeRewards(slot uint64) (*SyncCommitteeRewardsApiResponse, error) {
	url := fmt.Sprintf("/eth/v1/beacon/rewards/sync_committee/%d", slot)
	bts, status, err := c.postRequest(url, []string{})
	if err != nil {
		return nil, err
	}

	if status != 200 {
		if status == 404 {
			return nil, ErrBlockNotFound
		}
		if status == 500 {
			return nil, ErrSlotPreSyncCommittees
		}
		return nil, fmt.Errorf("http request error: %d", status)
	}

	r := &SyncCommitteeRewardsApiResponse{}

	err = json.Unmarshal(bts, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *StandardHttpClient) BlockRewards(slot uint64) (*BlockRewardsApiResponse, error) {
	url := fmt.Sprintf("/eth/v1/beacon/rewards/blocks/%d", slot)

	bts, status, err := c.getRequest(url)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		if status == 404 {
			return nil, ErrBlockNotFound
		}
		return nil, fmt.Errorf("http request error: %d", status)
	}

	r := &BlockRewardsApiResponse{}

	err = json.Unmarshal(bts, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *StandardHttpClient) ProposerAssignments(epoch uint64) (*EpochProposerAssignmentsApiResponse, error) {
	url := fmt.Sprintf("/eth/v1/validator/duties/proposer/%d", epoch)

	bts, status, err := c.getRequest(url)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("http request error: %d", status)
	}

	r := &EpochProposerAssignmentsApiResponse{}

	err = json.Unmarshal(bts, r)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *StandardHttpClient) ExecutionBlockNumber(slot uint64) (uint64, error) {
	url := fmt.Sprintf("/eth/v1/beacon/blocks/%d", slot)

	bts, status, err := c.getRequest(url)
	if err != nil {
		return 0, err
	}

	if status != 200 {
		if status == 404 {
			return 0, ErrBlockNotFound
		}
		return 0, fmt.Errorf("http request error: %d", status)
	}

	type internal struct {
		Data struct {
			Message struct {
				Body struct {
					ExecutionPayload struct {
						BlockNumber  string   `json:"block_number"`
						Transactions []string `json:"transactions"`
					} `json:"execution_payload"`
				} `json:"body"`
			} `json:"message"`
		} `json:"data"`
	}
	var r internal

	err = json.Unmarshal(bts, &r)
	if err != nil {
		return 0, err
	}

	if r.Data.Message.Body.ExecutionPayload.BlockNumber == "" { // slot if pre merge
		return 0, ErrSlotPreMerge
	}

	return strconv.ParseUint(r.Data.Message.Body.ExecutionPayload.BlockNumber, 10, 64)
}
