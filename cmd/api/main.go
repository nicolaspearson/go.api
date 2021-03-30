package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/cmd/api/config"
	controllers "github.com/nicolaspearson/go.api/cmd/api/controllers"
	"github.com/nicolaspearson/go.api/cmd/api/db"
	repositories "github.com/nicolaspearson/go.api/cmd/api/db/repositories"
	_ "github.com/nicolaspearson/go.api/cmd/api/docs"
	services "github.com/nicolaspearson/go.api/cmd/api/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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

	// Set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router with logger and recovery middleware attached
	r := gin.Default()

	d := db.Setup()
	sqlDatabase, err := d.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve generic database object: %v", err)
	}
	defer sqlDatabase.Close()

	initializeRoutes(r, d)

	r.Run(fmt.Sprintf("%s:%s", config.Vars.ServerHost, config.Vars.ServerPort))
}

func initializeRoutes(r *gin.Engine, d *gorm.DB) {
	userRepository := repositories.NewUserRepository(d)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", config.Vars.ServerHost, config.Vars.ServerPort))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", userController.GetById)
	}
}
