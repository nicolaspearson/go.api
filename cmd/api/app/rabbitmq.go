package app

import (
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/brokerconsts"
	"github.com/nicolaspearson/go.api/pkg/rabbitmq"
)

func (application *Application) AddRabbitMq(opts rabbitmq.Opts) *Application {
	broker, err := rabbitmq.New(opts)
	if err != nil {
		application.logger.Error("An error occurred while connecting to rabbitmq! ", err)
		return application
	}
	application.logger.Infoln("Broker connection successfully established!")

	application.broker = broker
	application.InitUserCreatedEvent()

	return application
}

func (application *Application) InitUserCreatedEvent() {
	err := application.broker.Bind(brokerconsts.UserCreatedExchangeName, brokerconsts.UserCreatedQueueName)
	if err != nil {
		application.logger.Error("An error occurred while binding to the exchange! ", err)
	}
}
