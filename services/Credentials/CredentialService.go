package services

import (
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/user_response"
)

type CredentialService interface {
	RegisterCredential(request.RegisterRequest) (*UserResponse.RegisterResponse, error)
}
