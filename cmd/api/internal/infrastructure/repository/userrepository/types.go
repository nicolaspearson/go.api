package userrepository

import (
	"github.com/google/uuid"
	"github.com/nicolaspearson/go.api/cmd/api/internal/domain/userentity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(entity userentity.Entity) (*userentity.CreatedEvent, error)
	GetAll() ([]userentity.Entity, error)
	GetByUuid(uuid uuid.UUID) (*userentity.Entity, error)
}

type UserRepository struct {
	db *gorm.DB
}
