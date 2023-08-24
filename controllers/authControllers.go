package controllers

import (
	"errors"
	"net/http"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/brunohradec/go-webstore/shared"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context) {
	var userDTO dtos.UserDTO

	err := c.BindJSON(&userDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to user DTO",
		})

		return
	}

	id, err := repository.SaveNewUser(dtos.UserDTOToModel(&userDTO))

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

func LoginUser(c *gin.Context) {
	secret := shared.Env.JWT.AccessTokenSecret
	tokenTTL := shared.Env.JWT.AccessTokenTTL

	var loginDTO dtos.LoginDTO

	err := c.BindJSON(&loginDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": "Could not bind JSON to DTO",
		})

		return
	}

	user, err := repository.FindUserByUseraname(loginDTO.Username)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mesage": "Could not find user with the given username",
		})

		return
	}

	err = auth.VerifyPassword(loginDTO.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mesage": "Provided password is incorrect",
		})

		return
	}

	token, err := auth.GenerateToken(user.ID, secret, tokenTTL)

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

func GetCurrentUser(c *gin.Context) {
	token, err := auth.ExtractTokenFromRequest(c)
	secret := shared.Env.JWT.AccessTokenSecret

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mesage": "Could not extract JSON web token from request headers or query",
		})

		return
	}

	userID, err := auth.ExtractUserIDFromToken(token, secret)

	user, err := repository.FindUserByID(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mesage": "Could not find user with the ID extracted from JSON web token",
		})

		return
	}

	c.JSON(http.StatusOK, dtos.UserModelToResponseDto(user))
}
