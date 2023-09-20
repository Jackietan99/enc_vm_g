package main

import (
	"crypto/ecdsa"
	"enc_vm_g/pkg"
	"fmt"
	"log"
	"math/big"

	"enc_vm_g/internal/constant"
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
	body []byte) ([]byte, error) {

	// 通过
	var (
		s = data.V1Signer{
			Salt: constant.V1,
		}
		i = data.DepositTx{
			ChainID:   big.NewInt(1),
			Value:     decimal.NewFromFloat(1.0),
			Data:      body,
			ID:        ID,
			Header:    pkg.BytesToHash(header),
			Timestamp: Timestamp}
	)

	return crypto.Sign(s.Hash(data.NewTx(i)).Bytes(), produceKey(key))
}

func EncV1Verify(data []byte) bool {
	return true
}

func produceKey(key []byte) *ecdsa.PrivateKey {

	// 计算需要填充多少个字节

	var (
		paddingNeeded = pkg.HashLength - len(key)
		paddedKey     = make([]byte, pkg.HashLength)
	)

	// 创建一个新切片用于存储混淆的填充字节
	// 使用原始密钥的第一个字节（或任何其他字节）来生成混淆字节
	for i := 0; i < paddingNeeded; i++ {
		paddedKey[i] = obfuscate(key[0]) // 或者obfuscate(originalKey[i % len(originalKey)])
	}

	// 将原始密钥拷贝到新的切片中，形成完整的32字节密钥
	copy(paddedKey[paddingNeeded:], key)

	// 输出填充后的密钥
	fmt.Printf("Padded Key: %x\n", paddedKey)

	// 使用paddedKey创建一个ECDSA私钥
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	priv.D = new(big.Int).SetBytes(paddedKey)

	// 你现在可以使用这个私钥进行以太坊或其他基于ECDSA的操作
	// ...

	// 验证这是否是一个有效的ECDSA私钥（可选）
	if priv.D.Cmp(secp256k1.S256().Params().N) >= 0 || priv.D.Sign() <= 0 {
		log.Fatalf("The private key is not valid.")
	}
	return priv
}

func obfuscate(b byte) byte {
	return b ^ 0xFF
}
