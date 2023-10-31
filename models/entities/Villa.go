package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Villa struct {
	Id              *uuid.UUID `gorm:"default:uuid_generate_v4()"`
	Location_id     *uuid.UUID
	Location        *VillaLocation
	Facility        []Facility `gorm:"many2many:villa_has_facility;foreignKey:Id;references:Id;"`
	Gallery         []Gallery
	Name            string
	Slug            string
	Description     string
	Max_capacity    uint
	Price_per_night *decimal.Decimal
	Check_in        *time.Time
	Check_out       *time.Time
	Address         string
	Status          string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *gorm.DeletedAt
}

func (Villa) TableName() string {
	return "properties_villa"
}
