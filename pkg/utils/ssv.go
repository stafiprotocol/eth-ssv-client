package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetAllSsvOperatorsFromApi(network string) ([]OperatorFromApi, error) {
	rsp, err := http.Get(fmt.Sprintf("https://api.ssv.network/api/v4/%s/operators?page=1&perPage=5000&ordering=id:asc", network))
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
	rspSsv := RspSsvOperators{}
	err = json.Unmarshal(bodyBytes, &rspSsv)
	if err != nil {
		return nil, err
	}

	return rspSsv.Operators, nil
}

type RspSsvOperators struct {
	Pagination struct {
		Total   int `json:"total"`
		Page    int `json:"page"`
		Pages   int `json:"pages"`
		PerPage int `json:"per_page"`
	} `json:"pagination"`
	Operators []OperatorFromApi `json:"operators"`
	Error     RspError          `json:"error"`
}

type OperatorFromApi struct {
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
	IsValid          bool   `json:"is_valid"`
	IsDeleted        bool   `json:"is_deleted"`
	IsActive         int    `json:"is_active"`
	Status           string `json:"status"`
	ValidatorsCount  int    `json:"validators_count"`
	Version          string `json:"version"`
	Network          string `json:"network"`
	MevRelays        string `json:"mev_relays,omitempty"`
	DkgAddress       string `json:"dkg_address,omitempty"`
	AddressWhitelist string `json:"address_whitelist,omitempty"`
}

type RspSsvOperator struct {
	OperatorFromApi
	Error RspError `json:"error"`
}

type RspError struct {
	Code    int `json:"code"`
	Message struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	} `json:"message"`
}

var apiOfSsvOperator = "https://api.ssv.network/api/v4/%s/operators/%d"

func GetOperatorFromApi(network string, id uint64) (*OperatorFromApi, error) {
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
	operator := RspSsvOperator{}
	err = json.Unmarshal(bodyBytes, &operator)
	if err != nil {
		return nil, err
	}

	if operator.Error.Code != 0 {
		return nil, fmt.Errorf("err code: %d, err: %s", operator.Error.Code, operator.Error.Message.Error)
	}

	return &operator.OperatorFromApi, nil

}

func MustGetOperatorDetail(network string, id uint64) (*OperatorFromApi, error) {
	retry := 0
	var operatorDetail *OperatorFromApi
	var err error
	for {
		if retry > RetryLimit {
			return nil, fmt.Errorf("GetOperatorDetail reach retry limit")
		}
		operatorDetail, err = GetOperatorFromApi(network, id)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return &OperatorFromApi{
					ID:      int(id),
					Network: network,
				}, nil
			}

			time.Sleep(RetryInterval)
			retry++
			continue
		}
		break
	}

	return operatorDetail, nil
}
