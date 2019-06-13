package atendance

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
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
	loggingWorkHistory(action, t)
	return t, nil
}

func postAtendance(action Action) error {
	postURL := "http://compweb01.gmo.local/cws/srwtimerec"
	values := url.Values{}
	values.Add("user_id", os.Getenv("WORK_USER"))
	values.Add("password", os.Getenv("WORK_PASSWORD"))
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
