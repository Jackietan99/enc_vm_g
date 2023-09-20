package data

import "math/big"

type Signer interface {
	Sender(tx *Transaction) (Address, error)
	SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error)
	ChainID() *big.Int
	Hash(tx *Transaction) Hash
	Equal(Signer) bool
}


type 