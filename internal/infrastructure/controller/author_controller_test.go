package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/internal/application"
	"library/internal/domain/service"
	"library/internal/infrastructure/controller"
	"library/internal/test/builder"
	mockAuthorRepo "library/internal/test/mock"

	"github.com/gin-gonic/gin"
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
