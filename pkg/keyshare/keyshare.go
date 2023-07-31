package keyshare

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/bls"
	"github.com/stafiprotocol/eth-ssv-client/pkg/crypto/rsa"
)

type IShare struct {
	privateKey string
	publicKey  string
	id         interface{}
}

// tx payload
type Payload struct {
	ValidatorPublicKey  string   `json:"validatorPublicKey"`
	OperatorIds         string   `json:"operatorIds"`
	SharePublicKeys     []string `json:"sharePublicKeys"`
	encryptedKeys       []string
	AbiSharePrivateKeys []string `json:"sharePrivateKey"`
	SsvAmount           string   `json:"ssvAmount"`
}

type Operator struct {
	Id        int             `json:"id"`
	PublicKey string          `json:"publicKey"`
	Fee       decimal.Decimal `json:"fee"`
}

type Shares struct {
	PublicKeys    []string `json:"publicKeys"`
	EncryptedKeys []string `json:"encryptedKeys"`
}

type KeystoreShareData struct {
	PublicKey string      `json:"publicKey"`
	Operators []*Operator `json:"operators"`
	Shares    *Shares     `json:"shares"`
}

type KeystoreShareRes struct {
	Version string `json:"version"`

	Data *KeystoreShareData `json:"data"`

	Payload struct {
		Readable *Payload `json:"readable"`
		Raw      string   `json:"raw"`
	} `json:"payload"`

	CreatedAt time.Time `json:"createdAt"`
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
func CreateThreshold(skBytes []byte, operators []*Operator) (map[uint64]*IShare, error) {
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
	shares := make(map[uint64]*IShare)
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

		shares[i] = &IShare{
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

func buildPayload(validatorPublicKey, ssvAmount string, operators []*Operator, keystoreShareHelpers []*KeystoreShareInfo) *Payload {
	operatorIds := ""
	for _, operator := range operators {
		operatorIds += strconv.Itoa(operator.Id) + ","
	}
	operatorIds = strings.TrimRight(operatorIds, ",")

	count := len(keystoreShareHelpers)
	sharePublicKeys := make([]string, 0, count)
	abiSharePrivateKeys := make([]string, 0, count)
	encryptedSharePrivateKeys := make([]string, 0, count)

	for _, helper := range keystoreShareHelpers {
		sharePublicKeys = append(sharePublicKeys, helper.PublicKey)
		abiSharePrivateKeys = append(abiSharePrivateKeys, helper.AbiEncryptedKey)
		encryptedSharePrivateKeys = append(encryptedSharePrivateKeys, helper.EncryptedKey)
	}

	payload := &Payload{
		ValidatorPublicKey:  validatorPublicKey,
		OperatorIds:         operatorIds,
		SharePublicKeys:     sharePublicKeys,
		AbiSharePrivateKeys: abiSharePrivateKeys,
		SsvAmount:           ssvAmount,
		encryptedKeys:       encryptedSharePrivateKeys,
	}

	return payload
}

func buildPayloadRaw(payload *Payload) string {
	sharePubkeysStr := strings.Replace(strings.Trim(fmt.Sprint(payload.SharePublicKeys), "[]"), " ", ",", -1)
	abiShareSecretsStr := strings.Replace(strings.Trim(fmt.Sprint(payload.AbiSharePrivateKeys), "[]"), " ", ",", -1)

	raw := fmt.Sprintf("%s,%s,%s,%s,%s", payload.ValidatorPublicKey, payload.OperatorIds, sharePubkeysStr, abiShareSecretsStr, payload.SsvAmount)
	return raw
}

func buildKeystoreShareRes(validatorPublicKey, version, ssvAmount string, operators []*Operator, keystoreShareHelpers []*KeystoreShareInfo) *KeystoreShareRes {
	payload := buildPayload(validatorPublicKey, ssvAmount, operators, keystoreShareHelpers)
	raw := buildPayloadRaw(payload)

	keystoreShareData := &KeystoreShareData{
		PublicKey: validatorPublicKey,
		Operators: operators,
		Shares: &Shares{
			PublicKeys:    payload.SharePublicKeys,
			EncryptedKeys: payload.encryptedKeys,
		},
	}

	keystoreShareRes := &KeystoreShareRes{
		Version:   version,
		CreatedAt: time.Now(),
		Data:      keystoreShareData,
	}
	keystoreShareRes.Payload.Readable = payload
	keystoreShareRes.Payload.Raw = raw

	return keystoreShareRes
}

func KeystoreShare(validatorPublicKey, version, ssvAmount string, skBytes []byte, operators []*Operator) (*KeystoreShareRes, error) {
	keystoreShareInfos, err := EncryptShares(skBytes, operators)
	if err != nil {
		return nil, errors.Unwrap(err)
	}
	keystoreSharesRes := buildKeystoreShareRes(validatorPublicKey, version, ssvAmount, operators, keystoreShareInfos)

	return keystoreSharesRes, nil
}

func KeystoreShareForJson(validatorPublicKey, version, ssvAmount string, skBytes []byte, operators []*Operator) (string, error) {
	keystoreSharesRes, err := KeystoreShare(validatorPublicKey, version, ssvAmount, skBytes, operators)
	if err != nil {
		return "", errors.Unwrap(err)
	}

	jsonRes, err := json.Marshal(keystoreSharesRes)
	if err != nil {
		return "", errors.Wrap(err, "failed to json marshal.")
	}

	return string(jsonRes), nil
}
