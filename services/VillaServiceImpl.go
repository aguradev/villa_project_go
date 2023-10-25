package services

import (
	"errors"
	"villa_go/exceptions"
	"villa_go/models/schemas"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	uuid "github.com/satori/go.uuid"
)

type VillaService interface {
	VillaLists() ([]response.VillaListResponse, error)
	VillaDataDetail(slug string) (*response.VillaListResponse, error)
	CreateNewVilla(request.VillaRequest) (*response.VillaListResponse, []exceptions.ValidationMessage, error)
	DeleteDataVilla(uuid.UUID) (bool, error)
	UpdateDataVilla(request.VillaRequest, uuid.UUID) (*response.VillaListResponse, error)
}

type VillaServiceImpl struct {
	VillaRepository    repositories.VillaRepository
	LocationRepository repositories.VillaLocationRepository
	validation         validator.Validate
	translate          ut.Translator
}

func NewVillaServiceImplement(villaRepo repositories.VillaRepository, locationRepo repositories.VillaLocationRepository, validate validator.Validate, trans ut.Translator) VillaService {
	return &VillaServiceImpl{
		VillaRepository:    villaRepo,
		LocationRepository: locationRepo,
		validation:         validate,
		translate:          trans,
	}
}

func (v *VillaServiceImpl) VillaLists() ([]response.VillaListResponse, error) {

	GetVillaList, QueryException := v.VillaRepository.GetAllVilla()

	if QueryException != nil {
		return nil, QueryException
	}

	return GetVillaList, nil
}

func (v *VillaServiceImpl) VillaDataDetail(slug string) (*response.VillaListResponse, error) {

	GetVillaDetail, QueryException := v.VillaRepository.GetVillaBySlug(slug)

	if QueryException != nil {
		return nil, errors.New("Villa records not found")
	}

	return GetVillaDetail, nil
}

func (v *VillaServiceImpl) DeleteDataVilla(id uuid.UUID) (bool, error) {
	return false, nil
}

func (v *VillaServiceImpl) CreateNewVilla(requestData request.VillaRequest) (*response.VillaListResponse, []exceptions.ValidationMessage, error) {

	var VillaReq schemas.Villa

	VillaReq = schemas.Villa{
		Name:            requestData.Name,
		Slug:            slug.Make(requestData.Name),
		Description:     requestData.Description,
		Address:         requestData.Address,
		Max_capacity:    requestData.Max_capacity,
		Price_per_night: requestData.Price_per_night,
		Check_in:        utils.ConvertClockTime(requestData.Check_in),
		Check_out:       utils.ConvertClockTime(requestData.Check_out),
		Status:          "available",
	}

	if requestData.Location_id != nil {

		LocationRecord, Exists := v.LocationRepository.GetLocationById(*requestData.Location_id)

		if Exists != nil {
			return nil, nil, Exists
		}

		VillaReq.Location_id = &LocationRecord.Id
	}

	QueryCreate, QueryErrException := v.VillaRepository.CreateVilla(VillaReq)

	if QueryErrException != nil {
		return nil, nil, errors.New("Error when create new villa")
	}

	return QueryCreate, nil, nil
}

func (v *VillaServiceImpl) UpdateDataVilla(request request.VillaRequest, id uuid.UUID) (*response.VillaListResponse, error) {
	return nil, nil
}
