/*
Copyright © 2021 michaelpeterswa

*/
package cmd

import (
	"github.com/michaelpeterswa/alpine/internal/wallet"
	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Open the Alpenglow Wallet",
	Long:  `Open the wallet tool used for holding and transacting Alpenglow tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		wallet.Run()
	},
}

func init() {
	rootCmd.AddCommand(walletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// walletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// walletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
