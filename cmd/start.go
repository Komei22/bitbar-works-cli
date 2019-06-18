package cmd

import (
	// "github.com/Komei22/bitbar-works-go/atendance"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start work command",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: check already start work

		// start work request and logging
		// _, err := atendance.StampAtendance(atendance.StartWork)
		// if err != nil {
		// 	printMacNotificationCenter("Fail start work")
		// 	return
		// }
		printMacNotificationCenter("Start work")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
