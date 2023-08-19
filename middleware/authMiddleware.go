package middleware

import (
	"net/http"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/shared"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessSecret := shared.Env.JWT.AccessTokenSecret
		token, err := auth.ExtractTokenFromRequest(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Could not extract token from request header or query",
			})
			c.Abort()

			return
		}

		err = auth.ValidateToken(token, accessSecret)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Provided JSON web token is not valid",
			})
			c.Abort()

			return
		}

		c.Next()
	}
}
