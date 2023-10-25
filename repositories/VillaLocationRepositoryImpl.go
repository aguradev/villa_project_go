package repositories

import (
	"errors"
	"villa_go/models/schemas"
	"villa_go/payloads/response"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VillaLocationRepository interface {
	GetAllLocation() ([]response.VillaLocationResponse, error)
	GetLocationById(id uuid.UUID) (*schemas.VillaLocation, error)
	CreateNewLocation(schemas.VillaLocation) (*schemas.VillaLocation, error)
	DeleteLocation(id uuid.UUID)
	UpdateLocation()
}

type VillaLocationRepositoryImpl struct {
	db *gorm.DB
}

func NewVillaLocationRepositoryImplement(Db *gorm.DB) VillaLocationRepository {
	return &VillaLocationRepositoryImpl{
		db: Db,
	}
}

func (l *VillaLocationRepositoryImpl) CreateNewLocation(location schemas.VillaLocation) (*schemas.VillaLocation, error) {

	QueryException := l.db.Table("location").Create(&location)

	if QueryException.Error != nil {
		return nil, errors.New("Error when create location")
	}

	return &location, nil

}

func (l *VillaLocationRepositoryImpl) GetAllLocation() ([]response.VillaLocationResponse, error) {

	var (
		Locations        []schemas.VillaLocation
		MappingLocations []response.VillaLocationResponse
	)

	LocationRecordException := l.db.Table("location").Find(&Locations)

	if LocationRecordException.RowsAffected == 0 {
		return nil, errors.New("Location records is empty")
	}

	if LocationRecordException.Error != nil {
		return nil, LocationRecordException.Error
	}

	MappingLocations = response.VillaLocationsReponses(Locations)

	return MappingLocations, nil

}

func (l *VillaLocationRepositoryImpl) GetLocationById(id uuid.UUID) (*schemas.VillaLocation, error) {
	var VillaLocation schemas.VillaLocation

	QueryRecordException := l.db.Table("location").Find(&VillaLocation, "id = ?", id)

	if QueryRecordException.RowsAffected == 0 {
		return nil, errors.New("Location does not exist")
	}

	return &VillaLocation, nil
}

func (l *VillaLocationRepositoryImpl) DeleteLocation(id uuid.UUID) {

}
func (l *VillaLocationRepositoryImpl) UpdateLocation() {

}
