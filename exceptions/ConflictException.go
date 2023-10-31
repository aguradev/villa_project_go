package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConflictException(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusConflict, Error{
		Code:    http.StatusConflict,
		Message: message,
		Errors:  nil,
	})
}
