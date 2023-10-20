package routes

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ApiRoutes(db *gorm.DB) {
	e := echo.New()

	Users := e.Group("/user")

	RoutesCredentials(db, Users)

	e.Logger.Fatal(e.Start(":8080"))
}
