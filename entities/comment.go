package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
}
