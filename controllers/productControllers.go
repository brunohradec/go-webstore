package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/paging"
	"github.com/brunohradec/go-webstore/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController interface {
	Save(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	FindByUserID(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func InitProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Save(c *gin.Context) {
	var productDTO dtos.ProductDTO

	err := c.BindJSON(&productDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to product DTO",
		})

		return
	}

	userID, err := auth.ExtractUserIDFromRequestToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not get current user ID",
		})

		return
	}

	newProduct := dtos.ProductDTOToModel(&productDTO)
	newProduct.UserID = userID

	id, err := controller.ProductService.Save(newProduct)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save new product",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (controller *ProductControllerImpl) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get product ID from path params",
		})

		return
	}

	product, err := controller.ProductService.FindByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Could not find product with the given ID",
			})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not find product by ID",
			})

			return
		}
	}

	c.JSON(http.StatusOK, dtos.ProductModelToResponseDTO(product))
}

func (controller *ProductControllerImpl) FindAll(c *gin.Context) {
	page := paging.ParsePageFromQuery(c)

	products := controller.ProductService.FindAll(page)
	productDTOs := make([]*dtos.ProductResponseDTO, len(products))

	for i, product := range products {
		productDTOs[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func (controller *ProductControllerImpl) FindByUserID(c *gin.Context) {
	page := paging.ParsePageFromQuery(c)

	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get user ID form path params",
		})

		return
	}

	products := controller.ProductService.FindByUserID(uint(userID), page)
	productDTOs := make([]*dtos.ProductResponseDTO, len(products))

	for i, product := range products {
		productDTOs[i] = dtos.ProductModelToResponseDTO(&product)
	}

	c.JSON(http.StatusOK, productDTOs)
}

func (controller *ProductControllerImpl) UpdateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get product ID from path params",
		})

		return
	}

	product, err := controller.ProductService.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find product with the given ID",
		})

		return
	}

	var productDTO dtos.ProductDTO

	if err := c.BindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON to product",
		})

		return
	}

	userID, err := auth.ExtractUserIDFromRequestToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not get current user ID",
		})

		return
	}

	if product.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Product user ID and logged in user ID do not match",
		})

		return
	}

	updatedProduct := dtos.ProductDTOToModel(&productDTO)
	updatedProduct.UserID = userID

	err = controller.ProductService.UpdateByID(uint(id), updatedProduct)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating product",
		})

		return
	}

	c.Status(http.StatusOK)
}

func (controller *ProductControllerImpl) DeleteByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get product ID from path params",
		})

		return
	}

	product, err := controller.ProductService.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find product with the given ID",
		})

		return
	}

	userID, err := auth.ExtractUserIDFromRequestToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not get current user ID",
		})

		return
	}

	if product.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Product user ID and logged in user ID do not match",
		})

		return
	}

	err = controller.ProductService.DeleteByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting product",
		})

		return
	}

	c.Status(http.StatusOK)
}
