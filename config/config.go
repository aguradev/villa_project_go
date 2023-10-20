package config

import "github.com/joho/godotenv"

func InitEnv() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Error to connect ENV")
	}

}
