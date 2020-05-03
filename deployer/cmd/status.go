package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.AddCommand(statusWebCmd)
	statusCmd.AddCommand(statusAPICmd)
	statusCmd.AddCommand(statusDatabaseCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of deployed artifacts",
	Long:  `status of deployed artifacts`,
}

var statusWebCmd = &cobra.Command{
	Use:   "web",
	Short: "status of web deployment",
	Long:  `status of web deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer status web' placeholder command")
	},
}

var statusAPICmd = &cobra.Command{
	Use:   "api",
	Short: "status api",
	Long:  `status api`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer status api' placeholder command")
	},
}

var statusDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "status of database deployment",
	Long:  `status of database deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer status database' placeholder command")
	},
}
