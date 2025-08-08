package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Missing-Minimus/projects-template/internal/infra/thirdparty/logger"
)

func InitRoutes(s *http.ServeMux) {
	logger.Info("Setting up routes")

	s.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})
}
