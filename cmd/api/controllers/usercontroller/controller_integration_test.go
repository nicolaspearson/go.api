package usercontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/cmd/api/containers"
	"github.com/nicolaspearson/go.api/cmd/api/internal/application/userservice"
	"github.com/nicolaspearson/go.api/cmd/api/internal/domain/userentity"
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/brokerconsts"
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/repository/userrepository"
	"github.com/nicolaspearson/go.api/pkg/rabbitmq"
	"github.com/nicolaspearson/go.api/pkg/util"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type IntegrationSuite struct {
	suite.Suite
	engine     *gin.Engine
	controller IUserController

	rabbitmqContainer   containers.IRabbitMqContainer
	postgresqlContainer containers.IPostgreSqlContainer

	db     *gorm.DB
	broker rabbitmq.IRabbitMq
}

func TestInit(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupSuite() {
	pool, err := dockertest.NewPool("")
	require.NoError(s.T(), err)

	s.rabbitmqContainer = containers.NewRabbitMqContainer(pool)
	s.postgresqlContainer = containers.NewPostgresqlContainer(pool)

	err = s.rabbitmqContainer.Create()
	require.NoError(s.T(), err)

	err = s.postgresqlContainer.Create()
	require.NoError(s.T(), err)

	s.broker = s.rabbitmqContainer.Connect()
	err = s.broker.Bind(brokerconsts.UserCreatedExchangeName, brokerconsts.UserCreatedQueueName)
	require.NoError(s.T(), err)

	s.db = s.postgresqlContainer.Connect()

	err = s.postgresqlContainer.AutoMigrate(s.db)
	require.NoError(s.T(), err)

	userRepository := userrepository.New(s.db)
	userService := userservice.New(s.broker, userRepository)

	s.controller = New(userService)

	s.engine = gin.Default()
	s.controller.Init(s.engine)
}

func (s *IntegrationSuite) AfterTest(_, _ string) {
	s.postgresqlContainer.Flush(s.db)

	err := s.rabbitmqContainer.Flush(brokerconsts.UserCreatedQueueName)
	require.NoError(s.T(), err)
}

func (s *IntegrationSuite) Test_GetAllUsers_ReturnsSuccess() {
	req, err := http.NewRequest(http.MethodGet, "/api/users", nil)
	require.Nil(s.T(), err)

	util.HttpRecorder(s.T(), s.engine, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func (s *IntegrationSuite) Test_CreateUser_ReturnsSuccess() {
	userCreateRequestDto := userservice.UserCreateRequestDto{
		FirstName: "Nicolas",
		LastName:  "Pearson",
		Email:     "nic.s.pearson@gmail.com",
		Password:  "!Test1234!",
	}
	jsonUser, err := json.Marshal(userCreateRequestDto)
	require.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(jsonUser))
	require.Nil(s.T(), err)

	util.HttpRecorder(s.T(), s.engine, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusCreated
		if statusOK {
			err = s.broker.Consume(brokerconsts.UserCreatedQueueName, 1, func(message []byte) bool {
				var event userentity.CreatedEvent
				err = json.Unmarshal(message, &event)
				require.NoError(s.T(), err)

				var entity userentity.Entity
				result := s.db.First(&entity, event.Uuid)
				require.NoError(s.T(), result.Error)

				require.Equal(s.T(), entity.Uuid.String(), event.Uuid.String())
				require.Equal(s.T(), entity.Email, userCreateRequestDto.Email)
				return false // Stop consuming
			})

			require.NoError(s.T(), err)
		}

		return statusOK
	})
}
