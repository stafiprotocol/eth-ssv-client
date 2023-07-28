package bls

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"sync"
)

// Size of an Ethereum BLS public key, in bytes.
var blsPubKeySize = 48

// PublicKey used in the BLS signature scheme.
type PublicKey struct {
	key        *bls.PublicKey
	serialized []byte
	accessMu   sync.RWMutex
}

// PublicKeyFromBytes creates a BLS public key from a byte slice.
func PublicKeyFromBytes(pub []byte) (*PublicKey, error) {
	if len(pub) != blsPubKeySize {
		return nil, fmt.Errorf("public key must be %d bytes", blsPubKeySize)
	}
	var key bls.PublicKey
	if err := key.Deserialize(pub); err != nil {
		return nil, fmt.Errorf("failed to deserialize public key: %v", err)
	}
	return &PublicKey{
		key: &key,
	}, nil
}

func EmptyPublicKey() *PublicKey {
	return &PublicKey{
		key: &bls.PublicKey{},
	}
}

// Aggregate two public keys.  This updates the value of the existing key.
func (k *PublicKey) Aggregate(other *PublicKey) {
	k.accessMu.Lock()
	k.key.Add(other.key)
	k.serialized = nil
	k.accessMu.Unlock()
}

// Marshal a BLS public key into a byte slice.
func (k *PublicKey) Marshal() []byte {
	k.accessMu.Lock()
	if k.serialized == nil {
		k.serialized = k.key.Serialize()
	}
	res := make([]byte, blsPubKeySize)
	copy(res, k.serialized)
	k.accessMu.Unlock()

	return res
}

// Copy creates a copy of the public key.
func (k *PublicKey) Copy() *PublicKey {
	k.accessMu.Lock()

	if k.serialized == nil {
		k.serialized = k.key.Serialize()
	}

	var newKey bls.PublicKey
	//#nosec G104
	_ = newKey.Deserialize(k.serialized)

	key := &PublicKey{
		key: &newKey,
	}

	if k.serialized != nil {
		key.serialized = make([]byte, blsPubKeySize)
		copy(key.serialized, k.serialized)
	}

	k.accessMu.Unlock()

	return key
}

func (p *PublicKey) Set(pkeys []*PublicKey, id *ID) error {
	msk := make([]bls.PublicKey, len(pkeys))
	for i, k := range pkeys {
		if k != nil {
			msk[i] = *k.key
		}
	}

	return p.key.Set(msk, &id.Id)
}

func (p *PublicKey) SerializeToHexStr() string {
	return p.key.SerializeToHexStr()
}
