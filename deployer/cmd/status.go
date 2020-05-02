package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.AddCommand(statusWebCmd)
	statusCmd.AddCommand(statusDatabaseCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of deployed artifacts",
	Long:  `status of deployed artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing status command")
	},
}

var statusWebCmd = &cobra.Command{
	Use:   "web",
	Short: "status of web deployment",
	Long:  `status of web deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing status:web")
	},
}

var statusDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "status of database deployment",
	Long:  `status of database deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing status:database")
	},
}
