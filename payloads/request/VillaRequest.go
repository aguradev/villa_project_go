package request

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type VillaRequest struct {
	Name            string          `json:"name" validate:"required"`
	Slug            string          `json:"slug" validate:"required"`
	Description     string          `json:"description" validate:"required"`
	Address         string          `json:"address" validate:"required"`
	Max_capacity    uint            `json:"max_capacity" validate:"required"`
	Price_per_night decimal.Decimal `json:"price_per_night" validate:"required"`
	Check_in        string          `json:"check_in" validate:"required"`
	Check_out       string          `json:"check_out" validate:"required"`
	Status          string          `json:"status"`
	Location_id     *uuid.UUID      `json:"location_id"`
}
