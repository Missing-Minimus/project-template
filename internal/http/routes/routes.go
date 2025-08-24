package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Missing-Minimus/projects-template/internal/http/middlewares"
	"github.com/Missing-Minimus/projects-template/internal/infra/controllers"
	"github.com/Missing-Minimus/projects-template/internal/infra/repositories"
	"github.com/Missing-Minimus/projects-template/internal/infra/thirdparty/database"
	"github.com/Missing-Minimus/projects-template/internal/infra/thirdparty/logger"
	"github.com/Missing-Minimus/projects-template/internal/services"
)

func InitRoutes(s *http.ServeMux) {
	logger.Info("Setting up routes")
	
	UserController := controllers.NewUserController(
		services.NewUserService(
			repositories.NewPostgresUserRepository(
				database.NewGormAdapter().Connect(),
			),
		),
	)

	s.Handle("POST /user", HandlerChain(
		UserController.CreateUser,
		middlewares.ApiKeyMiddleware,
		middlewares.LogMiddleware,
	))

	s.Handle("GET /user/{id}", HandlerChain(
		UserController.,
		middlewares.ApiKeyMiddleware,
		middlewares.LogMiddleware,
	))

	s.Handle("DELETE /user/{id}", HandlerChain(
		UserController.,
		middlewares.ApiKeyMiddleware,
		middlewares.LogMiddleware,
	))

	s.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})
}
