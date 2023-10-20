package config

import "github.com/spf13/viper"

type Enviroment struct {
	AppName   string      `mapstructure:"appName"`
	Server    Server      `mapstructure:"server"`
	Databases DatabaseEnv `mapstructure:"db"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type DatabaseEnv struct {
	Env        string `mapstructure:"env"`
	DBhost     string `mapstructure:"DB_HOST"`
	DBuser     string `mapstructure:"DB_USER"`
	DBpassword string `mapstructure:"DB_PASS"`
	DBport     int    `mapstructure:"DB_PORT"`
	DBname     string `mapstructure:"DB_NAME"`
}

func InitEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	if ViperException := viper.ReadInConfig(); ViperException != nil {
		panic(ViperException)
	}
}
