package routes

import (
	"github.com/gin-gonic/gin"
	"simple-sso-service/modules/ext/controller"
)

func RegisterSsoRoutes(router *gin.Engine) {
	router.GET("/login", controller.Login)
	router.GET("/callback", controller.Callback)
}
