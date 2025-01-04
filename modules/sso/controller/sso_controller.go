package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-sso-service/modules/sso/model"
	"simple-sso-service/modules/sso/repository"
	"simple-sso-service/modules/sso/service"
)

var userService = service.CreateUserService(repository.CreateSQLiteUserRepository())

func Login(c *gin.Context) {
	// TODO: implement me
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

	responseBody := responseBodyStruct{code: 302, redirectUri: "uri to login page"}

	c.JSON(http.StatusOK, responseBody)
}

func Token(c *gin.Context) {
	// TODO: implement me
}

func AuthInfo(c *gin.Context) {
	// TODO: implement me
}
