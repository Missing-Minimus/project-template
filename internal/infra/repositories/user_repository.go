package repositories

import (
	"errors"

	"github.com/Missing-Minimus/projects-template/internal/core/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindById(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindAll(params []string, values []string) ([]entities.User, error)
	Delete(id string) error
}

type PostgresUserRepository struct { // deixar o nome do banco antes do nome do repository pois aí a interface não fica
	// com o mesmo nome da struct
	db *gorm.DB // se estamos usando o ORM, o correto eh deixar ele como a conexao do nosso repositorio
	// pureDB *sql.DB  // mas se estivessemos usando puro, seria assim
}

func NewPostgresUserRepository(db *gorm.DB) UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (pgur *PostgresUserRepository) Create(u *entities.User) error {
	res := pgur.db.Create(&u)
	if res.RowsAffected == 0 {
		return errors.New("Could not create user")
	}

	return nil
}

func (pgur *PostgresUserRepository) FindById(id string) (*entities.User, error) {
	usr := entities.User{}
	res := pgur.db.First(&usr, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New("User was not found")
	}

	return &usr, nil
}

func (pgur *PostgresUserRepository) FindByEmail(email string) (*entities.User, error) {
	usr := entities.User{}
	res := pgur.db.First(&usr, "email = ?", email)
	if res.RowsAffected == 0 {
		return nil, errors.New("User was not found")
	}

	return &usr, nil
}

func (pgur *PostgresUserRepository) FindAll(params []string, values []string) ([]entities.User, error) {
	users := []entities.User{}
	res := pgur.db.Find(&users)
	if res.RowsAffected == 0 {
		return nil, errors.New("User was not found")
	} else if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (pgur *PostgresUserRepository) Delete(id string) error {
	res := pgur.db.Delete(id)
	if res.RowsAffected == 0 {
		return errors.New("User was not found")
	} else if res.Error != nil {
		return res.Error
	}

	return nil
}
