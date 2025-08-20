package application_test

import (
	"testing"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/mock"

	"github.com/stretchr/testify/assert"
)

func TestAppCreateAuthor(t *testing.T) {
	mockRepo := new(mock.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorApp(svc)

	author := &model.Author{FirstName: "Jane", LastName: "Smith"}

	mockRepo.On("Save", author).Return(nil)

	err := app.CreateAuthor(author)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
