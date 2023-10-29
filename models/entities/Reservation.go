package entities

import (
	"time"
	"villa_go/payloads/request"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Reservation struct {
	Id                    *uuid.UUID `gorm:"default:uuid_generate_v4()"`
	User_id               *uuid.UUID
	Reservation_detail_id *uuid.UUID
	User                  *Users
	Reservation_detail    *ReservationDetail
	Transaction_date      time.Time
	Status                string
	CreatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}

func (r *Reservation) GetReservationRequest(request request.ReservationRequest, price_per_day decimal.Decimal, userId *uuid.UUID, villaId *uuid.UUID, check_in_date *time.Time, check_out_date *time.Time) {

	r.User_id = userId
	r.Status = "pending"
	r.Transaction_date = time.Now().Local()
	r.Reservation_detail = &ReservationDetail{
		Villa_id:       villaId,
		Check_in_date:  check_in_date,
		Check_out_date: check_out_date,
		Guest_count:    request.Guest_count,
		Tax:            decimal.NewFromInt(50000),
	}
	r.Reservation_detail.Total = r.Reservation_detail.Tax.Add(price_per_day)

}
