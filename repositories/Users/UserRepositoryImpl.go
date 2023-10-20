package repositories

import (
	"villa_go/entities"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepositoryImplement struct {
	Db *gorm.DB
}

func CreateNewUserRepositoryImplment(db *gorm.DB) UserRepository {
	return &UserRepositoryImplement{
		Db: db,
	}
}

func (User *UserRepositoryImplement) GetAllUsers() ([]entities.Users, error) {
	return nil, nil
}

func (User *UserRepositoryImplement) GetUserById(id uuid.UUID) (entities.Users, error) {
	return entities.Users{}, nil
}

func (User *UserRepositoryImplement) CreateUser() (bool, error) {
	return false, nil
}

func (User *UserRepositoryImplement) DeleteUser() (bool, error) {
	return false, nil
}

func (User *UserRepositoryImplement) UpdateUser(id uuid.UUID) (bool, error) {
	return false, nil
}
