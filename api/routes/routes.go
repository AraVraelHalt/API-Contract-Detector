package routes

import (
	"net/http"

	"github.com/AraVraelHalt/API-Contract-Detector/api/controllers"
	"github.com/AraVraelHalt/API-Contract-Detector/services/collector"
	"github.com/rs/cors"
)

// Initializes all HTTP routes and middleware
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Initialize controllers
	healthController := controllers.NewHealthController()
	changesController := controllers.NewChangesController()

	// Register routes
	mux.HandleFunc("/health", healthController.HealthCheck)
	mux.HandleFunc("/changes", changesController.GetChanges)

	// Wrap with middleware (CORS + request collector)
	handler := cors.Default().Handler(collector.CaptureRequest(mux))

	return handler
}
