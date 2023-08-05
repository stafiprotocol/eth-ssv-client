package utils_test

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	reth "github.com/stafiprotocol/eth-ssv-client/bindings/Reth"
	stafi_ether "github.com/stafiprotocol/eth-ssv-client/bindings/StafiEther"
	storage "github.com/stafiprotocol/eth-ssv-client/bindings/Storage"
	user_deposit "github.com/stafiprotocol/eth-ssv-client/bindings/UserDeposit"
	withdraw "github.com/stafiprotocol/eth-ssv-client/bindings/Withdraw"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func TestAppendFile(t *testing.T) {
	path := "../../log_data/append_test2.txt"
	lastLine, err := utils.ReadLastLine(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(lastLine)
	err = utils.AppendToFile(path, "\ntest1")
	if err != nil {
		t.Fatal(err)
	}
	err = utils.AppendToFile(path, "\ntest1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilterLogs(t *testing.T) {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/3whje5yFZZxg9BqsldHTRku-VXWuf88E")
	if err != nil {
		t.Fatal(err)
	}

	rethContracts, err := reth.NewReth(common.HexToAddress("0x9559aaa82d9649c7a7b220e7c461d2e74c9a3593"), client)
	if err != nil {
		t.Fatal(err)
	}

	iter, err := rethContracts.FilterTokensMinted(&bind.FilterOpts{
		Context: context.Background(),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	for iter.Next() {
		t.Log(iter.Event.Raw.BlockNumber)
	}
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0x9559aaa82d9649c7a7b220e7c461d2e74c9a3593")},
		Topics:    [][]common.Hash{{common.HexToHash("0x6155cfd0fd028b0ca77e8495a60cbe563e8bce8611f0aad6fedbdaafc05d44a2")}},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(logs))
}

func TestDecodeInputData(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.zhejiang.ethpandaops.io")
	if err != nil {
		t.Fatal(err)
	}
	type params struct {
		DealedHeight              *big.Int `json:"_dealedHeight"`
		UserAmount                *big.Int `json:"_userAmount"`
		NodeAmount                *big.Int `json:"_nodeAmount"`
		PlatformAmount            *big.Int `json:"_platformAmount"`
		MaxClaimableWithdrawIndex *big.Int `json:"_maxClaimableWithdrawIndex"`
	}

	abiBts, err := os.ReadFile("../../bindings/Withdraw/withdraw_abi.json")
	if err != nil {
		t.Log(err)
	}
	withdrawal, err := abi.JSON(bytes.NewReader(abiBts))
	if err != nil {
		t.Fatal(err)
	}
	hashs := []string{
		"0x68e297c767a3f65c44a97c464618b064a851b0c18b79c2da44d1f3723fdf35e3",
		"0x07a03487c382f31b34431ffb1ec8a5ee19998bb3a938ea62921d541edd18ec43",
		"0xa97d54212df21092bfc95375be4eee72cd148e514943f3f875a9eb48e9d0cc9b",
		"0x34490e517c0ce6345937b0c594165b1943296e64bfdec073006227b3a16319e3",
		"0x53bbcec2c9efdb8a0067d1336f21253a9874c8ee0c739a2f0461e8e29a6f8470",
		"0xfdf97b3502f4456f2712c29589639ea93233dd01736b8f40c0a1ee787d2867ff",
		"0x1ee51390948a7ee05510011112064864f1cc6fe8c2efd6ce43ec08b229f049d5",
		"0x8b99a8c3ee20d040576ba3f3bb403a4c9eae37198ad69a90176c2cdaad98b96c",
		"0x0aeefd268dfe4b601ff47adc7aa4d1f262391662e638c6bfdfdba8f6ffd3d323",
		"0xf6068990323225c704a87d059da0813ca2b9673b4fecf49ecf01d1f5aa36baef",
		"0x2a2b98b7203240e0289e5e8622b0adde445d9f5a6e526efc147010516831170b",
		"0x8ff158c187e933774471d8612faf96f2ee9cf12e244705adaf3ec700c029824e",
		"0xb4bc1dac092ccc9d3885f79823db28749346d0087c442980895f4c82a5c0e3c1",
		"0xc930066c98dad05903057cfaf7b3dd0eda3fbab47cf3257364754d5c4492e171",
		"0xab513904267fc3e78aa6fdcdd09e6e93051fadb3dd6b1733022a40997f3dc7b5",
		"0xa80f1fa39b67a0485bbc69a74309a70f7063ee845882bf20d63c9d1b3cc37366",
		"0x366ce8a336427d7609c231de4dc076e92419018626c28ff18aa70982a2a96a8f",
		"0x7ea63e6c41cc1ed2447c69e09f409b893d8680856ded4175533354a7164fe4eb",
		"0x236b5c8327b3328de76aa903fb8a2d7db41627f5ddf3e5c4eda2a74858b1fe4b",
		"0x7347fc4aeda8b5064234979f1271df6c33e832c535b2442bd6210704dbf9ac7d",
		"0x509d551269a0c9e85d38a7ec3add46aa21921f28fdf124de2a92db31684ef0ef",
		"0xa654a80cf78983a75e8f168df4d3686aaad960c30415d6574cb8cfbc25e6b348",
		"0x4683eb073ba24faf781ed85eefad74955c2d9dcc578d04b8194b4523aacff404",
		"0xf61d4b0d65336bf9c746a28e05a20fff55d545363b5d9459f54bda59a8def92b",
		"0x88705401927ff936229b153cd3516c8cad9e16173ed51b06f16876658babfc89",
		"0x7a43d8ac64d1435239400be138f96f5b476477941dfe40d52de3b9767f8722ea",
		"0x8554b6785c99f207eb4691da08327f158ed398e925f5e340b5178da238ae3002",
		"0xe32ecdbba7b37499aa9fa9a59989a0be20d0b96cc785eb81a273248a1269348c",
	}

	totalUserAmount := big.NewInt(0)
	totalNodeAmount := big.NewInt(0)
	totalPlatformAmount := big.NewInt(0)

	totalUserAmount1 := big.NewInt(0)
	totalNodeAmount1 := big.NewInt(0)
	totalPlatformAmount1 := big.NewInt(0)

	for _, hash := range hashs {

		tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
		if err != nil {
			t.Fatal(err)
		}
		receipient, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
		if err != nil {
			t.Fatal(err)
		}

		t.Log("iiiiiii")
		for _, log := range receipient.Logs {
			t.Log("log ", log.Data)
		}
		method, err := withdrawal.MethodById(tx.Data()[:4])
		if err != nil {
			t.Fatal(err)
		}

		p := params{}

		inputMap := make(map[string]interface{}, 0)
		err = method.Inputs.UnpackIntoMap(inputMap, tx.Data()[4:])
		if err != nil {
			t.Log(err)
		}
		bts, _ := json.Marshal(inputMap)
		json.Unmarshal(bts, &p)
		total := new(big.Int).Add(new(big.Int).Add(p.UserAmount, p.NodeAmount), p.PlatformAmount)
		t.Log(p, total)
		totalUserAmount1 = new(big.Int).Add(totalUserAmount1, p.UserAmount)
		totalNodeAmount1 = new(big.Int).Add(totalNodeAmount1, p.NodeAmount)
		totalPlatformAmount1 = new(big.Int).Add(totalPlatformAmount1, p.PlatformAmount)

		// totalUserAmount = new(big.Int).Add(totalUserAmount, user.BigInt())
		// totalNodeAmount = new(big.Int).Add(totalNodeAmount, node.BigInt())
		// totalPlatformAmount = new(big.Int).Add(totalPlatformAmount, platform.BigInt())
	}
	t.Log(totalUserAmount, totalNodeAmount, totalPlatformAmount, new(big.Int).Add(totalNodeAmount, totalPlatformAmount))
	t.Log(totalUserAmount1, totalNodeAmount1, totalPlatformAmount1, new(big.Int).Add(totalNodeAmount1, totalPlatformAmount1))
	// {255680 4215416512500000 1328222037500000 291770450000000 3} 5835409000000000
	// utils_test.go:222: 4595384587500000 948253962500000 291770450000000
	// utils_test.go:220: {257920 3525303037500000 1477145212500000 263286750000000 3} 5265735000000000
	// utils_test.go:222: 4146766312500000 855681937500000 263286750000000
}

func TestStorage(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet-rpc.wetez.io/eth/v1/601083a01bf2f40729c5f75e62042208")
	if err != nil {
		t.Fatal(err)
	}

	withdrawPoolContract, err := withdraw.NewWithdraw(common.HexToAddress("0x27d64Dd9172E4b59a444817D30F7af8228F174CC"), client)
	if err != nil {
		t.Fatal(err)
	}
	reserveEthProposalId := utils.ReserveEthForWithdrawProposalId(big.NewInt(19473))

	iter, err := withdrawPoolContract.FilterProposalExecuted(&bind.FilterOpts{
		Context: context.Background(),
	}, [][32]byte{reserveEthProposalId})

	if err != nil {
		t.Log(err)
	}

	for iter.Next() {
		t.Log(iter.Event.ProposalId)
	}
	return
	s, err := storage.NewStorage(common.HexToAddress("0x126d3C08Fb282d5417793B7677E3F7DA8347A384"), client)
	if err != nil {
		t.Fatal(err)
	}
	stafiEtherAddress, err := s.GetAddress(&bind.CallOpts{
		Context: context.Background(),
	}, utils.ContractStorageKey("stafiEther"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("stafiEtherAddress: ", stafiEtherAddress)
	stafiEtherContract, err := stafi_ether.NewStafiEther(stafiEtherAddress, client)
	if err != nil {
		t.Fatal(err)
	}

	stafiDistributorAddress, err := s.GetAddress(&bind.CallOpts{
		Context: context.Background(),
	}, utils.ContractStorageKey("stafiDistributor"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("stafiDistributorAddress: ", stafiDistributorAddress)
	distributorBalance, err := stafiEtherContract.BalanceOf(&bind.CallOpts{}, stafiDistributorAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("stafiDistributor balance: ", decimal.NewFromBigInt(distributorBalance, -18))

	// ------ withdrawal pool
	withdrawPoolAddress, err := s.GetAddress(&bind.CallOpts{
		Context: context.Background(),
	}, utils.ContractStorageKey("stafiWithdraw"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("withdrawPoolAddress: ", withdrawPoolAddress)

	withdrawPoolBalance, err := client.BalanceAt(context.Background(), withdrawPoolAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("withdrawPoolBalance: ", decimal.NewFromBigInt(withdrawPoolBalance, -18))

	withdrawPoolContract, err = withdraw.NewWithdraw(withdrawPoolAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	totalMissingAmountForWithdraw, err := withdrawPoolContract.TotalMissingAmountForWithdraw(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("totalMissingAmountForWithdraw: ", decimal.NewFromBigInt(totalMissingAmountForWithdraw, -18))
	latestDistributeHeight, err := withdrawPoolContract.LatestDistributeHeight(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("latestDistributeWithdrawalHeight: ", latestDistributeHeight)

	maxClaimableWithdrawIndex, err := withdrawPoolContract.MaxClaimableWithdrawIndex(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("maxClaimableWithdrawIndex: ", maxClaimableWithdrawIndex)

	NextWithdrawIndex, err := withdrawPoolContract.NextWithdrawIndex(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("NextWithdrawIndex: ", NextWithdrawIndex)

	unclaimedWithdrawals, err := withdrawPoolContract.GetUnclaimedWithdrawalsOfUser(&bind.CallOpts{}, common.HexToAddress("0x99C6a3B0d131C996D9f65275fB5a196a8B57B583"))
	if err != nil {
		t.Fatal(err)
	}

	for _, w := range unclaimedWithdrawals {
		withdrawal, err := withdrawPoolContract.WithdrawalAtIndex(&bind.CallOpts{}, w)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("withdrawal ", withdrawal.Address, withdrawal.Amount)
	}

	//---------user deposit pool
	userDepositPoolAddress, err := s.GetAddress(&bind.CallOpts{
		Context: context.Background(),
	}, utils.ContractStorageKey("stafiUserDeposit"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("userDepositPoolAddress: ", userDepositPoolAddress)
	userDepositContract, err := user_deposit.NewUserDeposit(userDepositPoolAddress, client)
	if err != nil {
		t.Fatal(err)
	}

	userDepositPoolBalance, err := userDepositContract.GetBalance(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("userDepositPoolBalance: ", decimal.NewFromBigInt(userDepositPoolBalance, -18))
}

// 0x04df80 319360
// 0x039fc6d02bbbc0 1020101175000000

func TestGetOperatorDetail(t *testing.T) {
	detail, err := utils.GetOperatorDetail("prater", 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("detail %+v", detail)
}

func TestGetOperators(t *testing.T) {
	detail, err := utils.GetOperators("prater")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("detail %+v", detail)
}
func TestGetGas(t *testing.T) {
	base, err := utils.GetGaspriceFromBeacon()
	if err != nil {
		t.Log(err)
	}
	t.Log(base)
	client, err := ethclient.Dial("https://mainnet-rpc.wetez.io/eth/v1/601083a01bf2f40729c5f75e62042208")
	if err != nil {
		t.Fatal(err)
	}
	gasTip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		t.Log(err)
	}

	t.Log(gasTip)
}
