package middlewares

import (
	"villa_go/exceptions"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Roles string

const (
	Admin Roles = "Admin"
	User  Roles = "User"
)

func AccessbilityRole(role string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			Payloaduser := c.Get("user").(*jwt.Token)
			Claims, ok := Payloaduser.Claims.(jwt.MapClaims)

			if !ok {
				return exceptions.AuthorizationException(c, "Token Invalid")
			}

			Roles := Claims["role"].(string)

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

			return exceptions.AuthorizationException(c, "Role has not have accessbility")
		}
	}

}
