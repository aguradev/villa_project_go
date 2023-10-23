package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Villa struct {
	Id              uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name            string          `gorm:"type:varchar(50)"`
	Slug            string          `gorm:"type:varchar(50)"`
	Description     string          `gorm:"type:text"`
	Price_per_night decimal.Decimal `gorm:"type:decimal(10,2)"`
	Check_in        time.Time       `gorm:"type:time"`
	Check_out       time.Time       `gorm:"type:time"`
	Status          string          `gorm:"type:varchar(50)"`
	CreatedAt       time.Time       `gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime:milli"`
	DeletedAt       gorm.DeletedAt  `gorm:"index"`
}

func (Villa) TableName() string {
	return "properties_villa"
}
