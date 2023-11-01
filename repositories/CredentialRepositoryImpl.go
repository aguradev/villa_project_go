package repositories

import (
	"errors"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CredentialRepository interface {
	GetRoleUserForRegister(role string) (entities.Roles, error)
	CheckEmailExists(email string) (bool, error)
	CheckAuthCredential(request.AuthRequest) (*entities.Users, bool, error)
	RegisterUserCredential(entities.Users) (*resources.RegisterResponse, error)
	UserLoginProfile(echo.Context) (*entities.Users, error)
}

type CredentialRepositoryImplement struct {
	Db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) CredentialRepository {
	return &CredentialRepositoryImplement{
		Db: db,
	}
}

func (account *CredentialRepositoryImplement) RegisterUserCredential(User entities.Users) (*resources.RegisterResponse, error) {

	RegisterResponse := &resources.RegisterResponse{}

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

func (account *CredentialRepositoryImplement) CheckAuthCredential(request request.AuthRequest) (*entities.Users, bool, error) {

	var User entities.Users

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

func (account *CredentialRepositoryImplement) UserLoginProfile(ctx echo.Context) (*entities.Users, error) {

	var GetUser entities.Users

	GetClaimUser, ErrClaims := utils.ClaimToken(ctx)

	if ErrClaims != nil {
		return nil, ErrClaims
	}

	Id := GetClaimUser.Id

	if QueryUserException := account.Db.First(&GetUser, "id = ?", Id); QueryUserException.Error != nil {
		return nil, errors.New("User not found")
	}

	return &GetUser, nil

}

func (account *CredentialRepositoryImplement) CheckEmailExists(email string) (bool, error) {

	var User entities.Users

	EmailExistsErr := account.Db.First(&User, "email = ?", email)

	if EmailExistsErr.RowsAffected > 0 {
		return true, errors.New("Email already exists")
	}

	return false, nil

}
