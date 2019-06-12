package cmd

import (
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start work command",
	Run: func(cmd *cobra.Command, args []string) {
		printNotificationCenter("call start")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
