package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:   "get",
	Short: "get objects from the api",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get objects from the api")
		// Do Stuff Here
	},
}

func init() {
	CMD.AddCommand(getTask)
}
