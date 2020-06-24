package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "hoya-test.db")
	if err != nil {
		panic("Failed to conncet to database")
	}

	database.AutoMigrate(&Plant{})

	DB = database
}
