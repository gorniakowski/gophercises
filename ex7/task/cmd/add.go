package cmd

import (
	"gophercises/ex7/task/db"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task",
	Long:  `add adds a task to TODO`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.Add(task)
		if err != nil {
			panic(err)
		}
	},
}
