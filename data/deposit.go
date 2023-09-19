package data

import (
	"bytes"
	"math/big"

	"github.com/shopspring/decimal"
)

type DepositTx struct {
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

func (d DepositTx) txType() byte {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) data() []byte {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) nonce() int64 {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) copy() TxData {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) rawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) setSignatureValues(typeId, v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) encode(buffer *bytes.Buffer) error {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) decode(i []byte) error {
	//TODO implement me
	panic("implement me")
}
