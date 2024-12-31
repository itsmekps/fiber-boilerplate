// service/mongo/user_service.go
package mongodb

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"
	// "log"
)

// var PlayerServiceInstance *PlayerService

type PlayerService struct {
	repo *mongodb.PlayerRepository
}

func NewPlayerService(repo *mongodb.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

// User-related methods
func (s *PlayerService) GetPlayersList(page, limit int) ([]models.PlayerList, error) {
	return s.repo.GetPlayersList(page, limit)
}

func (s *PlayerService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}
