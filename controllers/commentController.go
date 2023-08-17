package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/services"
	"github.com/brunohradec/go-webstore/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewComment(c *gin.Context) {
	var newComment dtos.CommentDTO

	err := c.BindJSON(&newComment)

	if err != nil {
		msg := "Error binding JSON while saving new comment"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	// TODO - add reading of currently logged in user ID here
	userId := uint(1)

	newCommentId, err := services.SaveNewComment(dtos.CommentDTOToModel(&newComment, userId))

	if err != nil {
		msg := "Error while saving new comment"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": newCommentId,
	})
}

func FindCommentsByProductID(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	foundComments := services.FindCommentsByProductID(uint(productID), page)

	mappedComments := make([]*dtos.CommentResponseDto, len(foundComments))

	for i, comment := range foundComments {
		mappedComments[i] = dtos.CommentModelToResponseDTO(&comment)
	}

	c.JSON(http.StatusOK, mappedComments)
}

func UpdateCommentByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	var updatedComment dtos.CommentDTO

	if err := c.BindJSON(&updatedComment); err != nil {
		msg := "Error binding JSON while updating comment"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	// TODO - add reading of currently logged in user ID here
	userId := uint(1)

	err = services.UpdateCommentByID(uint(id), dtos.CommentDTOToModel(&updatedComment, userId))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "Error updating comment. Comment with the given id does not exist"
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		} else {
			msg = "Error updating comment"
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		}
	}

}

func DeleteCommentByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.DeleteCommentByID(uint(id))

	if err != nil {
		msg := "Error while deleting comment"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}
}
