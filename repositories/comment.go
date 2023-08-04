package repositories

import (
	"gorm.io/gorm"
	"inktrail/models"
)

func SaveComment(db *gorm.DB, comment *models.Comment) (*models.Comment, error) {
	err := db.Create(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// FindCommentsByBlogID gets all comments for a blog by its ID from the database
func FindCommentsByBlogID(db *gorm.DB, blogID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := db.Where("blog_id = ?", blogID).Find(&comments).Error
	return comments, err
}
