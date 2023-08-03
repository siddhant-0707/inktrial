package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(app *gin.Engine) {
	app.LoadHTMLGlob("templates/*")

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			gin.H{
				"title": "Home Page",
			})
	})
}
