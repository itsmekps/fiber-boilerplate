package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRepositories(client *mongo.Client) (map[string]interface{}, error) {
	db := client.Database("clg")
	repos := map[string]interface{}{
		"userRepo": NewUserRepository(db.Collection("users")),
		// Add other repository instances here...
	}
	return repos, nil
}
