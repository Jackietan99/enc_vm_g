package data

import (
	"bytes"
	"math/big"

	"github.com/shopspring/decimal"
)

type WithdrawTx struct {
	ChainID *big.Int
	Value   decimal.Decimal
	Data    []byte
	// Signature values
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (w WithdrawTx) txType() byte {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) data() []byte {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) nonce() int64 {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) copy() TxData {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) rawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) setSignatureValues(typeId, v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) encode(buffer *bytes.Buffer) error {
	//TODO implement me
	panic("implement me")
}

func (w WithdrawTx) decode(i []byte) error {
	//TODO implement me
	panic("implement me")
}
