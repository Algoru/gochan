package gochan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const ThreadsEndPoint = "https://a.4cdn.org/po/threads.json"

type Thread struct {
	Number       int64 `json:"no"`
	LastModified int64 `json:"last_modified"`
	Replies      uint8 `json:"replies"`
}

type Threads struct {
	Page    uint8    `json:"page"`
	Threads []Thread `json:"threads"`
}

func GetThreads() ([]Threads, error) {
	resp, err := http.Get(ThreadsEndPoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	threads := make([]Threads, 0)
	if err := json.Unmarshal(body, &threads); err != nil {
		return nil, err
	}

	return threads, nil
}
