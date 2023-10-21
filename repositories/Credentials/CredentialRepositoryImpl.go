package repositories

import (
	"errors"
	"villa_go/entities/models"
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/user_response"
	"villa_go/utils"

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

func (account *CredentialRepositoryImplement) RegisterUserCredential(User models.Users) (*UserResponse.RegisterResponse, error) {

	RegisterResponse := &UserResponse.RegisterResponse{}

	if RegisterException := account.Db.Create(&User).Error; RegisterException != nil {
		return nil, RegisterException
	}

	RegisterResponse.GetRegisterResponse(User)

	return RegisterResponse, nil
}

func (account *CredentialRepositoryImplement) GetRoleUserForRegister(role string) (models.Roles, error) {
	var Roles models.Roles

	RolesException := account.Db.Where("role = ?", role).First(&Roles)

	if RolesException.Error != nil {
		return models.Roles{}, errors.New("Roles not found")
	}

	return Roles, nil
}

func (account *CredentialRepositoryImplement) CheckAuthCredential(request request.AuthRequest) (*models.Users, bool, error) {

	var User models.Users

	// cara lain
	// account.Db.InnerJoins("Credential").InnerJoins("Credential.Roles").Find(&User, "username = ?", request.Username)

	if checkUserNameExist := account.Db.Preload("Credential").Preload("Credential.Roles").Joins("INNER JOIN credentials c ON c.id = users.credential_id").Joins("INNER JOIN roles r ON r.id = c.roles_id").Find(&User, "c.username = ?", request.Username); checkUserNameExist.Error != nil {

		if checkUserNameExist.Error == gorm.ErrRecordNotFound {
			return nil, false, errors.New("Authentication Failed")
		}

		return nil, false, checkUserNameExist.Error
	}

	if User.Credential == nil {
		return nil, false, errors.New("Authentication Failed")
	}

	isTrue, Exception := utils.ComparePasswordHashing(request.Password, User.Credential.Password)

	if !isTrue {
		return nil, false, errors.New("Password not correct")
	}

	if Exception != nil {
		return nil, false, Exception
	}

	return &User, true, nil
}
