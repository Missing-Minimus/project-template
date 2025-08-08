package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

func BasicAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		username, password, ok := r.BasicAuth()
		if !ok || strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
			restErr := rest_err.NewUnauthorizedError("Basic auth header not found or invalid")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}
		next.ServeHTTP(w, r)
	}
}
