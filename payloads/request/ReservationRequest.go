package request

type ReservationRequest struct {
	Villa_id          string                    `json:"villa_id,omitempty" validate:"required"`
	Guest_count       uint                      `json:"guest_count,omitempty" validate:"required"`
	Check_in_date     string                    `json:"check_in_date" validate:"required"`
	Check_out_date    string                    `json:"check_out_date" validate:"required"`
	ReservationDetail *ReservationDetailRequest `json:"reservation_detail,omitempty"`
}

type ReservationDetailRequest struct {
	SnapURL string `json:"snap_url,omitempty"`
}

type ReservationEmailRequest struct {
	Name_customer  string
	Villa_name     string
	Check_in_date  string
	Check_out_date string
	Guest_count    uint
	Total          int
}
