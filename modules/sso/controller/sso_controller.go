package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderLoginPage(c *gin.Context) {
	session := sessions.Default(c)
	//var data struct {
	//	RedirectLink string `json:"redirect-link"`
	//}

	//c.ShouldBindJSON(&data)
	//fmt.Println(data)
	//session.Set("redirect-link", data.RedirectLink)

	session.Set("redirect-link", "http://localhost:8081/somePageToRedirect")
	session.Save()
	c.HTML(http.StatusOK, "login.html", nil)
}

func RenderRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func HandleLoginRequest(c *gin.Context) {
	session := sessions.Default(c)
	redirectLink := session.Get("redirect-link").(string)
	session.Delete("redirect-link")
	c.Redirect(http.StatusFound, redirectLink)
}
