package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/AraVraelHalt/API-Contract-Detector/services/collector"
	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
)

func main() {
	storage.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Contract Break Detector backend running")
	})

	mux.HandleFunc("/changes", func(w http.ResponseWriter, r *http.Request) {
		rows, err := storage.DB.Query("SELECT endpoint, change, created_at FROM changes")

		if err != nil {
			http.Error(w, "Error fetching changes", http.StatusInternalServerError)
			return
		}
		
		defer rows.Close()

		var results []map[string]interface{}

		for rows.Next() {
			var endpoint, change string
			var createdAt string

			err := rows.Scan(&endpoint, &change, &createdAt)

			if err != nil {
				continue
			}

			results = append(results, map[string]interface{} {
				"endpoint": endpoint,
				"change": change,
				"created_at": createdAt,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	handler := collector.CaptureRequest(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
