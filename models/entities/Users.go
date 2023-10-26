package entities

import (
	"time"
	"villa_go/payloads/request"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id              *uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Credential_id   uuid.UUID
	Credential      *Credentials
	First_name      string
	Last_name       string
	Email           string
	Address         string
	Profile_picture string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (user *Users) RegisterUser(userRequest request.UserRequest, credentialRequest request.CredentialRequest) {
	user.First_name = userRequest.First_name
	user.Last_name = userRequest.Last_name
	user.Email = userRequest.Email
	user.Address = userRequest.Address
	user.Credential = &Credentials{
		Username: credentialRequest.Username,
		Password: credentialRequest.Password,
		Roles_id: uuid.UUID(credentialRequest.Roles_id),
	}

}
