package middlewares

import (
	"fmt"
	"villa_go/exceptions"
	userresponse "villa_go/payloads/response/user_response"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Roles string

const (
	admin Roles = "Admin"
	User  Roles = "User"
)

func AccessbilityRole(role string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			Payloaduser := c.Get("user").(*jwt.Token)
			Claims, ok := Payloaduser.Claims.(userresponse.JWTProfile)

			if !ok {
				return exceptions.AuthorizationException(c, "Token Invalid")
			}

			fmt.Println(Claims)

			return nil
		}
	}

}
