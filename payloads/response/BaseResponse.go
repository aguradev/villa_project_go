package response

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx echo.Context, data interface{}, message string, status int) error {
	return ctx.JSON(status, BaseResponse{
		Status:  uint(status),
		Message: message,
		Data:    data,
	})
}
