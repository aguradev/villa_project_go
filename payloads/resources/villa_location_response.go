package resources

import "villa_go/models/schemas"

type VillaLocationResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func VillaLocationsReponses(locations []schemas.VillaLocation) []VillaLocationResponse {

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

func (L *VillaLocationResponse) VillaDetailLocationResponse(location schemas.VillaLocation) {
	L.Id = location.Id.String()
	L.Name = location.Name
}
