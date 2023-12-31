package bls

import (
	"errors"
	"fmt"

	"github.com/herumi/bls-eth-go-binary/bls"
)

// PrivateKey is a private key in Ethereum 2.
// It is a point on the BLS12-381 curve.
type PrivateKey struct {
	key bls.SecretKey
}

// PrivateKeyFromBytes creates a BLS private key from a byte slice.
func PrivateKeyFromBytes(priv []byte) (*PrivateKey, error) {
	if len(priv) != 32 {
		return nil, errors.New("private key must be 32 bytes")
	}
	var sec bls.SecretKey
	if err := sec.Deserialize(priv); err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}
	return &PrivateKey{key: sec}, nil
}

// GeneratePrivateKey generates a random BLS private key.
func GeneratePrivateKey() (*PrivateKey, error) {
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	return &PrivateKey{key: sec}, nil
}

func EmptyPrivateKey() *PrivateKey {
	return &PrivateKey{key: bls.SecretKey{}}
}

// Marshal a secret key into a byte slice.
func (p *PrivateKey) Marshal() []byte {
	return p.key.Serialize()
}

// PublicKey obtains the public key corresponding to the BLS secret key.
func (p *PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{key: p.key.GetPublicKey()}
}

// Sign a message using a secret key.
func (p *PrivateKey) Sign(msg []byte) *Signature {
	sig := p.key.SignHash(msg)
	return &Signature{sig}
}

func (p *PrivateKey) Set(pkeys []*PrivateKey, id *ID) error {
	msk := make([]bls.SecretKey, len(pkeys))
	for i, k := range pkeys {
		if k != nil {
			msk[i] = k.key
		}
	}

	return p.key.Set(msk, &id.Id)
}

func (p *PrivateKey) SerializeToHexStr() string {
	return p.key.SerializeToHexStr()
}
