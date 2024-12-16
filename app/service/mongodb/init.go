package mongodb

import (
	"fiber-boilerplate/app/repository"
)

type Services struct {
	UserService UserService
}

func InitServices(repos map[string]interface{}) (*Services, error) {
	services := &Services{}

	userRepository := repos["UserRepository"].(repository.UserRepository)
	services.UserService = NewUserService(userRepository)

	return services, nil
}
