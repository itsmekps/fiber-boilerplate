package mysql

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mysql"
)

type UserService struct {
	userRepo *mysql.UserRepository
}

func NewUserService(userRepo *mysql.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	// Use the user repository to retrieve the user
	return s.userRepo.GetUser(id)
}
