package testdata

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// PROJECT URL:  https://github.com/sonatard/noctx
func noctx() {
	const url = "http://example.com"
	http.Get(url)

	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close() // OK
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(len(body))
}
