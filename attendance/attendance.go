package attendance

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Action is attendance action
type Action int

const (
	// ClockIn is starting work action
	ClockIn = iota
	// ClockOut is ending work action
	ClockOut
)

// RecordAttendance post and record attendance information.
func RecordAttendance(action Action) (time.Time, error) {
	t := time.Now()
	err := postAttendance(action)
	if err != nil {
		return t, err
	}
	err = loggingWorkHistory(action, t)
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
	if action == ClockIn {
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
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	filepath := home + "/.work_history"
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	outStr := fmt.Sprintf("%d,%s\n", action, t.Format(time.RFC3339))
	fmt.Fprintf(f, outStr)

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
	if string(out) == "" {
		a.SwTime = time.Time{}
		a.FwTime = time.Time{}
		return nil
	}

	r := csv.NewReader(bytes.NewReader(out))
	record, _ := r.Read()

	// If the end line in .work_history is Finish work action, FwTime set the time of the end line and SwTime set second line of the end.
	// If the end line in .work_history is Start work action, SwTime set the time of the end line and FwTime set ZeroTime.
	if record[0] == strconv.Itoa(ClockOut) {
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
