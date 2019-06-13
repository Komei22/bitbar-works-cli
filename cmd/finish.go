package cmd

import (
	"github.com/spf13/cobra"
)

// finishCmd represents the start command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish work command",
	Run: func(cmd *cobra.Command, args []string) {
		printMacNotificationCenter("call finish")
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)
}
