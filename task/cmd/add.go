package cmd

import (
	"fmt"
	"strings"

	"github.com/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		// db.CreateTask(task)
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

//init is a function which runs before the main (in the main.go) function
func init() {
	RootCmd.AddCommand(addCmd)
}
