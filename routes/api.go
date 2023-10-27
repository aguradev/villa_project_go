package routes

import (
	"villa_go/config"
	"villa_go/middlewares"
	"villa_go/modules"

	"github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func ApiRoutes(db *gorm.DB) {
	e := echo.New()
	validate, trans := config.InitValidation()
	en.RegisterDefaultTranslations(validate, trans)

	api := e.Group("/api", middlewares.LoggerAccess())

	verifyJWT := api.Group("", middlewares.LoginSignedIn())

	modules.BindingDependencyCredentials(db, api, validate, trans)
	modules.BindingDepedencyVilla(db, verifyJWT, validate, trans)
	modules.BindingDependencyReservation(db, api, verifyJWT, validate, trans)
	modules.BindingDepedencyVillaLocation(db, verifyJWT, validate, trans)

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
