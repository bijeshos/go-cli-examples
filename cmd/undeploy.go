package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undeployCmd)
}

var undeployCmd = &cobra.Command{
	Use:     "undeploy",
	Aliases: []string{"undep", "undepl"},
	Short:   "undeploy apps",
	Long:    `undeploy apps using the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing undeploy")
	},
}
