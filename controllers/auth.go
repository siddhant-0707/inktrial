package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"inktrail/config"
	"inktrail/models"
	"net/http"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CreateUser new user
func CreateUser(c *gin.Context) {
	/*	type NewUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}*/

	db := config.DB
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: newUser.Username,
		Password: hashedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

// Login User
/*func Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	for _, user := range models.Users {
		if user.Username == loginReq.Username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				return
			}

			// User authenticated, return token (if using JWT)

			c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
}*/
