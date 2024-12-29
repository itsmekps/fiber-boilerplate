// service/mongo/user_service.go
package mongodb

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"
	// "log"
)

// var UserServiceInstance *UserService

type UserService struct {
	repo *mongodb.UserRepository
}

func NewUserService(repo *mongodb.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// User-related methods
func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

// func (s *UserService) CreateUser(user *User) (*User, error) {
// 	// Implement logic using s.repo
// }

// func (s *UserService) UpdateUser(user *User) (*User, error) {
// 	// Implement logic using s.repo
// }

// func (s *UserService) DeleteUser(username string) error {
// 	// Implement logic using s.repo
// }

type User struct {
	// Define user fields (e.g., ID, Username, Email)
}
