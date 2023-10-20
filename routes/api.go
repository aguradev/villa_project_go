package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func ApiRoutes(db *gorm.DB) {
	e := echo.New()

	Users := e.Group("/user")

	RoutesCredentials(db, Users)

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
