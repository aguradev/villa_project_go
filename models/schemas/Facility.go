package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Facility struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(255);index:unique"`
	Villa     []Villa   `gorm:"many2many:villa_has_facility;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
