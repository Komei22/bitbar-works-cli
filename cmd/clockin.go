package cmd

import (
	"github.com/Komei22/bitbar-works-go/attendance"
	"github.com/spf13/cobra"
)

// clockinCmd represents the clockin command
var clockinCmd = &cobra.Command{
	Use:   "clockin",
	Short: "clockout command",
	RunE: func(cmd *cobra.Command, args []string) error {
		a := attendance.Attendance{}
		a.SetAttendanceInfo()
		if isStampAttendance(a.SwTime) {
			return nil
		}

		_, err := attendance.StampAttendance(attendance.StartWork)
		if err != nil {
			printMacNotificationCenter("Fail clock in")
			return err
		}
		printMacNotificationCenter("Clock in")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clockinCmd)
}
