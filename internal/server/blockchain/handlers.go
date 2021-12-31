package blockchain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "alpine")
}

func (txp *TransactionPool) NewTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var tx Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	h.Err("failed to JSON decode request", err)

	// TODO: set up for house wallet
	if tx.Sender == "alpine" {
		h.Logger.Warn("sender is disallowed... dropping...", zap.Any("transaction", tx))
	} else {
		txp.Push(tx)
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
