package services

import (
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	repositories "github.com/nicolaspearson/go.api/cmd/api/db/repositories"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

type IUserService interface {
	GetById(id uint) (*models.User, error)
}

// NewUserService creates a new UserService with the provided user repository.
func NewUserService(r repositories.IUserRepository) IUserService {
	return &UserService{userRepository: r}
}

// GetById retrieves the user by id using the user repository.
func (s *UserService) GetById(id uint) (*models.User, error) {
	return s.userRepository.GetById(id)
}
