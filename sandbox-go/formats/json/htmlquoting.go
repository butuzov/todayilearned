package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("testdata/package.json")
	defer f.Close()

	b, _ := io.ReadAll(f)
	tmp := map[string]interface{}{}

	if err := json.Unmarshal(b, &tmp); err != nil {
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
