package mongodb

import (
	"fiber-boilerplate/app/repository/mongodb"
	"log"
)

var UserServiceInstancem *UserService

func InitServices(mongoRepos map[string]interface{}) {
	// Accessing the "mongo" repository group from repos

	userRepo, ok := mongoRepos["userRepo"].(*mongodb.UserRepository)

	if !ok {
		log.Fatal("Invalid user repository instance")
	}

	// Initializing user service
	UserServiceInstancem = NewUserService(userRepo)
}