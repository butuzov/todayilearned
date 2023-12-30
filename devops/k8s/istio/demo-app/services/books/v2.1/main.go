package main

import (
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := NewAuthorsService()

	r := mux.NewRouter()
	group := r.PathPrefix("/api/v1/books").Subrouter().StrictSlash(true)

	// GetAllAuthors
	group.Methods("GET").Path("/").Handler(httptransport.NewServer(getAllBooksEndpoint(svc),
		decodeGetAllBooks,
		encodeGetAllBooks,
	))

	// GetAuthor
	group.Methods("GET").Path("/{id}/").Handler(httptransport.NewServer(getAuthorEndpoint(svc),
		decodeGetBook,
		encodeGetBook,
	))

	hc := r.PathPrefix("/healthcheck").Subrouter().StrictSlash(true)
	hc.Methods("GET").Path("/readiness").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	hc.Methods("GET").Path("/liveness").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(":"+defaultEnvVal("PORT", "3000"), r))
}

func defaultEnvVal(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
