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
	mockAuthorRepo "library/internal/test/mock"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAnAuthorWhenCreateInControllerThenReturnStatusCreated(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("Save", mock.Anything).Return(nil)
	body, _ := json.Marshal(author)
	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongAuthorWhenCreateInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)
	author := builder.NewAuthorBuilder().Build()
	errorExpected := errors.New("error saving DB")
	mockRepo.On("Save", mock.Anything).Return(errorExpected)
	body, _ := json.Marshal(author)
	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	ctrl.Create(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenAuthorWheUpdateInDataBaseInControllerThenReturnNilError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(author, nil)
	body, _ := json.Marshal(author)
	req, _ := http.NewRequest("PUT", "/authors/"+author.ID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: author.ID.String()}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongAuthorWhenUpdateInDataBaseInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	body, _ := json.Marshal(author)
	req, _ := http.NewRequest("PUT", "/authors/"+author.ID.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: author.ID.String()}}
	c.Request = req

	ctrl.Update(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenAuthorsInDBWhenGetAuthorsInControllerThenReturnList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)
	firstAuthor := *builder.NewAuthorBuilder().Build()
	secondAuthor := *builder.UpdateAuthorBuilder().Build()
	expectedAuthors := []model.Author{
		firstAuthor,
		secondAuthor,
	}
	mockRepo.On("FindAll").Return(expectedAuthors, nil)

	req, _ := http.NewRequest("GET", "/authors", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenGetAuthorsInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)

	mockRepo.On("FindAll").Return(nil, errors.New("error fetching authors"))

	req, _ := http.NewRequest("GET", "/authors", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenGetAuthorInControllerThenReturnAuthor(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)
	author := builder.NewAuthorBuilder().Build()

	mockRepo.On("FindById", author.ID).Return(author, nil)

	req, _ := http.NewRequest("GET", "/authors/"+author.ID.String(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: author.ID.String()}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenGetAuthorInControllerThenReturnError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	id := uuid.New()
	authorID := id.String()
	mockRepo := new(mockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	ctrl := controller.NewAuthorController(app)

	mockRepo.On("FindById", id).Return((*model.Author)(nil), errors.New("not found"))

	req, _ := http.NewRequest("GET", "/authors/"+authorID, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: authorID}}
	c.Request = req

	ctrl.GetById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertExpectations(t)
}
