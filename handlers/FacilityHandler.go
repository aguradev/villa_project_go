package handlers

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/services"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type FacilityHandler interface {
	GetAllFacilityHandler(echo.Context) error
	CreateFacilityHandler(echo.Context) error
	AddFacilityToVillaHandler(echo.Context) error
	RemoveFacilityToVillaHandler(echo.Context) error
}

type FacilityHandlerImpl struct {
	FacilityService services.VIllaFacilityService
}

func NewFacilityHandlerImpl(facility services.VIllaFacilityService) FacilityHandler {
	return &FacilityHandlerImpl{
		FacilityService: facility,
	}
}

func (f *FacilityHandlerImpl) GetAllFacilityHandler(ctx echo.Context) error {

	GetListFacility, ErrMessage := f.FacilityService.GetListFacility()

	if ErrMessage != nil {
		return exceptions.NotFoundException(ctx, "Location record is empty")
	}

	return response.HandleSuccess(ctx, GetListFacility, "Retrieved data facility", http.StatusOK)

}

func (f *FacilityHandlerImpl) CreateFacilityHandler(ctx echo.Context) error {

	var Request request.FacilityRequest

	BindingErr := ctx.Bind(&Request)

	if BindingErr != nil {
		return exceptions.BadRequestException(ctx, BindingErr.Error())
	}

	CreateResponse, ValidationMessage, ErrMessage := f.FacilityService.CreateNewFacility(Request, ctx)

	if ValidationMessage != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ValidationMessage)
	}

	if ErrMessage != nil {
		return exceptions.AppException(ctx, ErrMessage.Error())
	}

	return response.HandleSuccess(ctx, CreateResponse, "Facility created", http.StatusCreated)

}

func (f *FacilityHandlerImpl) AddFacilityToVillaHandler(ctx echo.Context) error {

	GetVillaId := ctx.Param("villa_id")
	var request request.FacilityToVillaRequest

	ParseToUuid, UuidException := uuid.FromString(GetVillaId)

	if UuidException != nil {
		return exceptions.BadRequestException(ctx, "Invalid format uuid")
	}

	BindingErr := ctx.Bind(&request)

	if BindingErr != nil {
		return exceptions.BadRequestException(ctx, BindingErr.Error())
	}

	QueryResults, ValidationErr, ErrException := f.FacilityService.AddFacilityToVilla(ctx, ParseToUuid, request)

	if ValidationErr != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ValidationErr)
	}

	if ErrException != nil {
		return exceptions.AppException(ctx, ErrException.Error())
	}

	return response.HandleSuccess(ctx, QueryResults, "Success add facilities", http.StatusCreated)

}

func (f *FacilityHandlerImpl) RemoveFacilityToVillaHandler(ctx echo.Context) error {

	GetVillaId := ctx.Param("villa_id")
	var request request.FacilityToVillaRequest

	ParseToUuid, UuidException := uuid.FromString(GetVillaId)

	if UuidException != nil {
		return exceptions.BadRequestException(ctx, "Invalid format uuid")
	}

	BindingErr := ctx.Bind(&request)

	if BindingErr != nil {
		return exceptions.BadRequestException(ctx, BindingErr.Error())
	}

	Deleted, ValidationErr, ErrException := f.FacilityService.RemoveFacilityToVilla(ctx, ParseToUuid, request)

	if ValidationErr != nil && !Deleted {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ValidationErr)
	}

	if ErrException != nil && !Deleted {
		return exceptions.AppException(ctx, ErrException.Error())
	}

	return response.HandleResponseDelete(ctx, "Facilities Removed")

}
