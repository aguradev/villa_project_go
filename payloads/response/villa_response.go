package response

import (
	"strconv"
	"time"
	"villa_go/models/schemas"
)

type VillaListResponse struct {
	Id              string                 `json:"id ,omitempty"`
	Name            string                 `json:"name"`
	Slug            string                 `json:"slug"`
	Description     string                 `json:"description,omitempty"`
	Address         string                 `json:"address,omitempty"`
	Max_capacity    uint                   `json:"max_capacity,omitempty"`
	Price_per_night int                    `json:"price_per_night,omitempty"`
	Check_in        *time.Time             `json:"check_in,omitempty"`
	Check_out       *time.Time             `json:"check_out,omitempty"`
	Status          string                 `json:"status"`
	Location        *VillaLocationResponse `json:"location,omitempty"`
}

func SetVillaResponse(Villa []schemas.Villa) []VillaListResponse {

	var SetVillaResponse []VillaListResponse

	for _, item := range Villa {

		setPrice, err := strconv.Atoi(item.Price_per_night.String())

		if err != nil {
			setPrice = 0
		}

		VillaResponse := VillaListResponse{
			Name:            item.Name,
			Slug:            item.Slug,
			Price_per_night: setPrice,
			Status:          item.Status,
			Location: &VillaLocationResponse{
				Name: item.Location.Name,
			},
		}

		SetVillaResponse = append(SetVillaResponse, VillaResponse)
	}

	return SetVillaResponse

}

func (v *VillaListResponse) SetVillaDetailResponse(villa schemas.Villa) {
	setPrice, err := strconv.Atoi(villa.Price_per_night.String())

	if err != nil {
		setPrice = 0
	}

	v.Id = villa.Id.String()
	v.Name = villa.Name
	v.Description = villa.Description
	v.Address = villa.Address
	v.Max_capacity = villa.Max_capacity
	v.Price_per_night = setPrice
	v.Slug = villa.Slug
	v.Check_in = &villa.Check_in
	v.Check_out = &villa.Check_out
	v.Status = villa.Status

	if villa.Location != nil {
		v.Location = &VillaLocationResponse{
			Id:   villa.Location.Id.String(),
			Name: villa.Location.Name,
		}
	}
}
