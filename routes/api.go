package routes

import (
	"net/http"
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

	api := e.Group("api", middlewares.LoggerAccess())

	users := api.Group("/user", middlewares.VerifyToken(), middlewares.AccessbilityRole("User"))
	admin := api.Group("/admin")

	domain.BindingDependencyCredentials(db, api, validate, trans)

	users.GET("/greeting", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello User")
	})

	admin.GET("/greeting", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Admin")
	})

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
