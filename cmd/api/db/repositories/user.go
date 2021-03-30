package db

import (
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

type IUserRepository interface {
	GetById(id uint) (*models.User, error)
}

func NewUserRepository(d *gorm.DB) IUserRepository {
	return &UserRepository{
		database: d,
	}
}

func (r *UserRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	err := r.database.Where("id = ?", id).
		First(&user).
		Error
	return &user, err
}
