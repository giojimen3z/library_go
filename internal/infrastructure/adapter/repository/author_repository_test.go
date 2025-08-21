package repository_test

import (
	"errors"
	"strings"
	"testing"

	"library/internal/infrastructure/adapter/repository"
	"library/internal/test/builder"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	authorQuery = "INSERT INTO authors"
)

func TestGivenValidAuthor_WhenSave_ThenReturnsNil(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("no se pudo crear sqlmock: %v", err)
	}
	defer func() { _ = sqlDB.Close() }()

	dialectic := postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialectic, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("no se pudo abrir gorm con postgres sqlmock: %v", err)
	}

	repo := repository.NewAuthorRepository(gdb)
	author := builder.NewAuthorBuilder().Build()
	expectedRegex := "^" + strings.ReplaceAll(authorQuery, "authors", `"authors"`)
	mock.ExpectBegin()
	mock.ExpectQuery(expectedRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(author.ID))
	mock.ExpectCommit()

	err = repo.Save(author)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGivenDBError_WhenSave_ThenReturnsError(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("no se pudo crear sqlmock: %v", err)
	}
	defer func() { _ = sqlDB.Close() }()
	dialectic := postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialectic, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("no se pudo abrir gorm con postgres sqlmock: %v", err)
	}
	repo := repository.NewAuthorRepository(gdb)
	author := builder.NewAuthorBuilder().Build()
	expectedErr := errors.New("inserción fallida")
	expectedRegex := "^" + strings.ReplaceAll(authorQuery, "authors", `"authors"`)
	mock.ExpectBegin()
	mock.ExpectQuery(expectedRegex).
		WillReturnError(expectedErr)
	mock.ExpectRollback()

	err = repo.Save(author)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	assert.NoError(t, mock.ExpectationsWereMet(), "no se cumplieron todas las expectativas de SQL")
}
