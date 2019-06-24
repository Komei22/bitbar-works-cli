package cmd

import (
	"github.com/Komei22/bitbar-works-go/atendance"
	"github.com/spf13/cobra"
)

// finishCmd represents the start command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish work command",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: check already start work

		_, err := atendance.StampAtendance(atendance.FinishWork)
		if err != nil {
			return err
		}
		printMacNotificationCenter("Finish work")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(finishCmd)
}
