package cmd

import (
	"github.com/spf13/cobra"
)

// endCmd represents the start command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End work command",
	Run: func(cmd *cobra.Command, args []string) {
		printMacNotificationCenter("call end")
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
