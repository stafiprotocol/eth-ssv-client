package keyshare

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/bls"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/rsa"
)

type Share struct {
	privateKey string
	publicKey  string
	id         string
}

type Operator struct {
	Id        uint64          `json:"id"`
	PublicKey string          `json:"publicKey"`
	Fee       decimal.Decimal `json:"fee"`
}

type KeystoreShareInfo struct {
	ID              uint64
	PublicKey       string
	secretKey       string
	EncryptedKey    string
	AbiEncryptedKey string
	Operator        *Operator
}

// CreateThreshold receives a bls.SecretKey hex and count.
// Will split the secret key into count shares
func CreateThreshold(skBytes []byte, operators []*Operator) (map[uint64]*Share, error) {
	threshold := uint64(len(operators))

	// master key Polynomial
	msk := make([]*bls.PrivateKey, threshold)
	mpk := make([]*bls.PublicKey, threshold)

	sk, err := bls.PrivateKeyFromBytes(skBytes)
	if err != nil {
		return nil, err
	}
	msk[0] = sk
	mpk[0] = sk.PublicKey()

	_F := (threshold - 1) / 3

	// Receives list of operators IDs. len(operator IDs) := 3 * F + 1
	// construct poly
	for i := uint64(1); i < threshold-_F; i++ {
		sk, err := bls.GeneratePrivateKey()
		if err != nil {
			return nil, err
		}
		msk[i] = sk
		mpk[i] = sk.PublicKey()
	}

	// evaluate shares - starting from 1 because 0 is master key
	shares := make(map[uint64]*Share)
	for i := uint64(1); i <= threshold; i++ {
		blsID := bls.ID{}
		// not equal to ts ?
		operatorId := operators[i-1].Id
		err := blsID.SetDec(operatorId)
		if err != nil {
			return nil, err
		}

		sk := bls.EmptyPrivateKey()
		err = sk.Set(msk, &blsID)
		if err != nil {
			return nil, err
		}

		pk := bls.EmptyPublicKey()
		err = pk.Set(mpk, &blsID)
		if err != nil {
			return nil, err
		}

		shares[i] = &Share{
			privateKey: sk.SerializeToHexStr(),
			publicKey:  pk.SerializeToHexStr(),
			id:         blsID.GetHexString(),
		}
	}
	return shares, nil
}

func EncryptShares(skBytes []byte, operators []*Operator) ([]*KeystoreShareInfo, error) {
	shareCount := uint64(len(operators))
	shares, err := CreateThreshold(skBytes, operators)
	if err != nil {
		return nil, errors.Wrap(err, "creating threshold err")
	}

	keystoreShareInfos := make([]*KeystoreShareInfo, 0, shareCount)
	for i := 0; i < int(shareCount); i++ {
		share := shares[uint64(i)+1]

		operator := operators[i]
		opk, err := base64.StdEncoding.DecodeString(operator.PublicKey)
		if err != nil {
			return nil, errors.Wrapf(err, "operator pubkey decode err. pubkey: %s", operator.PublicKey)
		}

		shareSk := "0x" + share.privateKey
		sharePk := "0x" + share.publicKey

		decryptShareSecret, err := rsa.PublicEncrypt(shareSk, string(opk))
		if err != nil {
			return nil, err
		}
		abiShareSecret, err := AbiCoder([]string{"string"}, []interface{}{decryptShareSecret})
		if err != nil {
			return nil, err
		}

		keystoreShareInfo := &KeystoreShareInfo{
			ID:              uint64(i),
			PublicKey:       sharePk,
			secretKey:       shareSk,
			EncryptedKey:    decryptShareSecret,
			AbiEncryptedKey: "0x" + hex.EncodeToString(abiShareSecret),
			Operator:        operator,
		}
		keystoreShareInfos = append(keystoreShareInfos, keystoreShareInfo)
	}
	return keystoreShareInfos, nil
}
