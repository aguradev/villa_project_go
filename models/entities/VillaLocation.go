package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VillaLocation struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
