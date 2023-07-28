package credential_test

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stafiprotocol/eth-ssv-client/pkg/constants"
	"github.com/stafiprotocol/eth-ssv-client/pkg/credential"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/mnemonic"
)

func TestNewCredential(t *testing.T) {
	mnemonicStr := "emotion problem need nice museum proud room sell field impact ankle poet child video luggage awful next trophy engine forest price average husband exotic"
	seed := mnemonic.NewSeed(mnemonicStr, "")

	credential, err := credential.NewCredential(seed, 0, big.NewInt(12e9), constants.GetChain(constants.ChainMAINNET), common.HexToAddress("0x27d64dd9172e4b59a444817d30f7af8228f174cc"))
	if err != nil {
		t.Fatal(err)
	}

	targetPubkey := hexutil.MustDecode("0x907b4c58ab0de67fc9db9e91063ab7042dcb471fa751ee1294c81bf9190a40d74d04731c6e1b09e8ed1214f85341f112")
	targetSig := hexutil.MustDecode("0xb9ddbb9c54c278c86a99f088cc56bcc73ac8a71bc3d988d0f8f92191d5fd452be88c882707d0cfd7953364daf95c4e2013f625b3e726b101a78e2c41924107980b79c5e69d0f0f4dcededfcef50d652e867509ade0c1442df2c5031309085571")
	targetDataRoot := hexutil.MustDecode("0xe39085854b63145f6cf52d73be61a78e4676d8548273593119bdcd2281f9e937")
	targetMessageRoot := hexutil.MustDecode("0x700c2d1179bc6b4a19b02efe724a46f1edbb293db3b009a41510f54b55c30f91")
	if !bytes.Equal(credential.SigningPK().Marshal(), targetPubkey) {
		t.Fatal("pubkey not match")
	}
	depositData, err := credential.SigningDepositData()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(hexutil.MustDecode("0x"+depositData.Signature), targetSig) {
		t.Fatal("sig not match")
	}
	if !bytes.Equal(hexutil.MustDecode("0x"+depositData.DepositDataRoot), targetDataRoot) {
		t.Fatal("sig not match")
	}
	if !bytes.Equal(hexutil.MustDecode("0x"+depositData.DepositMessageRoot), targetMessageRoot) {
		t.Fatal("sig not match")
	}
}
