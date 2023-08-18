package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/brunohradec/go-webstore/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewUser(c *gin.Context) {
	var userDTO dtos.UserDTO

	err := c.BindJSON(&userDTO)

	if err != nil {
		utils.RejectResponseAndLog(
			"Error binding JSON while saving new user",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	id, err := repository.SaveNewUser(dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.RejectResponseAndLog(
				"Error while saving new user. User with the given username already exists",
				http.StatusConflict,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
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
		utils.RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	user, err := repository.FindUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RejectResponseAndLog(
				"Error finding user. User with the given ID not found.",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
				"Error finding user by ID",
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
		utils.RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	var userDTO dtos.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		utils.RejectResponseAndLog(
			"Error binding JSON while updating user",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	err = repository.UpdateUserByID(uint(id), dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.RejectResponseAndLog(
				"Error updating user. User with the given username already exists.",
				http.StatusInternalServerError,
				err,
				c,
			)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RejectResponseAndLog(
				"Error updating user. User with the given ID does not exist.",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
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
		utils.RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	err = repository.DeleteUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RejectResponseAndLog(
				"Error deleting user. User with the given ID does not exist",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
				"Error deleting user",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}
}
