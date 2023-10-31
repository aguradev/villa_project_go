package utils

import (
	"errors"
	"net/http"
	"strings"
	"villa_go/payloads/resources"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func SetTokenCookie(ctx echo.Context, tokenJWT string) {

	SetCookieToken := &http.Cookie{
		Name:     "token",
		Value:    tokenJWT,
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(SetCookieToken)

}

func DeleteTokenCookie(ctx echo.Context) {

	SetCookieToken := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(SetCookieToken)

}

func ClaimToken(ctx echo.Context) (*resources.JWTProfile, error) {
	TokenCookie, ErrCookie := ctx.Cookie("token")

	if ErrCookie != nil {
		return nil, errors.New("Unauthorized")
	}

	TokenVal := TokenCookie.Value
	Claims := &resources.JWTProfile{}

	TokenClaims, ErrClaims := jwt.ParseWithClaims(TokenVal, Claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	if ErrClaims != nil {
		return nil, errors.New("Unauthorized, token invalid or expired")
	}

	if !TokenClaims.Valid {
		return nil, errors.New("Token invalid")
	}

	return Claims, nil
}

func CheckCookieSignatured(ctx echo.Context) (*jwt.Token, *resources.JWTProfile, bool, error) {
	GetCookie, ErrCookie := ctx.Cookie("token")

	if ErrCookie != nil {
		return nil, nil, false, errors.New("unauthroized")
	}

	TokenString := GetCookie.Value
	GetTokenAuthorization := ctx.Request().Header.Get("Authorization")

	if !strings.HasPrefix(GetTokenAuthorization, "Bearer ") {
		return nil, nil, false, errors.New("Invalid or missing Bearer token in Authorization header")
	}

	if GetTokenAuthorization[7:] != TokenString {
		// AnotherTokenAvailable, TokenClaim, ErrReplace := ParsingJWTFromAuthorizationHeader(ctx)

		// if ErrReplace != nil {
		// 	return nil, nil, false, ErrReplace
		// }

		// DeleteTokenCookie(ctx)
		// SetTokenCookie(ctx, GetTokenAuthorization[7:], TokenClaim.ExpiresAt)

		return nil, nil, false, errors.New("Token access dosent match")
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
