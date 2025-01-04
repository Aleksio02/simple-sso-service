package routes

import (
	"github.com/gin-gonic/gin"
	"simple-sso-service/modules/sso/controller"
)

func RegisterSsoRoutes(router *gin.Engine) {

	router.POST("/login", controller.HandleLoginRequest)
}
