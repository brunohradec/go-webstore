package services

import (
	"github.com/brunohradec/go-webstore/authutils"
	"github.com/brunohradec/go-webstore/entities"
	"github.com/brunohradec/go-webstore/infrastructure"
)

type AuthService interface {
	Register(user *entities.User) (uint, error)
	Login(username string, password string) (string, error)
}

type AuthServiceImpl struct {
	UserService UserService
	Env         *infrastructure.Env
}

func InitAuthService(
	userService UserService,
	env *infrastructure.Env,
) AuthService {
	return &AuthServiceImpl{
		UserService: userService,
		Env:         env,
	}
}

func (service *AuthServiceImpl) Register(user *entities.User) (uint, error) {
	return service.UserService.Save(user)
}

func (service *AuthServiceImpl) Login(username string, password string) (string, error) {
	secret := service.Env.JWT.AccessTokenSecret
	tokenTTL := service.Env.JWT.AccessTokenTTL

	user, err := service.UserService.FindByUseraname(username)

	if err != nil {
		return "", err
	}

	err = authutils.VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	token, err := authutils.GenerateToken(user.ID, secret, tokenTTL)

	if err != nil {
		return "", err
	}

	return token, nil
}
