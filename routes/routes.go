package routes

import (
	"github.com/gin-gonic/gin"
	"inktrail/controllers"
	"inktrail/handler"
)

func SetupRoutes(app *gin.Engine) {
	app.LoadHTMLGlob("templates/*")

	app.GET("/", handler.ShowIndexPage)
	app.POST("/api/register", controllers.CreateUser)
	//app.POST("/api/login", controllers.Login)
}
