package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brunohradec/go-webstore/controllers"
	"github.com/brunohradec/go-webstore/infrastructure"
	"github.com/brunohradec/go-webstore/middleware"
	"github.com/brunohradec/go-webstore/repositories"
	"github.com/brunohradec/go-webstore/services"
	"github.com/gin-gonic/gin"
)

func main() {
	env, err := infrastructure.Environment()

	if err != nil {
		log.Fatal("Error loading dotenv file")
		os.Exit(1)
	}

	DB, err := infrastructure.ConnectToDB(
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

	infrastructure.AutomigrateDB(DB)

	userRepository := repositories.InitUserRepository(DB)
	productRepository := repositories.InitProductRepository(DB)
	commentRepository := repositories.InitCommentRepository(DB)

	userService := services.InitUserService(userRepository)
	productService := services.InitProductService(productRepository)
	commentService := services.InitCommentService(commentRepository)
	authService := services.InitAuthService(userRepository, env)

	userController := controllers.InitUserController(userService)
	productController := controllers.InitProductController(productService)
	commentController := controllers.InitCommentController(commentService, userService)
	authController := controllers.InitAuthController(authService, userService)

	r := gin.Default()

	authMiddleware := middleware.JwtAuthMiddleware(env)

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)

			me := auth.Group("/me")
			me.Use(authMiddleware)

			me.GET("/", authController.Me)
		}

		users := api.Group("/users")
		users.Use(authMiddleware)

		{
			users.GET("/:id", userController.FindByID)
			users.PUT("/", userController.UpdateCurrent)
		}

		products := api.Group("/products")
		products.Use(authMiddleware)

		{
			products.POST("/", productController.Save)
			products.GET("/", productController.FindAll)
			products.GET("/:id", productController.FindByID)
			products.GET("/user/:userId", productController.FindByUserID)
			products.PUT("/:id", productController.UpdateByID)
			products.DELETE("/:id", productController.DeleteByID)
		}

		comments := api.Group("/comments")
		comments.Use(authMiddleware)

		{
			comments.POST("/", commentController.Save)
			comments.GET("/product/:productId", commentController.FindByProductID)
			comments.PUT("/:id", commentController.UpdateByID)
			comments.DELETE("/:id", commentController.DeleteByID)
		}
	}

	r.Run(":" + env.Port)
}
