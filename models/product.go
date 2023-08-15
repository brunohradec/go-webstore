package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       int64
	UserID      uint `gorm:"not null"`
	Comments    []Comment
}
