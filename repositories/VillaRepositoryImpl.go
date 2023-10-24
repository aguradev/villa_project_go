package repositories

import (
	"errors"
	"villa_go/entities/models"
	"villa_go/payloads/request"
	"villa_go/payloads/response"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VillaRepository interface {
	GetAllVilla() ([]response.VillaListResponse, error)
	GetVillaBySlug(slug string) (*models.Villa, error)
	CreateVilla(models.Villa) (*response.VillaListResponse, error)
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

func (v *VillaRepositoryImpl) GetAllVilla() ([]response.VillaListResponse, error) {

	var items []models.Villa

	VillaRecordException := v.db.Table("properties_villa").Find(&items)

	if VillaRecordException.RowsAffected == 0 {
		return nil, errors.New("Villa record is empty")
	}

	MappingItems := response.SetVillaResponse(items)

	return MappingItems, nil
}

func (v *VillaRepositoryImpl) GetVillaBySlug(slug string) (*models.Villa, error) {

	var items models.Villa

	VillaRecordException := v.db.Table("properties_villa").First(&items, "slug = ?", slug)

	if VillaRecordException.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("Villa record not found")
	}

	return &items, nil

}

func (v *VillaRepositoryImpl) CreateVilla(request models.Villa) (*response.VillaListResponse, error) {

	var Result response.VillaListResponse

	VillaRecordException := v.db.Table("properties_villa").Create(&request)

	if VillaRecordException.Error != nil {
		return nil, errors.New("Failed to create villa")
	}

	Result.SetVillaDetailResponse(request)

	return &Result, nil
}

func (v *VillaRepositoryImpl) DeleteVilla(id uuid.UUID) (bool, error) {

	var items models.Villa

	VillaRecordException := v.db.Table("properties_villa").Delete(&items, id)

	if VillaRecordException.Error == gorm.ErrRecordNotFound {
		return false, errors.New("Villa record not found")
	}

	return true, nil
}

func (v *VillaRepositoryImpl) UpdateVilla(request request.VillaRequest, id uuid.UUID) (bool, error) {
	return false, nil
}