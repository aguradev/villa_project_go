package modules

import (
	"villa_go/config"
	"villa_go/handlers"
	"villa_go/middlewares"
	"villa_go/repositories"
	"villa_go/services"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindingDependencyReservation(db *gorm.DB, api *echo.Group, route *echo.Group, validate *validator.Validate, trans ut.Translator) {

	VillaRepo := repositories.NewVillaRepositoryImplement(db)
	ReservationRepo := repositories.NewReservationRepositoryImpl(db)
	CredentialRepo := repositories.NewCredentialRepository(db)

	MidtransConfig := config.InitPaymentENV()

	MidtransService := services.NewMidtransServiceImpl(MidtransConfig, ReservationRepo)
	ReservationService := services.NewReservationServiceImplement(ReservationRepo, MidtransService, VillaRepo, CredentialRepo)

	ReservationHandler := handlers.NewReservationHandler(ReservationService, MidtransService)

	route.POST("/reservation", ReservationHandler.CreateReservationHandler, middlewares.AccessbilityRole("User"))
	api.POST("/reservation/callback", ReservationHandler.NotificationReservationHandler)
}
