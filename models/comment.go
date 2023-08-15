package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string
	UserID    uint
	ProductID uint
}
