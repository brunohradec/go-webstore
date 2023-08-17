package services

import (
	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/repositories"
)

func SaveNewUser(user *models.User) (uint, error) {
	return repositories.SaveNewUser(user)
}

func FindUserByID(ID uint) (*models.User, error) {
	return repositories.FindUserByID(ID)
}

func FindUserByUseraname(username string) (*models.User, error) {
	return repositories.FindUserByUseraname(username)
}

func UpdateUserByID(ID uint, updatedUser *models.User) error {
	return repositories.UpdateUserByID(ID, updatedUser)
}

func DeleteUserByID(ID uint) error {
	return repositories.DeleteUserById(ID)
}
