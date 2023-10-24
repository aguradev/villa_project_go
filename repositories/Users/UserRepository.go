package repositories

import (
	"villa_go/entities/models"

	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	GetUserById(uuid.UUID) (models.Users, error)
	CreateUser() (bool, error)
	DeleteUser() (bool, error)
	UpdateUser(uuid.UUID) (bool, error)
}
