package services

import (
	"errors"
	"time"

	"github.com/Missing-Minimus/projects-template/internal/core/entities"
	"github.com/Missing-Minimus/projects-template/internal/infra/controllers/model/request"
	"github.com/Missing-Minimus/projects-template/internal/infra/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repositories repositories.UserRepository
}

func NewUserService(repositories repositories.UserRepository) *UserService {
	return &UserService{repositories: repositories}
}

func (us *UserService) CreateUser(req *request.CreateUserRequest) (*entities.User, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email e senha são obrigatórios")
	}

	existing, _ := us.repositories.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email já cadastrado")
	}

	user := &entities.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password, // Em prod alterar para hash
		UpdatedAt: time.Now(),
	}

	err := us.repositories.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUser(id string) (*entities.User, error) {
	return us.repositories.FindById(id)
}

func (us *UserService) ListUsers(params []string, values []string) ([]entities.User, error) {
	return us.repositories.FindAll(params, values)
}

func (us *UserService) DeleteUser(id string) error {
	return us.repositories.Delete(id)
}
