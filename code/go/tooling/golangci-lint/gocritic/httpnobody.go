package testdata_gocritic

import "net/http"

// Example for the rule: hexLiteral

func httpNoBody() {
	url := "https://google.com"
	http.NewRequest("GET", url, nil)
}
