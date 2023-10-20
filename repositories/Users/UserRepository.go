package repositories

import (
	"villa_go/entities"

	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	GetAllUsers() ([]entities.Users, error)
	GetUserById(uuid.UUID) (entities.Users, error)
	CreateUser() (bool, error)
	DeleteUser() (bool, error)
	UpdateUser(uuid.UUID) (bool, error)
}
