package dtos

import (
	"time"

	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/services"
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

func CommentDTOToModel(dto *CommentDTO) *models.Comment {
	return &models.Comment{
		Content:   dto.Content,
		ProductID: dto.ProductID,
	}
}

func CommentModelToResponseDTO(model *models.Comment) *CommentResponseDto {
	/* As userID commes from Comment entity and userID is a foreign key,
	 * the user with the given ID must always exist and no error handling is
	 * necessary. */
	user, _ := services.FindUserByID(model.UserID)

	return &CommentResponseDto{
		ID:        model.ID,
		Content:   model.Content,
		UserID:    model.UserID,
		Username:  user.Username,
		ProductID: model.ProductID,
		CreatedAt: model.CreatedAt,
	}
}
