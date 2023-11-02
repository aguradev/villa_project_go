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

func BindingDependencyReservation(db *gorm.DB, api *echo.Group, validate *validator.Validate, trans ut.Translator) {
	MidtransConfig := config.InitPaymentENV()

	VillaRepo := repositories.NewVillaRepositoryImplement(db)
	ReservationRepo := repositories.NewReservationRepositoryImpl(db)
	CredentialRepo := repositories.NewCredentialRepository(db)

	MidtransService := services.NewMidtransServiceImpl(MidtransConfig, ReservationRepo)
	ReservationService := services.NewReservationServiceImplement(ReservationRepo, MidtransService, VillaRepo, CredentialRepo, validate, trans)

	ReservationHandler := handlers.NewReservationHandler(ReservationService, MidtransService)

	verifyJWT := api.Group("", middlewares.VerifiyTokenByCookie())
	AdminAccess := verifyJWT.Group("", middlewares.AccessbilityRole("Admin"))
	UserAccess := verifyJWT.Group("", middlewares.AccessbilityRole("User"))

	AdminAccess.GET("/reservation", ReservationHandler.GetAllReservationHandler)
	AdminAccess.GET("/reservation/:villa_id", ReservationHandler.GetReservationByIdHandler)
	UserAccess.POST("/reservation", ReservationHandler.CreateReservationHandler)
	UserAccess.GET("/reservation/user", ReservationHandler.GetReservationDataByUserLogin)
	api.POST("/reservation/callback", ReservationHandler.NotificationReservationHandler)
}
