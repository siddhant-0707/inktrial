package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"inktrail/models"
)

func ConnectDB() {
	dsn := "host=localhost user=siddhant password=sarthak1995 dbname=siddhant port=5433 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Connection Unsuccessfull")
	} else {
		println("Connection Successfull")
	}

	err = db.AutoMigrate(&models.User{}, &models.Blog{}, &models.Comments{})
	if err != nil {
		return
	}
}
