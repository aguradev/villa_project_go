package services

import (
	"errors"
	"net/http"
	"strings"
	"villa_go/exceptions"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type VillaService interface {
	VillaLists() ([]resources.VillaListResponse, error)
	VillaDataDetail(slug string) (*resources.VillaListResponse, error)
	CreateNewVilla(echo.Context, request.VillaRequest) (*resources.VillaListResponse, []exceptions.ValidationMessage, error, int)
	DeleteDataVilla(uuid.UUID) (bool, error)
	UpdateDataVilla(request.VillaRequest, uuid.UUID) (*resources.VillaListResponse, error)
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

func (v *VillaServiceImpl) VillaLists() ([]resources.VillaListResponse, error) {

	GetVillaList, QueryException := v.VillaRepository.GetAllVilla()

	if QueryException != nil {
		return nil, QueryException
	}

	return GetVillaList, nil
}

func (v *VillaServiceImpl) VillaDataDetail(slug string) (*resources.VillaListResponse, error) {

	GetVillaDetail, QueryException := v.VillaRepository.GetVillaBySlug(slug)

	if QueryException != nil {
		return nil, errors.New("Villa records not found")
	}

	return GetVillaDetail, nil
}

func (v *VillaServiceImpl) DeleteDataVilla(id uuid.UUID) (bool, error) {

	isDeleted, QueryException := v.VillaRepository.DeleteVilla(id)

	if !isDeleted {
		if strings.Contains(QueryException.Error(), "Villa record not exist") {
			return false, errors.New("Villa record not exist")
		}

		return false, QueryException
	}

	return true, nil
}

func (v *VillaServiceImpl) CreateNewVilla(ctx echo.Context, requestData request.VillaRequest) (*resources.VillaListResponse, []exceptions.ValidationMessage, error, int) {

	var VillaReq entities.Villa

	ValidationMessage := v.validation.Struct(requestData)

	if ValidationMessage != nil {
		return nil, utils.ValidationError(ctx, v.translate, ValidationMessage), nil, http.StatusUnprocessableEntity
	}

	VillaReq = entities.Villa{
		Name:            requestData.Name,
		Slug:            slug.Make(requestData.Name),
		Description:     requestData.Description,
		Address:         requestData.Address,
		Max_capacity:    requestData.Max_capacity,
		Price_per_night: &requestData.Price_per_night,
		Check_in:        utils.ConvertClockTime(requestData.Check_in),
		Check_out:       utils.ConvertClockTime(requestData.Check_out),
		Status:          "available",
	}

	IsExists, ExistMessage := v.VillaRepository.CheckVillaNameExists(VillaReq.Name)

	if IsExists {
		return nil, nil, ExistMessage, http.StatusConflict
	}

	if requestData.Location_id != nil {

		LocationRecord, Exists := v.LocationRepository.GetLocationById(*requestData.Location_id)

		if Exists != nil {
			return nil, nil, Exists, http.StatusNoContent
		}

		VillaReq.Location_id = &LocationRecord.Id
	}

	QueryCreate, QueryErrException := v.VillaRepository.CreateVilla(VillaReq)

	if QueryErrException != nil {
		return nil, nil, errors.New("Error when create new villa"), http.StatusInternalServerError
	}

	return QueryCreate, nil, nil, http.StatusCreated
}

func (v *VillaServiceImpl) UpdateDataVilla(request request.VillaRequest, id uuid.UUID) (*resources.VillaListResponse, error) {
	return nil, nil
}
