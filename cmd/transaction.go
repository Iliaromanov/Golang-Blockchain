package cmd

import (
	"github.com/Iliaromanov/Golang-Blockchain/database"
	"github.com/spf13/cobra"
	"fmt"
	"os"
)


func transactionCmd() *cobra.Command {
	var txCmd = &cobra.Command{
		Use:   "tx",
		Short: "Interact with transactions (add ...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
	}

	return txCmd
}


// Allows users to add transactions to state using CLI
func transactionAddCmd() * cobra.Command {
	var txAddCmd = & cobra.Command{
		Use:   "add",
		Short: "Adds new transaction to database",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve transaction details from cmd flags
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)

			// Create account structs for from Account and to Account
			fromAcc := database.NewAccount(from)
			toAcc := database.NewAccount(to)
			
			// Create transaction struct for the transaction
			tx := database.NewTransaction(fromAcc, toAcc, value, "")
			
			// Get current state
			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			defer state.Close()

			//Add transaction to mempool
			err = state.Add(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	return txAddCmd
}

