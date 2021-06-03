package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
)

func main() {
	http.HandleFunc("/", CountingMidleware(home))
	http.HandleFunc("/articles", CountingMidleware(articles))
	http.HandleFunc("/metrics", metricsPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func articles(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("nothing")); err != nil {
		log.Println(err)
	}
}

func metricsPage(w http.ResponseWriter, r *http.Request) {
	metrics.WritePrometheus(w, true)
}

func CountingMidleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf(`requests_total{path=%q}`, r.URL.Path)
		metrics.GetOrCreateCounter(s).Inc()
		h.ServeHTTP(w, r)
	}
}
