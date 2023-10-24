package request

type LocationRequest struct {
	Name []string `json:"location" validate:"required"`
}
