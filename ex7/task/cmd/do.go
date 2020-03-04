package cmd

import (
	"fmt"
	"gophercises/ex7/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark task as done",
	Long:  `removes task from the  TODO`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Faild to parse argument: %s\n", arg)
			} else {
				ids = append(ids, id)
			}

		}
		for _, id := range ids {
			err := db.Remove(id)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
