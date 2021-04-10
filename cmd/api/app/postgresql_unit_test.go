package app

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var (
		database *sql.DB
		err      error
	)

	database, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}), &gorm.Config{})

	require.NoError(s.T(), err)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func Test_Init(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_RunMigrations() {
	application := New()
	assert.Equal(s.T(), application.RunMigrations(), true)
}
