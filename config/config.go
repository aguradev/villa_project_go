package config

import (
	"github.com/spf13/viper"
)

type AuthConfig struct {
	SecretKey []byte
}

type PaymentGatewayConfig struct {
	MidtransClientKey string
}

func InitEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	if ViperException := viper.ReadInConfig(); ViperException != nil {
		panic(ViperException)
	}
}

func InitPaymentENV() *PaymentGatewayConfig {
	payment := &PaymentGatewayConfig{}

	viper.SetConfigName("")

	PaymentClientKey := viper.UnmarshalKey("payment.SERVER_KEY", &payment.MidtransClientKey)

	if PaymentClientKey != nil {
		panic(PaymentClientKey)
	}

	return payment
}

func InitAuthConfig() *AuthConfig {

	AuthConfig := &AuthConfig{}

	AuthConfig.SecretKey = []byte(viper.GetString("SECRET_KEY"))

	return AuthConfig

}
