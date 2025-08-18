package builder

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/infrastructure/controller"
	"library/internal/test/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthorController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mock.MockAuthorRepo)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorApp(svc)
	ctrl := controller.NewAuthorController(app)

	author := model.Author{FirstName: "Alice", LastName: "Wonder"}
	mockRepo.On("Save", &author).Return(nil)

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
