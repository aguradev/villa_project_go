package request

type ReservationRequest struct {
	Villa_id          string                    `json:"villa_id,omitempty"`
	Guest_count       uint                      `json:"guest_count,omitempty"`
	Check_in_date     string                    `json:"check_in_date"`
	ReservationDetail *ReservationDetailRequest `json:"reservation_detail,omitempty"`
}

type ReservationDetailRequest struct {
	SnapURL string `json:"snap_url,omitempty"`
}
