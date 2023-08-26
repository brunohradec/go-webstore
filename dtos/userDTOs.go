package dtos

import "github.com/brunohradec/go-webstore/entities"

type UserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserResponseDto struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

func UserDTOToModel(dto *UserDTO) *entities.User {
	return &entities.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Username:  dto.Username,
		Password:  dto.Password,
	}
}

func UserModelToResponseDto(model *entities.User) *UserResponseDto {
	return &UserResponseDto{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.LastName,
		Username:  model.Username,
	}
}
