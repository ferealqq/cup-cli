package cmd

import (
	"fmt"
	"os"

	get "github.com/ferealqq/cup-util/cmd/get"
	"github.com/ferealqq/cup-util/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	token := os.Getenv("TOKEN")
	api.NewClient(token)

	cmdRoot.AddCommand(get.CMD)
}

var cmdRoot = &cobra.Command{
	Use:   "cup",
	Short: "Clickup cli",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
