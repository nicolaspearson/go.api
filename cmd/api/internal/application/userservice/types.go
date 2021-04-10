package userservice

import (
	"github.com/google/uuid"
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/repository/userrepository"
	"github.com/nicolaspearson/go.api/pkg/rabbitmq"
)

type IUserService interface {
	Create(entity UserCreateRequestDto) error
	GetAll() ([]UserDto, error)
	GetByUuid(uuid uuid.UUID) (*UserDto, error)
}

type UserService struct {
	broker     rabbitmq.IRabbitMq
	repository userrepository.IUserRepository
}
