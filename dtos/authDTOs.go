package dtos

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReponseDTO struct {
	Token string `json:"token"`
}
