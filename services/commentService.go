package services

import (
	"github.com/brunohradec/go-webstore/entities"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/repositories"
)

type CommentService interface {
	Save(comment *entities.Comment) (uint, error)
	FindByID(ID uint) (*entities.Comment, error)
	FindByProductID(productID uint, page paging.Page) []entities.Comment
	UpdateByID(ID uint, updatedComment *entities.Comment) error
	DeleteByID(ID uint) error
}

type CommentServiceImpl struct {
	CommentRepository repositories.CommentRepository
}

func InitCommentService(commentRepository repositories.CommentRepository) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
	}
}

func (service *CommentServiceImpl) Save(comment *entities.Comment) (uint, error) {
	return service.CommentRepository.Save(comment)
}

func (service *CommentServiceImpl) FindByID(ID uint) (*entities.Comment, error) {
	return service.CommentRepository.FindByID(ID)
}

func (service *CommentServiceImpl) FindByProductID(productID uint, page paging.Page) []entities.Comment {
	return service.CommentRepository.FindByProductID(productID, page)
}

func (service *CommentServiceImpl) UpdateByID(ID uint, updatedComment *entities.Comment) error {
	return service.CommentRepository.UpdateByID(ID, updatedComment)
}

func (service *CommentServiceImpl) DeleteByID(ID uint) error {
	return service.CommentRepository.DeleteByID(ID)
}
