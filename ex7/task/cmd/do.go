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
		tasks := db.List()
		for _, id := range ids {
			if id < 1 || id > len(tasks) {
				fmt.Printf("Invalid Id: %v\n", id)
				continue
			}
			task := tasks[id-1]
			err := db.Remove(task.Key)
			if err != nil {
				fmt.Print("Something went wrong:", err)
				continue
			}
			fmt.Println("Removed task: ", task.Value)
		}
	},
}
