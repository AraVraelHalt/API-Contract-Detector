package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AraVraelHalt/API-Contract-Detector/api/routes"
	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
)

func main() {
	// Initialize database
	storage.InitDB()

	// Setup routes + middleware
	handler := routes.SetupRoutes()

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
