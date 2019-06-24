package cmd

import (
	"github.com/Komei22/bitbar-works-go/attendance"
	"github.com/spf13/cobra"
)

// clockoutCmd represents the start command
var clockoutCmd = &cobra.Command{
	Use:   "clockout",
	Short: "clockout command",
	RunE: func(cmd *cobra.Command, args []string) error {
		a := attendance.Attendance{}
		a.SetAttendanceInfo()
		if isClockInOut(a.FwTime) {
			return nil
		}

		_, err := attendance.RecordAttendance(attendance.ClockOut)
		if err != nil {
			printMacNotificationCenter("Fail clock out")
			return err
		}
		printMacNotificationCenter("Clock out")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clockoutCmd)
}
