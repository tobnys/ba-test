package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes the program",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No arguments specified, exiting...")
			fmt.Println("USAGE: bat exec [TARGET URL]")
		} else {
			fmt.Printf("EXECUTING WITH ARGS: %s...", args)
		}
	},
}
