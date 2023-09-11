package data

import "math/big"

type TxData interface {
	txType() byte // returns the type ID of the transaction
	data() []byte
	nonce() int64
	copy() TxData
	rawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(typeId, v, r, s *big.Int)
}
