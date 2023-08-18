package dtos

import "github.com/brunohradec/go-webstore/models"

type ProductDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type ProductResponseDTO struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	UserID      uint   `json:"userID"`
}

func ProductDTOToModel(dto *ProductDTO) *models.Product {
	return &models.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func ProductModelToResponseDTO(model *models.Product) *ProductResponseDTO {
	return &ProductResponseDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		UserID:      model.UserID,
	}
}
