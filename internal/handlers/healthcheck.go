package handlers

import "net/http"

type HealthCheckHandler struct{}

func NewHealthHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
