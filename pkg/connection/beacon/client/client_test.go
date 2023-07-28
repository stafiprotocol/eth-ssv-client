package client_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/signing"
	"github.com/prysmaticlabs/prysm/v3/config/params"
	"github.com/prysmaticlabs/prysm/v3/contracts/deposit"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/beacon/client"
	"github.com/stafiprotocol/eth-ssv-client/pkg/connection/types"
	"github.com/stafiprotocol/eth-ssv-client/pkg/utils"
)

func TestStatus(t *testing.T) {
	// c, err := client.NewStandardHttpClient("https://27Y0WDKrX1dYIkBXOugsSLh9hfr:a7c3849eba862fdd67382dab42e2a23c@eth2-beacon-mainnet.infura.io")
	// c, err := client.NewStandardHttpClient("https://beaconcha-rpc2.stafi.io")
	// c, err := client.NewStandardHttpClient("https://beacon.zhejiang.ethpandaops.io")
	c, err := client.NewStandardHttpClient("https://beacon-lighthouse-zhejiang.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	status, err := c.GetSyncStatus()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(status)

	head, err := c.GetBeaconHead()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(head.Epoch)
	duties, err := c.GetValidatorProposerDuties(14107)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(duties)
	// pubkey, err := types.HexToValidatorPubkey("af93696b857fb621048539d0f9ee7722d801e05cf3be3039decd17b937dd9d69f4450c407f5ae4e96d875cb754840c1c")
	// pubkey, err := types.HexToValidatorPubkey("b427ea30366336e4632d327428fac24ac3016534b18e0e39f5c2c4fffaa35656f982fba8e636599ae54b6f148a90a8e9")
	// pubkey, err := types.HexToValidatorPubkey("ae9d34a72d6d16c17e3703a12eeaa45063128046055516f0611a337caaea7cf823e1ae8c9298154c325fc10fcb279d42")
	// pubkey, err := types.HexToValidatorPubkey("b3ea762c11ef4548d7c2a1a707f69cf68a1f1b7fc63c7dcb414d6a7ab722e2155d7e3ac3b601abdb98e158ca6035e9c4")
	pubkey, err := types.HexToValidatorPubkey("b44b3b379d7e8bdb043336ba901c5c4a685041d48bc783b351ad3873b6cfca7f9142465ed56c82018cbde52d5ca709be")
	// pubkey, err := types.HexToValidatorPubkey("ad90505f19a31915940316ba5178984ae7e9164628eae689b3f99f2e50079ff421fd09edf46b8080322b4b0b7a5d2b26")
	if err != nil {
		t.Fatal(err)
	}
	epoch1 := uint64(13199)
	// epoch3 := uint64(167678)
	startStatus1, err := c.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{
		Epoch: &epoch1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v %v, %v,%d", startStatus1.Balance, startStatus1.Slashed, startStatus1.Status, startStatus1.ExitEpoch)
	// startStatus2, err := c.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{
	// 	Epoch: &epoch2,
	// })

	// t.Logf("%+v %v", startStatus2.Balance, startStatus2.Slashed)
	// startStatus3, err := c.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{
	// 	Epoch: &epoch3,
	// })

	// t.Logf("%+v", startStatus3.Balance)
	// reward1 := startStatus1.Balance - startStatus2.Balance
	// reward2 := startStatus2.Balance - startStatus3.Balance
	// t.Log(reward1, reward2, (reward1-reward2)/6)

}

func TestTx(t *testing.T) {
	ethClient, err := ethclient.Dial("https://test-eth-node.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	blockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(blockNumber)
	receipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash("0x06456bdf80482c3e0b59e53720438630f152f4fbcc7e02ab614e83198b1805be"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", receipt)

}

func TestBlock(t *testing.T) {
	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/4d058381a4d64d31b00a4e15df3ddb94")
	if err != nil {
		t.Fatal(err)
	}
	// blockNumber, err := ethClient.BlockNumber(context.Background())
	// if err != nil {
	// 	t.Fatal(err)
	// }
	blockNumber := 15541242
	t.Log(blockNumber)
	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", block.Header())

}
func TestBeaconBlock(t *testing.T) {
	// c, err := client.NewStandardHttpClient("https://27Y0WDKrX1dYIkBXOugsSLh9hfr:a7c3849eba862fdd67382dab42e2a23c@eth2-beacon-mainnet.infura.io")
	c, err := client.NewStandardHttpClient("https://beacon.zhejiang.ethpandaops.io")
	if err != nil {
		t.Fatal(err)
	}
	block, exists, err := c.GetBeaconBlock(263205)
	// block, _, err := c.GetBeaconBlock("5071581")
	// block, _, err := c.GetBeaconBlock("3339591")
	if err != nil {
		t.Fatal(err)
	}

	config, err := c.GetEth2Config()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", block)
	t.Logf("%v", exists)
	t.Logf("%+v", utils.StartSlotOfEpoch(config, 4652))
	t.Logf("%+v", block.Withdrawals)

}
func TestBeaconHead(t *testing.T) {
	// c, err := client.NewStandardHttpClient("https://27Y0WDKrX1dYIkBXOugsSLh9hfr:a7c3849eba862fdd67382dab42e2a23c@eth2-beacon-mainnet.infura.io")
	c, err := client.NewStandardHttpClient("https://beaconcha-rpc2.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	head, err := c.GetBeaconHead()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", head)

}
func TestConfig(t *testing.T) {
	// c, err := client.NewStandardHttpClient("https://27Y0WDKrX1dYIkBXOugsSLh9hfr:a7c3849eba862fdd67382dab42e2a23c@eth2-beacon-mainnet.infura.io")
	c, err := client.NewStandardHttpClient("https://beaconcha-rpc2.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	head, err := c.GetEth2Config()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", head)

	duties, err := c.GetValidatorProposerDuties(167578)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(duties)
	sc, err := c.GetSyncCommitteesForEpoch(167578)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sc)

	block, exist, err := c.GetBeaconBlock(5362523)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(block.SyncAggregate)
	t.Log(exist)

	// com,err:=c.GetCommitteesForEpoch(167578)
	// if err!=nil{
	// 	t.Fatal(err)
	// }
	// t.Log(com)

}

func TestSigs(t *testing.T) {
	withdrawBts, err := hexutil.Decode("0x003cd051a5757b82bf2c399d7476d1636473969af698377434af1d6c54f2bee9")
	if err != nil {
		t.Fatal(err)
	}
	sigBts, err := hexutil.Decode("0xaf6a1644b29ed4e8c012804dd1f507828a6001d776c3b026eca4eec8a82aa9d410603906c392891b5a2e53e0d16f0a7505080818eeaaba6f8caecf57ebc99c0b0bfe1a0b756bb3b5b2f4346bfb8d7c1c40e17f515cdca28e5526fda328fc68f4")
	if err != nil {
		t.Fatal(err)
	}
	validatorPubkey, err := types.HexToValidatorPubkey("b9eb2b1215aa1933d6d7361e7cf1182fef12c5d6643f8bb9fc373c059de7a066d9a6eb893cf355268b39980977331967")
	if err != nil {
		t.Fatal(err)
	}

	dp := ethpb.Deposit_Data{
		PublicKey:             validatorPubkey.Bytes(),
		WithdrawalCredentials: withdrawBts,
		Amount:                1e9,
		Signature:             sigBts,
	}

	domain, err := signing.ComputeDomain(
		params.BeaconConfig().DomainDeposit,
		params.BeaconConfig().GenesisForkVersion,
		params.BeaconConfig().ZeroHash[:],
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := deposit.VerifyDepositSignature(&dp, domain); err != nil {
		t.Fatal(err)
	}
}
