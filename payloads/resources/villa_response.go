package resources

import (
	"strconv"
	"time"
	"villa_go/models/entities"
)

type VillaListResponse struct {
	Id              string                  `json:"id,omitempty"`
	Name            string                  `json:"name,omitempty"`
	Slug            string                  `json:"slug,omitempty"`
	Description     string                  `json:"description,omitempty"`
	Address         string                  `json:"address,omitempty"`
	Max_capacity    uint                    `json:"max_capacity,omitempty"`
	Price_per_night int                     `json:"price_per_night,omitempty"`
	Check_in        *time.Time              `json:"check_in,omitempty"`
	Check_out       *time.Time              `json:"check_out,omitempty"`
	Status          string                  `json:"status,omitempty"`
	Location        *VillaLocationResponse  `json:"location,omitempty"`
	Facility        []VillaFacilityResponse `json:"facilities,omitempty"`
	Galleries       []GalleryResource       `json:"gallery,omitempty"`
}

func SetVillaResponse(Villa []entities.Villa) []VillaListResponse {

	var SetVillaResponse []VillaListResponse

	for _, item := range Villa {

		var GetFacility []VillaFacilityResponse

		setPrice, err := strconv.Atoi(item.Price_per_night.String())

		if err != nil {
			setPrice = 0
		}

		for _, facility := range item.Facility {
			Facility := VillaFacilityResponse{
				Name: facility.Name,
			}

			GetFacility = append(GetFacility, Facility)
		}

		VillaResponse := VillaListResponse{
			Name:            item.Name,
			Slug:            item.Slug,
			Price_per_night: setPrice,
			Status:          item.Status,
			Location: &VillaLocationResponse{
				Name: item.Location.Name,
			},
			Facility: GetFacility,
		}

		SetVillaResponse = append(SetVillaResponse, VillaResponse)
	}

	return SetVillaResponse

}

func (v *VillaListResponse) GetVillaFacilitiesResponse(villa entities.Villa) {

	var GetFacility []VillaFacilityResponse

	v.Id = villa.Id.String()
	v.Name = villa.Name
	v.Description = villa.Description

	for _, facility := range villa.Facility {
		Facility := VillaFacilityResponse{
			Id:   facility.Id.String(),
			Name: facility.Name,
		}

		GetFacility = append(GetFacility, Facility)
	}

	v.Facility = GetFacility

}

func (v *VillaListResponse) SetVillaDetailResponse(villa entities.Villa) {
	setPrice, err := strconv.Atoi(villa.Price_per_night.String())
	var GetFacility []VillaFacilityResponse
	var GetGallery []GalleryResource

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
	v.Check_in = villa.Check_in
	v.Check_out = villa.Check_out
	v.Status = villa.Status

	if villa.Location != nil {
		v.Location = &VillaLocationResponse{
			Id:   villa.Location.Id.String(),
			Name: villa.Location.Name,
		}
	}

	for _, facility := range villa.Facility {
		Facility := VillaFacilityResponse{
			Id:   facility.Id.String(),
			Name: facility.Name,
		}

		GetFacility = append(GetFacility, Facility)
	}

	for _, gallery := range villa.Gallery {
		Gallery := GalleryResource{
			Id:      gallery.Id.String(),
			Fileurl: gallery.Fileurl,
		}
		GetGallery = append(GetGallery, Gallery)
	}

	v.Facility = GetFacility
	v.Galleries = GetGallery
}
