package utils

import (
	"time"
	"villa_go/entities/models"
	UserResponse "villa_go/payloads/response/user_response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(User models.Users) (*UserResponse.AuthToken, error) {

	var Auth UserResponse.AuthToken

	claims := &UserResponse.JWTProfile{
		User.Credential_id,
		User.Credential.Username,
		User.Credential.Roles.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 300)),
		},
	}

	GetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, errException := GetToken.SignedString([]byte(viper.GetString("SECRET_KEY")))

	if errException != nil {
		return nil, errException
	}

	Auth.Token = result
	Auth.Email = User.Email
	Auth.Fullname = User.First_name + " " + User.Last_name

	return &Auth, nil

}
