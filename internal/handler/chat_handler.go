package handler

import (
	"net/http"
)

// @Summary Health check
// @Description Returns OK if server is running
// @Tags Health
// @Success 200
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
