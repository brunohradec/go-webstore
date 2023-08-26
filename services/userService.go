package services

import (
	"github.com/brunohradec/go-webstore/entities"
	"github.com/brunohradec/go-webstore/repositories"
)

type UserService interface {
	Save(user *entities.User) (uint, error)
	FindByID(ID uint) (*entities.User, error)
	FindByUseraname(username string) (*entities.User, error)
	UpdateByID(ID uint, updatedUser *entities.User) error
	DeleteByID(ID uint) error
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func InitUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Save(user *entities.User) (uint, error) {
	return service.UserRepository.Save(user)
}

func (service *UserServiceImpl) FindByID(ID uint) (*entities.User, error) {
	return service.UserRepository.FindByID(ID)
}

func (service *UserServiceImpl) FindByUseraname(username string) (*entities.User, error) {
	return service.UserRepository.FindByUseraname(username)
}

func (service *UserServiceImpl) UpdateByID(ID uint, updatedUser *entities.User) error {
	return service.UserRepository.UpdateByID(ID, updatedUser)
}

func (service *UserServiceImpl) DeleteByID(ID uint) error {
	return service.UserRepository.DeleteByID(ID)
}
