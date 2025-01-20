package main

import (
	"github.com/gin-gonic/gin"
	"simple-sso-service/modules/ext/routes"
)

func main() {
	app := gin.Default()

	routes.RegisterSsoRoutes(app)

	app.Run(":8081")
}
