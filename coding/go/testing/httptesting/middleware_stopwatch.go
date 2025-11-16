package httptesting

import (
	"log"
	"net/http"
	"time"
)

type stopwatch struct {
	req   *http.Request
	start time.Time
}

func start(req *http.Request) *stopwatch {
	return &stopwatch{
		req:   req,
		start: time.Now(),
	}
}

func (s *stopwatch) stop() {
	log.Printf("[%s] %q %v\n", s.req.Method, s.req.URL.String(), time.Now().Sub(s.start))
}

func StopWatch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer start(r).stop()

		next.ServeHTTP(w, r)
	})
}
