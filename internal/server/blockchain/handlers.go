package blockchain

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaelpeterswa/alpine/internal/aerror"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "alpine")
}

func (txp *TransactionPool) NewTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var tx SignedTransaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	h.Err("failed to JSON decode request", err)

	if len(*tx.PublicKey) != ed25519.PublicKeySize {
		aerror.NewErrorResponse(w, aerror.PublicKeyBadLength)
		return
	}
	if len(tx.Signature) != ed25519.SignatureSize {
		aerror.NewErrorResponse(w, aerror.SignatureBadLength)
		return
	}

	verified, err := VerifyTransaction(tx)
	h.Err("failed to verify transaction", err)

	if verified {
		txp.Push(tx.Tx)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
		aerror.NewErrorResponse(w, aerror.FailedVerification)
	}
}

func (txp *TransactionPool) TransactionPoolHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(txp.List())
	h.Err("failed to JSON encode PendingTransactions", err)
}

func (ab *AlpineBlockchain) BlockchainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(ab)
	h.Err("failed to JSON encode AlpineBlockchain", err)
}
