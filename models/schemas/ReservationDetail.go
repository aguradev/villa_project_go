package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ReservationDetail struct {
	Id                 uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Villa_id           *uuid.UUID      `gorm:"type:uuid"`
	Villa              *Villa          `gorm:"foreignKey:Villa_id;references:Id;"`
	Guest_count        uint            `gorm:"type:int"`
	Check_in_date      *time.Time      `gorm:"type:time"`
	Check_out_date     *time.Time      `gorm:"type:time"`
	Duration_day_price decimal.Decimal `gorm:"type:decimal(10,2)"`
	Tax                decimal.Decimal `gorm:"type:decimal(10,2)"`
	Total              decimal.Decimal `gorm:"type:decimal(10,2)"`
	SnapURL            string          `gorm:"text"`
	CreatedAt          time.Time       `gorm:"autoCreateTime"`
	UpdatedAt          time.Time       `gorm:"autoUpdateTime:milli"`
	DeletedAt          gorm.DeletedAt  `gorm:"index"`
}
