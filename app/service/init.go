// service/init.go
package service

import (
	"fiber-boilerplate/app/repository"
	"fiber-boilerplate/app/service/mysql"
)

type Services struct {
	UserService *service.UserService
}

func InitServices(repos *repository.Repositories) (*Services, error) {
	services := &Services{}

	if repos.MySQL != nil {
		mysqlServices, err := mysql.InitServices(repos.MySQL)
		if err != nil {
			return nil, err
		}
		services.UserService = mysqlServices.UserService
	}

	if repos.MongoDB != nil {
		mongoServices, err := mongodb.InitServices(repos.MongoDB)
		if err != nil {
			return nil, err
		}
		services.UserService = mongoServices.UserService
	}

	return services, nil
}
