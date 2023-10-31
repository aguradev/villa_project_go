package services

import (
	"errors"
	"villa_go/exceptions"
	"villa_go/models/entities"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type VIllaFacilityService interface {
	CreateNewFacility(request.FacilityRequest, echo.Context) ([]resources.VillaFacilityResponse, []exceptions.ValidationMessage, error)
	GetListFacility() ([]resources.VillaFacilityResponse, error)
	AddFacilityToVilla(echo.Context, uuid.UUID, request.FacilityToVillaRequest) (*resources.VillaListResponse, []exceptions.ValidationMessage, error)
	RemoveFacilityToVilla(echo.Context, uuid.UUID, request.FacilityToVillaRequest) (bool, []exceptions.ValidationMessage, error)
}

type VillaFacilityServiceImpl struct {
	FacilityRepo repositories.VillaFacilityRepository
	VillaRepo    repositories.VillaRepository
	Validator    *validator.Validate
	Trans        ut.Translator
}

func NewVillaFacilityServiceImplement(facility repositories.VillaFacilityRepository, villa repositories.VillaRepository, validator *validator.Validate, trans ut.Translator) VIllaFacilityService {
	return &VillaFacilityServiceImpl{
		FacilityRepo: facility,
		VillaRepo:    villa,
		Validator:    validator,
		Trans:        trans,
	}
}

func (f *VillaFacilityServiceImpl) CreateNewFacility(facilityRequest request.FacilityRequest, ctx echo.Context) ([]resources.VillaFacilityResponse, []exceptions.ValidationMessage, error) {

	var Facilities []entities.Facility

	ValidationException := f.Validator.Struct(facilityRequest)

	if ValidationException != nil {
		return nil, utils.ValidationError(ctx, f.Trans, ValidationException), nil
	}

	for index := range facilityRequest.Name {

		Facility := entities.Facility{
			Name: facilityRequest.Name[index],
		}

		GetResultCreateFacility, ErrCreate := f.FacilityRepo.CreateNewFacility(Facility)

		if ErrCreate != nil {
			return nil, nil, errors.New("Error when create new facility")
		}

		Facilities = append(Facilities, *GetResultCreateFacility)

	}

	MappingFacilities := resources.VillaFacilityListResponse(Facilities)

	return MappingFacilities, nil, nil

}

func (f *VillaFacilityServiceImpl) GetListFacility() ([]resources.VillaFacilityResponse, error) {

	GetFacilities, QueryListErr := f.FacilityRepo.GetAllFacility()

	if QueryListErr != nil {
		return nil, QueryListErr
	}

	return GetFacilities, nil

}

func (f *VillaFacilityServiceImpl) AddFacilityToVilla(ctx echo.Context, villaId uuid.UUID, requestFacility request.FacilityToVillaRequest) (*resources.VillaListResponse, []exceptions.ValidationMessage, error) {

	var Facilities []entities.Facility

	ValidationMessage := f.Validator.Struct(ctx)

	if ValidationMessage != nil {
		return nil, utils.ValidationError(ctx, f.Trans, ValidationMessage), nil
	}

	VillaData, ErrExists := f.VillaRepo.CheckVillaIsExists(villaId)

	if ErrExists != nil {
		return nil, nil, ErrExists
	}

	for index := range requestFacility.Id {

		ParseToUuid, ErrParse := uuid.FromString(requestFacility.Id[index])

		if ErrParse != nil {
			return nil, nil, errors.New("Invalid format uuid")
		}

		FacilityExists, ErrExists := f.FacilityRepo.GetFacilityById(ParseToUuid)

		if ErrExists != nil {
			return nil, nil, ErrExists
		}

		Facilities = append(Facilities, *FacilityExists)

	}

	GetVillaWithFacility, ErrAddFacility := f.VillaRepo.AddFacilities(*VillaData, Facilities)

	if ErrAddFacility != nil {
		return nil, nil, ErrAddFacility
	}

	return GetVillaWithFacility, nil, nil

}

func (f *VillaFacilityServiceImpl) RemoveFacilityToVilla(ctx echo.Context, id uuid.UUID, requestFacility request.FacilityToVillaRequest) (bool, []exceptions.ValidationMessage, error) {

	var Facilities []entities.Facility

	ValidationMessage := f.Validator.Struct(ctx)

	if ValidationMessage != nil {
		return false, utils.ValidationError(ctx, f.Trans, ValidationMessage), nil
	}

	VillaData, ErrExists := f.VillaRepo.CheckVillaIsExists(id)

	if ErrExists != nil {
		return false, nil, ErrExists
	}

	if len(requestFacility.Id) == 0 {
		return false, nil, errors.New("No id in request")
	}

	for index := range requestFacility.Id {

		ParseToUuid, ErrParse := uuid.FromString(requestFacility.Id[index])

		if ErrParse != nil {
			return false, nil, errors.New("Invalid format uuid")
		}

		FacilityExists, ErrExists := f.FacilityRepo.GetFacilityById(ParseToUuid)

		if ErrExists != nil {
			return false, nil, ErrExists
		}

		Facilities = append(Facilities, *FacilityExists)

	}

	IsDeleted, ErrAddFacility := f.VillaRepo.RemoveFacilities(*VillaData, Facilities)

	if ErrAddFacility != nil && !IsDeleted {
		return false, nil, ErrAddFacility
	}

	return true, nil, nil

}
