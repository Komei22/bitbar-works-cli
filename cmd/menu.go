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
			submenu.Line("Clock in").Bash(ex).Params([]string{"clockin"}).Terminal(false).Refresh()
			submenu.Line("Clock out").Bash(ex).Params([]string{"clockout"}).Terminal(false).Refresh()
		case working:
			startStatus := a.SwTime.Format("15:04")
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus))
			submenu.Line("Already clock in")
			submenu.Line("Clock out").Bash(ex).Params([]string{"clockout"}).Terminal(false).Refresh()
		case afterWork:
			startStatus := a.SwTime.Format("15:04")
			finishStatus := a.FwTime.Format("15:04")
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus))
			submenu.Line("Already clock in")
			submenu.Line("Already clock out")
		default:
		}

		plugin.Render()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}

func isClockInOut(t time.Time) bool {
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
	if !isClockInOut(a.SwTime) && !isClockInOut(a.FwTime) {
		return beforeWork
	} else if isClockInOut(a.SwTime) && !isClockInOut(a.FwTime) {
		return working
	} else {
		return afterWork
	}
}
