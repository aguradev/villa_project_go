package services

import (
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/UserResponse"
)

type CredentialService interface {
	RegisterCredential(request.RegisterRequest) (*UserResponse.RegisterResponse, error)
}
