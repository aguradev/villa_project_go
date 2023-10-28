package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetTokenCookie(ctx echo.Context, tokenJWT string, expired *jwt.NumericDate) {

	SetCookieToken := &http.Cookie{
		Name:     "token",
		Value:    tokenJWT,
		Expires:  expired.Time,
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(SetCookieToken)

}
