package repo

import (
	"io/ioutil"
	"net/http"

	log "gopkg.in/clog.v1"
)

type Notifier struct {
	targetURL string
}

func NewNotifier(targetURL string) Notifier {
	return Notifier{targetURL}
}

func (n Notifier) Notify() {
	client := http.Client{}
	request, err := http.NewRequest("POST", n.targetURL, nil)
	if err != nil {
		log.Info("Failed to create request: %+v", err)
	}
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("charset", "utf-8")
	resp, err := client.Do(request)
	if err != nil {
		log.Info("Failed to trigger action after pull request: %+v", err)
	}
	log.Info("Response status: %d", resp.StatusCode)
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Failed to get data from response: %+v", err)
	}
	log.Info("Response: %s", string(output))
}
