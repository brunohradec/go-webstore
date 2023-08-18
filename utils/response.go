package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RejectResponseAndLog(msg string, status int, err error, c *gin.Context) {
	log.Println(msg, err)
	c.JSON(status, gin.H{
		"message": msg,
	})
}
