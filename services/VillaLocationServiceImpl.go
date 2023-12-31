package services

import (
	"errors"
	"villa_go/exceptions"
	"villa_go/models/schemas"
	"villa_go/payloads/request"
	"villa_go/payloads/resources"
	"villa_go/repositories"
	"villa_go/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type VillaLocationService interface {
	ListsDataLocation() ([]resources.VillaLocationResponse, error)
	CreateNewLocation(echo.Context, request.LocationRequest) ([]resources.VillaLocationResponse, []exceptions.ValidationMessage, error)
}

type VillaLocationServiceImpl struct {
	location   repositories.VillaLocationRepository
	validation validator.Validate
	trans      ut.Translator
}

func NewVillaLocationServiceImplement(VillaLocation repositories.VillaLocationRepository, Validation validator.Validate, Translator ut.Translator) VillaLocationService {
	return &VillaLocationServiceImpl{
		location:   VillaLocation,
		validation: Validation,
		trans:      Translator,
	}
}

func (l *VillaLocationServiceImpl) ListsDataLocation() ([]resources.VillaLocationResponse, error) {

	LocationRecords, QueryException := l.location.GetAllLocation()

	if QueryException != nil {
		return nil, QueryException
	}

	return LocationRecords, nil
}

func (l *VillaLocationServiceImpl) CreateNewLocation(ctx echo.Context, request request.LocationRequest) ([]resources.VillaLocationResponse, []exceptions.ValidationMessage, error) {

	var Locations []schemas.VillaLocation
	var MappingLocations []resources.VillaLocationResponse

	ValidationException := l.validation.Struct(request)

	if ValidationException != nil {
		return nil, utils.ValidationError(ctx, l.trans, ValidationException), nil
	}

	for index := range request.Name {

		var VillaLocation schemas.VillaLocation
		VillaLocation.Name = request.Name[index]

		LocationRecord, CreateQueryException := l.location.CreateNewLocation(VillaLocation)

		if CreateQueryException != nil {
			return nil, nil, errors.New("Error when create new location")
		}

		Locations = append(Locations, *LocationRecord)

	}

	MappingLocations = resources.VillaLocationsReponses(Locations)

	return MappingLocations, nil, nil

}
