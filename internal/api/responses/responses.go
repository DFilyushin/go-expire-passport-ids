package responses

type HealthCheckResponse struct {
	Success bool `json:"success"`
}

type MetricsResponse struct {
	Value int `json:"value"`
}

type PassportCheckerResponse struct {
	PassportSeries string `json:"PassportSeries"`
	PassportNumber string `json:"PassportNumber"`
	Result         bool   `json:"Result"`
}
