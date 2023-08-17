package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewUser(c *gin.Context) {
	var userDTO dtos.UserDTO

	err := c.BindJSON(&userDTO)

	if err != nil {
		RejectResponseAndLog(
			"Error binding JSON while saving new user",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	id, err := services.SaveNewUser(dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			RejectResponseAndLog(
				"Error while saving new user. User with the given username already exists",
				http.StatusConflict,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error while saving new user",
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

func FindUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	user, err := services.FindUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Error finding user. User with the given ID not found.",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error finding user by ID",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

	c.JSON(http.StatusOK, user)
}

func FindUserByUseraname(c *gin.Context) {
	username := c.Query("username")

	user, err := services.FindUserByUseraname(username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Error finding user. User with the given username not found.",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error finding user by username",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	var userDTO dtos.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		RejectResponseAndLog(
			"Error binding JSON while updating user",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	err = services.UpdateUserByID(uint(id), dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			RejectResponseAndLog(
				"Error updating user. User with the given username already exists.",
				http.StatusInternalServerError,
				err,
				c,
			)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Error updating user. User with the given ID does not exist.",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error updating user",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

}

func DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	err = services.DeleteUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Error deleting user. User with the given ID does not exist",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error deleting user",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}
}
