package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/**/
var (
	tblCmd = &cobra.Command{ // This command will serve as the programs entry point.
		Use:   "tbl",
		Short: "The Blockchain Ledger CLI",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Describes version.",
		Run: func (cmd *cobra.Command, args []string) {
			fmt.Println("Version 0.1.0-beta")
		},
	}
)

func init() {
	tblCmd.AddCommand(versionCmd)
}

func main() {
	if err := tblCmd.Execute(); err != nil { // Execute root command and check for errors
		fmt.Fprintln(os.Stderr, err) // Print error to stderr instead of stdio
		os.Exit(1) // exit with code 1 (failed)
	}
}