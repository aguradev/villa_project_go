package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	Id                    uuid.UUID          `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	User_id               *uuid.UUID         `gorm:"type:uuid;"`
	Reservation_detail_id *uuid.UUID         `gorm:"type:uuid;"`
	User                  *Users             `gorm:"foreignKey:User_id;references:id;"`
	Reservation_detail    *ReservationDetail `gorm:"foreignKey:Reservation_detail_id;references:id;"`
	Transaction_date      time.Time          `gorm:"type:time"`
	Status                string             `gorm:"varchar(50)"`
	CreatedAt             time.Time          `gorm:"autoCreateTime"`
	UpdatedAt             time.Time          `gorm:"autoUpdateTime:milli"`
	DeletedAt             gorm.DeletedAt     `gorm:"index"`
}
