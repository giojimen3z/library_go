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
	mockBookRepo "library/internal/test/mock"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenABookWhenCreateInControllerThenReturnStatusCreated(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	book := builder.NewBookBuilder().Build()

	mockRepo.On("Save", mock.AnythingOfType("*model.Book")).Return(nil).Once()

	body, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongBookWhenCreateInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	book := builder.NewBookBuilder().Build()
	expected := errors.New("error saving DB")

	mockRepo.On("Save", mock.Anything).Return(expected).Once()

	body, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenBookWhenUpdateInControllerThenReturnOK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	book := builder.NewBookBuilder().Build()
	patch := &model.Book{Description: "updated description"} // sin UpdateBookBuilder

	updated := *book
	updated.Description = patch.Description

	mockRepo.
		On("Update", book.ID, mock.AnythingOfType("*model.Book")).
		Return(&updated, nil).Once()

	body, _ := json.Marshal(patch)
	req, _ := http.NewRequest("PUT", "/books/"+book.ID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: book.ID.String()}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongBookWhenUpdateInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	book := builder.NewBookBuilder().Build()
	patch := &model.Book{Description: "updated description"}
	expected := errors.New("db error")

	mockRepo.
		On("Update", book.ID, mock.AnythingOfType("*model.Book")).
		Return(nil, expected).Once()

	body, _ := json.Marshal(patch)
	req, _ := http.NewRequest("PUT", "/books/"+book.ID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: book.ID.String()}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenBooksInDBWhenGetBooksInControllerThenReturnList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	first := *builder.NewBookBuilder().Build()
	second := *builder.NewBookBuilder().WithID(uuid.New()).WithTitle("Another").Build()
	expectedBooks := []model.Book{first, second}

	mockRepo.On("FindAll").Return(expectedBooks, nil).Once()

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenGetBooksInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	mockRepo.On("FindAll").Return(nil, errors.New("error fetching books")).Once()

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenGetBookInControllerThenReturnBook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	book := builder.NewBookBuilder().Build()

	mockRepo.On("FindById", book.ID).Return(book, nil).Once()

	req, _ := http.NewRequest("GET", "/books/"+book.ID.String(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: book.ID.String()}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenGetBookInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	id := uuid.New()
	bookID := id.String()

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	mockRepo.On("FindById", id).Return((*model.Book)(nil), errors.New("not found")).Once()

	req, _ := http.NewRequest("GET", "/books/"+bookID, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: bookID}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidJSONWhenCreateBookInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBufferString("{invalid"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGivenInvalidUUIDWhenGetBookInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	req, _ := http.NewRequest("GET", "/books/invalid-uuid", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid-uuid"}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGivenInvalidUUIDWhenUpdateBookInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	req, _ := http.NewRequest("PUT", "/books/invalid-uuid", bytes.NewBufferString("{}"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid-uuid"}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGivenInvalidJSONWhenUpdateBookInControllerThenReturnBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	ctrl := controller.NewBookController(app)

	validID := uuid.New().String()
	req, _ := http.NewRequest("PUT", "/books/"+validID, bytes.NewBufferString("{invalid"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: validID}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
