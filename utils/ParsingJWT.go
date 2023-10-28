package utils

import (
	"errors"
	"strings"
	"villa_go/payloads/resources"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func ParsingJWTFromAuthorizationHeader(ctx echo.Context) (*jwt.Token, *resources.JWTProfile, error) {

	AuthorizationHeader := ctx.Request().Header.Get("Authorization")

	if AuthorizationHeader == "" || !strings.HasPrefix(AuthorizationHeader, "Bearer ") {
		return nil, nil, errors.New("Invalid or missing Bearer token in Authorization header")
	}

	tokenString := AuthorizationHeader[7:]

	claims := &resources.JWTProfile{}

	token, errTokenClaims := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	if errTokenClaims != nil {
		return nil, nil, errors.New("Unauthorized")
	}

	if !token.Valid {
		return nil, nil, errors.New("Token invalid")
	}

	return token, claims, nil
}
