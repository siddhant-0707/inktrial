package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	content string
	UserID  uint
	BlogID  uint
}
