package domain

import (
	"villa_go/controllers"
	repositories "villa_go/repositories/Credentials"
	services "villa_go/services/Credentials"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindingDependencyCredentials(db *gorm.DB, route *echo.Group) {

	CredentialRepository := repositories.NewCredentialRepository(db)
	CredentialService := services.CreateCredentialServiceImplement(CredentialRepository)
	CredentialController := controllers.CreateCredentialRoutes(CredentialService)

	route.POST("/register", CredentialController.RegisterUser)
	route.POST("/auth", CredentialController.AuthenticationUser)
}
