package modules

import (
	"villa_go/handlers"
	"villa_go/repositories"
	"villa_go/services"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindingDependencyCredentials(db *gorm.DB, route *echo.Group, validate *validator.Validate, trans ut.Translator) {
	CredentialRepository := repositories.NewCredentialRepository(db)
	CredentialService := services.CreateCredentialServiceImplement(CredentialRepository, validate, trans)
	CredentialController := handlers.CreateCredentialRoutes(CredentialService)

	route.POST("/register", CredentialController.RegisterUser)
	route.POST("/auth", CredentialController.AuthenticationUser)
}
