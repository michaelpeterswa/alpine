package blockchain

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var (
	once       sync.Once
	blockchain *AlpineBlockchain
)

// Singleton Pattern using sync.Once to ensure that only one AlpineBlockchain is created.
//
// Returns: *AlpineBlockchain
func InitAlpine() *AlpineBlockchain {
	once.Do(func() {
		blockchain = &AlpineBlockchain{}
	})
	return blockchain
}

// NewTransaction() takes a sender "s" and reciever "r" as well as an amount "a" and records it into a transaction object.
//
// Returns: Transaction
func NewTransaction(s string, r string, a float64) Transaction {
	return Transaction{
		Sender:   s,
		Receiver: r,
		Amount:   1.0,
	}
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
func (b *Block) AddTransaction(t Transaction) {
	b.Transactions = append(b.Transactions, t)
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

// Gets the last block on *AlpineBlockchain ab, used to calculate previous hash
//
// Returns: *Block
func (ab *AlpineBlockchain) GetLastBlock() *Block {
	chainLen := len(ab.Chain)
	if chainLen == 0 {
		return nil
	}
	return &ab.Chain[chainLen-1]
}

// Add *Block b to *AlpineBlockchain ab
//
// Returns: error
func (ab *AlpineBlockchain) AddBlock(b *Block) error {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	b.Index = int64(len(ab.Chain) + 1)
	prevHash, err := ab.GetLastBlock().Hash()
	if err != nil {
		return err
	}
	b.PreviousHash = prevHash
	b.Proof = 0

	ab.Chain = append(ab.Chain, *b)
	return nil
}

// Prints indented JSON representing the current blockchain
//
// Returns: error
func (ab *AlpineBlockchain) PrintBlockchain() error {
	res, err := json.MarshalIndent(ab, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
