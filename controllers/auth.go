package controllers

import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"inktrail/config"
	"inktrail/models"
	"net/http"
	"time"
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

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsername(u string) (*models.User, error) {
	db := config.DB
	var user models.User
	if err := db.Where(&models.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Login get user and password
func Login(c *gin.Context) {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error on login request",
			"data":    err.Error(),
		})
		return
	}

	identity := input.Username
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	userModel, err = getUserByUsername(identity)

	if userModel == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "User not found",
			"data":    err.Error(),
		})
		return
	} else {
		userData = UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Password: userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid password",
			"data":    nil,
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Success login",
		"data":    t,
	})
	return
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
