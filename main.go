package main

import (
	"github.com/gin-gonic/gin"
	"inktrail/config"
	"inktrail/routes"
)

func main() {
	app := gin.Default()

	config.ConnectDB()

	routes.SetupRoutes(app)
	err := app.Run()
	if err != nil {
		return
	}
}
