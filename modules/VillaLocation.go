package modules

import (
	"villa_go/handlers"
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
	LocationHandler := handlers.NewLocationHandlerImpl(LocationService)

	AdminAccess := route.Group("", middlewares.AccessbilityRole("Admin"))

	AdminAccess.GET("/location", LocationHandler.ListsLocationHandler)
	AdminAccess.POST("/location", LocationHandler.CreateNewLocationHandler)
}
