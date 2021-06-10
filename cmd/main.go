package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	// This command will serve as the programs entry point.
	var tblCmd = &cobra.Command{
		Use:   "tbl",
		Short: "The Blockchain Ledger CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	if err := tblCmd.Execute(); err != nil { // Execute root command and check for errors
		fmt.Fprintln(os.Stderr, err) // Print error to stderr instead of stdio
		os.Exit(1) // exit with code 1 (failed)
	}
}