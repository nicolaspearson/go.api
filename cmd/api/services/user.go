package services

import (
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	repositories "github.com/nicolaspearson/go.api/cmd/api/db/repositories"
)

type userService struct {
	userRepository repositories.UserRepository
}

type UserService interface {
	GetById(id uint) (*models.User, error)
}

// NewUserService creates a new UserService with the provided user repository.
func NewUserService(r repositories.UserRepository) UserService {
	return &userService{userRepository: r}
}

// GetById retrieves the user by id using the user repository.
func (s *userService) GetById(id uint) (*models.User, error) {
	return s.userRepository.GetById(id)
}
