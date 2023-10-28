package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"
	"villa_go/exceptions"
	"villa_go/payloads/resources"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
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

func DeleteTokenCookie(ctx echo.Context) {

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()

	ctx.SetCookie(cookie)

}

func CheckCookieSignatured(ctx echo.Context) (*jwt.Token, *resources.JWTProfile, bool, error) {
	GetCookie, ErrCookie := ctx.Cookie("token")

	if ErrCookie != nil {
		return nil, nil, false, echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "No cookies is set"))
	}

	TokenString := GetCookie.Value
	GetTokenAuthorization := ctx.Request().Header.Get("Authorization")

	if !strings.HasPrefix(GetTokenAuthorization, "Bearer ") {
		return nil, nil, false, errors.New("Invalid or missing Bearer token in Authorization header")
	}

	if GetTokenAuthorization != "Bearer "+TokenString {
		return nil, nil, false, errors.New("Token in Authorization invalid")
	}

	JwtClaims := &resources.JWTProfile{}

	ParseToken, errParsing := jwt.ParseWithClaims(TokenString, JwtClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	if errParsing != nil {
		return nil, nil, false, errParsing
	}

	return ParseToken, JwtClaims, true, nil
}
