package config

import "github.com/spf13/viper"

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

	PaymentClientKey := viper.UnmarshalKey("payment.CLIENT_KEY", &payment.MidtransClientKey)

	if PaymentClientKey != nil {
		panic(PaymentClientKey)
	}

	return payment
}
