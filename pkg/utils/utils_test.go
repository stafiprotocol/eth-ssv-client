package utils_test

import (
	"strings"
	"testing"

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

func TestGetOperatorFromApi(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := utils.MustGetOperatorDetail("mainnet", uint64(i))
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				t.Log("operatorId :", i)
				continue
			} else {
				t.Fatal(err)
			}

		}

	}
}

func TestGetOperatorFromGraph(t *testing.T) {
	o, err := utils.GetOperatorFromGraph("mainnet", 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(o)

}
