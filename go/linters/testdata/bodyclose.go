package testdata

import (
	"io/ioutil"
	"net/http"
)

// PROJECT URL: https://github.com/timakin/bodyclose
func bodyclose() int {
	resp, err := http.Get("http://example.com/") // Wrong case
	if err != nil {
		// handle error
		return -1
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return -1
	}

	return len(body)
}
