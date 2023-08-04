package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"inktrail/config"
	"inktrail/models"
	"inktrail/repositories"
	"strings"
)

func ValidateJWT(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func CurrentUser(c *gin.Context) (models.User, error) {
	err := ValidateJWT(c)
	if err != nil {
		return models.User{}, err
	}
	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)

	// Add debug statement to check the claims
	fmt.Printf("Claims: %+v\n", claims)

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return models.User{}, errors.New("invalid user ID in claims")
	}

	// Add debug statement to check the userID
	fmt.Printf("UserID: %v\n", userId)

	userID := uint(userId)
	db := config.DB
	user, err := repositories.FindUserById(db, userID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Config("SECRET")), nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
