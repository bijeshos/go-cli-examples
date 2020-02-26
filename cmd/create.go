package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"cr", "crt"},
	Short:   "create new deployments",
	Long:    `create new deployments using the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing create")
	},
}
