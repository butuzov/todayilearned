package testdata

import "encoding/json"

var _, _ = json.Marshal(nil) // want "Error return value of `encoding/json.Marshal` is not checked"
