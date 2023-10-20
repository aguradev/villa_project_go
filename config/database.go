package config

import (
	"fmt"
	"os"
	"villa_go/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database failed connect " + err.Error())
	}

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(entities.Users{}, entities.Admin{}, entities.Credentials{}, entities.Roles{})
}
