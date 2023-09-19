package data

import (
	"bytes"
	"math/big"

	"github.com/shopspring/decimal"
)

type ActivityTx struct {
	ChainID   *big.Int
	Value     decimal.Decimal
	Data      []byte
	ID        string
	Timestamp int64
	Header    Hash
	// Signature values
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (d ActivityTx) txType() byte {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) data() []byte {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) nonce() int64 {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) copy() TxData {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) rawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) setSignatureValues(typeId, v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) encode(buffer *bytes.Buffer) error {
	//TODO implement me
	panic("implement me")
}

func (d ActivityTx) decode(i []byte) error {
	//TODO implement me
	panic("implement me")
}
