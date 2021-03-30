package db

import (
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

type UserRepository interface {
	GetById(id uint) (*models.User, error)
}

func NewUserRepository(d *gorm.DB) UserRepository {
	return &userRepository{
		database: d,
	}
}

func (r *userRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	err := r.database.Where("id = ?", id).
		First(&user).
		Error
	return &user, err
}
