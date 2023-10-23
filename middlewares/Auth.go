package middlewares

import (
	"villa_go/exceptions"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func VerifyToken() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(viper.GetString("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return exceptions.AuthorizationException(c, "Token invalid or expried")
		},
	})
}
