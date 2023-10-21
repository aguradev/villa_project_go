package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthorizationException(ctx echo.Context, message string) error {

	return ctx.JSON(http.StatusUnauthorized, Error{
		Code:    http.StatusUnauthorized,
		Message: message,
		Errors:  nil,
	})

}
