package cmd

import (
	"os"

	"github.com/johnmccabe/go-bitbar"
	"github.com/spf13/cobra"
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "menu function",
	Run: func(cmd *cobra.Command, args []string) {
		ex, _ := os.Executable()

		plugin := bitbar.New()
		plugin.StatusLine("🏢🚶‍")
		submenu := plugin.NewSubMenu()
		submenu.Line("Start work").Bash(ex).Params([]string{"start"}).Terminal(false)
		submenu.Line("Finish work").Bash(ex).Params([]string{"finish"}).Terminal(false)

		plugin.Render()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}
