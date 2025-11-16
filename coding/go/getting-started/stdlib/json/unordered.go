// Working with unordered maps json data
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type pkgjs struct {
	Order []string
	Map   map[string]interface{}
}

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

	var s pkgjs

	if err := json.Unmarshal([]byte(jsonStr), &s); err != nil {
		log.Fatal(err)
	}

	// adding new values
	key, value := "sss", struct {
		Data struct {
			Url     string `json:"url,omitempty"`
			Version string `json:"version,omitempty"`
			GitHash string `json:"git_commit,omitempty"`
		} `json:"data"`
	}{}

	value.Data.Url = "http://example.com"

	s.Replace(key, value)
	b1, _ := s.MarshalJSON()
	fmt.Printf("Result:\n %s\n", b1)

	b2, _ := json.Marshal(s)
	fmt.Printf("Result:\n %s\n", b2)
}

func (pj *pkgjs) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &pj.Map); err != nil {
		return err
	}

	//nolint: wsl
	index := make(map[string]int)
	for key := range pj.Map {
		pj.Order = append(pj.Order, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(b, esc)
	}

	sort.Slice(pj.Order, func(i, j int) bool {
		return index[pj.Order[i]] < index[pj.Order[j]]
	})

	return nil
}

func (pj pkgjs) Replace(key string, value interface{}) {
	pj.Map[key] = value
}

func (pj pkgjs) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "{\n")

	seen := map[string]struct{}{}

	// extra keys
	for i := range pj.Order {
		seen[pj.Order[i]] = struct{}{}
	}

	for k := range pj.Map {
		if _, ok := seen[k]; ok {
			continue
		}
		pj.Order = append(pj.Order, k)
	}

	// building json
	for i, key := range pj.Order {
		seen[key] = struct{}{}

		// keys
		km, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}

		fmt.Fprintf(&buf, "%s%s:", "  ", km)

		// values
		vm, err := json.MarshalIndent(pj.Map[key], "  ", "    ")
		if err != nil {
			return nil, err
		}

		buf.Write(vm)

		if i != len(pj.Map)-1 {
			fmt.Fprintf(&buf, ",")
		}
		fmt.Fprintf(&buf, "\n")
	}

	// extra added
	fmt.Fprintf(&buf, "}\n")

	return buf.Bytes(), nil
}
