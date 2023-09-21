package data

import (
	"bytes"
	"enc_vm_g/internal/rlp"
	"enc_vm_g/internal/types"
	"enc_vm_g/pkg"
	"errors"
	"math/big"
	"sync/atomic"
	"time"
)

var (
	ErrInvalidSig         = errors.New("invalid transaction v, r, s values")
	ErrInvalidTxType      = errors.New("transaction type not valid in this context")
	ErrTxTypeNotSupported = errors.New("transaction type not supported")
	errShortTypedTx       = errors.New("typed transaction too short")
)

const (
	TransactionType = 0x00

	/**/
	DepositTxType  = 0x01
	WithdrawTxType = 0x02
	ActivityTxType = 0x03
	RecordTxType   = 0x04
)

/*
 * Transaction data interface
 */
type TxData interface {
	txType() byte // returns the type ID of the transaction
	data() []byte
	nonce() int64
	copy() TxData
	rawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(typeId, v, r, s *big.Int)
	encode(*bytes.Buffer) error
	decode([]byte) error
}

type Transaction struct {
	inner TxData

	time time.Time
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

// NewTx creates a new transaction.
func NewTx(inner TxData) *Transaction {

	tx := new(Transaction)
	tx.setDecoded(inner.copy(), 0)
	return tx
}

// setDecoded sets the inner transaction and size after decoding.
func (tx *Transaction) setDecoded(inner TxData, size uint64) {

	tx.inner = inner
	tx.time = time.Now()
	if size > 0 {
		tx.size.Store(size)
	}
}

// Type returns the transaction type.
func (tx *Transaction) Type() uint8 {
	return tx.inner.txType()
}

func (tx *Transaction) decodeType(b []byte) (TxData, error) {

	if len(b) <= 1 {
		return nil, errShortTypedTx
	}

	var (
		t TxData
	)

	switch b[0] {
	case DepositTxType:
		t = new(DepositTx)
	case WithdrawTxType:
		t = new(WithdrawTx)
	case RecordTxType:
		t = new(RecordTx)
	default:
		return nil, ErrTxTypeNotSupported
	}

	err := t.decode(b[1:])
	return t, err

}

func (tx *Transaction) WithSignature(signer Signer, sig []byte) (*Transaction, error) {

	r, s, v, err := signer.SignatureValues(tx, sig)
	if err != nil {
		return nil, err
	}
	cpy := tx.inner.copy()
	cpy.setSignatureValues(signer.ChainID(), v, r, s)
	return &Transaction{inner: cpy, time: tx.time}, nil
}

// The return values should not be modified by the caller.
func (tx *Transaction) RawSignatureValues() (v, r, s *big.Int) {
	return tx.inner.rawSignatureValues()
}

// m Transaction
func (tx *Transaction) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}

// encodeTyped writes the canonical encoding of a typed transaction to w.
func (tx *Transaction) encodeTyped(w *bytes.Buffer) error {
	w.WriteByte(tx.Type())
	return rlp.Encode(w, tx.inner)
}

func (tx *Transaction) Hash() pkg.Hash {

	if hash := tx.hash.Load(); hash != nil {
		return hash.(pkg.Hash)
	}

	var h = types.RlpHash(tx.inner)
	tx.hash.Store(h)
	return h
}
