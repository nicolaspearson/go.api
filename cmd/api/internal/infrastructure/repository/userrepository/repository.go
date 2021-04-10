package userrepository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nicolaspearson/go.api/cmd/api/internal/domain/userentity"
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/brokerconsts"
	"gorm.io/gorm"
)

func New(db *gorm.DB) IUserRepository {
	return UserRepository{db: db}
}

func (repository UserRepository) Create(entity userentity.Entity) (*userentity.CreatedEvent, error) {
	transaction := repository.db.Begin()
	defer func() (*userentity.CreatedEvent, error) {
		if r := recover(); r != nil {
			transaction.Rollback()
		}

		return nil, errors.New("an error occurred while creating a new user")
	}()

	entity.Uuid = uuid.New()
	transaction.Create(&entity)
	if err := transaction.Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return &userentity.CreatedEvent{
		ExchangeName: brokerconsts.UserCreatedExchangeName,
		Uuid:         entity.Uuid,
	}, nil
}

func (repository UserRepository) GetAll() ([]userentity.Entity, error) {
	var users []userentity.Entity
	result := repository.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repository UserRepository) GetByUuid(uuid uuid.UUID) (*userentity.Entity, error) {
	var entity userentity.Entity
	result := repository.db.First(&entity, uuid)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}
