package collector

import (
	"fmt"
	"net/http"

	"github.com/AraVraelHalt/API-Contract-Detector/services/inference"
	"github.com/AraVraelHalt/API-Contract-Detector/services/storage"
)

func CaptureRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Captured request: %s %s\n", r.Method, r.URL.Path)

		// TODO: handle streaming and errors properly
		buf := make([]byte, r.ContentLength)
		r.Body.Read(buf)
		schema := inference.InferSchema(buf)
		storage.SaveSchema(r.URL.Path, schema)

		fmt.Println("Schema:", schema)

		next.ServeHTTP(w, r)
	})
}
