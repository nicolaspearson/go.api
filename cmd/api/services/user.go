package services

import (
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	repositories "github.com/nicolaspearson/go.api/cmd/api/db/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

// NewUserService creates a new UserService with the provided user repository.
func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{repository}
}

// GetById retrieves the user by id using the user repository.
func (s *UserService) GetById(id uint) (*models.User, error) {
	return s.repository.GetById(id)
}
