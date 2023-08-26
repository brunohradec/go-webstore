package repositories

import (
	"log"

	"github.com/brunohradec/go-webstore/entities"
	"github.com/brunohradec/go-webstore/paging"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Save(comment *entities.Comment) (uint, error)
	FindByID(ID uint) (*entities.Comment, error)
	FindByProductID(productID uint, page paging.Page) []entities.Comment
	UpdateByID(ID uint, updatedComment *entities.Comment) error
	DeleteByID(ID uint) error
}

type PostgresCommentRepository struct {
	DB *gorm.DB
}

func InitCommentRepository(DB *gorm.DB) CommentRepository {
	return &PostgresCommentRepository{
		DB: DB,
	}
}

func (repository *PostgresCommentRepository) Save(comment *entities.Comment) (uint, error) {
	result := repository.DB.Create(comment)

	if result.Error != nil {
		log.Println("ERROR: could not save new comment", result.Error)
		return 0, result.Error
	}

	return comment.ID, nil
}

func (repository *PostgresCommentRepository) FindByID(ID uint) (*entities.Comment, error) {
	var comment entities.Comment

	result := repository.DB.First(&comment, ID)

	if result.Error != nil {
		log.Println("ERROR: could not find comment with ID", ID, result.Error)
		return nil, result.Error
	}

	return &comment, nil
}

func (repository *PostgresCommentRepository) FindByProductID(productID uint, page paging.Page) []entities.Comment {
	var comments []entities.Comment

	repository.DB.
		Scopes(paging.Paginate(page)).
		Where("product_id = ?", productID).
		Find(&comments)

	return comments
}

func (repository *PostgresCommentRepository) UpdateByID(ID uint, updatedComment *entities.Comment) error {
	updatedComment.ID = ID

	result := repository.DB.Save(updatedComment)

	if result.Error != nil {
		log.Println("ERROR: could not update comment with ID", ID, result.Error)
		return result.Error
	}

	return nil
}

func (repository *PostgresCommentRepository) DeleteByID(ID uint) error {
	result := repository.DB.Delete(&entities.Comment{}, ID)

	if result.Error != nil {
		log.Println("ERROR: could not delete comment with ID", ID, result.Error)
		return result.Error
	}

	return nil
}
