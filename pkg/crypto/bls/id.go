package bls

import (
	"fmt"

	"github.com/herumi/bls-eth-go-binary/bls"
)

type ID struct {
	Id bls.ID
}

func (id *ID) SetDec(dec int) error {
	return id.Id.SetDecString(fmt.Sprintf("%d", dec))
}

func (id *ID) GetHexString() string {
	return id.Id.GetHexString()
}
