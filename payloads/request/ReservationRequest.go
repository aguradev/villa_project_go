package request

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type ReservationRequest struct {
	User_id     *uuid.UUID `json:"user_id"`
	Villa_id    *uuid.UUID `json:"villa_id"`
	Guest_count uint       `json:"guest_count"`
}

type ReservationDetailRequest struct {
	Villa_id *uuid.UUID `json:"villa_id"`
	Amount   decimal.Decimal
}
