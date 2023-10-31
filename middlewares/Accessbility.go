package middlewares

import (
	"villa_go/exceptions"
	"villa_go/utils"

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
			GetClaimsToken, ErrClaims := utils.ClaimToken(c)

			if ErrClaims != nil {
				return exceptions.AuthorizationException(c, ErrClaims.Error())
			}

			Roles := GetClaimsToken.Roles

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
