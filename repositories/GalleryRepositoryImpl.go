package repositories

import (
	"errors"
	"villa_go/models/entities"

	"gorm.io/gorm"
)

type GalleryRepository interface {
	CreateGallery(entities.Gallery) (*entities.Gallery, error)
}

type GalleryRepositoryImpl struct {
	Db *gorm.DB
}

func NewGalleryRepositoryImplement(db *gorm.DB) GalleryRepository {
	return &GalleryRepositoryImpl{
		Db: db,
	}
}

func (g *GalleryRepositoryImpl) CreateGallery(gallery entities.Gallery) (*entities.Gallery, error) {

	Transaction := g.Db.Begin()

	CreateErrException := Transaction.Table("galleries").Create(&gallery)

	if CreateErrException.Error != nil {
		Transaction.Rollback()
		return nil, errors.New("Error when create gallery")
	}

	Transaction.Commit()

	return &gallery, nil

}
