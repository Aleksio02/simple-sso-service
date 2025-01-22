package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"simple-sso-service/modules/sso/model"
	"simple-sso-service/modules/sso/repository"
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
		return errors.New("incorrect username or password")
	}
	return nil
}

func CreateUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		UserRepo: userRepo,
	}
}
