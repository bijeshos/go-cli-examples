package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "go-cli",
		Short: "A sample CLI application",
		Long:  `A sample CLI application written in Go`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Bijesh O S")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "MIT License")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Bijesh O S <@bijeshos>")
	viper.SetDefault("license", "MIT")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Print the version number of go-cli",
	Long:  `Print the version number of go-cli sample application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing add")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print the version number of go-cli",
	Long:  `Print the version number of go-cli sample application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing init")
	},
}
