package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	}
	`
	tmp := map[string]interface{}{}

	if err := json.Unmarshal([]byte(jsonStr), &tmp); err != nil {
		log.Fatal(err)
	}

	// hiding all fields in other field!
	var out struct {
		Name string          `json:"name"`
		Raw  json.RawMessage `json:"data,omitempty"`
	}

	out.Name = tmp["name"].(string)
	out.Raw = json.RawMessage(jsonStr)

	body, _ := json.MarshalIndent(out, "  ", "  ")

	fmt.Printf("%s", string(body))
}
