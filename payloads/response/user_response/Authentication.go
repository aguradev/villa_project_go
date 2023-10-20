package userresponse

import "villa_go/entities/models"

type AuthJWTResponse struct {
	Username string           `json:"username"`
	Token    string           `json:"token"`
	Profile  RegisterResponse `json:"profile"`
}

type RegisterResponse struct {
	Username   string `json:"username omiempty"`
	First_name string `json:"first_name omiempty"`
	Last_name  string `json:"last_name omiempty"`
	Email      string `json:"email omiempty"`
	Address    string `json:"address omiempty"`
}

func (user *RegisterResponse) GetRegisterResponse(User models.Users) {
	user.Username = User.Credential.Username
	user.First_name = User.First_name
	user.Last_name = User.Last_name
	user.Email = User.Email
	user.Address = User.Address
}
