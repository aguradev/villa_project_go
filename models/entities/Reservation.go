package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	Id                    uuid.UUID `gorm:"default:uuid_generate_v4()"`
	User_id               *uuid.UUID
	Reservation_detail_id *uuid.UUID
	User                  *Users
	Reservation_detail    *ReservationDetail
	Transaction_date      time.Time
	Status                string
	CreatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
