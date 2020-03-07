package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.AddCommand(getServicesCmd)
	getCmd.AddCommand(getJobsCmd)
}

var getCmd = &cobra.Command{
	Use:   "status",
	Short: "status of deployed artifacts",
	Long:  `status of deployed artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing status command")
	},
}

var getServicesCmd = &cobra.Command{
	Use:   "web",
	Short: "status of web deployment",
	Long:  `status of web deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing status:web")
	},
}

var getJobsCmd = &cobra.Command{
	Use:   "database",
	Short: "status of database deployment",
	Long:  `status of database deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get:jobs")
	},
}
