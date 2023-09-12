package data

import (
	"bytes"
	"math/big"
)

type DepositTx struct {
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
