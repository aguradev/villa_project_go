package utils

import (
	"villa_go/exceptions"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidationError(ctx echo.Context, trans ut.Translator, err error) []exceptions.ValidationMessage {

	validationError, validate := err.(validator.ValidationErrors)

	if validate {
		ValidationMessages := []exceptions.ValidationMessage{}
		for _, e := range validationError {

			Message := exceptions.ValidationMessage{
				Field:   e.Field(),
				Message: e.Translate(trans),
			}

			ValidationMessages = append(ValidationMessages, Message)
		}

		return ValidationMessages
	}

	return nil

}
