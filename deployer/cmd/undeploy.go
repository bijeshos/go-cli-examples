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
	Short:   "Undeploy artifacts (web, api or database)",
	Long:    `This command can be used together with web, api or database sub-commands to undeploy respective artifacts`,
}

var undeployWebCmd = &cobra.Command{
	Use:   "web",
	Short: "Undeploy web artifacts",
	Long:  `This command can be used to undeploy web artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer undeploy web' placeholder command")
	},
}

var undeployAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Uneploy API artifacts",
	Long:  `This command can be used to undeploy API artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer undeploy api' placeholder command")
	},
}

var undeployDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Undeploy database artifacts",
	Long:  `This command can be used to undeploy database artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer undeploy database' placeholder command")
	},
}
