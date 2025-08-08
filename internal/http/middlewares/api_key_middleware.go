package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoPolaski/go-rest-err/rest_err"
)

func ApiKeyMiddleware(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")

	if len(apiKey) == 0 {
		restErr := rest_err.NewUnauthorizedError("invalid credentials")
		json.NewEncoder(w).Encode(restErr)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// verify existence in DB
}
