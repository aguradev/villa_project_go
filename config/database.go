package config

import (
	"fmt"
	"villa_go/entities"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v dbname=%v port=%v sslmode=disable",
		viper.GetString("db.DB_HOST"),
		viper.GetString("db.DB_USER"),
		viper.GetString("db.DB_NAME"),
		viper.GetString("db.DB_PORT"),
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
