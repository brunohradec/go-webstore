package services

import (
	"github.com/brunohradec/go-webstore/models"
	"github.com/brunohradec/go-webstore/repositories"
	"github.com/brunohradec/go-webstore/utils"
)

func SaveNewProduct(product *models.Product) (uint, error) {
	return repositories.SaveNewProduct(product)
}

func FindProductByID(ID uint) (*models.Product, error) {
	return repositories.FindProductByID(ID)
}

func FindAllProducts(page utils.Page) []models.Product {
	return repositories.FindAllProducts(page)
}

func FindProductsByUserID(userID uint, page utils.Page) []models.Product {
	return repositories.FindProductsByUserID(userID, page)
}

func UpdateProductByID(ID uint, updatedProduct *models.Product) error {
	return repositories.UpdateProductByID(ID, updatedProduct)
}

func DeleteProductByID(ID uint) error {
	return repositories.DeleteProductByID(ID)
}
