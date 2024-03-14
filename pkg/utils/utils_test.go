package utils_test

import (
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
	list := []uint64{198, 45, 48, 121, 200, 156, 83, 88, 22, 55, 20, 90, 49, 33, 120, 71, 135, 179, 132, 73, 208, 187, 14, 21, 134, 7}
	for _, l := range list {

		op, err := utils.GetOperatorFromApi("mainnet", l)
		if err != nil {
			t.Fatal(err, l)
		}
		t.Logf("%d %+v", l, op)
	}
}
