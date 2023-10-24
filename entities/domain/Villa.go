package domain

import (
	"villa_go/controllers"
	"villa_go/middlewares"
	"villa_go/repositories"
	"villa_go/services"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindingDepedencyVilla(db *gorm.DB, route *echo.Group, validate *validator.Validate, trans ut.Translator) {

	VillaRepo := repositories.NewVillaRepositoryImplement(db)
	LocationRepo := repositories.NewVillaLocationRepositoryImplement(db)

	VillaService := services.NewVillaServiceImplement(VillaRepo, LocationRepo, *validate, trans)
	VillaHandler := controllers.NewVillaController(VillaService)

	route.GET("/villa", VillaHandler.VillaListsHandler)
	route.GET("/villa/:slug", VillaHandler.VillaDetailHandler)
	route.POST("/villa", VillaHandler.CreateNewVillaHandler, middlewares.AccessbilityRole("Admin"))
}
