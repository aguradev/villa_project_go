package request

import "github.com/google/uuid"

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles_id uuid.UUID
}

type RegisterRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
}
