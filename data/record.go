package data

import (
	"bytes"
	"math/big"
)

type RecordTx struct {
}

func (r RecordTx) txType() byte {
	//TODO implement me
	panic("implement me")
}

func (r RecordTx) data() []byte {
	//TODO implement me
	panic("implement me")
}

func (r RecordTx) nonce() int64 {
	//TODO implement me
	panic("implement me")
}

func (r RecordTx) copy() TxData {
	//TODO implement me
	panic("implement me")
}

func (x RecordTx) rawSignatureValues() (v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (x RecordTx) setSignatureValues(typeId, v, r, s *big.Int) {
	//TODO implement me
	panic("implement me")
}

func (r RecordTx) encode(buffer *bytes.Buffer) error {
	//TODO implement me
	panic("implement me")
}

func (r RecordTx) decode(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
