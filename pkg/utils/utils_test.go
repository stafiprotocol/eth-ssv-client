package utils_test

import (
	"context"
	"testing"

	"github.com/depocket/multicall-go/call"
	"github.com/ethereum/go-ethereum/ethclient"
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

func TestMultiCall(t *testing.T) {
	ethClient, err := ethclient.Dial("https://goerli.infura.io/v3/4190eadf65d1449cb6401fc6d5256d2c")
	if err != nil {
		t.Fatal(err)
	}

	caller := call.NewContractBuilder().WithClient(ethClient).AddMethod("getOperatorById(uint64)(address,uint256,uint32,address,bool,bool)")
	data, err := caller.AddCall("test", "0xAE2C84c48272F5a1746150ef333D5E5B51F68763", "getOperatorById", uint64(1)).FlexibleCall(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
