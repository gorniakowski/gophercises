package cmd

import (
	"fmt"
	"gophercises/ex7/task/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "print all tasks",
	Long:  `This prints all the tasks in the TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(db.List())
	},
}
