package middleware

import (
	"github.com/gin-gonic/gin"
	"inktrail/utils"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
