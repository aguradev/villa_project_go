package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ReservationDetail struct {
	Id                 uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Villa_id           *uuid.UUID
	Villa              *Villa
	Check_in_date      *time.Time
	Check_out_date     *time.Time
	Duration_day_price decimal.Decimal
	Tax                decimal.Decimal
	Total              decimal.Decimal
	Guest_count        uint
	SnapURL            string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}
