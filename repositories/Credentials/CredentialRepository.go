package repositories

import (
	"villa_go/entities"
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/user_response"
)

type CredentialRepository interface {
	GetRoleUserForRegister(role string) (entities.Roles, error)
	CheckAuthCredential(request.CredentialRequest) (bool, error)
	RegisterUserCredential(entities.Users) (*UserResponse.RegisterResponse, error)
}
