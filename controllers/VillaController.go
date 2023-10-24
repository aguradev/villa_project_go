package controllers

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/request"
	"villa_go/payloads/response"
	"villa_go/services"

	"github.com/labstack/echo/v4"
)

type VillaController interface {
	VillaListsHandler(echo.Context) error
	VillaDetailHandler(echo.Context) error
	CreateNewVillaHandler(echo.Context) error
}

type VillaControllerImpl struct {
	VillaService services.VillaService
}

func NewVillaController(villaService services.VillaService) VillaController {
	return &VillaControllerImpl{
		VillaService: villaService,
	}
}

func (v *VillaControllerImpl) VillaListsHandler(ctx echo.Context) error {

	DataVilla, RecordException := v.VillaService.VillaLists()

	if RecordException != nil {
		return exceptions.NotFoundException(ctx, "Villa record is empty")
	}

	return response.HandleSuccess(ctx, DataVilla, "Retrieved data villa", http.StatusOK)

}

func (v *VillaControllerImpl) VillaDetailHandler(ctx echo.Context) error {

	GetSlug := ctx.Param("slug")

	GetVillaDetail, RecordException := v.VillaService.VillaDataDetail(GetSlug)

	if RecordException != nil {
		return exceptions.NotFoundException(ctx, "Villa not found")
	}

	return response.HandleSuccess(ctx, GetVillaDetail, "Get Villa Detail", http.StatusOK)

}

func (v *VillaControllerImpl) CreateNewVillaHandler(ctx echo.Context) error {

	var Request request.VillaRequest

	BindingRequest := ctx.Bind(&Request)

	if BindingRequest != nil {
		return exceptions.BadRequestException(ctx, BindingRequest.Error())
	}

	ResponseNewVilla, Validation, QueryException := v.VillaService.CreateNewVilla(Request)

	if Validation != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", Validation)
	}

	if QueryException != nil {
		return exceptions.AppException(ctx, QueryException.Error())
	}

	return response.HandleSuccess(ctx, ResponseNewVilla, "Villa created", http.StatusCreated)
}
