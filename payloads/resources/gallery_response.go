package resources

import "villa_go/models/entities"

type GalleryResource struct {
	Id      string `json:"id"`
	Fileurl string `json:"image"`
}

func GetListsGalleryResponse(galleries []entities.Gallery) []GalleryResource {

	var GalleriesResponse []GalleryResource

	for _, val := range galleries {
		Gallery := GalleryResource{
			Id:      val.Id.String(),
			Fileurl: val.Fileurl,
		}

		GalleriesResponse = append(GalleriesResponse, Gallery)
	}

	return GalleriesResponse
}
