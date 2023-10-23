package domain

import (
	"villa_go/controllers"
	"villa_go/repositories"
	"villa_go/services"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindingDepedencyVilla(db *gorm.DB, route *echo.Group, validate *validator.Validate, trans ut.Translator) {

	VillaRepo := repositories.NewVillaRepositoryImplement(db)
	VillaService := services.NewVillaServiceImplement(VillaRepo)
	VillaHandler := controllers.NewVillaController(VillaService)

	route.GET("/villa", VillaHandler.VillaListsHandler)
	route.POST("/villa/create", VillaHandler.CreateNewVillaHandler)
}
