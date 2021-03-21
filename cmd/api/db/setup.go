package db

import (
	"fmt"
	"log"

	"github.com/nicolaspearson/go.api/cmd/api/config"
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup creates the database connection
func Setup() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		config.Vars.DbHost,
		config.Vars.DbPort,
		config.Vars.DbUser,
		config.Vars.DbName,
		config.Vars.DbPassword,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection successfully established")
	database.AutoMigrate(&models.User{})
	return database
}
