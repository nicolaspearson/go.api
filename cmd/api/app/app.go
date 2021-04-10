package app

import (
	"fmt"

	"github.com/nicolaspearson/go.api/cmd/api/config"
	"github.com/nicolaspearson/go.api/pkg/postgresql"
	"github.com/nicolaspearson/go.api/pkg/rabbitmq"
	"github.com/sirupsen/logrus"
)

func New() IApplication {
	return &Application{}
}

func (application *Application) Build() IApplication {
	application.logger = *logrus.New()

	if err := config.LoadConfig("./config"); err != nil {
		application.logger.Fatalf("Invalid application configuration: %v", err)
	}

	application.logger.Infof("Environment: %s", config.Vars.Environment)
	application.logger.Infof("ReleaseVersion: %s", config.Vars.ReleaseVersion)
	application.logger.Infof("Version: %s", config.Vars.Version)

	application.AddRabbitMq(rabbitmq.Opts{
		Username:    config.Vars.RabbitMqUsername,
		Password:    config.Vars.RabbitMqPassword,
		Host:        config.Vars.RabbitMqHost,
		VirtualHost: config.Vars.RabbitMqVirtualHost,
	})

	application.AddPostgreSql(postgresql.Opts{
		Host:     config.Vars.DbHost,
		User:     config.Vars.DbUser,
		Password: config.Vars.DbPassword,
		Database: config.Vars.DbName,
		Port:     config.Vars.DbPort,
	})

	application.AddRouter()
	application.AddControllers().InitMiddlewares().AddSwagger()

	return application
}

func (application *Application) Run() error {
	if application.broker != nil {
		defer application.broker.Close()
	}

	database, dbErr := application.db.DB()
	if dbErr != nil {
		application.logger.Fatalf("Failed to retrieve the database object: %v", dbErr)
	}
	defer database.Close()

	application.url = fmt.Sprintf("%s:%s", config.Vars.ServerHost, config.Vars.ServerPort)
	application.logger.Infof("Serving application: %s", application.url)
	err := application.engine.Run(application.url)
	if err != nil {
		application.logger.Fatalf("Failed to start application: %v", err)
	}

	return nil
}
