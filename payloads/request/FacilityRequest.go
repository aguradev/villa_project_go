package request

type FacilityRequest struct {
	Name []string `json:"name" validate:"required"`
}

type FacilityToVillaRequest struct {
	Id []string `json:"facility_id" validate:"required"`
}
