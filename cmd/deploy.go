package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.AddCommand(deployWebCmd)

}

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"dep", "depl"},
	Short:   "deploy apps",
	Long:    `deploy apps`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing deploy")
	},
}

var deployWebCmd = &cobra.Command{
	Use:   "web",
	Short: "deploy web",
	Long:  `deploy web`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing deploy:web")
	},
}
