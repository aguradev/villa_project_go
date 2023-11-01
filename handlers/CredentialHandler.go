package handlers

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/services"
	"villa_go/utils"

	"github.com/labstack/echo/v4"
)

type CredentialController interface {
	RegisterUser(echo.Context) error
	AuthenticationUser(echo.Context) error
	LogoutUser(echo.Context) error
}

type CredentialControllerImpl struct {
	CredentialService services.CredentialService
}

func CreateCredentialRoutes(Credential services.CredentialService) CredentialController {
	return &CredentialControllerImpl{Credential}
}

func (Crendetial *CredentialControllerImpl) RegisterUser(ctx echo.Context) error {

	var CredentialRequest request.RegisterRequest

	if CredentialbindingException := ctx.Bind(&CredentialRequest); CredentialbindingException != nil {
		return echo.NewHTTPError(http.StatusBadRequest, CredentialbindingException.Error())
	}

	RegisterUser, ValidationMessage, Err, EmailExists := Crendetial.CredentialService.RegisterCredential(ctx, CredentialRequest)

	if ValidationMessage != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ValidationMessage)
	}

	if EmailExists {
		return exceptions.ConflictException(ctx, Err.Error())
	}

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

func (Credential *CredentialControllerImpl) LogoutUser(ctx echo.Context) error {

	_, CookieErr := ctx.Cookie("token")

	if CookieErr != nil {
		return exceptions.AuthorizationException(ctx, "Unauthorized")
	}

	utils.DeleteTokenCookie(ctx)

	return response.HandleSuccess(ctx, nil, "Logout successfully", http.StatusCreated)

}
