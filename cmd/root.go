package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// var cfgFile string

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
		fmt.Println(err)
		os.Exit(1)
	}
}

func printMacNotificationCenter(i interface{}) {
	cmdstr := fmt.Sprintf(`osascript -e "display notification \"%s\" with title \"Works\""`, i)
	fmt.Println(i)
	exec.Command("sh", "-c", cmdstr).Run()
}

// func init() {
// 	cobra.OnInitialize(initConfig)

// Here you will define your flags and configuration settings.
// Cobra supports persistent flags, which, if defined here,
// will be global for your application.

// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bitbar-works.yaml)")

// Cobra also supports local flags, which will only run
// when this action is called directly.
// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name ".bitbar-works" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".bitbar-works")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
