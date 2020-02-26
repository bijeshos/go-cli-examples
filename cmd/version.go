package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-cli",
	Long:  `Print the version number of go-cli sample application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cli v0.0.1-alpha")
	},
}
