package controllers

import (
	"net/http"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	CredentialService "villa_go/services/Credentials"

	"github.com/labstack/echo"
)

type CredentialControllerImpl struct {
	CredentialService CredentialService.CredentialService
}

func CreateCredentialRoutes(Credential CredentialService.CredentialService, group *echo.Group) {
	CredentialHandler := CredentialControllerImpl{Credential}

	group.POST("/register", CredentialHandler.RegisterUser)
	group.POST("/auth", CredentialHandler.AuthenticationUser)
}

func (Crendetial *CredentialControllerImpl) RegisterUser(ctx echo.Context) error {

	var CredentialRequest request.RegisterRequest

	if CredentialbindingException := ctx.Bind(&CredentialRequest); CredentialbindingException != nil {
		return echo.NewHTTPError(http.StatusBadRequest, CredentialbindingException.Error())
	}

	RegisterUser, Err := Crendetial.CredentialService.RegisterCredential(ctx, CredentialRequest)

	if Err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Register Failed")
	}

	return response.HandleSuccess(ctx, RegisterUser, "Register User Success", http.StatusCreated)
}

func (Credential *CredentialControllerImpl) AuthenticationUser(ctx echo.Context) error {

	var Request request.CredentialRequest

	if RequestException := ctx.Bind(&Request); RequestException.Error != nil {
		return echo.NewHTTPError(http.StatusBadGateway, RequestException.Error)
	}

	return nil
}
