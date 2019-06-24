package cmd

import (
	// "github.com/Komei22/bitbar-works-go/atendance"
	"github.com/Komei22/bitbar-works-go/atendance"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start work command",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: check already start work

		// start work request and logging
		_, err := atendance.StampAtendance(atendance.StartWork)
		if err != nil {
			printMacNotificationCenter("Fail start work")
			return err
		}
		printMacNotificationCenter("Start work")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
