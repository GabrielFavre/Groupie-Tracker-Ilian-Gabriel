package api

import (
	"encoding/json"
	"net/http"
	"time"
)

var Client = &http.Client{
	Timeout: 10 * time.Second,
}

func FetchData(url string, target interface{}) error {
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
