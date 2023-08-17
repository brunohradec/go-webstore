package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/services"
	"github.com/brunohradec/go-webstore/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewComment(c *gin.Context) {
	var comment dtos.CommentDTO

	err := c.BindJSON(&comment)

	if err != nil {
		RejectResponseAndLog(
			"Error binding JSON while saving new comment",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	// TODO - add reading of currently logged in user ID here
	userId := uint(1)

	id, err := services.SaveNewComment(dtos.CommentDTOToModel(&comment, userId))

	if err != nil {
		RejectResponseAndLog(
			"Error while saving new comment",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func FindCommentsByProductID(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

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

	comments := services.FindCommentsByProductID(uint(id), page)
	commentDTOs := make([]*dtos.CommentResponseDto, len(comments))

	for i, comment := range comments {
		commentDTOs[i] = dtos.CommentModelToResponseDTO(&comment)
	}

	c.JSON(http.StatusOK, commentDTOs)
}

func UpdateCommentByID(c *gin.Context) {
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

	var commentDTO dtos.CommentDTO

	if err := c.BindJSON(&commentDTO); err != nil {
		RejectResponseAndLog(
			"Error binding JSON while updating comment",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	// TODO - add reading of currently logged in user ID here
	userId := uint(1)

	err = services.UpdateCommentByID(uint(id), dtos.CommentDTOToModel(&commentDTO, userId))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Error updating comment. Comment with the given ID does not exist",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error updating comment",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

}

func DeleteCommentByID(c *gin.Context) {
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

	err = services.DeleteCommentByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Error deleting comment. Comment with the given ID does not exist",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Error deleting comment",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}
}
