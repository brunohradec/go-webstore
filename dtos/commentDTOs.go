package dtos

import (
	"time"

	"github.com/brunohradec/go-webstore/models"
)

type CommentDTO struct {
	Content   string `json:"content" binding:"required"`
	ProductID uint   `json:"productID" binding:"required"`
}

type CommentResponseDto struct {
	ID        uint      `json:"ID"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userID"`
	Username  string    `json:"username"`
	ProductID uint      `json:"productID"`
	CreatedAt time.Time `json:"createdAt"`
}

func CommentDTOToModel(dto *CommentDTO) *models.Comment {
	return &models.Comment{
		Content:   dto.Content,
		ProductID: dto.ProductID,
	}
}

func CommentModelToResponseDTO(model *models.Comment) *CommentResponseDto {
	return &CommentResponseDto{
		ID:        model.ID,
		Content:   model.Content,
		UserID:    model.UserID,
		ProductID: model.ProductID,
		CreatedAt: model.CreatedAt,
	}
}
