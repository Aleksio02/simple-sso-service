package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"simple-sso-service/modules/sso/model"
	"simple-sso-service/modules/sso/repository"
	"simple-sso-service/modules/sso/service"
	"simple-sso-service/modules/sso/utils"
	"strconv"
)

// TODO: вынести объявление в другое место
var JWT_SECRET = []byte("someappsecret")

var userService = service.CreateUserService(repository.CreateSQLiteUserRepository())

var authCodes = map[string]string{}

func Login(c *gin.Context) {
	url, _ := url.ParseQuery(c.Request.URL.RawQuery)
	redirectUri := url["redirect_uri"][0]

	requestBody := model.AuthRequest{}
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorObject(strconv.Itoa(http.StatusBadRequest), "Ошибка чтения объекта из тела запроса"))
		return
	}

	err = userService.Login(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorObject(strconv.Itoa(http.StatusBadRequest), "Неверное имя пользователя или пароль"))
		return
	}
	code := utils.GenerateAuthCode()
	authCodes[code] = requestBody.Username
	c.Redirect(http.StatusPermanentRedirect, redirectUri+"?code="+code)
}

func Register(c *gin.Context) {
	var requestBody model.AuthRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		panic(err)
		return
	}

	type responseBodyStruct struct {
		Code        int    `json:"code"`
		RedirectUri string `json:"redirectUri"`
	}

	err = userService.Register(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorObject(strconv.Itoa(http.StatusBadRequest), "Пользователь с таким именем уже существует"))
		return
	}

	responseBody := responseBodyStruct{Code: 200, RedirectUri: "uri to login form page"}
	c.JSON(http.StatusOK, responseBody)
}

func Token(c *gin.Context) {
	url, _ := url.ParseQuery(c.Request.URL.RawQuery)
	code := url["code"][0]
	if username, exists := authCodes[code]; exists {
		delete(authCodes, code)
		signedToken := userService.CreateTokenForUsername(username, JWT_SECRET)
		c.JSON(http.StatusOK, map[string]string{
			"token": signedToken,
		})
		return
	}
	c.JSON(http.StatusBadRequest, getErrorObject(strconv.Itoa(http.StatusBadRequest), "Некорректный авторизационный код"))
}

func AuthInfo(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) == 0 {
		c.JSON(http.StatusBadRequest, getErrorObject(strconv.Itoa(http.StatusBadRequest), "Отсутствует токен"))
		return
	}

	claims, err := userService.ParseToken(tokenString, JWT_SECRET)

	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorObject(strconv.Itoa(http.StatusBadRequest), "Ошибка чтения токена"))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"username": claims["username"],
	})
}

func getErrorObject(code string, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}
