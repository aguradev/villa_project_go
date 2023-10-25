package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DeleteResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

func HandleResponseDelete(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusOK, DeleteResponse{
		Status:  http.StatusOK,
		Message: message,
	})
}
