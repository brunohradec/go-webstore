package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Products  []Product
	Comments  []Comment
}
