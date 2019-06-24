package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Komei22/bitbar-works-go/attendance"
	"github.com/johnmccabe/go-bitbar"
	"github.com/spf13/cobra"
)

const (
	beforeWork = iota
	working
	afterWork
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "menu function",
	Run: func(cmd *cobra.Command, args []string) {
		ex, _ := os.Executable()

		plugin := bitbar.New()
		submenu := plugin.NewSubMenu()

		startStatus := "Not yet"
		finishStatus := "Not yet"

		a := attendance.Attendance{}
		a.SetAttendanceInfo()

		switch checkWorkStatus(a) {
		case beforeWork:
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus)).Color("red")
			submenu.Line("Clock in").Bash(ex).Params([]string{"clockin"}).Terminal(false)
			submenu.Line("Clock out").Bash(ex).Params([]string{"clockout"}).Terminal(false)
		case working:
			startStatus := a.SwTime.Format("15:04")
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus))
			submenu.Line("Clock in")
			submenu.Line("Clock out").Bash(ex).Params([]string{"clockout"}).Terminal(false)
		case afterWork:
			startStatus := a.SwTime.Format("15:04")
			finishStatus := a.FwTime.Format("15:04")
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus))
			submenu.Line("Clock in")
			submenu.Line("Clock out")
		default:
		}

		plugin.Render()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}

func isStampAttendance(t time.Time) bool {
	if !isToday(t) || t.IsZero() {
		return false
	}
	return true
}

func isToday(t time.Time) bool {
	if t.Day() == time.Now().Day() {
		return true
	}
	return false
}

func checkWorkStatus(a attendance.Attendance) int {
	if !isStampAttendance(a.SwTime) && !isStampAttendance(a.FwTime) {
		return beforeWork
	} else if isStampAttendance(a.SwTime) && !isStampAttendance(a.FwTime) {
		return working
	} else {
		return afterWork
	}
}
