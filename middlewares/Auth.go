package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func LoginSignedIn() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("SECRET_KEY")),
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "missing or malformed jwt",
			})
		},
	})
}
