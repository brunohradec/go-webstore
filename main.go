package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brunohradec/go-webstore/controllers"
	"github.com/brunohradec/go-webstore/initializers"
	"github.com/brunohradec/go-webstore/middleware"
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
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.RegisterUser)
			auth.POST("/login", controllers.LoginUser)
			auth.GET("/me", controllers.GetCurrentUser)
		}

		users := api.Group("/users")
		users.Use(middleware.JwtAuthMiddleware())

		{
			users.GET("/:id", controllers.FindUserByID)
			users.PUT("/:id", controllers.UpdateUser)
		}

		products := api.Group("/products")
		users.Use(middleware.JwtAuthMiddleware())

		{
			products.POST("/", controllers.SaveNewProduct)
			products.GET("/", controllers.FindAllProducts)
			products.GET("/:id", controllers.FindProductByID)
			products.GET("/user/:userId", controllers.FindProductsByUserID)
			products.PUT("/:id", controllers.UpdateProductByID)
			products.DELETE("/:id", controllers.DeleteProductByID)
		}

		comments := api.Group("/comments")
		users.Use(middleware.JwtAuthMiddleware())

		{
			comments.POST("/", controllers.SaveNewComment)
			comments.GET("/product/:productId", controllers.FindCommentsByProductID)
			comments.PUT("/:id", controllers.UpdateCommentByID)
			comments.DELETE("/:id", controllers.DeleteCommentByID)
		}
	}

	r.Run(":" + env.Port)
}
