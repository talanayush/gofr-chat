package handler

import (
	"log"
	"net/http"
)

// @Summary Health check
// @Description Returns OK if server is running
// @Tags Health
// @Success 200
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Println("failed to write response:", err)
	}
}
