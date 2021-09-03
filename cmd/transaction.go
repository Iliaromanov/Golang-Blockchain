package cmd

import (
	"github.com/Iliaromanov/Golang-Blockchain/database"
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func transactionCmd() *cobra.Command {
	var txCmd = &cobra.Command{
		Use:   "tx",
		Short: "Interact with transactions (add ...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
	}
	txCmd.AddCommand(transactionAddCmd())

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
			data, _ := cmd.Flags().GetString(flagData)

			// Create account structs for from Account and to Account
			fromAcc := database.NewAccount(from)
			toAcc := database.NewAccount(to)
			
			// Create transaction struct for the transaction
			tx := database.NewTransaction(fromAcc, toAcc, value, data)
			
			// Get current state
			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			defer state.Close()

			//Add transaction to mempool
			err = state.AddTx(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			//Persist transaction to disk
			_, err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("Transaction successfully persisted to the ledger.")
		},
	}


	// Create flags to allow user to specify transaction details within the CLI
	txAddCmd.Flags().String(flagFrom, "", "Account from which tokens are sent")
	txAddCmd.MarkFlagRequired(flagFrom)

	txAddCmd.Flags().String(flagTo, "", "Account that recieves the tokens")
	txAddCmd.MarkFlagRequired(flagTo)

	txAddCmd.Flags().Uint(flagValue, 0, "Amount of tokens being transferred")
	txAddCmd.MarkFlagRequired(flagValue)

	txAddCmd.Flags().String(flagData, "", "Possible Values: 'reward'")

	return txAddCmd
}

