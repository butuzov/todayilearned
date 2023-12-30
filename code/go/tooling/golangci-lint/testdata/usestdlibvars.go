package testdata

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON marshals v to JSON, automatically escaping HTML,
// setting the Content-Type header as "application/json; charset=utf-8",
// sends an HTTP response header with the provided statusCode and
// writes the marshaled v as bytes to the connection as part of an HTTP reply.
func JSON(w http.ResponseWriter, statusCode int, v any) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
