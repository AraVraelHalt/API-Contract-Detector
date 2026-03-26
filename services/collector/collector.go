package collector

import (
	"fmt"
	"net/http"
)

func CaptureRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Captured request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
