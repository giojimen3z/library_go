package repository_test

import (
	"errors"
	"strings"
	"testing"

	"library/internal/domain/model"
	"library/internal/infrastructure/adapter/repository"
	"library/internal/test/builder"
	"library/internal/test/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	authorInsertQuery     = "INSERT INTO authors"
	authorsSelectAllRegex = `^SELECT .* FROM "authors"`
	authorSelectByIDRegex = `^SELECT .* FROM "authors" WHERE "authors"\."id" = \$1`
	authorUpdateRegex     = `^UPDATE "authors" SET .* WHERE id = \$\d+`
	errInsertFailed       = "insert failed"
	errSelectFailed       = "select failed"
	errUpdateFailed       = "update failed"
)

func TestGivenValidAuthorWhenSaveThenReturnsNil(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	author := builder.NewAuthorBuilder().Build()
	expectedRegex := "^" + strings.ReplaceAll(authorInsertQuery, "authors", `"authors"`)
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(expectedRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(author.ID))
	sqlMock.ExpectCommit()

	err := repo.Save(t.Context(), author)

	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenSaveThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	author := builder.NewAuthorBuilder().Build()
	expectedErr := errors.New(errInsertFailed)
	expectedRegex := "^" + strings.ReplaceAll(authorInsertQuery, "authors", `"authors"`)
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(expectedRegex).
		WillReturnError(expectedErr)
	sqlMock.ExpectRollback()

	err := repo.Save(t.Context(), author)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	assert.NoError(t, sqlMock.ExpectationsWereMet(), "not all SQL expectations were met")
}

func TestGivenAuthorsWhenFindAllThenReturnsSlice(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	a1 := builder.NewAuthorBuilder().Build()
	a2 := builder.NewAuthorBuilder().Build()
	sqlMock.ExpectQuery(authorsSelectAllRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "bio", "created_at", "updated_at"}).
			AddRow(a1.ID, a1.FirstName, a1.LastName, a1.Bio, a1.CreatedAt, a1.UpdatedAt).
			AddRow(a2.ID, a2.FirstName, a2.LastName, a2.Bio, a2.CreatedAt, a2.UpdatedAt))

	authors, err := repo.FindAll(t.Context())

	assert.NoError(t, err)
	assert.Len(t, authors, 2)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenFindAllThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	expectedErr := errors.New(errSelectFailed)
	sqlMock.ExpectQuery(authorsSelectAllRegex).
		WillReturnError(expectedErr)

	authors, err := repo.FindAll(t.Context())

	assert.Error(t, err)
	assert.Nil(t, authors)
	assert.EqualError(t, err, expectedErr.Error())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenExistingAuthorIDWhenFindByIdThenReturnsAuthor(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	a := builder.NewAuthorBuilder().Build()
	sqlMock.ExpectQuery(authorSelectByIDRegex).
		WithArgs(a.ID, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "bio", "created_at", "updated_at"}).
			AddRow(a.ID, a.FirstName, a.LastName, a.Bio, a.CreatedAt, a.UpdatedAt))

	got, err := repo.FindById(t.Context(), a.ID)

	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, a.ID, got.ID)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenUnknownAuthorIDWhenFindByIdThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	a := builder.NewAuthorBuilder().Build()

	sqlMock.ExpectQuery(authorSelectByIDRegex).
		WithArgs(a.ID, sqlmock.AnyArg()).
		WillReturnError(gorm.ErrRecordNotFound)

	got, err := repo.FindById(t.Context(), a.ID)

	assert.Error(t, err)
	assert.Nil(t, got)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenValidPatchWhenUpdateThenReturnsNilError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	original := builder.NewAuthorBuilder().Build()
	patch := builder.UpdateAuthorBuilder().Build()
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(authorUpdateRegex).
		WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()

	got, err := repo.Update(t.Context(), original.ID, patch)

	assert.NoError(t, err)
	assert.Nil(t, got)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenUpdateThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewAuthorRepository(gdb)
	orig := builder.NewAuthorBuilder().Build()
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(authorUpdateRegex).
		WillReturnError(errors.New(errUpdateFailed))
	sqlMock.ExpectRollback()

	got, err := repo.Update(t.Context(), orig.ID, &model.Author{Bio: "x"})

	assert.Error(t, err)
	assert.Nil(t, got)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}
