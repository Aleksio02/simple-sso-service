package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"simple-sso-service/modules/sso/controller"
)

func RegisterSsoRoutes(router *gin.Engine) {

	store := cookie.NewStore([]byte("secret")) // "secret" — ключ шифрования
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", controller.RenderLoginPage)
	router.GET("/register", controller.RenderRegisterPage)
	router.POST("/login", controller.HandleLoginRequest)
}
