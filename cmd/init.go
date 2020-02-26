package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize go-cli",
	Long:  `initialize go-cli with options`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cli v0.0.1-alpha")
	},
}
