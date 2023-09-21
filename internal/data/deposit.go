package data

import (
	"bytes"
	"enc_vm_g/pkg"
	"math/big"

	"github.com/shopspring/decimal"
)

type DepositTx struct {
	ChainID   *big.Int
	Value     decimal.Decimal
	Data      []byte
	ID        string
	Timestamp *big.Int
	Header    pkg.Hash
	// Signature values
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (d DepositTx) txType() byte {
	return DepositTxType
}

func (d DepositTx) data() []byte {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) nonce() int64 {
	//TODO implement me
	panic("implement me")
}

func (tx DepositTx) copy() TxData {
	cpy := &DepositTx{

		Data: pkg.CopyBytes(tx.Data),
		// These are initialized below.
		Value: decimal.Zero,
		V:     new(big.Int),
		R:     new(big.Int),
		S:     new(big.Int),
	}

	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

func (d DepositTx) rawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) setSignatureValues(typeId, v, r, s *big.Int) {
	d.ChainID, d.V, d.R, d.S = typeId, v, r, s
}

func (d DepositTx) encode(buffer *bytes.Buffer) error {
	//TODO implement me
	panic("implement me")
}

func (d DepositTx) decode(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
