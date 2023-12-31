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

type VillaHandler interface {
	VillaListsHandler(echo.Context) error
	VillaDetailHandler(echo.Context) error
	CreateNewVillaHandler(echo.Context) error
	DeleteVillaHandler(echo.Context) error
}

type VillaHandlerImpl struct {
	VillaService services.VillaService
}

func NewVillaHandler(villaService services.VillaService) VillaHandler {
	return &VillaHandlerImpl{
		VillaService: villaService,
	}
}

func (v *VillaHandlerImpl) VillaListsHandler(ctx echo.Context) error {

	DataVilla, RecordException := v.VillaService.VillaLists()

	if RecordException != nil {
		return exceptions.NotFoundException(ctx, "Villa record is empty")
	}

	return response.HandleSuccess(ctx, DataVilla, "Retrieved data villa", http.StatusOK)

}

func (v *VillaHandlerImpl) VillaDetailHandler(ctx echo.Context) error {

	GetSlug := ctx.Param("slug")

	GetVillaDetail, RecordException := v.VillaService.VillaDataDetail(GetSlug)

	if RecordException != nil {
		return exceptions.NotFoundException(ctx, "Villa not found")
	}

	return response.HandleSuccess(ctx, GetVillaDetail, "Get Villa detail", http.StatusOK)

}

func (v *VillaHandlerImpl) CreateNewVillaHandler(ctx echo.Context) error {

	var Request request.VillaRequest

	BindingRequest := ctx.Bind(&Request)

	if BindingRequest != nil {
		return exceptions.BadRequestException(ctx, BindingRequest.Error())
	}

	ResponseNewVilla, Validation, QueryException, HttpStatus := v.VillaService.CreateNewVilla(ctx, Request)

	if Validation != nil {
		return exceptions.ValidationException(ctx, "One or more validation errors occurred", Validation)
	}

	if QueryException != nil {
		if HttpStatus == http.StatusInternalServerError {
			return exceptions.AppException(ctx, QueryException.Error())
		}
		if HttpStatus == http.StatusConflict {
			return exceptions.ConflictException(ctx, QueryException.Error())
		}
		if HttpStatus == http.StatusNoContent {
			return exceptions.NotFoundException(ctx, QueryException.Error())
		}
	}

	return response.HandleSuccess(ctx, ResponseNewVilla, "Villa created", http.StatusCreated)
}

func (v *VillaHandlerImpl) DeleteVillaHandler(ctx echo.Context) error {

	GetId, ParsingException := uuid.FromString(ctx.Param("id"))

	if ParsingException != nil {
		return exceptions.AppException(ctx, ParsingException.Error())
	}

	Deleted, QueryException := v.VillaService.DeleteDataVilla(GetId)

	if !Deleted && QueryException != nil {
		return exceptions.NotFoundException(ctx, QueryException.Error())
	}

	return response.HandleResponseDelete(ctx, "Villa Deleted")
}
