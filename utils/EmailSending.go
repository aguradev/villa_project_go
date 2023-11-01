package utils

import (
	"bytes"
	"text/template"
	"villa_go/payloads/request"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendingEmail(to string, subject string, requestEmail request.ReservationEmailRequest) error {

	Results, ErrResult := ParseTemplateEmail("utils/html_template/transaction_success.html", requestEmail)

	if ErrResult != nil {
		return ErrResult
	}

	Message := gomail.NewMessage()
	Message.SetHeader("From", viper.GetString("email.USER"))
	Message.SetHeader("To", to)
	Message.SetHeader("Subject", subject)
	Message.SetBody("text/html", Results)

	c := gomail.NewDialer(viper.GetString("email.HOST"), viper.GetInt("email.PORT"), viper.GetString("email.USER"), viper.GetString("email.PASS"))

	if errSending := c.DialAndSend(Message); errSending != nil {
		return errSending
	}

	return nil

}

func ParseTemplateEmail(templateFileName string, data interface{}) (string, error) {

	t, errTemplate := template.ParseFiles(templateFileName)

	if errTemplate != nil {
		return "", errTemplate
	}

	buf := new(bytes.Buffer)

	if errExce := t.Execute(buf, data); errExce != nil {
		return "", errExce
	}

	return buf.String(), nil

}
