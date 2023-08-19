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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to DTO",
		})

		return
	}

	// TODO - add reading of currently logged in user ID here
	userID := uint(1)

	id, err := repository.SaveNewComment(dtos.CommentDTOToModel(&comment, userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new comment",
		})

		return
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get product ID from path params",
		})

		return
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get comment ID form path params",
		})

		return
	}

	comment, err := repository.FindCommentByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find comment with the given ID",
		})

		return
	}

	var commentDTO dtos.CommentDTO

	if err := c.BindJSON(&commentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to DTO",
		})

		return
	}

	userID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get current user ID",
		})

		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Comment user ID and logged in user ID do not match",
		})

		return
	}

	err = repository.UpdateCommentByID(uint(id), dtos.CommentDTOToModel(&commentDTO, userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update comment",
		})

		return
	}

	c.Status(http.StatusOK)
}

func DeleteCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get comment ID from path params",
		})

		return
	}

	comment, err := repository.FindCommentByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find comment with the given ID",
		})

		return
	}

	userID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get current user ID",
		})

		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Comment user ID and logged in user ID do not match",
		})

		return
	}

	err = repository.DeleteCommentByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting comment",
		})

		return
	}

	c.Status(http.StatusOK)
}
