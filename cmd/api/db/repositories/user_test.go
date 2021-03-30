package db

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	models "github.com/nicolaspearson/go.api/cmd/api/db/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository IUserRepository
	user       *models.User
}

func (s *Suite) SetupSuite() {
	var (
		database *sql.DB
		err      error
	)

	database, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}), &gorm.Config{})

	require.NoError(s.T(), err)

	s.repository = NewUserRepository(s.DB)

	var (
		email     = "john.doe@example.com"
		enabled   = false
		firstName = "John"
		lastName  = "Doe"
		password  = "secret"
	)
	s.user = &models.User{Email: email, Enabled: enabled, FirstName: firstName, LastName: lastName, Password: password}
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetById() {
	var id uint = 0
	rows := sqlmock.NewRows([]string{"email", "enabled", "firstName", "lastName", "password"}).
		AddRow(s.user.Email, s.user.Enabled, s.user.FirstName, s.user.LastName, s.user.Password)

	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deletedAt" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(rows)

	res, err := s.repository.GetById(id)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(s.user, res))
}
