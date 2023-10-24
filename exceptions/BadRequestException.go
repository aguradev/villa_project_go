package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequestException(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusBadRequest, Error{
		Code:    http.StatusBadRequest,
		Message: message,
		Errors:  nil,
	})
}
