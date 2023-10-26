package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ReservationDetail struct {
	Id        uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Villa_id  *uuid.UUID
	Villa     *Villa
	Tax       decimal.Decimal
	Total     decimal.Decimal
	Amount    decimal.Decimal
	SnapURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
