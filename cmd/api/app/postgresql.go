package app

import (
	"fmt"

	"github.com/nicolaspearson/go.api/cmd/api/internal/domain/userentity"
	"github.com/nicolaspearson/go.api/pkg/postgresql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (application *Application) AddPostgreSql(opts postgresql.Opts) *Application {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		opts.Host, opts.Port, opts.User, opts.Database, opts.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		application.logger.Fatalf("Failed to connect to the database: %v", err)
	}
	application.logger.Infoln("Database connection successfully established!")

	application.db = db

	application.RunMigrations()
	return application
}

func (application *Application) RunMigrations() bool {
	err := application.db.AutoMigrate(userentity.Entity{})
	if err != nil {
		application.logger.Fatalf("An error occurred while running migrations: %v", err)
	}
	return true
}
