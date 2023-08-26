package dtos

import "github.com/brunohradec/go-webstore/entities"

type ProductDTO struct {
	Name        string `json:"name" binding:"required"`
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

func ProductDTOToModel(dto *ProductDTO) *entities.Product {
	return &entities.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func ProductModelToResponseDTO(model *entities.Product) *ProductResponseDTO {
	return &ProductResponseDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		UserID:      model.UserID,
	}
}
