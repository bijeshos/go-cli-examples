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
	Short: "Check status of deployed artifacts (web, api or database)",
	Long:  `This command can be used together with web, api or database sub-commands to check status of respective deployments`,
}

var statusWebCmd = &cobra.Command{
	Use:   "web",
	Short: "Check status of web deployment",
	Long:  `This command can be used to status of web deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer status web' placeholder command")
	},
}

var statusAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Check status of API deployment",
	Long:  `This command can be used to status of API deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer status api' placeholder command")
	},
}

var statusDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Check status of database deployment",
	Long:  `This command can be used to status of database deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer status database' placeholder command")
	},
}
