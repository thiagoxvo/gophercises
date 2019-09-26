package cmd

import "github.com/spf13/cobra"

// RootCmd is our root command
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
