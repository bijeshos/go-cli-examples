package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.AddCommand(deployWebCmd)
	deployCmd.AddCommand(deployAPICmd)
	deployCmd.AddCommand(deployDatabaseCmd)

}

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"dep", "depl"},
	Short:   "deploy apps",
	Long:    `deploy apps`,
}

var deployWebCmd = &cobra.Command{
	Use:   "web",
	Short: "deploy web",
	Long:  `deploy web`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer deploy web' placeholder command")
	},
}

var deployAPICmd = &cobra.Command{
	Use:   "api",
	Short: "deploy api",
	Long:  `deploy api`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer deploy api' placeholder command")
	},
}

var deployDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "deploy database",
	Long:  `deploy database`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'deployer deploy database' placeholder command")
	},
}
