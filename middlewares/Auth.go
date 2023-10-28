package middlewares

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/resources"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func CheckTokenByCookie() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			GetCookie, ErrCookie := ctx.Cookie("token")

			if ErrCookie != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "No cookies is set"))
			}

			TokenString := GetCookie.Value
			JwtClaims := &resources.JWTProfile{}

			ParseToken, errParsing := jwt.ParseWithClaims(TokenString, JwtClaims, func(t *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("SECRET_KEY")), nil
			})

			if errParsing != nil {
				switch errParsing {
				case jwt.ErrSignatureInvalid:
					return ctx.JSON(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized"))
				case jwt.ErrTokenExpired:
					return ctx.JSON(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized token expired"))
				default:
					return ctx.JSON(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized Access"))
				}
			}

			if !ParseToken.Valid {
				return ctx.JSON(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized Access Token Invalid"))
			}

			return next(ctx)
		}
	}
}

func VerifyTokenSignature() echo.MiddlewareFunc {
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
