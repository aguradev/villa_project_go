package controllers

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	CredentialService "villa_go/services/Credentials"

	"github.com/labstack/echo/v4"
)

type CredentialController interface {
	RegisterUser(echo.Context) error
	AuthenticationUser(echo.Context) error
}

type CredentialControllerImpl struct {
	CredentialService CredentialService.CredentialService
}

func CreateCredentialRoutes(Credential CredentialService.CredentialService) CredentialController {
	return &CredentialControllerImpl{Credential}
}

func (Crendetial *CredentialControllerImpl) RegisterUser(ctx echo.Context) error {

	var CredentialRequest request.RegisterRequest

	if CredentialbindingException := ctx.Bind(&CredentialRequest); CredentialbindingException != nil {
		return echo.NewHTTPError(http.StatusBadRequest, CredentialbindingException.Error())
	}

	RegisterUser, Err := Crendetial.CredentialService.RegisterCredential(CredentialRequest)

	if Err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Register Failed")
	}

	return response.HandleSuccess(ctx, RegisterUser, "Register User Success", http.StatusCreated)
}

func (Credential *CredentialControllerImpl) AuthenticationUser(ctx echo.Context) error {

	var Request request.AuthRequest

	if RequestException := ctx.Bind(&Request); RequestException != nil {
		return echo.NewHTTPError(http.StatusBadGateway, RequestException)
	}

	GetAuthResponse, ValidationException, errException := Credential.CredentialService.AuthUser(ctx, Request)

	if ValidationException != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ValidationException)
	}

	if errException != nil {
		return exceptions.AuthorizationException(ctx, errException.Error())
	}

	return response.HandleSuccess(ctx, GetAuthResponse, "User Authentication Success", http.StatusOK)
}
