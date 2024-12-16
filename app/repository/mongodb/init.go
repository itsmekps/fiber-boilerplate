// repository/mongodb/init.go
package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRepositories(mongoDB *mongo.Database) (map[string]interface{}, error) {
	repos := make(map[string]interface{})

	userRepository := NewUserRepository(mongoDB.Collection("users"))
	repos["UserRepository"] = userRepository

	// Add more repositories as needed

	return repos, nil
}
