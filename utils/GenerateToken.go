package utils

import (
	"time"
	"villa_go/entities/models"
	UserResponse "villa_go/payloads/response/user_response"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GenerateToken(User models.Users, ctx echo.Context) (*UserResponse.AuthToken, error) {

	var Auth UserResponse.AuthToken
	var Payload UserResponse.JWTProfile

	Payload.Id = User.Credential_id
	Payload.Username = User.Credential.Username
	Payload.Roles = User.Credential.Roles.Role
	Payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 300))

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
