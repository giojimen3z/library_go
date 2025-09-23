package application_test

import (
	"errors"
	"testing"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/builder"
	mockLoanRepo "library/internal/test/mock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenLoanWhenAppCreateLoanThenReturnNilError(t *testing.T) {
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)

	loan := builder.NewLoanBuilder().Build()
	mockRepo.On("CantLoansByMember", mock.Anything, loan.MemberID).Return(int64(0), nil)
	mockRepo.On("FindByMemeberId", mock.Anything, loan.MemberID).Return([]model.Loan{}, nil)
	mockRepo.On("Save", mock.Anything, loan).Return(nil)

	err := app.CreateLoanUseCase(t.Context(), loan)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongLoanWhenAppCreateLoanThenReturnError(t *testing.T) {
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	loan := builder.NewLoanBuilder().Build()
	errorExpected := errors.New("error saving DB")
	mockRepo.On("CantLoansByMember", mock.Anything, loan.MemberID).Return(int64(0), nil)
	mockRepo.On("FindByMemeberId", mock.Anything, loan.MemberID).Return([]model.Loan{}, nil)
	mockRepo.On("Save", mock.Anything, loan).Return(errorExpected)

	err := app.CreateLoanUseCase(t.Context(), loan)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenLoanWhenAppUpdateInDataBaseThenReturnNilError(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	patch := builder.UpdateLoanBuilder().Build()
	existing := builder.NewLoanBuilder().Build()
	updated := *existing
	updated.ReturnedAt = patch.ReturnedAt

	mockRepo.On("Update", mock.Anything, loanID, patch).Return(&updated, nil).Once()
	result, err := app.UpdateLoanUseCase(t.Context(), loanID, patch)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, updated.ReturnedAt, result.ReturnedAt)
	mockRepo.AssertExpectations(t)

}

func TestGivenWrongLoanWhenAppUpdateInDataBaseThenReturnError(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	patch := builder.UpdateLoanBuilder().Build()
	errorExpected := errors.New("error updating into DB")
	mockRepo.On("Update", mock.Anything, loanID, patch).Return((*model.Loan)(nil), errorExpected).Once()

	result, err := app.UpdateLoanUseCase(t.Context(), loanID, patch)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenLoansInDBWhenAppGetLoansThenReturnList(t *testing.T) {
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	firstLoan := *builder.NewLoanBuilder().Build()
	secondLoan := *builder.UpdateLoanBuilder().Build()
	expectedLoans := []model.Loan{
		firstLoan,
		secondLoan,
	}
	mockRepo.On("FindAll").Return(expectedLoans, nil)

	result, err := app.GetLoansUseCase(t.Context())

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedLoans, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenAppGetLoansThenReturnError(t *testing.T) {
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	expectedError := errors.New("error fetching loans")
	mockRepo.On("FindAll").Return([]model.Loan{}, expectedError)

	result, err := app.GetLoansUseCase(t.Context())

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenAppGetLoanThenReturnLoan(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	expectedLoan := builder.NewLoanBuilder().Build()
	mockRepo.On("FindById", loanID).Return(expectedLoan, nil)

	result, err := app.GetLoanUseCase(t.Context(), loanID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedLoan, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenAppGetLoanThenReturnError(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	expectedError := errors.New("loan not found")
	mockRepo.On("FindById", loanID).Return(&model.Loan{}, expectedError)

	result, err := app.GetLoanUseCase(t.Context(), loanID)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
