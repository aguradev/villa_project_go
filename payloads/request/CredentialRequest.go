package request

import "github.com/google/uuid"

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=5"`
}

type CredentialRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles_id uuid.UUID
}

type RegisterRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	First_name string `json:"first_name" validate:"required"`
	Last_name  string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Address    string `json:"address" validate:"required"`
}
