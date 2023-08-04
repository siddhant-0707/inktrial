package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user-id"`
	BlogID  uint   `json:"blog-id"`
}
