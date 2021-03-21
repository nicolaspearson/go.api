package db

import (
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (*models.User, error)
}

type repo struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *repo {
	return &repo{
		Database: database,
	}
}

func (r *repo) GetById(id uint) (*models.User, error) {
	var user models.User

	err := r.Database.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
