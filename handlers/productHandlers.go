package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/brunohradec/go-webstore/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewProduct(c *gin.Context) {
	var productDTO dtos.ProductDTO

	err := c.BindJSON(&productDTO)

	if err != nil {
		utils.RejectResponseAndLog(
			"Error binding JSON while saving new product",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	id, err := repository.SaveNewProduct(dtos.ProductDTOToModel(&productDTO))

	if err != nil {
		utils.RejectResponseAndLog(
			"Error while saving new product",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func FindProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		utils.RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	product, err := repository.FindProductByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RejectResponseAndLog(
				"Error finding product. Product with the given ID not found.",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
				"Error finding product by ID",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

	c.JSON(http.StatusOK, product)
}

func FindAllProducts(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	products := repository.FindAllProducts(page)
	productDTOs := make([]*dtos.ProductResponseDTO, len(products))

	for i, product := range products {
		productDTOs[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func FindProductsByUserID(c *gin.Context) {
	page := utils.ParsePageFromQuery(c)

	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)

	if err != nil {
		utils.RejectResponseAndLog(
			"Error parsing user ID from query params while finding products by user ID",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	products := repository.FindProductsByUserID(uint(userID), page)
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
		utils.RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	var productDTO dtos.ProductDTO

	if err := c.BindJSON(&productDTO); err != nil {
		utils.RejectResponseAndLog(
			"Error binding JSON while updating product",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	err = repository.UpdateProductByID(uint(id), dtos.ProductDTOToModel(&productDTO))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RejectResponseAndLog(
				"Error updating product. Product with the given ID does not exist",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
				"Error updating product",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

}

func DeleteProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		utils.RejectResponseAndLog(
			"Error while parsing ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	err = repository.DeleteProductByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RejectResponseAndLog(
				"Error deleting product. Product with the given ID does not exist",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			utils.RejectResponseAndLog(
				"Error deleting product",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}
}
