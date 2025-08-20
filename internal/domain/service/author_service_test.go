package service_test

import (
	"errors"
	"testing"

	"library/internal/domain/service"
	"library/internal/test/builder"

	mmockAuthorRepo "library/internal/test/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAnAuthorShouldSaveInDBThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("SaveRepository", mock.AnythingOfType("*model.Author")).Return(nil)
	err := serviceAuthor.CreateAuthorService(author)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongAuthorShouldFailSaveInDBThenReturnError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	author := builder.NewAuthorBuilder().Build()
	errorExpected := errors.New("error saving into DB")
	mockRepo.On("SaveRepository", mock.Anything).Return(errorExpected)

	err := serviceAuthor.CreateAuthorService(author)

	assert.NotNil(t, err)
	assert.Error(t, errorExpected)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}
