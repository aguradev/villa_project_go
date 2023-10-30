package resources

import "villa_go/models/entities"

type VillaFacilityResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func VillaFacilityListResponse(facilities []entities.Facility) []VillaFacilityResponse {

	var Facilities []VillaFacilityResponse

	for _, val := range facilities {

		Facility := VillaFacilityResponse{
			Id:   val.Id.String(),
			Name: val.Name,
		}

		Facilities = append(Facilities, Facility)

	}

	return Facilities

}

func (F *VillaFacilityResponse) GetVillaFacility(facility entities.Facility) {
	F.Id = facility.Id.String()
	F.Name = facility.Name
}
