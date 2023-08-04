package repositories

import (
	"gorm.io/gorm"
	"inktrail/models"
)

func SaveBlog(db *gorm.DB, blog *models.Blog) (*models.Blog, error) {
	err := db.Create(blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}
