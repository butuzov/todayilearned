// This is an example of handling default values (zero values) with yaml

package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type count uint64

func (c *count) IsZero() bool {
	return c == nil
}

func main() {
	tests := map[string]string{
		"with zero": "  count: 0",
		"with none": "  some: none",
		"with nums": "  count: 1",
	}

	var keys []string
	for k, _ := range tests {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		var (
			lines_of_yaml = []string{
				"main:",
				tests[key],
				`  found: string`,
			}

			in  = []byte(strings.Join(lines_of_yaml, "\n"))
			out struct {
				Main struct {
					Num *count `yaml:"count,omitempty"`
					// Test int    `yaml:"testnum,omitempty"`
				} `yaml:"main"`
			}
		)

		fmt.Println(strings.Repeat("\n", 2))
		fmt.Println(strings.Repeat("*", 4) + " " + key + " " + strings.Repeat("*", 65))

		// with 0 -> we have 0
		fmt.Println("In....")
		fmt.Println(strings.Repeat("=", 80))
		if err := yaml.Unmarshal(in, &out); err == nil {
			spew.Dump(out)
			// spew.Println(out)
		} else {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Println("Out....")
		fmt.Println(strings.Repeat("=", 80))

		// out.Main.Test = 10
		if b, err := yaml.Marshal(out); err == nil {
			fmt.Println(string(b))
		}

		fmt.Println(strings.Repeat("\n", 2))
	}
}
