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

func BindingDepedencyVillaFacility(db *gorm.DB, route *echo.Group, validate *validator.Validate, trans ut.Translator) {

	FacilityRepo := repositories.NewVillaFacilaityRepositoryImpl(db)
	VillaRepo := repositories.NewVillaRepositoryImplement(db)

	FacilityService := services.NewVillaFacilityServiceImplement(FacilityRepo, VillaRepo, validate, trans)
	FacilityHandler := handlers.NewFacilityHandlerImpl(FacilityService)

	route.Group("", middlewares.AccessbilityRole("Admin"))
	route.GET("/facility", FacilityHandler.GetAllFacilityHandler)
	route.POST("/facility", FacilityHandler.CreateFacilityHandler)
	route.POST("/facility/:villa_id", FacilityHandler.AddFacilityToVillaHandler)
}
