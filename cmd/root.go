package cmd

import (
	"fmt"
	"log"
	"os"
	"task_manager/core"

	"github.com/spf13/cobra"
)

var taskManagerSrvc *core.TaskManagerService

func init() {
	taskManagerSrvc = core.NewTaskManagerService()
	rootCmd.AddCommand(saySomethingNice)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doCmd)
}

var saySomethingNice = &cobra.Command{
	Use:   "nice",
	Short: "Say something nice to kosi",
	Long:  "Please say somehting nice to kosi",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kosi you've got this")
		var output string

		for _, arg := range args {
			output = output + arg + " "
		}

		fmt.Println(output)
	},
}

var rootCmd = &cobra.Command{
	Use:   "taskManager",
	Short: "This is a CLI task manager application written in go",
	Long:  "This is a CLI task manager application written in golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello there")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		log.Println(err)
		os.Exit(1)
	}
}
