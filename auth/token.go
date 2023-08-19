package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/brunohradec/go-webstore/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, secret string, ttl int) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(ttl)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ValidateToken(token string, secret string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractTokenFromRequest(c *gin.Context) (string, error) {
	token := c.Query("token")

	if token != "" {
		return token, nil
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}

	return "", errors.New("Could not find token in request query or header")
}

func ExtractUserIDFromToken(token string, secret string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if ok && parsedToken.Valid {
		// JWT libraries usially repsresent all numbers as float64 type
		claimValue, found := claims["user_id"].(float64)

		if found {
			return uint(claimValue), nil
		} else {
			return 0, errors.New("Claim containing user ID not found in provided token")
		}
	} else {
		return 0, errors.New("Provided token could not be parsed")
	}
}

func ExtractUserIDFromRequest(c *gin.Context) (uint, error) {
	token, err := ExtractTokenFromRequest(c)
	secret := shared.Env.JWT.AccessTokenSecret

	if err != nil {
		return 0, err
	}

	return ExtractUserIDFromToken(token, secret)
}
