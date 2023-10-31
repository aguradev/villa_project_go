package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Facility struct {
	Id        uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
