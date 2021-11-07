package controllers

import (
	"encoding/json"
	"expired-passport-checker/internal/api/responses"
	"net/http"
)

type InfrastructureController struct{}

// GetHealthCheck godoc
// @Summary Get service state
// @Description Health check for service
// @Tags infrastructure
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthCheckResponse
// @Router /health [get]
func (ic InfrastructureController) GetHealthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(responses.HealthCheckResponse{Success: true})
	}
}

// GetMetrics godoc
// @Summary Get service metrics
// @Description Return service metrics
// @Tags infrastructure
// @Accept  json
// @Produce  json
// @Success 200 {object} MetricsResponse
// @Router /metrics [get]
func (ic InfrastructureController) GetMetrics() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(responses.MetricsResponse{Value: 1})
	}
}
