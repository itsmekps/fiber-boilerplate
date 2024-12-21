// service/init.go
package service

import (
	"fiber-boilerplate/app/service/mongodb"
	"fiber-boilerplate/app/service/mysql"
	"log"
)

func InitServices(repos map[string]interface{}) {
	// Initialize MySQL services
	mysqlRepos, ok := repos["mysql"].(map[string]interface{})
	if !ok {
		log.Fatal("MySQL repositories not found or invalid type")
	}

	mongoRepos, ok := repos["mongodb"].(map[string]interface{})
	if !ok {
		log.Fatal("MySQL repositories not found or invalid type")
	}

	mysql.InitServices(mysqlRepos)
	mongodb.InitServices(mongoRepos)

	// Initialize MongoDB services
	// mongo.InitServices(repos.(mapinterface{}))
}
