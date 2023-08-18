package dtos

import (
	"time"

	"github.com/brunohradec/go-webstore/models"
)

type CommentDTO struct {
	Content   string `json:"content" validate:"required"`
	ProductID uint   `json:"productID" validate:"required"`
}

type CommentResponseDto struct {
	ID        uint      `json:"ID"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userID"`
	Username  string    `json:"username"`
	ProductID uint      `json:"productID"`
	CreatedAt time.Time `json:"createdAt"`
}

// User ID usually commes from the currently logged in user
func CommentDTOToModel(dto *CommentDTO, userID uint) *models.Comment {
	return &models.Comment{
		Content:   dto.Content,
		ProductID: dto.ProductID,
		UserID:    userID,
	}
}

// User data is needed for adding usernames to the comment DTO
func CommentModelToResponseDTO(model *models.Comment, user *models.User) *CommentResponseDto {
	return &CommentResponseDto{
		ID:        model.ID,
		Content:   model.Content,
		UserID:    model.UserID,
		Username:  user.Username,
		ProductID: model.ProductID,
		CreatedAt: model.CreatedAt,
	}
}
