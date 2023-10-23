package response

import (
	"villa_go/entities/models"
)

func SetVillaResponse(Villa []models.Villa) []models.VillaListResponse {

	var SetVillaResponse []models.VillaListResponse

	for _, item := range Villa {
		VillaResponse := models.VillaListResponse{
			Name:            item.Name,
			Slug:            item.Slug,
			Description:     item.Description,
			Price_per_night: item.Price_per_night,
			Check_in:        item.Check_in,
			Check_out:       item.Check_out,
		}

		SetVillaResponse = append(SetVillaResponse, VillaResponse)
	}

	return SetVillaResponse

}
