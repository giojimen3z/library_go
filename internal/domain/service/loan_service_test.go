package service_test

import (
	"errors"
	"testing"

	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/builder"

	mmockLoanRepo "library/internal/test/mock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAnLoanWhenSaveInDBThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	loan := builder.NewLoanBuilder().Build()
	mockRepo.On("CantLoansByMember", mock.Anything, loan.MemberID).Return(int64(0), nil)
	mockRepo.On("FindByMemeberId", mock.Anything, loan.MemberID).Return([]model.Loan{}, nil)
	mockRepo.On("Save", mock.Anything, loan).Return(nil)

	err := serviceLoan.CreateLoan(t.Context(), loan)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongLoanWhenSaveInDBThenReturnError(t *testing.T) {
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	loan := builder.NewLoanBuilder().Build()
	errorExpected := errors.New("error saving into DB")
	mockRepo.On("CantLoansByMember", mock.Anything, loan.MemberID).Return(int64(0), nil)
	mockRepo.On("FindByMemeberId", mock.Anything, loan.MemberID).Return([]model.Loan{}, nil)
	mockRepo.On("Save", mock.Anything, loan).Return(errorExpected)

	err := serviceLoan.CreateLoan(t.Context(), loan)

	assert.NotNil(t, err)
	assert.Error(t, errorExpected)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenAnLoanWhenUpdateInDBThenReturnUpdatedEntity(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	patch := builder.UpdateLoanBuilder().Build()
	existing := builder.NewLoanBuilder().Build()
	updated := *existing
	updated.ReturnedAt = patch.ReturnedAt

	mockRepo.On("Update", mock.Anything, loanID, patch).Return(&updated, nil).Once()
	result, err := serviceLoan.UpdateLoan(t.Context(), loanID, patch)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updated.ReturnedAt, result.ReturnedAt)
	mockRepo.AssertExpectations(t)
}

func TestGivenUpdateErrorWhenUpdateLoanInDBThenReturnError(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	patch := builder.UpdateLoanBuilder().Build()
	errorExpected := errors.New("error updating into DB")

	mockRepo.On("Update", mock.Anything, loanID, patch).Return((*model.Loan)(nil), errorExpected).Once()

	result, err := serviceLoan.UpdateLoan(t.Context(), loanID, patch)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenLoansInDBWhenGetLoansThenReturnList(t *testing.T) {
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	firstLoan := *builder.NewLoanBuilder().Build()
	secondLoan := *builder.UpdateLoanBuilder().Build()
	expectedLoans := []model.Loan{
		firstLoan,
		secondLoan,
	}
	mockRepo.On("FindAll").Return(expectedLoans, nil)

	result, err := serviceLoan.GetLoans(t.Context())

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedLoans, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenGetLoansThenReturnError(t *testing.T) {
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	expectedError := errors.New("error fetching loans")
	mockRepo.On("FindAll").Return([]model.Loan{}, expectedError)

	result, err := serviceLoan.GetLoans(t.Context())

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenGetLoanThenReturnLoan(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	expectedLoan := builder.NewLoanBuilder().Build()
	mockRepo.On("FindById", loanID).Return(expectedLoan, nil)

	result, err := serviceLoan.GetLoanById(t.Context(), loanID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedLoan, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenGetLoanThenReturnError(t *testing.T) {
	loanID := uuid.New()
	mockRepo := new(mmockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	serviceLoan := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	expectedError := errors.New("loan not found")
	mockRepo.On("FindById", loanID).Return(&model.Loan{}, expectedError)

	result, err := serviceLoan.GetLoanById(t.Context(), loanID)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
