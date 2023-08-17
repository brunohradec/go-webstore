package dtos

type ProductCreateDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type ProductResponseDto struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	UserID      uint   `json:"userID"`
}
