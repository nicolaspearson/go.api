package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/pkg/rabbitmq"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IApplication interface {
	Build() IApplication
	Run() error
	RunMigrations() bool
}

type Application struct {
	broker rabbitmq.IRabbitMq
	db     *gorm.DB
	engine *gin.Engine
	logger logrus.Logger
	url    string
}
