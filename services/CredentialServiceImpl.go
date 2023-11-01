package services

import (
	"errors"
	"villa_go/exceptions"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CredentialService interface {
	RegisterCredential(echo.Context, request.RegisterRequest) (*resources.RegisterResponse, []exceptions.ValidationMessage, error, bool)
	AuthUser(echo.Context, request.AuthRequest) (*resources.AuthToken, []exceptions.ValidationMessage, error)
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

func (Credential *CredentialServiceImpl) RegisterCredential(ctx echo.Context, register request.RegisterRequest) (*resources.RegisterResponse, []exceptions.ValidationMessage, error, bool) {
	User := &entities.Users{}

	ValidationMessage := Credential.Validator.Struct(register)

	if ValidationMessage != nil {
		return nil, utils.ValidationError(ctx, Credential.TranslatorValidation, ValidationMessage), nil, false
	}

	CredentialRequest := request.CredentialRequest{
		Username: register.Username,
	}

	PasswordHash, ExceptionPass := utils.HashPassword(register.Password)
	GetRoles, RoleExists := Credential.CredentialRepository.GetRoleUserForRegister("User")

	if RoleExists != nil {
		return nil, nil, errors.New("Role not found"), false
	}

	CredentialRequest.Roles_id = uuid.UUID(GetRoles.Id)

	if ExceptionPass != nil {
		return nil, nil, ExceptionPass, false
	}

	CredentialRequest.Password = PasswordHash

	UserRequest := request.UserRequest{
		First_name: register.First_name,
		Last_name:  register.Last_name,
		Email:      register.Email,
		Address:    register.Address,
	}

	EmailExists, ErrMessage := Credential.CredentialRepository.CheckEmailExists(UserRequest.Email)

	if EmailExists {
		return nil, nil, ErrMessage, EmailExists
	}

	User.RegisterUser(UserRequest, CredentialRequest)

	UserRegister, Err := Credential.CredentialRepository.RegisterUserCredential(*User)

	if Err != nil {
		return nil, nil, Err, false
	}

	return UserRegister, nil, nil, false

}

func (Credential *CredentialServiceImpl) AuthUser(ctx echo.Context, request request.AuthRequest) (*resources.AuthToken, []exceptions.ValidationMessage, error) {

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
