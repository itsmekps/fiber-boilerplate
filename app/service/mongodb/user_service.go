package user

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"
)

type UserService struct {
	repository *mongodb.UserRepository
}

func NewUserService(repository *mongodb.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repository.GetUser(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repository.CreateUser(user)
}

func (s *UserService) UpdateUser(id int, user *models.User) error {
	return s.repository.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repository.DeleteUser(id)
}
