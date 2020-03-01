package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "very simple task manager",
	Long:  `task is cli for managing TODOS`,
}

func Execute() {
	rootCmd.Execute()
}
