package services

import (
	"github.com/brunohradec/go-webstore/entities"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/repositories"
)

type ProductService interface {
	Save(product *entities.Product) (uint, error)
	FindByID(ID uint) (*entities.Product, error)
	FindAll(page paging.Page) []entities.Product
	FindByUserID(userID uint, page paging.Page) []entities.Product
	UpdateByID(ID uint, updatedProduct *entities.Product) error
	DeleteByID(ID uint) error
}

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func InitProductService(productRepository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

func (service *ProductServiceImpl) Save(product *entities.Product) (uint, error) {
	return service.ProductRepository.Save(product)
}

func (service *ProductServiceImpl) FindByID(ID uint) (*entities.Product, error) {
	return service.ProductRepository.FindByID(ID)
}

func (service *ProductServiceImpl) FindAll(page paging.Page) []entities.Product {
	return service.ProductRepository.FindAll(page)
}

func (service *ProductServiceImpl) FindByUserID(userID uint, page paging.Page) []entities.Product {
	return service.ProductRepository.FindByUserID(userID, page)
}

func (service *ProductServiceImpl) UpdateByID(ID uint, updatedProduct *entities.Product) error {
	return service.ProductRepository.UpdateByID(ID, updatedProduct)
}

func (service *ProductServiceImpl) DeleteByID(ID uint) error {
	return service.ProductRepository.DeleteByID(ID)
}
