package mysql

import (
	"fiber-boilerplate/app/repository/mysql"
	"log"
)

var UserServiceInstance *UserService

func InitServices(mysqlRepos map[string]interface{}) {

	// Accessing the "mysql" repository group from repos
	// mysqlRepos, ok := repos["mysql"].(map[string]interface{})
	// if !ok {
	// 	log.Fatal("MySQL repositories not found or invalid type")
	// }

	// Extracting the user repository
	userRepo, ok := mysqlRepos["userRepo"].(*mysql.UserRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}

	// Initializing user service
	UserServiceInstance = NewUserService(userRepo)

	// Initialize other services...
}
