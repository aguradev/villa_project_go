package utils

import (
	"time"
	"villa_go/models/entities"
	"villa_go/payloads/resources"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GenerateToken(User entities.Users, ctx echo.Context) (*resources.AuthToken, error) {

	var Auth resources.AuthToken
	var Payload resources.JWTProfile

	Payload.Id = User.Credential_id
	Payload.Username = User.Credential.Username
	Payload.Roles = User.Credential.Roles.Role
	Payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 48))

	GetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Payload)

	result, errException := GetToken.SignedString([]byte(viper.GetString("SECRET_KEY")))

	if errException != nil {
		return nil, errException
	}

	Auth.Token = result
	Auth.Email = User.Email
	Auth.Fullname = User.First_name + " " + User.Last_name

	return &Auth, nil

}
