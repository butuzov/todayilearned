package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, _ := os.Open("testdata/package.json")
	defer f.Close()

	b, _ := io.ReadAll(f)
	tmp := map[string]interface{}{}

	if err := json.Unmarshal(b, &tmp); err != nil {
		log.Fatal(err)
	}

	// hiding all fields in other field!
	var out struct {
		Name string          `json:"name"`
		Raw  json.RawMessage `json:"data,omitempty"`
	}

	out.Name = tmp["name"].(string)
	out.Raw = json.RawMessage(b)

	body, _ := json.MarshalIndent(out, "  ", "  ")

	fmt.Printf("%s", string(body))
}
