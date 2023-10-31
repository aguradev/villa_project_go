package repositories

import (
	"errors"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VillaRepository interface {
	GetAllVilla() ([]resources.VillaListResponse, error)
	CheckVillaIsExists(uuid.UUID) (*entities.Villa, error)
	GetVillaBySlug(slug string) (*resources.VillaListResponse, error)
	CreateVilla(entities.Villa) (*resources.VillaListResponse, error)
	DeleteVilla(uuid.UUID) (bool, error)
	UpdateVilla(request.VillaRequest, uuid.UUID) (bool, error)
	AddFacilities(entities.Villa, []entities.Facility) (*resources.VillaListResponse, error)
	CheckVillaNameExists(name string) (bool, error)
}

type VillaRepositoryImpl struct {
	db *gorm.DB
}

func NewVillaRepositoryImplement(Db *gorm.DB) VillaRepository {
	return &VillaRepositoryImpl{
		db: Db,
	}
}

func (v *VillaRepositoryImpl) GetAllVilla() ([]resources.VillaListResponse, error) {

	var items []entities.Villa

	VillaRecordException := v.db.Table("properties_villa").Joins("Location").Preload("Facility").Find(&items)

	if VillaRecordException.RowsAffected == 0 {
		return nil, errors.New("Villa records is empty")
	}

	MappingItems := resources.SetVillaResponse(items)

	return MappingItems, nil
}

func (v *VillaRepositoryImpl) GetVillaBySlug(slug string) (*resources.VillaListResponse, error) {

	var items entities.Villa
	var LocationDetail resources.VillaListResponse

	VillaRecordException := v.db.Table("properties_villa").Joins("Location").Preload("Facility").First(&items, "slug = ?", slug)

	if VillaRecordException.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("Villa record not found")
	}

	LocationDetail.SetVillaDetailResponse(items)

	return &LocationDetail, nil

}

func (v *VillaRepositoryImpl) CreateVilla(request entities.Villa) (*resources.VillaListResponse, error) {

	var Result resources.VillaListResponse

	VillaRecordException := v.db.Table("properties_villa").Create(&request)

	if VillaRecordException.Error != nil {
		return nil, errors.New("Failed to create villa")
	}

	Result.SetVillaDetailResponse(request)

	return &Result, nil
}

func (v *VillaRepositoryImpl) DeleteVilla(id uuid.UUID) (bool, error) {

	var items entities.Villa

	if VillaRecordException := v.db.Table("properties_villa").First(&items, "id = ?", id); VillaRecordException.Error != nil {
		if VillaRecordException.Error == gorm.ErrRecordNotFound {
			return false, errors.New("Villa record not exist")
		}

		return false, VillaRecordException.Error
	}

	if DeleteException := v.db.Delete(&items); DeleteException.Error != nil {
		return false, DeleteException.Error
	}

	return true, nil
}

func (v *VillaRepositoryImpl) UpdateVilla(request request.VillaRequest, id uuid.UUID) (bool, error) {
	return false, nil
}

func (v *VillaRepositoryImpl) CheckVillaIsExists(id uuid.UUID) (*entities.Villa, error) {

	var items entities.Villa

	if CheckIsExist := v.db.Table("properties_villa").First(&items, "id = ?", id); CheckIsExist.Error != nil {
		return nil, errors.New("Villa not found")
	}

	return &items, nil

}

func (v *VillaRepositoryImpl) AddFacilities(villa entities.Villa, facilities []entities.Facility) (*resources.VillaListResponse, error) {

	var VillaFacility entities.Villa
	var Response resources.VillaListResponse

	Transcation := v.db.Begin()

	ErrAddedFacility := Transcation.Model(&villa).Association("Facility").Append(facilities)

	if ErrAddedFacility != nil {
		Transcation.Rollback()
		return nil, errors.New("Error when add facility")
	}

	Transcation.Commit()

	GetVillaByIdException := v.db.Table("properties_villa").Preload("Facility").First(&VillaFacility, "id = ?", villa.Id)

	if GetVillaByIdException.Error != nil {
		return nil, errors.New("Villa does not exist")
	}

	Response.GetVillaFacilitiesResponse(VillaFacility)

	return &Response, nil

}

func (v *VillaRepositoryImpl) CheckVillaNameExists(name string) (bool, error) {

	var villa entities.Villa

	CheckVillaIsExists := v.db.First(&villa, "name = ?", name)

	if CheckVillaIsExists.RowsAffected > 0 {
		return true, errors.New("Villa name already exists")
	}

	return false, nil

}
