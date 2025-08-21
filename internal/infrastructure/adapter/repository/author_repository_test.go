package repository_test

import (
	"errors"
	"strings"
	"testing"

	"library/internal/infrastructure/adapter/repository"
	"library/internal/test/builder"
	"library/internal/test/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	authorQuery = "INSERT INTO authors"
)

func TestGivenValidAuthorWhenSaveThenReturnsNil(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	author := builder.NewAuthorBuilder().Build()
	expectedRegex := "^" + strings.ReplaceAll(authorQuery, "authors", `"authors"`)
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(expectedRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(author.ID))
	sqlMock.ExpectCommit()

	err := repo.Save(author)

	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenSaveThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	author := builder.NewAuthorBuilder().Build()
	expectedErr := errors.New("inserción fallida")
	expectedRegex := "^" + strings.ReplaceAll(authorQuery, "authors", `"authors"`)
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(expectedRegex).
		WillReturnError(expectedErr)
	sqlMock.ExpectRollback()

	err := repo.Save(author)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	assert.NoError(t, sqlMock.ExpectationsWereMet(), "no se cumplieron todas las expectativas de SQL")
}
