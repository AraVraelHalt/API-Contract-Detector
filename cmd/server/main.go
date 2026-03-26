package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AraVraelHalt/API-Contract-Detector/services/collector"
	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
)

func main() {
	storage.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Contract Break Detector backend running")
	})

	handler := collector.CaptureRequest(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
