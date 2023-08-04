package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"inktrail/models"
	"os"
)

func ConnectDB() {
	dsn := "host=localhost user=siddhant password=sarthak1995 dbname=siddhant port=5433 sslmode=disable TimeZone=Asia/Kolkata"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Connection Unsuccessfull")
	} else {
		println("Connection Successfull")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Blog{}, &models.Comment{})
	if err != nil {
		return
	}
}

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
