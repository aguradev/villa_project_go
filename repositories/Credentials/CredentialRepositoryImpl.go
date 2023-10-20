package repositories

import (
	"errors"
	"villa_go/entities"
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/UserResponse"

	"gorm.io/gorm"
)

type CredentialRepositoryImplement struct {
	Db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) CredentialRepository {
	return &CredentialRepositoryImplement{
		Db: db,
	}
}

func (account *CredentialRepositoryImplement) RegisterUserCredential(User entities.Users) (*UserResponse.RegisterResponse, error) {

	RegisterResponse := &UserResponse.RegisterResponse{}

	if RegisterException := account.Db.Create(&User).Error; RegisterException != nil {
		return nil, RegisterException
	}

	RegisterResponse.GetRegisterResponse(User)

	return RegisterResponse, nil
}

func (account *CredentialRepositoryImplement) GetRoleUserForRegister(role string) (entities.Roles, error) {
	var Roles entities.Roles

	RolesException := account.Db.Where("role = ?", role).First(&Roles)

	if RolesException.Error != nil {
		return entities.Roles{}, errors.New("Roles not found")
	}

	return Roles, nil
}

func (account *CredentialRepositoryImplement) CheckAuthCredential(request request.CredentialRequest) (bool, error) {
	return false, nil
}
