package controllers

import (
	"errors"
	"fmt"
	"log"
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
		msg := fmt.Sprintf(
			"Error binding JSON while saving user with the username %s",
			userDTO.Username,
		)

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	id, err := services.SaveNewUser(dtos.UserDTOToModel(&userDTO))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			msg = fmt.Sprintf(
				"Error while saving user. User with the username %s already exists",
				userDTO.Username,
			)

			log.Println(msg, err)

			c.JSON(http.StatusConflict, gin.H{
				"message": msg,
			})
		} else {
			msg = fmt.Sprintf(
				"Error while saving user with the username %s",
				userDTO.Username,
			)

			log.Println(msg, err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
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
		msg := fmt.Sprintf(
			"Error while parsing ID from path params while finding user with the ID %s",
			idStr,
		)

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	user, err := services.FindUserByID(uint(id))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = fmt.Sprintf("User with the ID %d not found", id)
		} else {
			msg = fmt.Sprintf("Error while finding user with the ID %d", id)
		}

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, user)
}

func FindUserByUseraname(c *gin.Context) {
	username := c.Query("username")

	user, err := services.FindUserByUseraname(username)

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = fmt.Sprintf("User with the username %s not found", username)
		} else {
			msg = fmt.Sprintf("Error while finding user with the username %s", username)
		}

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params while updating user by ID"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	var userDTO dtos.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		msg := fmt.Sprintf("Error binding JSON while updating user with the ID %d", id)
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.UpdateUserByID(uint(id), dtos.UserDTOToModel(&userDTO))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			msg = fmt.Sprintf(
				"Error while updating user. User with the username %s already exists.",
				userDTO.Username,
			)

			log.Println(msg, err)

			c.JSON(http.StatusConflict, gin.H{
				"message": msg,
			})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = fmt.Sprintf(
				"Error updating user. User with the ID %d does not exist",
				id,
			)

			log.Println(msg, err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		} else {
			msg = fmt.Sprintf("Error updating user with the ID %d", id)
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		}
	}

}

func DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params while deleting user by ID"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.DeleteUserByID(uint(id))

	if err != nil {
		msg := fmt.Sprintf("Error while deleting user with the ID %d", id)
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}
}
