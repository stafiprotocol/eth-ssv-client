package utils_test

import (
	"os"
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

func TestGetOperatorFromGraph(t *testing.T) {
	operators := []uint64{695, 82, 749, 626, 459, 48, 474, 356, 1133, 1131, 117, 159}
	for operator := range operators {
		d, err := utils.GetOperatorFromApi("mainnet", uint64(operator))
		if err != nil {
			t.Log(err)
		}
		t.Log("-----api", d)
		t.Log("+++++++")
	}
	apikey := os.Getenv("apikey")
	operatorss, err := utils.BatchGetOperatorFromGraph(apikey, operators)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(operatorss)

}
