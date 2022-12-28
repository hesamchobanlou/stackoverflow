package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rebootCmd)
}

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot the charger",
	Long:  "reboot the charger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rebooting charger")
	},
}
