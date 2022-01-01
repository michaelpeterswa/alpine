package blockchain

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"
)

// Block is a structure that contains an index, timestamp, slice of transactions, proof value, and the hash of the previous block.
type Block struct {
	Index        int64         `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int64         `json:"proof"`
	PreviousHash string        `json:"previous-hash"`
}

// Create new *Block with current timestamp
//
// Returns: *Block
func NewBlock() *Block {
	return &Block{
		Timestamp: time.Now(),
	}
}

// Append Transaction t to *Block b
func (b *Block) AddTransaction(t Transaction) error {
	if len(b.Transactions) < int(rs.TPB) {
		b.Transactions = append(b.Transactions, t)
		return nil
	}
	return errors.New("block full")
}

// SHA512 Hash of *Block b encoded to hexidecimal string for readability
//
// Returns: string (of hash) and error
func (b *Block) Hash() (string, error) {
	res, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	hash := sha512.Sum512(res)
	strHash := hex.EncodeToString(hash[:])
	return strHash, nil
}
