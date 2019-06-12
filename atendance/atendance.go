package atendance

import(
	"os"

	"net/http"
	"net/url"
)

func PostEndWork() error {
	postURL := "http://compweb01.gmo.local/cws/srwtimerec"
	values := url.Values{}
	values.Add("dakoku", "taisya")
	values.Add("user_id", os.Getenv("WORK_USER"))
	values.Add("password", os.Getenv("WORK_PASSWORD"))

	_, err := http.PostForm(postURL, values)
	if err != nil {
		return err
	}
	return nil
}


