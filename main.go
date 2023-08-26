package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brunohradec/go-webstore/controllers"
	"github.com/brunohradec/go-webstore/initializers"
	"github.com/brunohradec/go-webstore/middleware"
	"github.com/brunohradec/go-webstore/repositories"
	"github.com/brunohradec/go-webstore/services"
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

	userRepository := repositories.InitUserRepository(shared.DB)
	productRepository := repositories.InitProductRepository(shared.DB)
	commentRepository := repositories.InitCommentRepository(shared.DB)

	userService := services.InitUserService(userRepository)
	productService := services.InitProductService(productRepository)
	commentService := services.InitCommentService(commentRepository)

	userController := controllers.InitUserController(userService)
	productController := controllers.InitProductController(productService)
	commentController := controllers.InitCommentController(commentService, userService)
	authController := controllers.InitAuthController(userService)

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
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.GET("/me", authController.Me)
		}

		users := api.Group("/users")
		users.Use(middleware.JwtAuthMiddleware())

		{
			users.GET("/:id", userController.FindByID)
			users.PUT("/", userController.UpdateCurrent)
		}

		products := api.Group("/products")
		users.Use(middleware.JwtAuthMiddleware())

		{
			products.POST("/", productController.Save)
			products.GET("/", productController.FindAll)
			products.GET("/:id", productController.FindByID)
			products.GET("/user/:userId", productController.FindByUserID)
			products.PUT("/:id", productController.UpdateByID)
			products.DELETE("/:id", productController.DeleteByID)
		}

		comments := api.Group("/comments")
		users.Use(middleware.JwtAuthMiddleware())

		{
			comments.POST("/", commentController.Save)
			comments.GET("/product/:productId", commentController.FindByProductID)
			comments.PUT("/:id", commentController.UpdateByID)
			comments.DELETE("/:id", commentController.DeleteByID)
		}
	}

	r.Run(":" + env.Port)
}
