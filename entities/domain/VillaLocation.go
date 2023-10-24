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

func BindingDepedencyVillaLocation(db *gorm.DB, route *echo.Group, validate *validator.Validate, trans ut.Translator) {
	LocationRepo := repositories.NewVillaLocationRepositoryImplement(db)
	LocationService := services.NewVillaLocationServiceImplement(LocationRepo, *validate, trans)
	LocationHandler := controllers.NewLocationHandlerImpl(LocationService)

	route.Use(middlewares.AccessbilityRole("Admin"))
	route.GET("/location", LocationHandler.ListsLocationHandler)
	route.POST("/location", LocationHandler.CreateNewLocationHandler)
}
