package dtos

import "time"

type CommentCreateDto struct {
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
