package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Villa struct {
	Id              uuid.UUID       `gorm:"type:uuid;primaryKey;"`
	Name            string          `gorm:"type:varchar(50)"`
	Slug            string          `gorm:"type:varchar(50)"`
	Description     string          `gorm:"type:text"`
	Price_per_night decimal.Decimal `gorm:"type:decimal(10,2)"`
	Check_in        time.Time
	Check_out       time.Time
	Status          string         `gorm:"type:varchar(50)"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type VillaListResponse struct {
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	Description     string          `json:"description"`
	Price_per_night decimal.Decimal `json:"price_per_night"`
	Check_in        time.Time       `json:"check_in"`
	Check_out       time.Time       `json:"check_out"`
}
