package routes

import (
	"villa_go/controllers"
	repositories "villa_go/repositories/Credentials"
	services "villa_go/services/Credentials"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoutesCredentials(db *gorm.DB, route *echo.Group) {

	CredentialRepository := repositories.NewCredentialRepository(db)
	CredentialService := services.CreateCredentialServiceImplement(CredentialRepository)

	controllers.CreateCredentialRoutes(CredentialService, route)
}
