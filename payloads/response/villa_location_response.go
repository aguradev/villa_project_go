package response

import (
	"villa_go/entities/models"
)

type VillaLocationResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func VillaLocationsReponses(locations []models.VillaLocation) []VillaLocationResponse {

	var Listlocations []VillaLocationResponse

	for _, val := range locations {

		MappingLocation := VillaLocationResponse{
			Id:   val.Id.String(),
			Name: val.Name,
		}

		Listlocations = append(Listlocations, MappingLocation)

	}

	return Listlocations

}

func (L *VillaLocationResponse) VillaDetailLocationResponse(location models.VillaLocation) {
	L.Id = location.Id.String()
	L.Name = location.Name
}
