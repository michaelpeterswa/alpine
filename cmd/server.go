/*
Copyright © 2021 michaelpeterswa

*/
package cmd

import (
	"github.com/michaelpeterswa/alpine/internal/blockchain"
	"github.com/michaelpeterswa/alpine/internal/logging"
	"github.com/spf13/cobra"
	"nw.codes/handlerr"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Alpine Blockchain Server",
	Long:  `The server for processing the Alpine blockchain.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.InitZap()
		h := handlerr.Handlerr{
			Logger: logger,
		}

		logger.Info("alpine is starting...")

		alp, err := blockchain.InitAlpine(100000)
		h.Err("create alpine blockchain failed", err)

		alp.PrintBlockchain()

		tx := blockchain.NewTransaction("asdf", "sdfg", 1)

		block := blockchain.NewBlock()

		block.AddTransaction(tx)

		err = alp.AddBlock(block)
		h.Err("add block failed", err)

		alp.PrintBlockchain()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
