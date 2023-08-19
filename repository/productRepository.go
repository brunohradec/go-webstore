package repository

import (
	"log"

	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/shared"
)

func SaveNewProduct(product *models.Product) (uint, error) {
	result := shared.DB.Create(product)

	if result.Error != nil {
		log.Println("ERROR: could not save new product", result.Error)
		return 0, result.Error
	}

	return product.ID, nil
}

func FindProductByID(ID uint) (*models.Product, error) {
	var product models.Product

	result := shared.DB.First(&product, ID)

	if result.Error != nil {
		log.Println("ERROR: could not find product with ID", ID, result.Error)
		return nil, result.Error
	}

	return &product, nil
}

func FindAllProducts(page paging.Page) []models.Product {
	var products []models.Product

	shared.DB.Scopes(paging.Paginate(page)).Find(&products)

	return products
}

func FindProductsByUserID(userID uint, page paging.Page) []models.Product {
	var products []models.Product

	shared.DB.
		Scopes(paging.Paginate(page)).
		Where("user_id = ?", userID).
		Find(&products)

	return products
}

func UpdateProductByID(ID uint, updatedProduct *models.Product) error {
	updatedProduct.ID = ID

	result := shared.DB.Save(updatedProduct)

	if result.Error != nil {
		log.Println("ERROR: could not update product with ID", ID, result.Error)
		return result.Error
	}

	return nil
}

func DeleteProductByID(ID uint) error {
	result := shared.DB.Delete(&models.Product{}, ID)

	if result.Error != nil {
		log.Println("ERROR: could not delete product with ID", ID, result.Error)
		return result.Error
	}

	return nil
}
