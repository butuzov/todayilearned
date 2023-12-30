// Dealing with html quoting
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	jsonStr := `
	{
		"name": "foo",
		"dependencies": {
			"@serverless/cli": "^1.5.2",
			"yaml-ast-parser": "0.0.43"
		},
		"foobar": "barbar"
	}`

	tmp := map[string]any{}

	if err := json.Unmarshal([]byte(jsonStr), &tmp); err != nil {
		log.Fatal(err)
	}

	body, _ := json.MarshalIndent(tmp, "  ", "  ")

	// character to rune transition
	// https://golang.org/src/encoding/json/encode.go?s=6551:6569#L188
	bodyStr := string(body)
	bodyStr = strings.ReplaceAll(bodyStr, `\u003e`, ">")
	bodyStr = strings.ReplaceAll(bodyStr, `\u003c`, "<")
	bodyStr = strings.ReplaceAll(bodyStr, `\u0026`, "&")

	fmt.Printf("%s", bodyStr)
}
