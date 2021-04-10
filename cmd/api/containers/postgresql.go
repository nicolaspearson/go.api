package containers

import (
	"fmt"
	"log"

	"github.com/nicolaspearson/go.api/cmd/api/config"
	"github.com/nicolaspearson/go.api/cmd/api/internal/domain/userentity"
	"github.com/nicolaspearson/go.api/pkg/postgresql"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSqlContainer struct {
	pool      *dockertest.Pool
	resource  *dockertest.Resource
	imagename string
	opts      postgresql.Opts
}

type IPostgreSqlContainer interface {
	C() PostgreSqlContainer
	Create() error
	Connect() *gorm.DB
	AutoMigrate(db *gorm.DB) error
	Flush(db *gorm.DB)
}

func NewPostgresqlContainer(pool *dockertest.Pool) IPostgreSqlContainer {
	opts := postgresql.Opts{
		Host:     config.Vars.DbHost,
		User:     config.Vars.DbUser,
		Password: config.Vars.DbPassword,
		Database: config.Vars.DbName,
		Port:     config.Vars.DbPort,
	}

	return PostgreSqlContainer{pool: pool, opts: opts, imagename: "postgresql-integration-tests"}
}

func (container PostgreSqlContainer) C() PostgreSqlContainer {
	return container
}

func (container PostgreSqlContainer) Create() error {
	if IsRunning(*container.pool, container.imagename) {
		return nil
	}

	dockerOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=" + container.opts.User,
			"POSTGRES_PASSWORD=" + container.opts.Password,
			"POSTGRES_DB=" + container.opts.Database,
		},
		ExposedPorts: []string{container.opts.Port},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(container.opts.Port): {{HostIP: "0.0.0.0", HostPort: container.opts.Port}},
		},
		Name: container.imagename,
	}

	resource, err := container.pool.RunWithOptions(&dockerOpts)
	if err != nil {
		log.Fatalf("Could not start resource (Postgresql Integration Tests Container): %s", err.Error())
		return err
	}

	container.resource = resource
	return nil
}

func (container PostgreSqlContainer) Connect() *gorm.DB {
	var db *gorm.DB
	if err := container.pool.Retry(func() error {
		defaultDsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
		dsn := fmt.Sprintf(defaultDsn, container.opts.Host, container.opts.User, container.opts.Password, container.opts.Database, container.opts.Port)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return db
}

func (container PostgreSqlContainer) AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(userentity.Entity{})
	if err != nil {
		return err
	}

	return nil
}

func (container PostgreSqlContainer) Flush(db *gorm.DB) {
	db.Exec("truncate table public.users")
}
