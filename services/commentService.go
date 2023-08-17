package services

import (
	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/repositories"
	"github.com/brunohradec/go-webstore/utils"
)

func SaveNewComment(comment *models.Comment) (uint, error) {
	return repositories.SaveNewComment(comment)
}

func FindCommentByID(ID uint) (*models.Comment, error) {
	return repositories.FindCommentByID(ID)
}

func FindCommentsByProductID(productID uint, page utils.Page) []models.Comment {
	return repositories.FindCommentsByProductID(productID, page)
}

func UpdateCommentByID(ID uint, updatedComment *models.Comment) error {
	return repositories.UpdateCommentByID(ID, updatedComment)
}

func DeleteCommentByID(ID uint) error {
	return repositories.DeleteCommentByID(ID)
}
