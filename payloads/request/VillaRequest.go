package request

import (
	"time"

	"github.com/shopspring/decimal"
)

type VillaRequest struct {
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	Description     string          `json:"description"`
	Price_per_night decimal.Decimal `json:"price_per_night"`
	Check_in        time.Time       `json:"check_in"`
	Check_out       time.Time       `json:"check_out"`
}
