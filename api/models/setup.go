package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDataBase() {
	dbHost := viper.Get("POSTGRES_HOST")
	dbPort := viper.Get("POSTGRES_PORT")
	dbUser := viper.Get("POSTGRES_USER")
	dbPassword := viper.Get("POSTGRES_PASSWORD")
	dbName := viper.Get("POSTGRES_NAME")

	dbStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	database, err := gorm.Open("postgres", dbStr)
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Plant{})

	DB = database
}
