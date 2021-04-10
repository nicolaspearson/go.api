package userservice

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/nicolaspearson/go.api/cmd/api/internal/domain/userentity"
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/repository/userrepository"
	"github.com/nicolaspearson/go.api/pkg/rabbitmq"
)

func New(broker rabbitmq.IRabbitMq, repository userrepository.IUserRepository) IUserService {
	return UserService{broker: broker, repository: repository}
}

func (service UserService) Create(userDto UserCreateRequestDto) error {
	var entity userentity.Entity
	err := mapstructure.Decode(userDto, &entity)
	if err != nil {
		return err
	}

	event, err := service.repository.Create(entity)
	if err != nil {
		return err
	}

	jsonEvent, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	err = service.broker.Publish(event.ExchangeName, jsonEvent)
	if err != nil {
		fmt.Printf("An error occurred while throwing the event! Event: %+v, Error: %v+\n", event, err)
		return err
	}
	return nil
}

func (service UserService) GetAll() ([]UserDto, error) {
	dto := make([]UserDto, 0)

	users, err := service.repository.GetAll()
	if err != nil {
		return dto, err
	}

	err = mapstructure.Decode(users, &dto)
	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (service UserService) GetByUuid(uuid uuid.UUID) (*UserDto, error) {
	user, err := service.repository.GetByUuid(uuid)
	if err != nil {
		return nil, err
	}

	var dto UserDto
	err = mapstructure.Decode(user, &dto)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}
