package repositories

import (
	"errors"
	"villa_go/models/entities"
	"villa_go/payloads/resources"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VillaFacilityRepository interface {
	CreateNewFacility(entities.Facility) (*entities.Facility, error)
	GetFacilityById(uuid.UUID) (*entities.Facility, error)
	GetAllFacility() ([]resources.VillaFacilityResponse, error)
}

type VillaFacilityRepositoryImpl struct {
	Db *gorm.DB
}

func NewVillaFacilaityRepositoryImpl(db *gorm.DB) VillaFacilityRepository {
	return &VillaFacilityRepositoryImpl{
		Db: db,
	}
}

func (f *VillaFacilityRepositoryImpl) CreateNewFacility(facility entities.Facility) (*entities.Facility, error) {

	Transaction := f.Db.Begin()

	CreateQueryErr := Transaction.Create(&facility)

	if CreateQueryErr.Error != nil {
		Transaction.Rollback()
		return nil, errors.New("Error when create facility")
	}

	Transaction.Commit()

	return &facility, nil

}

func (f *VillaFacilityRepositoryImpl) GetAllFacility() ([]resources.VillaFacilityResponse, error) {

	var Facility []entities.Facility

	GetAllQueryErr := f.Db.Table("facilities").Find(&Facility)

	if GetAllQueryErr.RowsAffected == 0 {
		return nil, errors.New("Facility records is empty")
	}

	GetListsFacilities := resources.VillaFacilityListResponse(Facility)

	return GetListsFacilities, nil

}

func (f *VillaFacilityRepositoryImpl) GetFacilityById(id uuid.UUID) (*entities.Facility, error) {
	var facility entities.Facility

	FindQueryException := f.Db.Table("facilities").First(&facility, "id = ?", id)

	if FindQueryException.Error != nil {
		return nil, errors.New("facility does not exists")
	}

	return &facility, nil

}
