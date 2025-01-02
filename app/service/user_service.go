package service

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (s *UserService) GetUser(id primitive.ObjectID) (*models.User, error) {
	fmt.Println("objID service: ", id)

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
