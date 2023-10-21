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

	e.Use(middlewares.LoggerAccess())
	users := e.Group("/user")

	domain.BindingDependencyCredentials(db, users, validate, trans)

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
