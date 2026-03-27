package collector

import (
	"fmt"
	"net/http"

	"github.com/AraVraelHalt/API-Contract-Detector/services/inference"
	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
	"github.com/AraVraelHalt/API-Contract-Detector/services/diff-engine"
)

func CaptureRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Captured request: %s %s\n", r.Method, r.URL.Path)

		buf := make([]byte, r.ContentLength)
		r.Body.Read(buf)

		oldSchema, err := storage.GetLastSchema(r.URL.Path)
		newSchema := inference.InferSchema(buf)

		if err == nil {
			diffengine.DetectBreakingChanges(r.URL.Path, oldSchema, newSchema)
		}

		storage.SaveSchema(r.URL.Path, newSchema)

		next.ServeHTTP(w, r)
	})
}
