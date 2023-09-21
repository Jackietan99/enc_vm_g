package main

import (
	"crypto/ecdsa"
	"enc_vm_g/pkg"
	"fmt"
	"log"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"

	"enc_vm_g/internal/crypto"
	"enc_vm_g/internal/crypto/secp256k1"
	"enc_vm_g/internal/data"

	"github.com/shopspring/decimal"
)

// 混淆函数，基于原始密钥的某个字节生成一个混淆字节

func init() {
	println("enc_vm plugin.....")
}

/*
key组成为 id,account,created_at
*/
func EncV1Sign(
	key []byte,
	header []byte,
	Timestamp int64,
	ID string,
	txType int,
	body []byte) (string, error) {

	// 通过
	var (
		s = data.NewSigner()
		i = data.DepositTx{
			ChainID:   big.NewInt(1),
			Value:     decimal.NewFromFloat(1.0),
			Data:      body,
			ID:        ID,
			Header:    pkg.BytesToHash(header),
			Timestamp: big.NewInt(Timestamp)}
		tx = data.NewTx(i)
	)

	sig, err := crypto.Sign(s.Hash(data.NewTx(i)).Bytes(), produceKey(key))
	if err != nil {
		return "", err
	}

	t, err := tx.WithSignature(s, sig)
	if err != nil {
		return "", err
	}

	return t.Hash().Hex(), nil
}

func EncV1Verify(data []byte) bool {
	return true
}

func produceKey(key []byte) *ecdsa.PrivateKey {

	var (
		paddingNeeded = 32 - len(key) // Assuming pkg.HashLength is 32
		paddedKey     = make([]byte, 32)
	)

	for i := 0; i < paddingNeeded; i++ {
		paddedKey[i] = obfuscate(key[0])
	}

	copy(paddedKey[paddingNeeded:], key)
	fmt.Printf("Padded Key: %x\n", paddedKey)

	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: btcec.S256(),
		},
		D: new(big.Int).SetBytes(paddedKey),
	}

	// Compute the public key
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(paddedKey)

	if priv.D.Cmp(secp256k1.S256().Params().N) >= 0 || priv.D.Sign() <= 0 {
		log.Fatalf("The private key is not valid.")
	}
	return priv
}

func obfuscate(b byte) byte {
	return b ^ 0xFF
}
