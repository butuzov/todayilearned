package main

import (
	"log"
	"net/http"
	"os"
)

const AppPort = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requested: %s", r.RequestURI)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		log.Printf("Pong")
		w.Write([]byte("pong"))
	})

	port := port()
	log.Printf("Preparing to run on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func port() string {
	port := os.Getenv("APP_PORT")
	if port != "" {
		return port
	}
	return AppPort
}
