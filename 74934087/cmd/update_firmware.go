package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateFirmwareCmd)
}

var updateFirmwareCmd = &cobra.Command{
	Use:   "update_firmware",
	Short: "Updates the firmware",
	Long:  "Update the firmware",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updating firmware")
	},
}
