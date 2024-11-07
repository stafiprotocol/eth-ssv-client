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
	for i := 0; i < 1000; i++ {
		detail, err := utils.MustGetOperatorDetail("mainnet", uint64(i))
		if err != nil {
			continue
			// if strings.Contains(err.Error(), "404") {
			// 	t.Log("operatorId :", i)
			// 	continue
			// } else {
			// 	t.Fatal(err)
			// }
		}
		t.Log(i, detail.Active)

	}
}

func TestGetOperatorFromGraph(t *testing.T) {

	for operator := range []uint64{695, 82, 749, 626, 459, 48, 474, 356, 1133, 1131, 117, 159} {
		o, err := utils.GetOperatorFromGraph("mainnet", uint64(operator))
		if err != nil {
			t.Log(err)
		}
		t.Log("graph", o)

		d, err := utils.GetOperatorFromApi("mainnet", uint64(operator))
		if err != nil {
			t.Log(err)
		}
		t.Log("api", d)
	}

}
