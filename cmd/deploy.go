package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"dep", "depl"},
	Short:   "deploy apps",
	Long:    `deploy apps using the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing deploy")
	},
}
