package blockchain

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/michaelpeterswa/alpine/internal/server/logging"
	"github.com/spf13/cobra"
	"nw.codes/handlerr"
)

var (
	h  handlerr.Handlerr
	rs *RunnerSettings
)

type RunnerSettings struct {
	TPB         int64
	Circulation int64
}

func Run(cmd *cobra.Command) {
	logger := logging.InitZap()
	h = handlerr.Handlerr{
		Logger: logger,
	}

	logger.Info("alpine is starting...")

	settings, err := getSettings(cmd)
	h.Err("get settings failed", err)

	if settings == nil {
		os.Exit(1)
	}

	rs = settings

	alp, err := InitAlpine(rs.Circulation, NewTransactionPool())
	h.Err("create alpine blockchain failed", err)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/transaction/new", alp.TxPool.NewTransactionHandler).Methods("POST")
	r.HandleFunc("/transaction/pool", alp.TxPool.TransactionPoolHandler).Methods("GET")
	r.HandleFunc("/blockchain", alp.BlockchainHandler).Methods("GET")
	r.HandleFunc("/blockchain/forge", alp.ForgeHandler).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}

func getSettings(cmd *cobra.Command) (*RunnerSettings, error) {
	tpb, err := cmd.Flags().GetInt64("tpb")
	if err != nil {
		return nil, err
	}

	circ, err := cmd.Flags().GetInt64("circulation")
	if err != nil {
		return nil, err
	}

	return &RunnerSettings{
		TPB:         tpb,
		Circulation: circ,
	}, nil

}
