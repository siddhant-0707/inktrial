package routes

import (
	"github.com/gin-gonic/gin"
	"inktrail/controllers"
	"inktrail/handler"
	"inktrail/middleware"
)

func SetupRoutes(app *gin.Engine) {
	app.LoadHTMLGlob("templates/*")

	app.GET("/", handler.ShowIndexPage)

	publicRoutes := app.Group("/auth")
	publicRoutes.POST("/register", controllers.CreateUser)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := app.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.POST("/blog", controllers.AddBlog)
	protectedRoutes.GET("/blog", controllers.GetAllBlogs)
	protectedRoutes.GET("/blog/:id", controllers.GetBlogByID)

	protectedRoutes.POST("/:id/comment", controllers.AddComment)
	protectedRoutes.GET("/:id/comment", controllers.GetCommentsByBlogID)
}
