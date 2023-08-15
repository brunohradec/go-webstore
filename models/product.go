package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       int64
	UserID      uint
	Comments    []Comment
}
