package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VillaLocation struct {
	Id        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	Name      string         `gorm:"type:varchar(50)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (VillaLocation) TableName() string {
	return "location"
}
