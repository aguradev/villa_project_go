package middlewares

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func VerifiyTokenByCookie() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			ParseToken, _, IsValid, errParsing := utils.CheckCookieSignatured(ctx)

			if errParsing != nil && !IsValid {
				switch errParsing {
				case jwt.ErrSignatureInvalid:
					return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized"))
				case jwt.ErrTokenExpired:
					return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized token expired"))
				default:
					return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, errParsing.Error()))
				}
			}

			if !ParseToken.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized Access Token Invalid"))
			}

			return next(ctx)
		}
	}
}

func VerifyTokenSignature() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("SECRET_KEY")),
		ErrorHandlerWithContext: func(err error, ctx echo.Context) error {
			return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized"))
		},
	})
}
