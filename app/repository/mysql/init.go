package mysql

import (
	"database/sql"
)

func InitRepositories(db *sql.DB) (map[string]interface{}, error) {
	repos := make(map[string]interface{})

	userRepository := NewUserRepository(db)
	repos["UserRepository"] = userRepository

	// Add more repositories as needed

	return repos, nil
}
