package blockchain

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelpeterswa/alpine/internal/server/logging"
	"nw.codes/handlerr"
)

var h handlerr.Handlerr

func Run() {
	logger := logging.InitZap()
	h = handlerr.Handlerr{
		Logger: logger,
	}

	logger.Info("alpine is starting...")

	alp, err := InitAlpine(100000)
	h.Err("create alpine blockchain failed", err)

	txPool := NewTransactionPool()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/transaction/new", txPool.NewTransactionHandler).Methods("POST")
	r.HandleFunc("/transaction/pool", txPool.TransactionPoolHandler).Methods("GET")
	r.HandleFunc("/blockchain", alp.BlockchainHandler).Methods("GET")

	http.ListenAndServe("localhost:8080", r)
}
