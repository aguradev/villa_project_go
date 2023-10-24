package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type ValidationMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidationException(ctx echo.Context, message string, form []ValidationMessage) error {
	return ctx.JSON(http.StatusUnprocessableEntity, Error{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
		Errors:  form,
	})
}
