package authutils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetPrincipalIDFromRequest(c *gin.Context) uint {
	principalID, _ := c.Get("request-principal-id")
	return principalID.(uint)
}
