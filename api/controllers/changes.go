package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
)

// Handles breaking changes endpoints
type ChangesController struct{}

// Creates a new ChangesController
func NewChangesController() *ChangesController {
	return &ChangesController{}
}

// Returns breaking changes in JSON
func (cc *ChangesController) GetChanges(w http.ResponseWriter, r *http.Request) {
	changes, err := storage.GetChanges()
	if err != nil {
		http.Error(w, "Error fetching changes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(changes)
}
