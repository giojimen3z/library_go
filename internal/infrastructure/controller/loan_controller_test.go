package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/infrastructure/controller"
	"library/internal/test/builder"
	mockLoanRepo "library/internal/test/mock"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAnLoanWhenCreateInControllerThenReturnStatusCreated(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	loan := builder.NewLoanBuilder().Build()

	mockRepo.On("CantLoansByMember", mock.Anything, loan.MemberID).Return(int64(0), nil)
	mockRepo.On("FindByMemeberId", mock.Anything, loan.MemberID).Return([]model.Loan{}, nil)
	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*model.Loan")).Return(nil)
	body, _ := json.Marshal(loan)
	req, _ := http.NewRequest("POST", "/loans", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongLoanWhenCreateInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	loan := builder.NewLoanBuilder().Build()
	errorExpected := errors.New("error saving DB")
	mockRepo.On("CantLoansByMember", mock.Anything, loan.MemberID).Return(int64(0), nil)
	mockRepo.On("FindByMemeberId", mock.Anything, loan.MemberID).Return([]model.Loan{}, nil)
	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*model.Loan")).Return(errorExpected)

	body, _ := json.Marshal(loan)
	req, _ := http.NewRequest("POST", "/loans", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenLoanWheUpdateInDataBaseInControllerThenReturnNilError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	loan := builder.NewLoanBuilder().Build()
	patch := builder.UpdateLoanBuilder().Build()
	updated := *loan
	updated.ReturnedAt = patch.ReturnedAt

	mockRepo.On("Update", mock.Anything, loan.ID, mock.AnythingOfType("*model.Loan")).Return(loan, nil).Once()

	body, _ := json.Marshal(patch)
	req, _ := http.NewRequest("PUT", "/loans/"+loan.ID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: loan.ID.String()}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongLoanWhenUpdateInDataBaseInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	loan := builder.NewLoanBuilder().Build()
	patch := builder.UpdateLoanBuilder().Build()

	expectedErr := errors.New("database update failed")
	mockRepo.On("Update", mock.Anything, loan.ID, mock.AnythingOfType("*model.Loan")).Return(&model.Loan{}, expectedErr).Once()

	body, _ := json.Marshal(patch)
	req, _ := http.NewRequest("PUT", "/loans/"+loan.ID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: loan.ID.String()}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenLoansInDBWhenGetLoansInControllerThenReturnList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	firstLoan := *builder.NewLoanBuilder().Build()
	secondLoan := *builder.UpdateLoanBuilder().Build()
	expectedLoans := []model.Loan{
		firstLoan,
		secondLoan,
	}
	mockRepo.On("FindAll").Return(expectedLoans, nil)

	req, _ := http.NewRequest("GET", "/loans", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenGetLoansInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)

	mockRepo.On("FindAll").Return([]model.Loan{}, errors.New("error fetching loans"))

	req, _ := http.NewRequest("GET", "/loans", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenGetLoanInControllerThenReturnLoan(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	loan := builder.NewLoanBuilder().Build()

	mockRepo.On("FindById", loan.ID).Return(loan, nil)

	req, _ := http.NewRequest("GET", "/loans/"+loan.ID.String(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: loan.ID.String()}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenGetLoanInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	id := uuid.New()
	loanID := id.String()
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)

	mockRepo.On("FindById", id).Return((*model.Loan)(nil), errors.New("not found"))

	req, _ := http.NewRequest("GET", "/loans/"+loanID, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: loanID}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidJSONWhenCreateLoanInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	req, _ := http.NewRequest("POST", "/loans", bytes.NewBufferString("{invalid"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGivenInvalidUUIDWhenGetLoanInControllerThenReturnInternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	req, _ := http.NewRequest("GET", "/loans/invalid-uuid", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid-uuid"}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGivenInvalidUUIDWhenUpdateLoanInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	req, _ := http.NewRequest("PUT", "/loans/invalid-uuid", bytes.NewBufferString("{}"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid-uuid"}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGivenInvalidJSONWhenUpdateLoanInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockLoanRepo.LoanRepoMock)
	serviceBookCopy := service.NewBookCopyService()
	serviceMember := service.NewMemberService()
	svc := service.NewLoanService(mockRepo, serviceMember, serviceBookCopy)
	app := application.NewLoanUseCase(svc)
	ctrl := controller.NewLoanController(app)
	validID := uuid.New().String()
	req, _ := http.NewRequest("PUT", "/loans/"+validID, bytes.NewBufferString("{invalid"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: validID}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
