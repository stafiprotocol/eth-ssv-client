package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apiOfSsvOperator = "https://api.ssv.network/api/v4/%s/operators/%d"

type OperatorDetail struct {
	ID             int    `json:"id"`
	IDStr          string `json:"id_str"`
	DeclaredFee    string `json:"declared_fee"`
	PreviousFee    string `json:"previous_fee"`
	Fee            string `json:"fee"`
	PublicKey      string `json:"public_key"`
	OwnerAddress   string `json:"owner_address"`
	Location       string `json:"location"`
	SetupProvider  string `json:"setup_provider"`
	Eth1NodeClient string `json:"eth1_node_client"`
	Eth2NodeClient string `json:"eth2_node_client"`
	Description    string `json:"description"`
	WebsiteURL     string `json:"website_url"`
	TwitterURL     string `json:"twitter_url"`
	LinkedinURL    string `json:"linkedin_url"`
	Logo           string `json:"logo"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	Performance    struct {
		Two4H   float64 `json:"24h"`
		Three0D float64 `json:"30d"`
	} `json:"performance"`
	IsValid         bool   `json:"is_valid"`
	IsDeleted       bool   `json:"is_deleted"`
	IsActive        int    `json:"is_active"`
	Status          string `json:"status"`
	ValidatorsCount int    `json:"validators_count"`
	Version         string `json:"version"`
	Network         string `json:"network"`
	Error           struct {
		Code    int `json:"code"`
		Message struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		} `json:"message"`
	} `json:"error"`
}

func GetOperatorDetail(network string, id int) (*OperatorDetail, error) {
	rsp, err := http.Get(fmt.Sprintf(apiOfSsvOperator, network, id))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status err %d", rsp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) == 0 {
		return nil, fmt.Errorf("bodyBytes zero err")
	}
	operator := OperatorDetail{}
	err = json.Unmarshal(bodyBytes, &operator)
	if err != nil {
		return nil, err
	}

	if operator.Error.Code != 0 {
		return nil, fmt.Errorf("err code: %d, err: %s", operator.Error.Code, operator.Error.Message.Error)
	}

	return &operator, nil

}
