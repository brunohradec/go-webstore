package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Item struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	SellerID    int64  `json:"sellerID"`
}

var users = []User{
	{
		ID:       1,
		Username: "iivic",
		Name:     "Ivan",
		LastName: "Ivic",
		Email:    "iivic@test.com",
		Password: "iivicpass",
	},
	{
		ID:       2,
		Username: "pperic",
		Name:     "Pero",
		LastName: "Peric",
		Email:    "pperic@test.com",
		Password: "ppericpass",
	},
}

var items = []Item{
	{
		ID:          1,
		Title:       "Jabuka",
		Description: "Jabuka Granny Smith",
		Price:       200,
		SellerID:    1,
	},
	{
		ID:          2,
		Title:       "Kruska",
		Description: "Kruska Williams",
		Price:       250,
		SellerID:    1,
	},
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/items", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, items)
	})

	r.Run()
}
