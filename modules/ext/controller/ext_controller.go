package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"simple-sso-service/modules/ext/connector"
	"simple-sso-service/modules/ext/service"
	"strconv"
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
	codes := url["code"]
	if codes == nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "Отсутствует авторизационный код",
		})
		return
	}
	code := codes[0]
	responseBody := connector.GetToken(service.GetService("own-sso"), code)

	// TODO: Возвращать информацию о пользователе в будущем (если будет необходимость)
	responseBodyCode, _ := strconv.Atoi(responseBody["code"])
	c.JSON(responseBodyCode, responseBody)
}
