package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Gallery struct {
	Id        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	Villa_id  *uuid.UUID `gorm:"type:uuid"`
	Fileurl   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

func (Gallery) TableName() string {
	return "galleries"
}
