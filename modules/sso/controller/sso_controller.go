package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/url"
	"simple-sso-service/modules/sso/model"
	"simple-sso-service/modules/sso/repository"
	"simple-sso-service/modules/sso/service"
	"simple-sso-service/modules/sso/utils"
	"strings"
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = userService.Login(requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
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
		code        int
		redirectUri string
	}

	err = userService.Register(requestBody)
	if err != nil {
		// TODO: aleksioi: выкинуть внятную ошибку
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	responseBody := responseBodyStruct{code: 302, redirectUri: "uri to login form page"}

	c.JSON(http.StatusOK, responseBody)
}

func Token(c *gin.Context) {
	url, _ := url.ParseQuery(c.Request.URL.RawQuery)
	code := url["code"][0]
	if username, exists := authCodes[code]; exists {
		delete(authCodes, code)
		payload := jwt.MapClaims{
			"username": username,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		signedToken, _ := token.SignedString(JWT_SECRET)
		c.JSON(http.StatusOK, map[string]string{
			"token": signedToken,
		})
		return
	}
	c.JSON(http.StatusBadRequest, "Invalid code")
}

func AuthInfo(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) == 0 {
		c.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	token := strings.Split(tokenString, " ")[1]

	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Убедимся, что метод подписи соответствует ожиданиям
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})
	c.JSON(http.StatusOK, map[string]interface{}{
		"username": claims["username"],
	})
}
