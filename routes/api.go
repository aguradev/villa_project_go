package routes

import (
	"villa_go/entities/domain"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func ApiRoutes(db *gorm.DB) {
	e := echo.New()

	users := e.Group("/user")

	domain.BindingDependencyCredentials(db, users)

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
