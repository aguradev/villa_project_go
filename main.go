package main

import (
	"villa_go/config"
	"villa_go/routes"
)

func main() {
	config.InitEnv()

	db := config.Database()

	config.Migration(db)

	routes.ApiRoutes(db)
}
