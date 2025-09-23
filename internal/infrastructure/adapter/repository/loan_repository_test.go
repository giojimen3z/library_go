package repository_test

import (
	"errors"
	"library/internal/domain/model"
	"library/internal/infrastructure/adapter/repository"
	"library/internal/test/builder"
	"library/internal/test/mock"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	loanInsertQuery     = "INSERT INTO loans"
	loansSelectAllRegex = `^SELECT .* FROM "loans"`
	loanSelectByIDRegex = `^SELECT .* FROM "loans" WHERE "loans"\."id" = \$1`
	loanUpdateRegex     = `^UPDATE "loans" SET .* WHERE id = \$\d+`
	errLoanInsertFailed = "insert failed"
	errLoanSelectFailed = "select failed"
	errLoanUpdateFailed = "update failed"
)

func TestGivenValidLoanWhenSaveThenReturnsNil(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	loan := builder.NewLoanBuilder().Build()
	expectedRegex := "^" + strings.ReplaceAll(loanInsertQuery, "loans", `"loans"`)
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(expectedRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(loan.ID))
	sqlMock.ExpectCommit()

	err := repo.Save(t.Context(), loan)

	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenSaveLoanThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	loan := builder.NewLoanBuilder().Build()
	expectedErr := errors.New(errLoanInsertFailed)
	expectedRegex := "^" + strings.ReplaceAll(loanInsertQuery, "loans", `"loans"`)
	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(expectedRegex).
		WillReturnError(expectedErr)
	sqlMock.ExpectRollback()

	err := repo.Save(t.Context(), loan)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	assert.NoError(t, sqlMock.ExpectationsWereMet(), "not all SQL expectations were met")
}

func TestGivenLoansWhenFindAllThenReturnsSlice(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	a1 := builder.NewLoanBuilder().Build()
	a2 := builder.NewLoanBuilder().Build()
	sqlMock.ExpectQuery(loansSelectAllRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_id", "copy_id", "loaned_at", "due_date", "returned_at", "fine_cents", "created_at"}).
			AddRow(a1.ID, a1.MemberID, a1.CopyID, a1.LoanedAt, a1.DueDate, a1.ReturnedAt, a1.FineCents, a1.CreatedAt).
			AddRow(a2.ID, a2.MemberID, a2.CopyID, a2.LoanedAt, a2.DueDate, a2.ReturnedAt, a2.FineCents, a2.CreatedAt))

	loans, err := repo.FindAll(t.Context())

	assert.NoError(t, err)
	assert.Len(t, loans, 2)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenFindAllLoansThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	expectedErr := errors.New(errLoanSelectFailed)
	sqlMock.ExpectQuery(loansSelectAllRegex).
		WillReturnError(expectedErr)

	loans, err := repo.FindAll(t.Context())

	assert.Error(t, err)
	assert.Nil(t, loans)
	assert.EqualError(t, err, expectedErr.Error())
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenExistingLoanIDWhenFindLoanByIdThenReturnsLoan(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	a := builder.NewLoanBuilder().Build()
	sqlMock.ExpectQuery(loanSelectByIDRegex).
		WithArgs(a.ID, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "member_id", "copy_id", "loaned_at", "due_date", "returned_at", "fine_cents", "created_at"}).
			AddRow(a.ID, a.MemberID, a.CopyID, a.LoanedAt, a.DueDate, a.ReturnedAt, a.FineCents, a.CreatedAt))

	got, err := repo.FindById(t.Context(), a.ID)

	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, a.ID, got.ID)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenUnknownLoanIDWhenFindLoanByIdThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	a := builder.NewLoanBuilder().Build()

	sqlMock.ExpectQuery(loanSelectByIDRegex).
		WithArgs(a.ID, sqlmock.AnyArg()).
		WillReturnError(gorm.ErrRecordNotFound)

	got, err := repo.FindById(t.Context(), a.ID)

	assert.Error(t, err)
	assert.Nil(t, got)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenValidPatchWhenUpdateLoanThenReturnsNilError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	original := builder.NewLoanBuilder().Build()
	patch := builder.UpdateLoanBuilder().Build()
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(loanUpdateRegex).
		WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()

	got, err := repo.Update(t.Context(), original.ID, patch)

	assert.NoError(t, err)
	assert.Nil(t, got)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGivenDBErrorWhenUpdateLoanThenReturnsError(t *testing.T) {
	gdb, sqlMock := mock.SetupGormWithSQLMock(t)
	repo := repository.NewLoanRepository(gdb)
	orig := builder.NewLoanBuilder().Build()
	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(loanUpdateRegex).
		WillReturnError(errors.New(errLoanUpdateFailed))
	sqlMock.ExpectRollback()

	got, err := repo.Update(t.Context(), orig.ID, &model.Loan{FineCents: 20})

	assert.Error(t, err)
	assert.Nil(t, got)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}
