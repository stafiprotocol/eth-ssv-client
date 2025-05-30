package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetAllSsvOperatorsFromApi(network string) ([]OperatorFromApi, error) {
	rsp := make([]OperatorFromApi, 0)
	firstPage, err := getAllSsvOperatorsFromApi(network, 1)
	if err != nil {
		return nil, err
	}
	rsp = append(rsp, firstPage.Operators...)
	for i := 1; i <= firstPage.Pagination.Pages; i++ {
		netxtPage, err := getAllSsvOperatorsFromApi(network, i)
		if err != nil {
			return nil, err
		}
		rsp = append(rsp, netxtPage.Operators...)
	}
	return rsp, nil
}

func getAllSsvOperatorsFromApi(network string, page int) (*RspSsvOperators, error) {
	rsp, err := http.Get(fmt.Sprintf("https://api.ssv.network/api/v4/%s/operators?page=%d&perPage=100&ordering=id:asc", network, page))
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

	return &rspSsv, nil
}

type Operator struct {
	Active bool
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

func GetOperatorFromApi(network string, id uint64) (*Operator, error) {
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
	active := false
	if operator.OperatorFromApi.IsActive == 1 {
		active = true
	}

	return &Operator{Active: active}, nil
}

const graphURL = "https://gateway.thegraph.com/api/subgraphs/id/7V45fKPugp9psQjgrGsfif98gWzCyC6ChN7CW98VyQnr"

type OperatorRsp struct {
	ID             string `json:"id"`
	Fee            string `json:"fee"`
	Removed        bool   `json:"removed"`
	TotalWithdrawn string `json:"totalWithdrawn"`
	IsPrivate      bool   `json:"isPrivate"`
}

type graphResponse struct {
	Data struct {
		Operators []OperatorRsp `json:"operators"`
	} `json:"data"`
}

func BatchGetOperatorFromGraph(apiKey string, ids []uint64) (map[uint64]*Operator, error) {
	idList := make([]string, len(ids))
	for i, id := range ids {
		idList[i] = fmt.Sprintf("\"%d\"", id)
	}

	joinedIDs := strings.Join(idList, ", ")

	query := fmt.Sprintf(`{ operators(where: { id_in: [%s] }) { id fee removed totalWithdrawn isPrivate } }`, joinedIDs)

	payload := map[string]interface{}{
		"query":         query,
		"operationName": "Subgraphs",
		"variables":     map[string]interface{}{},
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", graphURL, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response: %s", bodyBytes)
	}

	var res graphResponse
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, err
	}

	if len(res.Data.Operators) != len(ids) {
		return nil, fmt.Errorf("operator count mismatch: some IDs may not exist or were removed, req: %s, res: %s",
			query, string(bodyBytes))
	}

	operators := make(map[uint64]*Operator)
	for _, op := range res.Data.Operators {
		idUint, err := strconv.ParseUint(op.ID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid operator ID format: %v, v: %s", err, op.ID)
		}
		operators[idUint] = &Operator{Active: !op.Removed}
	}

	return operators, nil
}
