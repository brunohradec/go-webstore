package repositories

import (
	"log"

	"github.com/brunohradec/go-webstore/entities"
	"github.com/brunohradec/go-webstore/paging"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product *entities.Product) (uint, error)
	FindByID(ID uint) (*entities.Product, error)
	FindAll(page paging.Page) []entities.Product
	FindByUserID(userID uint, page paging.Page) []entities.Product
	UpdateByID(ID uint, updatedProduct *entities.Product) error
	DeleteByID(ID uint) error
}

type PostgresProductRepository struct {
	DB *gorm.DB
}

func InitProductRepository(DB *gorm.DB) ProductRepository {
	return &PostgresProductRepository{
		DB: DB,
	}
}

func (repository *PostgresProductRepository) Save(product *entities.Product) (uint, error) {
	result := repository.DB.Create(product)

	if result.Error != nil {
		log.Println("ERROR: could not save new product", result.Error)
		return 0, result.Error
	}

	return product.ID, nil
}

func (repository *PostgresProductRepository) FindByID(ID uint) (*entities.Product, error) {
	var product entities.Product

	result := repository.DB.First(&product, ID)

	if result.Error != nil {
		log.Println("ERROR: could not find product with ID", ID, result.Error)
		return nil, result.Error
	}

	return &product, nil
}

func (repository *PostgresProductRepository) FindAll(page paging.Page) []entities.Product {
	var products []entities.Product

	repository.DB.Scopes(paging.Paginate(page)).Find(&products)

	return products
}

func (repository *PostgresProductRepository) FindByUserID(userID uint, page paging.Page) []entities.Product {
	var products []entities.Product

	repository.DB.
		Scopes(paging.Paginate(page)).
		Where("user_id = ?", userID).
		Find(&products)

	return products
}

func (repository *PostgresProductRepository) UpdateByID(ID uint, updatedProduct *entities.Product) error {
	updatedProduct.ID = ID

	result := repository.DB.Save(updatedProduct)

	if result.Error != nil {
		log.Println("ERROR: could not update product with ID", ID, result.Error)
		return result.Error
	}

	return nil
}

func (repository *PostgresProductRepository) DeleteByID(ID uint) error {
	result := repository.DB.Delete(&entities.Product{}, ID)

	if result.Error != nil {
		log.Println("ERROR: could not delete product with ID", ID, result.Error)
		return result.Error
	}

	return nil
}
