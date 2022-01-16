package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "This command lists all the tasks on our task manager",
	Long:  "This command lists all the tasks stored on the database for our task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all tasks on our task manager")
		err := taskManagerSrvc.ListTodos()
		if err != nil {
			fmt.Println(err)
		}
	},
}
