package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Iliaromanov/Golang-Blockchain/database"
	"fmt"
	"os"
)

// Loads latest db state and prints it to stdout
func BalancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list ...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
	}
	balancesCmd.AddCommand(balancesListCmd)

	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all balances.",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer state.Close()

		fmt.Println("Account balances:")
		fmt.Println("-----------------")
		for account, balance := range state.Balances {
			fmt.Printf("%v: %v\n", account, balance)
		}
	},
}