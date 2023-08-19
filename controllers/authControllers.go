package controllers

import (
	"net/http"

	"github.com/brunohradec/go-webstore/auth"
	"github.com/brunohradec/go-webstore/dtos"
	"github.com/brunohradec/go-webstore/repository"
	"github.com/brunohradec/go-webstore/shared"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {

}

func GetCurrentUser(c *gin.Context) {
	token, err := auth.ExtractTokenFromRequest(c)
	secret := shared.Env.JWT.AccessTokenSecret

	if err != nil {
		RejectResponseAndLog(
			"Could not extract JSON web token from request headers or query",
			http.StatusUnauthorized,
			err,
			c,
		)
	}

	userID, err := auth.ExtractUserIDFromToken(token, secret)

	user, err := repository.FindUserByID(userID)

	if err != nil {
		RejectResponseAndLog(
			"Could not find user with the ID extracted from JSON web token",
			http.StatusUnauthorized,
			err,
			c,
		)
	}

	c.JSON(http.StatusOK, dtos.UserModelToResponseDto(user))
}
