package daos

import (
	"github.com/nicolaspearson/go.api/cmd/api/config"
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) GetById(id uint) (*models.User, error) {
	var user models.User

	err := config.Config.Database.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
