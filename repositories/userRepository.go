package repositories

import (
	"log"

	"github.com/brunohradec/go-webstore/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entities.User) (uint, error)
	FindByID(ID uint) (*entities.User, error)
	FindByUseraname(username string) (*entities.User, error)
	UpdateByID(ID uint, updatedUser *entities.User) error
	DeleteByID(ID uint) error
}

type PostgresUserRepository struct {
	DB *gorm.DB
}

func InitUserRepository(DB *gorm.DB) UserRepository {
	return &PostgresUserRepository{
		DB: DB,
	}
}

func (repository *PostgresUserRepository) Save(user *entities.User) (uint, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("ERROR: could not save new user. Error hashing password", err)
		return 0, err
	}

	user.Password = string(passwordHash)

	result := repository.DB.Create(user)

	if result.Error != nil {
		log.Println("ERROR: could not save new user", result.Error)
		return 0, result.Error
	}

	return user.ID, nil
}

func (repository *PostgresUserRepository) FindByID(ID uint) (*entities.User, error) {
	var user entities.User

	result := repository.DB.First(&user, ID)

	if result.Error != nil {
		log.Println("ERROR: could not find user with id", ID, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (repository *PostgresUserRepository) FindByUseraname(username string) (*entities.User, error) {
	var user entities.User

	result := repository.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		log.Println("ERROR: could not find user with username", username, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (repository *PostgresUserRepository) UpdateByID(ID uint, updatedUser *entities.User) error {
	updatedUser.ID = ID

	result := repository.DB.Save(updatedUser)

	if result.Error != nil {
		log.Println("ERROR: could not update user with ID", ID, result.Error)
		return result.Error
	}

	return nil
}

func (repository *PostgresUserRepository) DeleteByID(ID uint) error {
	result := repository.DB.Delete(&entities.User{}, ID)

	if result.Error != nil {
		log.Println("ERROR: could not delete user with ID", ID, result.Error)
		return result.Error
	}

	return nil
}
