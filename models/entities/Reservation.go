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

func (r *Reservation) GetReservationRequest(request request.ReservationRequest, price_per_day decimal.Decimal, UserId *uuid.UUID, VillaId *uuid.UUID) {

	r.User_id = UserId
	r.Status = "Pending"
	r.Transaction_date = time.Now().Local()
	r.Reservation_detail = &ReservationDetail{
		Villa_id: VillaId,
		Tax:      decimal.NewFromInt(10000),
	}
	r.Reservation_detail.Total = r.Reservation_detail.Tax.Add(price_per_day)

}
