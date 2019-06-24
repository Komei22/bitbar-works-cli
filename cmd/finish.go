package cmd

import (
	"github.com/Komei22/bitbar-works-go/attendance"
	"github.com/spf13/cobra"
)

// finishCmd represents the start command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish work command",
	RunE: func(cmd *cobra.Command, args []string) error {
		a := attendance.Attendance{}
		a.SetAttendanceInfo()
		if isStampAttendance(a.FwTime) {
			return nil
		}

		_, err := attendance.StampAttendance(attendance.FinishWork)
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
