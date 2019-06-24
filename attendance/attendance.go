package attendance

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Action is attendance action
type Action int

const (
	// StartWork is starting work action
	StartWork = iota
	// FinishWork is ending work action
	FinishWork
)

// StampAttendance post and record attendance information.
func StampAttendance(action Action) (time.Time, error) {
	t := time.Now()
	// err := postAttendance(action)
	// if err != nil {
	// 	return t, err
	// }
	err := loggingWorkHistory(action, t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func postAttendance(action Action) error {
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

func initWorkHistory() error {
	err := exec.Command("sh", "-c", "touch ~/.work_history").Run()
	return err
}

func loggingWorkHistory(action Action, t time.Time) error {
	cmdstr := fmt.Sprintf(`echo '%d,%s' >> ~/.work_history`, action, t.Format(time.RFC3339))
	err := exec.Command("sh", "-c", cmdstr).Run()
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

// Attendance have attendance information
type Attendance struct {
	SwTime time.Time
	FwTime time.Time
}

// SetAttendanceInfo set start and finish work time
func (a *Attendance) SetAttendanceInfo() error {
	cmdstr := `cat ~/.work_history | tail -1`
	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		return err
	}

	// If work history is nothing, SwTime and FwTime set ZeroDay.
	if string(out) == "\n" {
		a.SwTime = time.Time{}
		a.FwTime = time.Time{}
		return nil
	}

	r := csv.NewReader(bytes.NewReader(out))
	record, _ := r.Read()

	// If the end line in .work_history is Finish work action, FwTime set the time of the end line and SwTime set second line of the end.
	// If the end line in .work_history is Start work action, SwTime set the time of the end line and FwTime set ZeroTime.
	if record[0] == strconv.Itoa(FinishWork) {
		a.FwTime, err = time.Parse(time.RFC3339, record[1])
		cmdstr = `cat ~/.work_history | tail -2 | head -1`
		out, err = exec.Command("sh", "-c", cmdstr).Output()
		if err != nil {
			return err
		}
		r := csv.NewReader(bytes.NewReader(out))
		record, _ := r.Read()
		a.SwTime, err = time.Parse(time.RFC3339, record[1])
		if err != nil {
			return err
		}
	} else {
		a.SwTime, err = time.Parse(time.RFC3339, record[1])
		a.FwTime = time.Time{}
	}

	return nil
}
