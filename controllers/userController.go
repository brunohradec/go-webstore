package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewUser(c *gin.Context) {
	var newUser dtos.UserDTO

	err := c.BindJSON(&newUser)

	if err != nil {
		msg := "Error binding JSON while saving new user"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	newUserId, err := services.SaveNewUser(dtos.UserDTOToModel(&newUser))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			msg = "Error while saving user. User with the given username already exists"
			log.Println(msg, err)
			c.JSON(http.StatusConflict, gin.H{
				"message": msg,
			})
		} else {
			msg = "Error while saving new user"
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": newUserId,
	})
}

func FindUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	foundUser, err := services.FindUserByID(uint(id))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "User with the given id not found"
		} else {
			msg = "Error while finding user"
		}

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, foundUser)
}

func FindUserByUseraname(c *gin.Context) {
	username := c.Query("username")

	foundUser, err := services.FindUserByUseraname(username)

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "User with the given username not found"
		} else {
			msg = "Error while finding user with the given username"
		}

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, foundUser)
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	var updatedUser dtos.UserDTO

	if err := c.BindJSON(&updatedUser); err != nil {
		msg := "Error binding JSON while updating user"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.UpdateUserByID(uint(id), dtos.UserDTOToModel(&updatedUser))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			msg = "Error while updating user. User with the given username already exists."
			log.Println(msg, err)
			c.JSON(http.StatusConflict, gin.H{
				"message": msg,
			})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "Error updating user. User with the given username does not exist"
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		} else {
			msg = "Error updating user"
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
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.DeleteUserByID(uint(id))

	if err != nil {
		msg := "Error while deleting user"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}
}
