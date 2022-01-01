/*
Copyright © 2021 michaelpeterswa

*/
package cmd

import (
	"log"
	"os"

	"github.com/michaelpeterswa/alpine/internal/wallet"
	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Open the Alpenglow Wallet",
	Long:  `Open the wallet tool used for holding and transacting Alpenglow tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		wallet.Run(cmd)
	},
}

func init() {
	rootCmd.AddCommand(walletCmd)

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get working directory")
	}
	walletCmd.PersistentFlags().StringP("directory", "d", path, "Set the directory for file operations")
}
