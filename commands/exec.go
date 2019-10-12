package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tobnys/ba-test/worker"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes the program",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("USAGE: bat exec [TARGET URL]")
		} else {
			fmt.Printf("EXECUTING WITH ARGS: %s... \n", args)
			worker := worker.InitWorker()
			fmt.Println(worker)
		}
	},
}
