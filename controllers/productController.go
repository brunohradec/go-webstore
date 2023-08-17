package controllers

import (
	"errors"
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
	var newProduct dtos.ProductDTO

	err := c.BindJSON(&newProduct)

	if err != nil {
		msg := "Error binding JSON while saving new product"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	newProductId, err := services.SaveNewProduct(dtos.ProductDTOToModel(&newProduct))

	if err != nil {
		msg := "Error while saving new product"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": newProductId,
	})
}

func FindProductByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	foundUser, err := services.FindProductByID(uint(id))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "Product with the given id not found"
		} else {
			msg = "Error while finding product"
		}

		log.Println(msg, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, foundUser)
}

func FindAllProducts(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	foundProducts := services.FindAllProducts(page)

	mappedProducts := make([]*dtos.ProductResponseDTO, len(foundProducts))

	for i, product := range foundProducts {
		mappedProducts[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, mappedProducts)
}

func FindProductsByUserID(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	userIDStr := c.Param("id")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	foundProducts := services.FindProductsByUserID(uint(userID), page)

	mappedProducts := make([]*dtos.ProductResponseDTO, len(foundProducts))

	for i, product := range foundProducts {
		mappedProducts[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, mappedProducts)
}

func UpdateProductById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	var updatedProduct dtos.ProductDTO

	if err := c.BindJSON(&updatedProduct); err != nil {
		msg := "Error binding JSON while updating product"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.UpdateProductByID(uint(id), dtos.ProductDTOToModel(&updatedProduct))

	if err != nil {
		var msg string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg = "Error updating product. Product with the given id does not exist"
			log.Println(msg, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		} else {
			msg = "Error updating product"
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
		msg := "Error while parsing ID from path params"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	err = services.DeleteProductByID(uint(id))

	if err != nil {
		msg := "Error while deleting product"
		log.Println(msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}
}
