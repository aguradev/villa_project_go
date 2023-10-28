package middlewares

import (
	"villa_go/exceptions"
	"villa_go/payloads/resources"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Roles string

const (
	Admin Roles = "Admin"
	User  Roles = "User"
)

func AccessbilityRole(role string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			TokenCookie, ErrCookie := c.Cookie("token")

			if ErrCookie != nil {
				return exceptions.AuthorizationException(c, "Unauthorized")
			}

			TokenVal := TokenCookie.Value

			Claims := &resources.JWTProfile{}

			TokenClaims, ErrClaims := jwt.ParseWithClaims(TokenVal, Claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("SECRET_KEY")), nil
			})

			if ErrClaims != nil {
				return exceptions.AuthorizationException(c, "Unauthorized, token invalid or expired")
			}

			if !TokenClaims.Valid {
				return exceptions.AuthorizationException(c, "Token invalid")
			}

			Roles := Claims.Roles

			switch role {
			case "Admin":
				if Roles == string(Admin) {
					return next(c)
				}
			case "User":
				if Roles == string(User) {
					return next(c)
				}
			default:
				break
			}

			return exceptions.AuthorizationException(c, "You do not have permission to access this resource")
		}
	}

}
