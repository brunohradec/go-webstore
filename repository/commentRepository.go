package repository

import (
	"log"

	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/shared"
)

func SaveNewComment(comment *models.Comment) (uint, error) {
	result := shared.DB.Create(comment)

	if result.Error != nil {
		log.Println("ERROR: could not save new comment", result.Error)
		return 0, result.Error
	}

	return comment.ID, nil
}

func FindCommentByID(ID uint) (*models.Comment, error) {
	var comment models.Comment

	result := shared.DB.First(&comment, ID)

	if result.Error != nil {
		log.Println("ERROR: could not find comment with ID", ID, result.Error)
		return nil, result.Error
	}

	return &comment, nil
}

func FindCommentsByProductID(productID uint, page paging.Page) []models.Comment {
	var comments []models.Comment

	shared.DB.
		Scopes(paging.Paginate(page)).
		Where("product_id = ?", productID).
		Find(&comments)

	return comments
}

func UpdateCommentByID(ID uint, updatedComment *models.Comment) error {
	updatedComment.ID = ID

	result := shared.DB.Save(updatedComment)

	if result.Error != nil {
		log.Println("ERROR: could not update comment with ID", ID, result.Error)
		return result.Error
	}

	return nil
}

func DeleteCommentByID(ID uint) error {
	result := shared.DB.Delete(&models.Comment{}, ID)

	if result.Error != nil {
		log.Println("ERROR: could not delete comment with ID", ID, result.Error)
		return result.Error
	}

	return nil
}
