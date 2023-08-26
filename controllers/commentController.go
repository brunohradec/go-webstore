package controllers

import (
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/authutils"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/services"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	Save(c *gin.Context)
	FindByProductID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type CommentControllerImpl struct {
	CommentService services.CommentService
	UserService    services.UserService
}

func InitCommentController(
	commentService services.CommentService,
	userService services.UserService,
) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
		UserService:    userService,
	}
}

func (controller *CommentControllerImpl) Save(c *gin.Context) {
	var commentDTO dtos.CommentDTO

	err := c.BindJSON(&commentDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to DTO",
		})

		return
	}

	principalID := authutils.GetPrincipalIDFromRequest(c)

	newComment := dtos.CommentDTOToModel(&commentDTO)
	newComment.UserID = principalID

	id, err := controller.CommentService.Save(newComment)

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

func (controller *CommentControllerImpl) FindByProductID(c *gin.Context) {
	page := paging.ParsePageFromQuery(c)

	productIdStr := c.Param("productId")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get product ID from path params",
		})

		return
	}

	comments := controller.CommentService.FindByProductID(uint(productId), page)
	commentDTOs := make([]*dtos.CommentResponseDto, len(comments))

	for i, comment := range comments {
		/* As userID commes from Comment entity and userID is a foreign key,
		* the user with the given ID must always exist and no error handling is
		* necessary. */
		user, _ := controller.UserService.FindByID(comment.UserID)

		commentDTOs[i] = dtos.CommentModelToResponseDTO(&comment)
		commentDTOs[i].Username = user.Username
	}

	c.JSON(http.StatusOK, commentDTOs)
}

func (controller *CommentControllerImpl) UpdateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get comment ID form path params",
		})

		return
	}

	comment, err := controller.CommentService.FindByID(uint(id))

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

	principalID := authutils.GetPrincipalIDFromRequest(c)

	if comment.UserID != principalID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Comment user ID and logged in user ID do not match",
		})

		return
	}

	updatedComment := dtos.CommentDTOToModel(&commentDTO)
	updatedComment.UserID = principalID

	err = controller.CommentService.UpdateByID(uint(id), updatedComment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update comment",
		})

		return
	}

	c.Status(http.StatusOK)
}

func (controller *CommentControllerImpl) DeleteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get comment ID from path params",
		})

		return
	}

	comment, err := controller.CommentService.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find comment with the given ID",
		})

		return
	}

	principalID := authutils.GetPrincipalIDFromRequest(c)

	if comment.UserID != principalID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Comment user ID and logged in user ID do not match",
		})

		return
	}

	err = controller.UserService.DeleteByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting comment",
		})

		return
	}

	c.Status(http.StatusOK)
}
