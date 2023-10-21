package services

import (
	"villa_go/exceptions"
	"villa_go/payloads/request"
	UserResponse "villa_go/payloads/response/user_response"

	"github.com/labstack/echo/v4"
)

type CredentialService interface {
	RegisterCredential(request.RegisterRequest) (*UserResponse.RegisterResponse, error)
	AuthUser(echo.Context, request.AuthRequest) (*UserResponse.AuthToken, []exceptions.ValidationMessage, error)
}
