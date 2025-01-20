package routes

import (
	"github.com/gin-gonic/gin"
	"simple-sso-service/modules/sso/controller"
)

func RegisterSsoRoutes(router *gin.Engine) {

	router.GET("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.GET("/token", controller.Token)
	router.GET("/authInfo", controller.AuthInfo)
}
