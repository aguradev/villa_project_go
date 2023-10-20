package services

import (
	"errors"
	"villa_go/entities"
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/UserResponse"
	CredentialRepo "villa_go/repositories/Credentials"
	"villa_go/utils"

	"github.com/google/uuid"
)

type CredentialServiceImpl struct {
	CredentialRepository CredentialRepo.CredentialRepository
}

func CreateCredentialServiceImplement(Credential CredentialRepo.CredentialRepository) CredentialService {
	return &CredentialServiceImpl{
		CredentialRepository: Credential,
	}
}

func (Credential *CredentialServiceImpl) RegisterCredential(register request.RegisterRequest) (*UserResponse.RegisterResponse, error) {

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
