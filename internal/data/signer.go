package data

import (
	"enc_vm_g/internal/constant"
	"enc_vm_g/internal/crypto"
	"enc_vm_g/internal/types"
	"enc_vm_g/pkg"
	"errors"
	"fmt"
	"math/big"
)

type Signer interface {
	Sender(tx *Transaction) (pkg.Address, error)
	SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error)
	ChainID() *big.Int
	Hash(tx *Transaction) pkg.Hash
	Equal(Signer) bool
}

var (
	big8 = big.NewInt(8)
)

/*
初始版本的签名器
*/
type V1Signer struct {
	Salt                string // 用于生成签名器的盐
	chainId, chainIdMul *big.Int
}

func NewSigner() V1Signer {
	return V1Signer{
		Salt:       constant.V1,
		chainId:    big.NewInt(1),
		chainIdMul: big.NewInt(2),
	}
}

// m V1Signer
func (m V1Signer) Sender(tx *Transaction) (pkg.Address, error) {

	V, R, S := tx.RawSignatureValues()
	V = new(big.Int).Sub(V, m.chainIdMul)
	V.Sub(V, big8)
	return recoverPlain(m.Hash(tx), R, S, V)
}

// m V1Signer
func (s V1Signer) SignatureValues(tx *Transaction, sig []byte) (R, S, V *big.Int, err error) {

	R, S, V = decodeSignature(sig)
	if s.chainId.Sign() != 0 {
		V = big.NewInt(int64(sig[64] + 35))
		V.Add(V, s.chainIdMul)
	}
	return R, S, V, nil
}

// m V1Signer
func (m V1Signer) ChainID() *big.Int {
	return m.chainId
}

// m V1Signer
func (m V1Signer) Hash(tx *Transaction) pkg.Hash {
	return types.RlpHash([]interface{}{
		m.Salt,
		tx.inner,
		tx.time.UnixMilli(),
	})
}

// m V1Signer
func (m V1Signer) Equal(s Signer) bool {
	S, ok := s.(V1Signer)
	return ok && S.chainId.Cmp(m.chainId) == 0
}

func recoverPlain(sighash pkg.Hash, R, S, Vb *big.Int) (pkg.Address, error) {
	if Vb.BitLen() > 8 {
		return pkg.Address{}, ErrInvalidSig
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S) {
		return pkg.Address{}, ErrInvalidSig
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return pkg.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return pkg.Address{}, errors.New("invalid public key")
	}
	var addr pkg.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}

func decodeSignature(sig []byte) (r, s, v *big.Int) {

	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}

	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})

	return r, s, v
}
