package services

import (
	"mime/multipart"
	"villa_go/config"
	"villa_go/models/entities"
	"villa_go/payloads/resources"
	"villa_go/repositories"
	"villa_go/utils"

	uuid "github.com/satori/go.uuid"
)

type GalleryService interface {
	UploadImages([]*multipart.FileHeader, uuid.UUID) ([]resources.GalleryResource, error)
}

type GalleryServiceImpl struct {
	GalleryRepo repositories.GalleryRepository
	VillaRepo   repositories.VillaRepository
}

func GalleryServiceImplement(gallery repositories.GalleryRepository, villa repositories.VillaRepository) GalleryService {
	return &GalleryServiceImpl{
		GalleryRepo: gallery,
		VillaRepo:   villa,
	}
}

func (g *GalleryServiceImpl) UploadImages(images []*multipart.FileHeader, villaId uuid.UUID) ([]resources.GalleryResource, error) {

	var galleries []entities.Gallery

	Villas, ErrNotExists := g.VillaRepo.CheckVillaIsExists(villaId)

	if ErrNotExists != nil {
		return nil, ErrNotExists
	}

	for _, file := range images {

		var gallery entities.Gallery

		client := config.ConfigStorage()

		uploadImagesURL, errUpload := utils.UploadImageFile(file, client)

		if errUpload != nil {
			return nil, errUpload
		}

		gallery.Villa_id = Villas.Id
		gallery.Fileurl = uploadImagesURL

		CreateGalleryData, ErrCreateMessage := g.GalleryRepo.CreateGallery(gallery)

		if ErrCreateMessage != nil {
			return nil, ErrCreateMessage
		}

		galleries = append(galleries, *CreateGalleryData)
	}

	ResponseGallery := resources.GetListsGalleryResponse(galleries)

	return ResponseGallery, nil

}
