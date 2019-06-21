package atendance

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Action is atendance action
type Action int

const (
	// StartWork is starting work action
	StartWork = iota
	// FinishWork is ending work action
	FinishWork
)

// StampAtendance post and record atendance information.
func StampAtendance(action Action) (time.Time, error) {
	t := time.Now()
	err := postAtendance(action)
	if err != nil {
		return t, err
	}
	err = loggingWorkHistory(action, t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func postAtendance(action Action) error {
	if err := loadConfig(); err != nil {
		return err
	}

	postURL := "http://compweb01.gmo.local/cws/srwtimerec"
	values := url.Values{}
	values.Add("user_id", viper.GetString("WORK_USER"))
	values.Add("password", viper.GetString("WORK_PASSWORD"))
	if action == StartWork {
		values.Add("dakoku", "syussya")
	} else {
		values.Add("dakoku", "taisya")
	}

	_, err := http.PostForm(postURL, values)
	if err != nil {
		return err
	}

	return nil
}

func loggingWorkHistory(action Action, t time.Time) error {
	cmdstr := fmt.Sprintf(`echo '%d, %s' >> ~/.work_history`, action, t.Format(time.RFC3339))
	_, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		return err
	}
	return nil
}

func loadConfig() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	viper.SetConfigType("toml")
	viper.AddConfigPath(home)
	viper.SetConfigName(".bitbar-works")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
