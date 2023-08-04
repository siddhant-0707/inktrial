package repositories

import (
	"gorm.io/gorm"
	"inktrail/models"
)

func FindUserById(db *gorm.DB, id uint) (models.User, error) {
	var user models.User

	err := db.Preload("Blogs").Where("ID = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
