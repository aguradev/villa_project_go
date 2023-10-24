package controllers

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/services"

	"github.com/labstack/echo/v4"
)

type LocationController interface {
	ListsLocationHandler(echo.Context) error
	CreateNewLocationHandler(echo.Context) error
}

type LocationHandlerImpl struct {
	locationService services.VillaLocationService
}

func NewLocationHandlerImpl(LocationService services.VillaLocationService) LocationController {
	return &LocationHandlerImpl{
		locationService: LocationService,
	}
}

func (h *LocationHandlerImpl) ListsLocationHandler(ctx echo.Context) error {

	ListLocations, RecordException := h.locationService.ListsDataLocation()

	if RecordException != nil {
		return exceptions.NotFoundException(ctx, "Location record is empty")
	}

	return response.HandleSuccess(ctx, ListLocations, "Retrieved data location", http.StatusOK)

}

func (h *LocationHandlerImpl) CreateNewLocationHandler(ctx echo.Context) error {

	var request request.LocationRequest

	BindingRequest := ctx.Bind(&request)

	if BindingRequest != nil {
		return exceptions.BadRequestException(ctx, BindingRequest.Error())
	}

	CreateResponse, ValidationException, ErrException := h.locationService.CreateNewLocation(ctx, request)

	if ValidationException != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", ValidationException)
	}

	if ErrException != nil {
		return exceptions.AppException(ctx, ErrException.Error())
	}

	return response.HandleSuccess(ctx, CreateResponse, "Location created", http.StatusCreated)

}
