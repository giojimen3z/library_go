package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"library/internal/application"
	"library/internal/domain/service"
	"library/internal/infrastructure/app"
	"library/internal/infrastructure/controller"
	mockAuthorRepo "library/internal/test/mock"
	mockLoanRepo "library/internal/test/mock"
)

func TestGivenAuthorControllerWhenInitHandlersThenSetOnHandlers(t *testing.T) {
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	useCase := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(useCase)

	mockLoanRepo := new(mockLoanRepo.LoanRepoMock)
	svcMember := service.NewMemberService()
	svcBookCopyService := service.NewBookCopyService()
	svcLoan := service.NewLoanService(mockLoanRepo, svcMember, svcBookCopyService)
	useCaseLoan := application.NewLoanUseCase(svcLoan)
	ctrlLoan := controller.NewLoanController(useCaseLoan)

	h := app.NewHandlers(ctrl, ctrlLoan)

	assert.Equal(t, ctrl, h.Author)
}

func TestGivenAuthorControllerWhenPOSTAuthorsThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	useCase := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(useCase)

	mockLoanRepo := new(mockLoanRepo.LoanRepoMock)
	svcMember := service.NewMemberService()
	svcBookCopyService := service.NewBookCopyService()
	svcLoan := service.NewLoanService(mockLoanRepo, svcMember, svcBookCopyService)
	useCaseLoan := application.NewLoanUseCase(svcLoan)
	ctrlLoan := controller.NewLoanController(useCaseLoan)

	h := app.NewHandlers(ctrl, ctrlLoan)
	app.RegisterRoutes(r, h)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/library/authors", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

func TestGivenNoAuthorControllerWhenGETAuthorsThenReturnNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := app.NewHandlers(nil, nil)
	app.RegisterRoutes(r, h)
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/library/authors", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGivenNoAuthorControllerWhenGETAuthorByIDThenReturnNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := app.NewHandlers(nil, nil)
	app.RegisterRoutes(r, h)
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/library/authors/123", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGivenNoAuthorControllerWhenPUTAuthorByIDThenReturnNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	h := app.NewHandlers(nil, nil)
	app.RegisterRoutes(r, h)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/library/authors/123", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
