package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	svc := NewDashboardService(
		strings.TrimRight(
			defaultEnvVal("URL_BASE_BOOKS", "http://localhost:8081/api/v1"), "/")+"/",
		strings.TrimRight(
			defaultEnvVal("URL_BASE_AUTHORS", "http://localhost:8082/api/v1"), "/")+"/",
		logger,
	)

	r := mux.NewRouter()
	group := r.PathPrefix("/api/v1/dashboard").Subrouter().StrictSlash(true)
	group.Methods("GET").Path("/").Handler(httptransport.NewServer(getDashboardEndpoint(svc),
		decodeRequest,
		encodeResponse,
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
