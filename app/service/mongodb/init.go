package mongodb

import (
	"fiber-boilerplate/app/repository/mongodb"
	"log"
)

var UserServiceInstance *UserService
var PlayerServiceInstance *PlayerService

func InitServices(mongoRepos map[string]interface{}) {
	// Accessing the "mongo" repository group from repos

	userRepo, ok := mongoRepos["userRepo"].(*mongodb.UserRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}
	playerRepo, ok := mongoRepos["playerRepo"].(*mongodb.PlayerRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}

	// Initializing user service
	UserServiceInstance = NewUserService(userRepo)
	PlayerServiceInstance = NewPlayerService(playerRepo)
}
