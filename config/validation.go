package config

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func InitValidation() (*validator.Validate, ut.Translator) {

	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	return validate, trans
}
