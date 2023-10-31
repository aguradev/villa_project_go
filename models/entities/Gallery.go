package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Gallery struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Villa_id  *uuid.UUID
	Fileurl   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Gallery) TableName() string {
	return "galleries"
}
