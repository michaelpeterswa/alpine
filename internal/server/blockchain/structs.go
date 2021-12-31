package blockchain

import (
	"crypto/ed25519"
	"sync"
	"time"
)

type SignedTransaction struct {
	PublicKey *ed25519.PublicKey `json:"public-key"`
	Signature []byte             `json:"signature"`
	Tx        Transaction        `json:"transaction"`
}

// Transaction is a simple structure that has a sender, reciever, and amount.
type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

// Block is a structure that contains an index, timestamp, slice of transactions, proof value, and the hash of the previous block.
type Block struct {
	Index        int64         `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int64         `json:"proof"`
	PreviousHash string        `json:"previous-hash"`
}

// AlpineBlockchain is a structure that conains a mutex lock and the blockchain itself.
type AlpineBlockchain struct {
	mu          sync.Mutex
	Circulation int64   `json:"-"`
	Chain       []Block `json:"blockchain"`
}
