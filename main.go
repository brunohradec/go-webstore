package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brunohradec/go-webstore/initializers"
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

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + env.Port)
}
