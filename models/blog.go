package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	title   string
	content string
	UserID  uint
}
