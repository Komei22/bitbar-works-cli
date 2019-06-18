package cmd

import (
	"github.com/Komei22/bitbar-works-go/atendance"
	"github.com/spf13/cobra"
)

// finishCmd represents the start command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish work command",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := atendance.StampAtendance(atendance.FinishWork)
		if err != nil {
			printMacNotificationCenter(err)
			return
		}
		printMacNotificationCenter("call finish")
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)
}
