package repository

import (
	"log"

	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/shared"
	"golang.org/x/crypto/bcrypt"
)

func SaveNewUser(user *models.User) (uint, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("ERROR: could not save new user. Error hashing password", err)
		return 0, err
	}

	user.Password = string(passwordHash)

	result := shared.DB.Create(user)

	if result.Error != nil {
		log.Println("ERROR: could not save new user", result.Error)
		return 0, result.Error
	}

	return user.ID, nil
}

func FindUserByID(ID uint) (*models.User, error) {
	var user models.User

	result := shared.DB.First(&user, ID)

	if result.Error != nil {
		log.Println("ERROR: could not find user with id", ID, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func FindUserByUseraname(username string) (*models.User, error) {
	var user models.User

	result := shared.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		log.Println("ERROR: could not find user with username", username, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func UpdateUserByID(ID uint, updatedUser *models.User) error {
	updatedUser.ID = ID

	result := shared.DB.Save(updatedUser)

	if result.Error != nil {
		log.Println("ERROR: could not update user with ID", ID, result.Error)
		return result.Error
	}

	return nil
}

func DeleteUserByID(ID uint) error {
	result := shared.DB.Delete(&models.User{}, ID)

	if result.Error != nil {
		log.Println("ERROR: could not delete user with ID", ID, result.Error)
		return result.Error
	}

	return nil
}
