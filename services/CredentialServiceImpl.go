package services

import (
	"errors"
	"villa_go/exceptions"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CredentialService interface {
	RegisterCredential(request.RegisterRequest) (*response.RegisterResponse, error)
	AuthUser(echo.Context, request.AuthRequest) (*response.AuthToken, []exceptions.ValidationMessage, error)
}

type CredentialServiceImpl struct {
	CredentialRepository repositories.CredentialRepository
	Validator            *validator.Validate
	TranslatorValidation ut.Translator
}

func CreateCredentialServiceImplement(Credential repositories.CredentialRepository, validate *validator.Validate, trans ut.Translator) CredentialService {
	return &CredentialServiceImpl{
		CredentialRepository: Credential,
		Validator:            validate,
		TranslatorValidation: trans,
	}
}

func (Credential *CredentialServiceImpl) RegisterCredential(register request.RegisterRequest) (*response.RegisterResponse, error) {
	User := &entities.Users{}

	CredentialRequest := request.CredentialRequest{
		Username: register.Username,
	}

	PasswordHash, ExceptionPass := utils.HashPassword(register.Password)
	GetRoles, RoleExists := Credential.CredentialRepository.GetRoleUserForRegister("User")

	if RoleExists != nil {
		return nil, errors.New("Role not found")
	}

	CredentialRequest.Roles_id = uuid.UUID(GetRoles.Id)

	if ExceptionPass != nil {
		return nil, ExceptionPass
	}

	CredentialRequest.Password = PasswordHash

	UserRequest := request.UserRequest{
		First_name: register.First_name,
		Last_name:  register.Last_name,
		Email:      register.Email,
		Address:    register.Address,
	}

	User.RegisterUser(UserRequest, CredentialRequest)

	UserRegister, Err := Credential.CredentialRepository.RegisterUserCredential(*User)

	if Err != nil {
		return nil, Err
	}

	return UserRegister, nil

}

func (Credential *CredentialServiceImpl) AuthUser(ctx echo.Context, request request.AuthRequest) (*response.AuthToken, []exceptions.ValidationMessage, error) {

	ValidationException := Credential.Validator.Struct(request)

	if ValidationException != nil {
		return nil, utils.ValidationError(ctx, Credential.TranslatorValidation, ValidationException), nil
	}

	findUser, isExist, _ := Credential.CredentialRepository.CheckAuthCredential(request)

	if !isExist {
		return nil, nil, errors.New("Authentication Failed Incorrect Credential")
	}

	GenerateToken, errGenerate := utils.GenerateToken(*findUser, ctx)

	if errGenerate != nil {
		return nil, nil, errors.New("Failed Generate Token")
	}

	return GenerateToken, nil, nil
}
