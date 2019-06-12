package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "menu function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("menu called")
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}
