package controllers

import (
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/gin-gonic/gin"
)

func SaveNewComment(c *gin.Context) {
	var comment dtos.CommentDTO

	err := c.BindJSON(&comment)

	if err != nil {
		RejectResponseAndLog(
			"Could not bind JSON to DTO",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	// TODO - add reading of currently logged in user ID here
	userID := uint(1)

	id, err := repository.SaveNewComment(dtos.CommentDTOToModel(&comment, userID))

	if err != nil {
		RejectResponseAndLog(
			"Could not save new comment",
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
	page := paging.ParsePageFromQuery(c)

	productIdStr := c.Param("productId")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get product ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	comments := repository.FindCommentsByProductID(uint(productId), page)
	commentDTOs := make([]*dtos.CommentResponseDto, len(comments))

	for i, comment := range comments {
		/* As userID commes from Comment entity and userID is a foreign key,
		* the user with the given ID must always exist and no error handling is
		* necessary. */
		user, _ := repository.FindUserByID(comment.UserID)
		commentDTOs[i] = dtos.CommentModelToResponseDTO(&comment, user)
	}

	c.JSON(http.StatusOK, commentDTOs)
}

func UpdateCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get comment ID form path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	comment, err := repository.FindCommentByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Could not find comment with the given ID",
			http.StatusNotFound,
			err,
			c,
		)
	}

	var commentDTO dtos.CommentDTO

	if err := c.BindJSON(&commentDTO); err != nil {
		RejectResponseAndLog(
			"Could not bind JSON to DTO",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	userID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		RejectResponseAndLog(
			"Could not get current user ID",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	if comment.UserID != userID {
		RejectResponseAndLog(
			"Comment user ID and logged in user ID do not match",
			http.StatusForbidden,
			err,
			c,
		)
	}

	err = repository.UpdateCommentByID(uint(id), dtos.CommentDTOToModel(&commentDTO, userID))

	if err != nil {
		RejectResponseAndLog(
			"Could not update comment",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	c.Status(http.StatusOK)
}

func DeleteCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get comment ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	comment, err := repository.FindCommentByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Could not find comment with the given ID",
			http.StatusNotFound,
			err,
			c,
		)
	}

	userID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		RejectResponseAndLog(
			"Could not get current user ID",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	if comment.UserID != userID {
		RejectResponseAndLog(
			"Comment user ID and logged in user ID do not match",
			http.StatusForbidden,
			err,
			c,
		)
	}

	err = repository.DeleteCommentByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Error deleting comment",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	c.Status(http.StatusOK)
}
