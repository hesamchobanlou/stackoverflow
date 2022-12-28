package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "charger",
	Short: "Some description",
	Long:  "Some long description",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// read config files
}
