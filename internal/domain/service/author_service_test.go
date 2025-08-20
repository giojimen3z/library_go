package service_test

import (
	"errors"
	"testing"

	"library/internal/domain/service"
	"library/internal/test/builder"

	mmockAuthorRepo "library/internal/test/mock"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnAuthorShouldSaveInDBThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("Save", author).Return(nil)

	err := serviceAuthor.CreateAuthor(author)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongAuthorShouldFailSaveInDBThenReturnError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	author := builder.NewAuthorBuilder().Build()
	errorExpected := errors.New("error saving into DB")
	mockRepo.On("Save", author).Return(errorExpected)

	err := serviceAuthor.CreateAuthor(author)

	assert.NotNil(t, err)
	assert.Error(t, errorExpected)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}
