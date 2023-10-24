package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AppException(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusInternalServerError, Error{
		Code:    http.StatusInternalServerError,
		Message: message,
		Errors:  nil,
	})
}
