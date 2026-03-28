package controllers

import (
	"fmt"
	"net/http"
)

// Handles health-check endpoints
type HealthController struct{}

// Returns a new HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// Responds with simple message to indicate the server is running
func (hc *HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "API Contract Break Detector backend running")
}
