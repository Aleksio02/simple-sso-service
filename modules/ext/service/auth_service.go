package service

type AuthService interface {
	getServiceLink() string
	GetLoginLink() string
	GetTokenLink() string
}

type SsoAuthService struct {
}

func (s SsoAuthService) getServiceLink() string {
	return "http://localhost:8080"
}

func (s SsoAuthService) GetLoginLink() string {
	// TODO: указать url страницы авторизации в будущем
	return s.getServiceLink() + "/login"
}

func (s SsoAuthService) GetTokenLink() string {
	return s.getServiceLink() + "/token"
}

func GetService(serviceName string) AuthService {
	return SsoAuthService{}
}
