package middleware

import (
	"net/http"

	"github.com/brunohradec/go-webstore/authutils"
	"github.com/brunohradec/go-webstore/infrastructure"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(env *infrastructure.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := env.JWT.AccessTokenSecret
		token, err := authutils.ExtractTokenFromRequest(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Could not extract token from request header or query",
			})
			c.Abort()

			return
		}

		err = authutils.ValidateToken(token, secret)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Provided JSON web token is not valid",
			})
			c.Abort()

			return
		}

		principalID, err := authutils.ExtractUserIDFromToken(token, secret)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Provided JSON web token is not valid",
			})
			c.Abort()

			return
		}

		c.Set("request-token", token)
		c.Set("request-principal-id", principalID)

		c.Next()
	}
}
