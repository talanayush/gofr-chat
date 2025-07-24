package handler

import (
    "gofr.dev/pkg/gofr"
)

// @Summary Health check
// @Description Returns OK if server is running
// @Tags Health
// @Success 200
// @Router /health [get]
func HealthCheck(ctx *gofr.Context) (interface{}, error) {
    return "OK", nil
}