package controllers

import (
	"errors"
	"net/http"

	"github.com/brunohradec/go-webstore/authutils"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Me(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthService services.AuthService
	UserService services.UserService
}

func InitAuthController(
	authService services.AuthService,
	userService services.UserService) AuthController {

	return &AuthControllerImpl{
		AuthService: authService,
		UserService: userService,
	}
}

func (controller *AuthControllerImpl) Register(c *gin.Context) {
	var userDTO dtos.UserDTO

	err := c.BindJSON(&userDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to user DTO",
		})

		return
	}

	id, err := controller.AuthService.Register(dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{
				"message": "User with the given username already exists",
			})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"mesage": "Could not save new user",
			})

			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (controller *AuthControllerImpl) Login(c *gin.Context) {
	var loginDTO dtos.LoginDTO

	err := c.BindJSON(&loginDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": "Could not bind JSON to DTO",
		})

		return
	}

	token, err := controller.AuthService.Login(loginDTO.Username, loginDTO.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"mesage": "Could not generate access token",
		})

		return
	}

	c.JSON(http.StatusOK, dtos.LoginReponseDTO{
		AccessToken: token,
	})
}

func (controller *AuthControllerImpl) Me(c *gin.Context) {
	principalID := authutils.GetPrincipalIDFromRequest(c)

	user, err := controller.UserService.FindByID(principalID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mesage": "Could not find user with the ID extracted from JSON web token",
		})

		return
	}

	c.JSON(http.StatusOK, dtos.UserModelToResponseDto(user))
}
