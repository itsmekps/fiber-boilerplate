package service

import (
	"fiber-boilerplate/app/repository"
	"log"
)

var UserServiceInstance *UserService

func InitServices(repos map[string]interface{}) {
	// Initializing user service
	userRepo, ok := repos["userRepo"].(*repository.UserRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}
	UserServiceInstance = NewUserService(userRepo)

	// Initialize other services...
}
