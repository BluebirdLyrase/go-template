package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"my-api/internal/modules/auth/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// --- Test Suite Setup ---

type UserRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	s.Require().NoError(err)

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	s.db, err = gorm.Open(dialector, &gorm.Config{})
	s.Require().NoError(err)

	s.repo = NewUserRepository(s.db)
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	// Ensure all expectations were met
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

// --- Tests ---

func (s *UserRepositoryTestSuite) TestCreate_Success() {
	user := &models.User{
		Email: "test@example.com",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()

	err := s.repo.Create(user)
	assert.NoError(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestCreate_Error() {
	user := &models.User{Email: "test@example.com"}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WillReturnError(fmt.Errorf("db error"))
	s.mock.ExpectRollback()

	err := s.repo.Create(user)
	assert.Error(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestGetByID_Found() {
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at"}).
		AddRow(1, "test@example.com", now, now)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE`)).
		WithArgs(1, 1). // GORM passes id twice for First()
		WillReturnRows(rows)

	user, err := s.repo.GetByID(1)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), user)
	assert.Equal(s.T(), "test@example.com", user.Email)
}

func (s *UserRepositoryTestSuite) TestGetByID_NotFound() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE`)).
		WithArgs(99, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	user, err := s.repo.GetByID(99)
	assert.NoError(s.T(), err) // not found returns nil, nil
	assert.Nil(s.T(), user)
}

func (s *UserRepositoryTestSuite) TestGetByEmail_Found() {
	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at"}).
		AddRow(1, "test@example.com", now, now)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1`)).
		WithArgs("test@example.com", 1).
		WillReturnRows(rows)

	user, err := s.repo.GetByEmail("test@example.com")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), user)
	assert.Equal(s.T(), "test@example.com", user.Email)
}

func (s *UserRepositoryTestSuite) TestGetByEmail_NotFound() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1`)).
		WithArgs("notfound@example.com", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	user, err := s.repo.GetByEmail("notfound@example.com")
	assert.NoError(s.T(), err)
	assert.Nil(s.T(), user)
}
