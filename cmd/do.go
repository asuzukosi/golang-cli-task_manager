package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "This command will do a task on our task manager",
	Long:  "This command will do a task on our task manager and store it in our database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the do command")
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument: " + arg)
				log.Fatal(err)
			}
			err = taskManagerSrvc.DoTodo(id)
			if err != nil {
				fmt.Println(err)
			}

		}

	},
}
