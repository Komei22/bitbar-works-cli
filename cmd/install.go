package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install setup bitbar-works",
	Long:  `install setup bitbar-works`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("install called")

		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		createConfig(home + "/.bitbar-works.toml")
		createWorkHistory(home + "/.work_history")

		pluginDir, err := getBitbarPluginDir()
		if err != nil {
			return err
		}

		pluginSymlink := pluginDir + "/bitbar-works.200s"
		ex, _ := os.Executable()

		err = os.Symlink(ex, pluginSymlink)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func createWorkHistory(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(([]byte)(""))

	return err
}

func createConfig(filepath string) error {
	template := `
WORK_USER = "Your user id"
WORK_PASSWORD = "Your password"`

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(([]byte)(template))
	return nil
}

func getBitbarPluginDir() (string, error) {
	cmdName := "/usr/bin/defaults"
	cmdArgs := []string{"read", "com.matryer.BitBar", "pluginsDirectory"}

	out, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("unable to determine pluginsDirectory: %v, %s", err, string(out))
	}

	dir := strings.TrimRight(string(out), "\n")

	if !dirExists(dir) {
		return "", fmt.Errorf("unable to check if dir exists: %v, %s", err, dir)
	}
	return dir, nil
}

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return true
	}
	return false
}
