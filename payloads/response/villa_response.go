package response

import (
	"strconv"
	"time"
	"villa_go/entities/models"
)

type VillaListResponse struct {
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Description     string    `json:"description"`
	Price_per_night int       `json:"price_per_night"`
	Check_in        time.Time `json:"check_in"`
	Check_out       time.Time `json:"check_out"`
}

func SetVillaResponse(Villa []models.Villa) []VillaListResponse {

	var SetVillaResponse []VillaListResponse

	for _, item := range Villa {

		setPrice, err := strconv.Atoi(item.Price_per_night.String())

		if err != nil {
			setPrice = 0
		}

		VillaResponse := VillaListResponse{
			Name:            item.Name,
			Slug:            item.Slug,
			Description:     item.Description,
			Price_per_night: setPrice,
			Check_in:        item.Check_in,
			Check_out:       item.Check_out,
		}

		SetVillaResponse = append(SetVillaResponse, VillaResponse)
	}

	return SetVillaResponse

}

func (v *VillaListResponse) SetVillaDetailResponse(villa models.Villa) {
	setPrice, err := strconv.Atoi(villa.Price_per_night.String())

	if err != nil {
		setPrice = 0
	}

	v.Name = villa.Name
	v.Description = villa.Description
	v.Price_per_night = setPrice
	v.Slug = villa.Slug
	v.Check_in = villa.Check_in
	v.Check_out = villa.Check_out
}
