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
	Long:  `status of deployed artifacts details`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get")
	},
}

var getServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "get services details",
	Long:  `get services details and configs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get:services")
	},
}

var getJobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "get job details",
	Long:  `get job details and configs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing get:jobs")
	},
}
