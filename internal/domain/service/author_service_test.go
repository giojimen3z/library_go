package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/builder"
	mmockAuthorRepo "library/internal/test/mock"
)

func TestGivenAnAuthorWhenSaveInDBThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("Save", mock.AnythingOfType("*model.Author")).Return(nil)
	err := serviceAuthor.CreateAuthor(t.Context(), author)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongAuthorWhenSaveInDBThenReturnError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	author := builder.NewAuthorBuilder().Build()
	errorExpected := errors.New("error saving into DB")
	mockRepo.On("Save", mock.Anything).Return(errorExpected)

	err := serviceAuthor.CreateAuthor(t.Context(), author)

	assert.NotNil(t, err)
	assert.Error(t, errorExpected)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenAnAuthorWhenUpdateInDBThenReturnUpdatedEntity(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	patch := builder.UpdateAuthorBuilder().Build()
	existing := builder.NewAuthorBuilder().Build()
	updated := *existing
	updated.Bio = patch.Bio

	mockRepo.On("Update", authorID, patch).Return(patch, nil).Once()

	result, err := serviceAuthor.UpdateAuthor(t.Context(), authorID, patch)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updated.Bio, result.Bio)
	mockRepo.AssertExpectations(t)
}

func TestGivenUpdateErrorWhenUpdateInDBThenReturnError(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	patch := builder.UpdateAuthorBuilder().Build()
	errorExpected := errors.New("error updating into DB")
	mockRepo.On("Update", authorID, patch).Return(&model.Author{}, errorExpected).Once()

	result, err := serviceAuthor.UpdateAuthor(t.Context(), authorID, patch)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenAuthorsInDBWhenGetAuthorsThenReturnList(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	firstAuthor := *builder.NewAuthorBuilder().Build()
	secondAuthor := *builder.UpdateAuthorBuilder().Build()
	expectedAuthors := []model.Author{
		firstAuthor,
		secondAuthor,
	}
	mockRepo.On("FindAll").Return(expectedAuthors, nil)

	result, err := serviceAuthor.GetAuthors(t.Context())

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedAuthors, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenGetAuthorsThenReturnError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	expectedError := errors.New("error fetching authors")
	mockRepo.On("FindAll").Return([]model.Author{}, expectedError)

	result, err := serviceAuthor.GetAuthors(t.Context())

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenGetAuthorThenReturnAuthor(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	expectedAuthor := builder.NewAuthorBuilder().Build()
	mockRepo.On("FindById", authorID).Return(expectedAuthor, nil)

	result, err := serviceAuthor.GetAuthor(t.Context(), authorID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedAuthor, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenGetAuthorThenReturnError(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	serviceAuthor := service.NewAuthorService(mockRepo)
	expectedError := errors.New("author not found")
	mockRepo.On("FindById", authorID).Return(&model.Author{}, expectedError)

	result, err := serviceAuthor.GetAuthor(t.Context(), authorID)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
