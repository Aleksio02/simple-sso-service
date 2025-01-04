package service

import (
	"simple-sso-service/modules/sso/model"
	"simple-sso-service/modules/sso/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (us *UserService) Register(request model.AuthRequest) error {
	// TODO: aleksioi: тут должна быть валидация полей и хэширование пароля
	err := us.UserRepo.SaveUser(request.Username, request.Password, "USER")
	return err
}

func CreateUserService(userRepo repository.UserRepository) UserService {
	return UserService{
		UserRepo: userRepo,
	}
}
