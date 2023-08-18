package dtos

import "github.com/brunohradec/go-webstore/models"

type UserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UserResponseDto struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

func UserDTOToModel(dto *UserDTO) *models.User {
	return &models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Username:  dto.Username,
		Password:  dto.Password,
	}
}

func UserModelToResponseDto(model *models.User) *UserResponseDto {
	return &UserResponseDto{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.LastName,
		Username:  model.Username,
	}
}
