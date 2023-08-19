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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get user ID from path params",
		})

		return
	}

	user, err := repository.FindUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Could not find user with the given ID",
			})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not find user by ID",
			})

			return
		}
	}

	c.JSON(http.StatusOK, dtos.UserModelToResponseDto(user))
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get user ID from path params",
		})

		return
	}

	user, err := repository.FindUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find user with the given ID",
		})

		return
	}

	var userDTO dtos.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to DTO",
		})

		return
	}

	tokenUserID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get current user ID",
		})

		return
	}

	if user.ID != tokenUserID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Given user ID and logged in user ID do not match",
		})

		return
	}

	err = repository.UpdateUserByID(uint(id), dtos.UserDTOToModel(&userDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "User with the given username already exists.",
			})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not update user",
			})

			return
		}
	}

	c.Status(http.StatusOK)
}
