package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFoundException(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusNotFound, Error{
		Code:    http.StatusNoContent,
		Message: message,
		Errors:  nil,
	})
}
