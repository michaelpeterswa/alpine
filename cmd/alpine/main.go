package main

import (
	"github.com/michaelpeterswa/alpine/internal/blockchain"
	"github.com/michaelpeterswa/alpine/internal/logging"
	"nw.codes/handlerr"
)

func main() {
	logger := logging.InitZap()
	h := handlerr.Handlerr{
		Logger: logger,
	}

	logger.Info("alpine is starting...")

	alp := blockchain.InitAlpine()

	tx := blockchain.NewTransaction("asdf", "sdfg", 1)

	block := blockchain.NewBlock()

	block.AddTransaction(tx)

	err := alp.AddBlock(block)
	h.Err("add block failed", err)

	alp.PrintBlockchain()

}
