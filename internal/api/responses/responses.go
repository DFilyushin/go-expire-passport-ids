package responses

//HealthCheckResponse response for health check
type HealthCheckResponse struct {
	Success bool `json:"success"`
}

//MetricsResponse response from metrics
type MetricsResponse struct {
	Value int `json:"value"`
}

//PassportCheckerResponse response from passport checker
type PassportCheckerResponse struct {
	PassportSeries string `json:"PassportSeries"`
	PassportNumber string `json:"PassportNumber"`
	Result         bool   `json:"Result"`
}
