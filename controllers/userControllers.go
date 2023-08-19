package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get user ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	user, err := repository.FindUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Could not find user with the given ID",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Could not find user by ID",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

	c.JSON(http.StatusOK, dtos.UserModelToResponseDto(user))
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get user ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	user, err := repository.FindUserByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Could not find user with the given ID",
			http.StatusNotFound,
			err,
			c,
		)
	}

	var userDTO dtos.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		RejectResponseAndLog(
			"Could not bind JSON to DTO",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	tokenUserID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		RejectResponseAndLog(
			"Could not get current user ID",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	if user.ID != tokenUserID {
		RejectResponseAndLog(
			"Given user ID and logged in user ID do not match",
			http.StatusForbidden,
			err,
			c,
		)
	}

	err = repository.UpdateUserByID(uint(id), dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			RejectResponseAndLog(
				"User with the given username already exists.",
				http.StatusInternalServerError,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Could not update user",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

	c.Status(http.StatusOK)
}
