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

func UpdateUser(c *gin.Context) {
	userID, err := auth.ExtractUserIDFromRequestToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not get current user ID",
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

	updatedUser := dtos.UserDTOToModel(&userDTO)

	err = repository.UpdateUserByID(userID, updatedUser)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "User with the given ID not found.",
			})

			return
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
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
