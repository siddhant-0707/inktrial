package main

import (
	"github.com/gin-gonic/gin"
	"inktrail/config"
	"inktrail/routes"
)

func main() {
	/*	r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		err := r.Run()
		if err != nil {
			return
		} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")*/

	app := gin.Default()

	config.ConnectDB()

	routes.SetupRoutes(app)
	err := app.Run()
	if err != nil {
		return
	}
}
