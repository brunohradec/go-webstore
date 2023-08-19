package dtos

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginReponseDTO struct {
	Token string `json:"token"`
}
