package resources

import (
	"strconv"
	"time"
	"villa_go/models/entities"
)

type ReservationResource struct {
	Id                 string                     `json:"id,omitempty"`
	Transaction_date   *time.Time                 `json:"transaction_date,omitempty"`
	Status             string                     `json:"status,omitempty"`
	User               *UserResource              `json:"user,omitempty"`
	Reservation_detail *ReservationDetailResource `json:"reservation_detail,omitempty"`
}

type ReservationDetailResource struct {
	Id             string             `json:"id,omitempty"`
	Check_in_date  *time.Time         `json:"check_in_date,omitempty"`
	Check_out_date *time.Time         `json:"check_out_date,omitempty"`
	Villa          *VillaListResponse `json:"villa,omitempty"`
	Tax            int                `json:"tax,omitempty"`
	Total          int                `json:"total,omitempty"`
	SnapURL        string             `json:"transaction_url,omitempty"`
}

func (r *ReservationResource) GetDetailReservationResponse(reservation entities.Reservation) {
	SetTotal, ErrException := strconv.Atoi(reservation.Reservation_detail.Total.String())

	if ErrException != nil {
		SetTotal = 0
	}

	r.Id = reservation.Id.String()
	r.Transaction_date = &reservation.Transaction_date
	r.Status = reservation.Status
	r.User = &UserResource{
		First_name: reservation.User.First_name,
		Last_name:  reservation.User.Last_name,
		Email:      reservation.User.Email,
	}
	r.Reservation_detail = &ReservationDetailResource{
		Id:             reservation.Reservation_detail.Id.String(),
		Check_in_date:  reservation.Reservation_detail.Check_in_date,
		Check_out_date: reservation.Reservation_detail.Check_out_date,
		Villa: &VillaListResponse{
			Name:    reservation.Reservation_detail.Villa.Name,
			Address: reservation.Reservation_detail.Villa.Address,
			Status:  reservation.Reservation_detail.Villa.Status,
			Location: &VillaLocationResponse{
				Name: reservation.Reservation_detail.Villa.Location.Name,
			},
		},
		Tax:     int(reservation.Reservation_detail.Tax.IntPart()),
		Total:   SetTotal,
		SnapURL: reservation.Reservation_detail.SnapURL,
	}
}

func GetListReservationResponse(reservations []entities.Reservation) []ReservationResource {

	var ReservationResponse []ReservationResource

	for _, ResVal := range reservations {

		Reservation := ReservationResource{
			Transaction_date: &ResVal.Transaction_date,
			Status:           ResVal.Status,
			User: &UserResource{
				First_name: ResVal.User.First_name,
				Last_name:  ResVal.User.Last_name,
			},
			Reservation_detail: &ReservationDetailResource{
				Check_in_date:  ResVal.Reservation_detail.Check_in_date,
				Check_out_date: ResVal.Reservation_detail.Check_out_date,
				Villa: &VillaListResponse{
					Name: ResVal.Reservation_detail.Villa.Name,
				},
			},
		}

		ReservationResponse = append(ReservationResponse, Reservation)

	}

	return ReservationResponse

}
