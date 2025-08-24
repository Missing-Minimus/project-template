package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Missing-Minimus/projects-template/internal/infra/controllers/model/request"
	"github.com/Missing-Minimus/projects-template/internal/infra/controllers/model/response"
	"github.com/Missing-Minimus/projects-template/internal/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		UserService: us,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	usr := request.CreateUserRequest{}
	json.NewDecoder(r.Body).Decode(&usr)

	createdUser, err := uc.UserService.CreateUser(&usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			map[string]string{
				"erro": err.Error(),
			},
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		response.UserResponse{
			Email:    createdUser.Email,
			Password: createdUser.Password,
			Username: createdUser.Username,
		},
	)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")

	user, err := uc.UserService.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Usuário não encontrado"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erro": err.Error()})
		return
	}

	responses := []response.UserResponse{}
	for _, u := range users {
		responses = append(responses, response.UserResponse{
			Email:    u.Email,
			Password: u.Password,
			Username: u.Username,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses)
}
