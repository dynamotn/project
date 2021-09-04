package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gitlab.com/dynamo.foss/projekt/pkg"
	"gitlab.com/dynamo.foss/projekt/pkg/projekt/cli/folder"
	"gitlab.com/dynamo.foss/projekt/pkg/projekt/cli/template"
	"gitlab.com/dynamo.foss/projekt/pkg/projekt/cli/boilerplate"
)

var (
	cfgFile string
	RootCmd = &cobra.Command{
		Use:   "projekt",
		Short: "A smart command to work with your project folder",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.projekt/config.yaml)")

	RootCmd.AddCommand(folder.Cmd)
	RootCmd.AddCommand(template.Cmd)
	RootCmd.AddCommand(boilerplate.Cmd)
	RootCmd.AddCommand(pkg.VersionCmd)

	pkg.SetColorAndStyles(RootCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "/.projekt")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
