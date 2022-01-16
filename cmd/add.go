package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Long:  "This is a command that adds a task to you task list, the task that will be added is the sentence specfied after the command name",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := taskManagerSrvc.AddTodo(task)
		if err != nil {
			fmt.Println(err)
		}
	},
}
