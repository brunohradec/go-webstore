package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brunohradec/go-webstore/handlers"
	"github.com/brunohradec/go-webstore/initializers"
	"github.com/brunohradec/go-webstore/shared"
	"github.com/gin-gonic/gin"
)

func main() {
	env, err := initializers.LoadDotenvVariables()

	if err != nil {
		log.Fatal("Error loading dotenv file")
		os.Exit(1)
	}

	db, err := initializers.ConnectToDB(
		env.DB.Host,
		env.DB.Port,
		env.DB.Name,
		env.DB.Username,
		env.DB.Password,
	)

	if err != nil {
		log.Fatal("Error connecting to the database")
		os.Exit(1)
	}

	initializers.AutomigrateDB(db)

	shared.Env = env
	shared.DB = db

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", handlers.SaveNewUser)
			users.GET("/:id", handlers.FindUserByID)
			users.PUT("/:id", handlers.UpdateUserByID)
			users.DELETE("/:id", handlers.DeleteUserByID)
		}

		products := api.Group("/products")
		{
			products.POST("/", handlers.SaveNewProduct)
			products.GET("/", handlers.FindAllProducts)
			products.GET("/:id", handlers.FindProductByID)
			products.GET("/user/:userId", handlers.FindProductsByUserID)
			products.PUT("/:id", handlers.UpdateProductByID)
			products.DELETE("/:id", handlers.DeleteProductByID)
		}

		comments := api.Group("/comments")
		{
			comments.POST("/", handlers.SaveNewComment)
			comments.GET("/product/:productId", handlers.FindCommentsByProductID)
			comments.PUT("/:id", handlers.UpdateCommentByID)
			comments.DELETE("/:id", handlers.DeleteCommentByID)
		}
	}

	r.Run(":" + env.Port)
}
