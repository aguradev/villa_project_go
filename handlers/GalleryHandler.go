package handlers

import (
	"net/http"
	"villa_go/exceptions"
	"villa_go/payloads/response"
	"villa_go/services"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type GalleryHandler interface {
	UploadHandler(echo.Context) error
}

type GalleryHandlerImpl struct {
	GalleryService services.GalleryService
}

func GalleryHandlerImplement(gallery services.GalleryService) GalleryHandler {

	return &GalleryHandlerImpl{
		GalleryService: gallery,
	}

}

func (g *GalleryHandlerImpl) UploadHandler(ctx echo.Context) error {

	GetId := ctx.Param("villa_id")

	IdUUID, errParsing := uuid.FromString(GetId)

	if errParsing != nil {
		return exceptions.BadRequestException(ctx, errParsing.Error())
	}

	formUploads, errFiles := ctx.MultipartForm()

	if errFiles != nil {
		return exceptions.BadRequestException(ctx, errFiles.Error())
	}

	ResultUpload, errUpload := g.GalleryService.UploadImages(formUploads.File["images"], IdUUID)

	if errUpload != nil {
		return exceptions.AppException(ctx, errUpload.Error())
	}

	return response.HandleSuccess(ctx, ResultUpload, "Success upload image", http.StatusOK)
}
