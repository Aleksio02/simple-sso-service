package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"simple-sso-service/modules/ext/connector"
	"simple-sso-service/modules/ext/service"
)

var REDIRECT_URI = "http://localhost:8081/callback"

func Login(c *gin.Context) {
	redirectLink := fmt.Sprintf("%s?redirect_uri=%s", service.GetService("own-sso").GetLoginLink(), REDIRECT_URI)
	c.JSON(http.StatusOK, map[string]any{
		"redirect_uri": redirectLink,
	})
}

func Callback(c *gin.Context) {
	url, _ := url.ParseQuery(c.Request.URL.RawQuery)
	code := url["code"][0]

	responseBody := connector.GetToken(service.GetService("own-sso"), code)

	// TODO: Возвращать информацию о пользователе в будущем (если будет необходимость)
	c.JSON(http.StatusOK, responseBody)
}
