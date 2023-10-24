package routes

import (
	"villa_go/config"
	"villa_go/entities/domain"
	"villa_go/middlewares"

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

	accessbility := api.Group("", middlewares.LoginSignedIn())

	domain.BindingDependencyCredentials(db, api, validate, trans)
	domain.BindingDepedencyVilla(db, accessbility, validate, trans)

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
