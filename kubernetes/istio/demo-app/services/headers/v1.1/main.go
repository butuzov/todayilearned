package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data := []string{}
		for k, v := range r.Header {
			data = append(data, fmt.Sprintf("%s: %s", k, strings.Join(v, "; ")))
		}

		sort.Strings(data)

		b, err := json.Marshal(data)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(b)
	})

	log.Fatal(http.ListenAndServe(":"+defaultEnvVal("PORT", "3000"), nil))
}

func defaultEnvVal(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
