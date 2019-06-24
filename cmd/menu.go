package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Komei22/bitbar-works-go/atendance"
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

		a := atendance.Atendance{}
		a.SetAtendanceInfo()

		switch checkWorkStatus(a) {
		case beforeWork:
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus)).Color("red")
			submenu.Line("Start work").Bash(ex).Params([]string{"start"}).Terminal(false)
			submenu.Line("Finish work").Bash(ex).Params([]string{"finish"}).Terminal(false)
		case working:
			startStatus := a.SwTime.Format("15:04")
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus))
			submenu.Line("Start work")
			submenu.Line("Finish work").Bash(ex).Params([]string{"finish"}).Terminal(false)
		case afterWork:
			startStatus := a.SwTime.Format("15:04")
			finishStatus := a.FwTime.Format("15:04")
			plugin.StatusLine(fmt.Sprintf("🏢 %s /🚶‍ %s", startStatus, finishStatus))
			submenu.Line("Start work")
			submenu.Line("Finish work")
		default:
		}

		plugin.Render()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}

func isStampAtendance(t time.Time) bool {
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

func checkWorkStatus(a atendance.Atendance) int {
	if !isStampAtendance(a.SwTime) && !isStampAtendance(a.FwTime) {
		return beforeWork
	} else if isStampAtendance(a.SwTime) && !isStampAtendance(a.FwTime) {
		return working
	} else {
		return afterWork
	}
}
