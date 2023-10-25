package resources

import (
	"villa_go/models/entities"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type AuthToken struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type JWTProfile struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Roles    string    `json:"role"`
	jwt.RegisteredClaims
}

type RegisterResponse struct {
	Username   string `json:"username,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Last_name  string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"address,omitempty"`
}

func (user *RegisterResponse) GetRegisterResponse(User entities.Users) {
	user.Username = User.Credential.Username
	user.First_name = User.First_name
	user.Last_name = User.Last_name
	user.Email = User.Email
	user.Address = User.Address
}
