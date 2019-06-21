package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bitbar-works",
	Short: "Attendance cli",
	Run: func(cmd *cobra.Command, args []string) {
		// ex, _ := os.Executable()
		// stat, _ := os.Lstat(ex)
		// if (stat.Mode() & os.ModeSymlink) != 0 {
		menuCmd.Run(cmd, args)
		// } else {
		// 	cmd.Help()
		// }
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func printMacNotificationCenter(i interface{}) {
	cmdstr := fmt.Sprintf(`osascript -e "display notification \"%s\" with title \"Works\""`, i)
	fmt.Println(i)
	exec.Command("sh", "-c", cmdstr).Run()
}
