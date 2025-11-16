package inliner

import (
	"encoding/json"
	"sort"
)

type Struct struct {
	Type string         `json:"__type"`
	Meta map[string]any // inlined field
}

func NewStruct(key string, meta map[string]any) *Struct {
	return &Struct{
		Type: key,
		Meta: meta,
	}
}

func (a Struct) MarshalJSON() ([]byte, error) {
	type Aspect_ Struct
	b, _ := json.Marshal(Aspect_(a))

	var m map[string]json.RawMessage
	_ = json.Unmarshal(b, &m)

	var keys []string
	for k := range a.Meta {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var key string
	for _, key = range keys {
		if _, ok := m[key]; ok {
			continue
		}
		b, _ = json.Marshal(a.Meta[key])
		m[key] = b
	}

	return json.Marshal(m)
}
