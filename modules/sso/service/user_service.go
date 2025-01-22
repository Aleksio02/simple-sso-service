package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"simple-sso-service/modules/sso/model"
	"simple-sso-service/modules/sso/repository"
	"strings"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (us *UserService) Register(request model.AuthRequest) error {
	// TODO: aleksioi: тут должна быть валидация полей
	request.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	err := us.UserRepo.SaveUser(request.Username, request.Password, "USER")
	return err
}

func (us *UserService) Login(request model.AuthRequest) error {
	user, err := us.UserRepo.GetUserByUsername(request.Username)
	if err != nil {
		return err
	}
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
	if hashedPassword != user.Password {
		return errors.New("неверное имя пользователя или пароль")
	}
	return nil
}

func (us *UserService) CreateTokenForUsername(username string, secret []byte) string {
	payload := jwt.MapClaims{
		"username": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, _ := token.SignedString(secret)
	return signedToken
}

func (us *UserService) ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token := strings.Split(tokenString, " ")[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Убедимся, что метод подписи соответствует ожиданиям
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return jwt.MapClaims{}, errors.New("ошибка чтения токена")
	}

	return claims, nil
}

func CreateUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		UserRepo: userRepo,
	}
}
