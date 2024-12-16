// repository/init.go
package repository

import (
	"database/sql"
	"fiber-boilerplate/app/repository/mysql"

	"fiber-boilerplate/app/repository/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	MySQL   map[string]interface{}
	MongoDB map[string]interface{}
}

func InitRepositories(dbType string, mysqlDB *sql.DB, mongoDB *mongo.Database) (*Repositories, error) {
	repos := &Repositories{
		MySQL:   make(map[string]interface{}),
		MongoDB: make(map[string]interface{}),
	}

	if dbType == "mysql" {
		mysqlRepos, err := mysql.InitRepositories(mysqlDB)
		if err != nil {
			return nil, err
		}
		repos.MySQL = mysqlRepos
	}

	if dbType == "mongodb" {
		mongoRepos, err := mongodb.InitRepositories(mongoDB)
		if err != nil {
			return nil, err
		}
		repos.MongoDB = mongoRepos
	}

	return repos, nil
}
