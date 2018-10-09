package main

import (
	"github.com/OpenOCC/OCC/cmd/ecli/cmd"
	"github.com/OpenOCC/OCC/cmd/ecli/param"
	"github.com/spf13/cobra"
)

var (
	cmds = []*cobra.Command{}
)

func init() {
	cmds = append(cmds, cmd.TransactionCmd, cmd.AccountCmd)
}

func main() {
	var RootCmd = &cobra.Command{
		Use: "ecli",
	}
	RootCmd.AddCommand(cmds...)
	RootCmd.PersistentFlags().BoolVar(&param.Localnet, "localnet", false, "localnet peers")
	RootCmd.PersistentFlags().BoolVar(&param.Testnet, "testnet", false, "testnet peers")
	RootCmd.PersistentFlags().BoolVar(&param.Mainnet, "mainnet", false, "mainnet peers")
	RootCmd.Execute()
}
