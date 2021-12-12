package testing

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update golden files")

type Person struct {
	Name string `json:"name,omitempty"`
}

func (p Person) Export() ([]byte, error) {
	return json.Marshal(p)
}

func TestPerson_Export(t *testing.T) {
	var (
		fixture = func(name string) (location string) {
			t.Helper()
			return filepath.Join("testdata", fmt.Sprintf("%s.json", name))
		}

		tests = map[string]struct {
			s         Person
			want      []byte
			assertion assert.ErrorAssertionFunc
		}{
			"Joe": {
				s:         Person{Name: "Joe"},
				want:      readFile(fixture("Joe")),
				assertion: assert.NoError,
			},
		}
	)
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.s.Export()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)

			if *update {
				_ = writeFile(fixture("Joe"), bytes.NewBuffer(got))
			}
		})
	}
}
