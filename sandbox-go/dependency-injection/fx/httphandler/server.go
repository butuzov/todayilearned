package httphandler

import (
	"net/http"

	"go.uber.org/zap"
)

// Handler for http requests
type Handler struct {
	mux *http.ServeMux
	log *zap.SugaredLogger
}

// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	h.log.Warn("request in")
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
	h.log.Warn("request out")
}

func New(s *http.ServeMux, l *zap.SugaredLogger) *Handler {
	h := Handler{s, l}
	h.registerRoutes()
	return &h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
