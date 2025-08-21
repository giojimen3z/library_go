package application_test

import (
	"testing"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	mmockAuthorRepo "library/internal/test/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAuthorWhenAppCreateAuthorThenReturnNoError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)

	author := &model.Author{FirstName: "Jane", LastName: "Smith"}

	mockRepo.On("Save", mock.Anything).Return(nil)

	err := app.CreateAuthorUseCase(author)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
