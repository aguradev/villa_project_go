package modules

import (
	"villa_go/handlers"
	"villa_go/middlewares"
	"villa_go/repositories"
	"villa_go/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindingDepedencyGallery(db *gorm.DB, route *echo.Group) {

	GalleryRepo := repositories.NewGalleryRepositoryImplement(db)
	VillaRepo := repositories.NewVillaRepositoryImplement(db)
	GalleryService := services.GalleryServiceImplement(GalleryRepo, VillaRepo)
	GalleryHandler := handlers.GalleryHandlerImplement(GalleryService)

	AdminAccess := route.Group("", middlewares.AccessbilityRole("Admin"))

	AdminAccess.POST("/gallery/upload-images/:villa_id", GalleryHandler.UploadHandler)

}
