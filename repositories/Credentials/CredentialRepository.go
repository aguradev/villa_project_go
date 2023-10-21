package repositories

import (
	"villa_go/entities/models"
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/user_response"
)

type CredentialRepository interface {
	GetRoleUserForRegister(role string) (models.Roles, error)
	CheckAuthCredential(request.AuthRequest) (*models.Users, bool, error)
	RegisterUserCredential(models.Users) (*UserResponse.RegisterResponse, error)
}
