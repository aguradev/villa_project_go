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

	VillaRecordException := v.db.Table("properties_villa").Joins("Location").Find(&items)

	if VillaRecordException.RowsAffected == 0 {
		return nil, errors.New("Villa records is empty")
	}

	MappingItems := resources.SetVillaResponse(items)

	return MappingItems, nil
}

func (v *VillaRepositoryImpl) GetVillaBySlug(slug string) (*resources.VillaListResponse, error) {

	var items entities.Villa
	var LocationDetail resources.VillaListResponse

	VillaRecordException := v.db.Table("properties_villa").Joins("Location").First(&items, "slug = ?", slug)

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
