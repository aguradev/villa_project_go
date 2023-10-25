package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Credentials struct {
	Id        uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Roles_id  uuid.UUID
	Roles     *Roles
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
