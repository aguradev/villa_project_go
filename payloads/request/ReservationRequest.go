package request

type ReservationRequest struct {
	Villa_id    string `json:"villa_id"`
	Guest_count uint   `json:"guest_count"`
}
