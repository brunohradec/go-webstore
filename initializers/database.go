package initializers

import (
	"fmt"

	"github.com/brunohradec/go-webstore/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(
	host string,
	port string,
	name string,
	username string,
	password string) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		username,
		password,
		name,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutomigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Comment{})
}
