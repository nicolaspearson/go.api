package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/cmd/api/controllers/indexcontroller"
	"github.com/nicolaspearson/go.api/cmd/api/controllers/usercontroller"
	"github.com/nicolaspearson/go.api/cmd/api/internal/application/userservice"
	"github.com/nicolaspearson/go.api/cmd/api/internal/infrastructure/repository/userrepository"
	_ "github.com/nicolaspearson/go.api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (application *Application) AddControllers() *Application {
	application.InitIndexController()
	application.InitUserController()
	return application
}

func (application *Application) InitIndexController() {
	indexcontroller.New().Init(application.engine)
}

func (application *Application) InitUserController() {
	userRepository := userrepository.New(application.db)
	userService := userservice.New(application.broker, userRepository)
	usercontroller.New(userService).Init(application.engine)
}

func (application *Application) InitMiddlewares() *Application {
	application.engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, PATCH, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	return application
}

func (application *Application) AddSwagger() *Application {
	application.engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))
	application.engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	})

	return application
}
