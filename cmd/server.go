/*
Copyright © 2021 michaelpeterswa

*/
package cmd

import (
	"github.com/michaelpeterswa/alpine/internal/server/blockchain"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Alpine Blockchain Server",
	Long:  `The server for processing the Alpine blockchain.`,
	Run: func(cmd *cobra.Command, args []string) {
		blockchain.Run(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().Int64P("tpb", "b", 2, "Transactions Per Block")
	serverCmd.PersistentFlags().Int64P("circulation", "c", 10, "Coins in Circulation")
}
