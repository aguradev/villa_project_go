package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Villa struct {
	Id              uuid.UUID
	Location_id     *uuid.UUID
	Location        *VillaLocation
	Name            string
	Slug            string
	Description     string
	Max_capacity    uint
	Price_per_night decimal.Decimal
	Check_in        time.Time
	Check_out       time.Time
	Address         string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
