package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Content string
	UserID  uint
	BlogID  uint
}
