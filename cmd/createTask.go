package cmd

import "github.com/spf13/cobra"

var createTask = &cobra.Command{
	Use:   "task",
	Short: "create task",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
