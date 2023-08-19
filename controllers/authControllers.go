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
		RejectResponseAndLog(
			"Could not bind JSON to user DTO",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	id, err := repository.SaveNewUser(dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			RejectResponseAndLog(
				"User with the given username already exists",
				http.StatusConflict,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Could not save new user",
				http.StatusInternalServerError,
				err,
				c,
			)
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
		RejectResponseAndLog(
			"Could not bind JSON to DTO",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	user, err := repository.FindUserByUseraname(loginDTO.Username)

	if err != nil {
		RejectResponseAndLog(
			"Could not find user with the given username",
			http.StatusNotFound,
			err,
			c,
		)
	}

	err = auth.VerifyPassword(loginDTO.Password, user.Password)

	if err != nil {
		RejectResponseAndLog(
			"Provided password is incorrect",
			http.StatusUnauthorized,
			err,
			c,
		)
	}

	token, err := auth.GenerateToken(user.ID, secret, tokenTTL)

	if err != nil {
		RejectResponseAndLog(
			"Could not generate access token",
			http.StatusUnauthorized,
			err,
			c,
		)
	}

	c.JSON(http.StatusOK, dtos.LoginReponseDTO{
		Token: token,
	})
}

func GetCurrentUser(c *gin.Context) {
	token, err := auth.ExtractTokenFromRequest(c)
	secret := shared.Env.JWT.AccessTokenSecret

	if err != nil {
		RejectResponseAndLog(
			"Could not extract JSON web token from request headers or query",
			http.StatusUnauthorized,
			err,
			c,
		)
	}

	userID, err := auth.ExtractUserIDFromToken(token, secret)

	user, err := repository.FindUserByID(userID)

	if err != nil {
		RejectResponseAndLog(
			"Could not find user with the ID extracted from JSON web token",
			http.StatusUnauthorized,
			err,
			c,
		)
	}

	c.JSON(http.StatusOK, dtos.UserModelToResponseDto(user))
}
