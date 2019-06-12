package atendance

import (
	"fmt"
	"os/exec"
	"time"
)

// StartWork post and record starting work time
func StartWork() error {
	// postURL := "http://compweb01.gmo.local/cws/srwtimerec"
	// values := url.Values{}
	// values.Add("dakoku", "syussya")
	// values.Add("user_id", os.Getenv("WORK_USER"))
	// values.Add("password", os.Getenv("WORK_PASSWORD"))

	// _, err := http.PostForm(postURL, values)
	// if err != nil {
	// 	return err
	// }
	err := loggingWorkHistory("start")
	if err != nil {
		return err
	}
	return nil
}

// EndWork post and record ending work time
func EndWork() error {
	// postURL := "http://compweb01.gmo.local/cws/srwtimerec"
	// values := url.Values{}
	// values.Add("dakoku", "taisya")
	// values.Add("user_id", os.Getenv("WORK_USER"))
	// values.Add("password", os.Getenv("WORK_PASSWORD"))

	// _, err := http.PostForm(postURL, values)
	// if err != nil {
	// 	return err
	// }
	err := loggingWorkHistory("end")
	if err != nil {
		return err
	}
	return nil
}

func loggingWorkHistory(action string) error {
	t := time.Now().Format(time.RFC3339)
	cmdstr := fmt.Sprintf(`echo '%s, %s' >> ~/.work_history`, action, t)
	_, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		return err
	}
	return nil
}
