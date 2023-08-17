package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/services"
	"github.com/brunohradec/go-webstore/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewProduct(c *gin.Context) {
	var productDTO dtos.ProductDTO

	err := c.BindJSON(&productDTO)

	if err != nil {
		msg := "Error binding JSON while saving product"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	id, err := services.SaveNewProduct(dtos.ProductDTOToModel(&productDTO))

	if err != nil {
		msg := "Error while saving product"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func FindProductByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params while finding product by ID"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	product, err := services.FindProductByID(uint(id))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = fmt.Sprintf("Product with the ID %d not found", id)
		} else {
			msg = fmt.Sprintf("Error while finding product with the ID %d", id)
		}

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, product)
}

func FindAllProducts(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	products := services.FindAllProducts(page)
	productDTOs := make([]*dtos.ProductResponseDTO, len(products))

	for i, product := range products {
		productDTOs[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func FindProductsByUserID(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	userIDStr := c.Param("id")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params while finding products by user ID"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	products := services.FindProductsByUserID(uint(userID), page)
	productDTOs := make([]*dtos.ProductResponseDTO, len(products))

	for i, product := range products {
		productDTOs[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func UpdateProductByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params while updating product by ID"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	var productDTO dtos.ProductDTO

	if err := c.BindJSON(&productDTO); err != nil {
		msg := fmt.Sprintf("Error binding JSON while updating product with the ID %d", id)
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.UpdateProductByID(uint(id), dtos.ProductDTOToModel(&productDTO))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = fmt.Sprintf("Error updating product. Product with the ID %d does not exist", id)
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		} else {
			msg = fmt.Sprintf("Error updating product with the ID %d", id)
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		}
	}

}

func DeleteProductByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params while deleting product by ID"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.DeleteProductByID(uint(id))

	if err != nil {
		msg := fmt.Sprintf("Error while deleting product with the ID %d", id)
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}
}
