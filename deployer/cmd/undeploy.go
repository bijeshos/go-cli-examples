package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undeployCmd)

	undeployCmd.AddCommand(undeployWebCmd)
	undeployCmd.AddCommand(undeployAPICmd)
	undeployCmd.AddCommand(undeployDatabaseCmd)

}

var undeployCmd = &cobra.Command{
	Use:     "undeploy",
	Aliases: []string{"undep", "undepl"},
	Short:   "undeploy apps",
	Long:    `undeploy apps`,
}

var undeployWebCmd = &cobra.Command{
	Use:   "web",
	Short: "undeploy web",
	Long:  `undeploy web`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer undeploy web' placeholder command")
	},
}

var undeployAPICmd = &cobra.Command{
	Use:   "api",
	Short: "undeploy api",
	Long:  `undeploy api`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer undeploy api' placeholder command")
	},
}

var undeployDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "undeploy database",
	Long:  `undeploy database`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer undeploy database' placeholder command")
	},
}
