package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	tblCmd.AddCommand(BalancesCmd())
}

func Execute() {
	cobra.CheckErr(tblCmd.Execute())
}

func IncorrectUsageErr() error {
	return fmt.Errorf("Incorrect usage")
}