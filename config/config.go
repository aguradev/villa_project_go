package config

import "github.com/spf13/viper"

func InitEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	if ViperException := viper.ReadInConfig(); ViperException != nil {
		panic(ViperException)
	}
}
