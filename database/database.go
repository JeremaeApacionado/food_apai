package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var location = "host=localhost user=postgres password=admin dbname=sample port=5432 sslmode=disable"

func DatabaseConnection() {
	DB, err = gorm.Open(postgres.Open(location), &gorm.Config{})
	if err != nil {
		fmt.Println("Not connected")
	}
	fmt.Println("Connected to  Database")
	DB.AutoMigrate(&User{})
}

type User struct {
	Name     string
	Password string
}
