package blockchain

import (
	"github.com/michaelpeterswa/alpine/internal/server/logging"
	"nw.codes/handlerr"
)

func Run() {
	logger := logging.InitZap()
	h := handlerr.Handlerr{
		Logger: logger,
	}

	logger.Info("alpine is starting...")

	alp, err := InitAlpine(100000)
	h.Err("create alpine blockchain failed", err)

	alp.PrintBlockchain()

	tx := NewTransaction("asdf", "sdfg", 1)

	block := NewBlock()

	block.AddTransaction(tx)

	err = alp.AddBlock(block)
	h.Err("add block failed", err)

	alp.PrintBlockchain()
}
