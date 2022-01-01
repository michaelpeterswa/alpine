package blockchain

import "crypto/ed25519"

// SignedTransaction is a structure that contains the ED25519 Key and Signature used in verification.
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
