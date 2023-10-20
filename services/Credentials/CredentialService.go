package services

import (
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/UserResponse"

	"github.com/labstack/echo"
)

type CredentialService interface {
	RegisterCredential(echo.Context, request.RegisterRequest) (*UserResponse.RegisterResponse, error)
}
