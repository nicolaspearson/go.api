package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/cmd/api/config"
	"github.com/nicolaspearson/go.api/cmd/api/controllers"
	"github.com/nicolaspearson/go.api/cmd/api/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Golang Starter API
// @version 1.0
// @description Swagger API documentation for the Golang Starter API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email nic.s.pearson@gmail.com
// @license.name MIT
// @license.url https://github.com/nicolaspearson/go.api/blob/master/LICENSE
// @BasePath /api/v1
func main() {
	if err := config.LoadConfig("./config"); err != nil {
		log.Fatalf("Invalid application configuration: %v", err)
	}
	log.Printf("Environment: %s", config.Vars.Environment)
	log.Printf("ReleaseVersion: %s", config.Vars.ReleaseVersion)
	log.Printf("Version: %s", config.Vars.Version)

	// Creates a router without any middleware by default
	e := gin.New()
	// Logger middleware will write the logs to gin.DefaultWriter even if you set GIN_MODE=release.
	e.Use(gin.Logger())
	// Recovery middleware recovers from any panics and returns a 500 if there was one.
	e.Use(gin.Recovery())

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users/:id", controllers.GetById)
	}

	config.Config.Database = db.Setup()
	sqlDatabase, err := config.Config.Database.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve generic database object: %v", err)
	}
	defer sqlDatabase.Close()

	e.Run(fmt.Sprintf("%s:%s", config.Vars.ServerHost, config.Vars.ServerPort))
}
