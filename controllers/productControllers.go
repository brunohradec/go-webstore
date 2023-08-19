package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNewProduct(c *gin.Context) {
	var productDTO dtos.ProductDTO

	err := c.BindJSON(&productDTO)

	if err != nil {
		RejectResponseAndLog(
			"Could not bind JSON to product DTO",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	id, err := repository.SaveNewProduct(dtos.ProductDTOToModel(&productDTO))

	if err != nil {
		RejectResponseAndLog(
			"Could not save new product",
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
		RejectResponseAndLog(
			"Could not get product ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	product, err := repository.FindProductByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RejectResponseAndLog(
				"Could not find product with the given ID",
				http.StatusNotFound,
				err,
				c,
			)
		} else {
			RejectResponseAndLog(
				"Could not find product by ID",
				http.StatusInternalServerError,
				err,
				c,
			)
		}
	}

	c.JSON(http.StatusOK, dtos.ProductModelToResponseDTO(product))
}

func FindAllProducts(c *gin.Context) {
	page := paging.ParsePageFromQuery(c)

	products := repository.FindAllProducts(page)
	productDTOs := make([]*dtos.ProductResponseDTO, len(products))

	for i, product := range products {
		productDTOs[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func FindProductsByUserID(c *gin.Context) {
	page := paging.ParsePageFromQuery(c)

	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get user ID form path params",
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
		RejectResponseAndLog(
			"Could not get product ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	product, err := repository.FindProductByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Could not find product with the given ID",
			http.StatusNotFound,
			err,
			c,
		)
	}

	var productDTO dtos.ProductDTO

	if err := c.BindJSON(&productDTO); err != nil {
		RejectResponseAndLog(
			"Could not bind JSON to product",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	userID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		RejectResponseAndLog(
			"Could not get current user ID",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	if product.UserID != userID {
		RejectResponseAndLog(
			"Product user ID and logged in user ID do not match",
			http.StatusForbidden,
			err,
			c,
		)
	}

	err = repository.UpdateProductByID(uint(id), dtos.ProductDTOToModel(&productDTO))

	if err != nil {
		RejectResponseAndLog(
			"Error updating product",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	c.Status(http.StatusOK)
}

func DeleteProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		RejectResponseAndLog(
			"Could not get product ID from path params",
			http.StatusBadRequest,
			err,
			c,
		)
	}

	product, err := repository.FindProductByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Could not find product with the given ID",
			http.StatusNotFound,
			err,
			c,
		)
	}

	userID, err := auth.ExtractUserIDFromRequest(c)

	if err != nil {
		RejectResponseAndLog(
			"Could not get current user ID",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	if product.UserID != userID {
		RejectResponseAndLog(
			"Product user ID and logged in user ID do not match",
			http.StatusForbidden,
			err,
			c,
		)
	}

	err = repository.DeleteProductByID(uint(id))

	if err != nil {
		RejectResponseAndLog(
			"Error deleting product",
			http.StatusInternalServerError,
			err,
			c,
		)
	}

	c.Status(http.StatusOK)
}
