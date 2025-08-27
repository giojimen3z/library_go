package application_test

import (
	"errors"
	"testing"

	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/builder"
	mmockAuthorRepo "library/internal/test/mock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAuthorWhenAppCreateAuthorThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	author := builder.NewAuthorBuilder().Build()
	mockRepo.On("Save", mock.AnythingOfType("*model.Author")).Return(nil)
	err := app.CreateAuthorUseCase(author)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongAuthorWhenAppCreateAuthorThenReturnError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	author := builder.NewAuthorBuilder().Build()
	errorExpected := errors.New("error saving DB")
	mockRepo.On("Save", mock.Anything).Return(errorExpected)

	err := app.CreateAuthorUseCase(author)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenAuthorWhenAppUpdateInDataBaseThenReturnNilError(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	patch := builder.UpdateAuthorBuilder().Build()
	existing := builder.NewAuthorBuilder().Build()
	updated := *existing
	updated.Bio = patch.Bio
	mockRepo.On("FindById", authorID).Return(existing, nil).Once()
	mockRepo.On("Update", authorID, patch).Return(nil, nil).Once()
	mockRepo.On("FindById", authorID).Return(&updated, nil).Once()
	result, err := app.UpdateAuthorUseCase(authorID, patch)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, updated.Bio, result.Bio)
	mockRepo.AssertExpectations(t)

}

func TestGivenWrongAuthorWhenAppUpdateInDataBaseThenReturnError(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	patch := builder.UpdateAuthorBuilder().Build()
	existing := builder.NewAuthorBuilder().Build()
	errorExpected := errors.New("error updating into DB")
	mockRepo.On("FindById", authorID).Return(existing, nil).Once()
	mockRepo.On("Update", authorID, mock.AnythingOfType("*model.Author")).Return(nil, errorExpected).Once()

	result, err := app.UpdateAuthorUseCase(authorID, patch)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenAuthorsInDBWhenAppGetAuthorsThenReturnList(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	firstAuthor := *builder.NewAuthorBuilder().Build()
	secondAuthor := *builder.UpdateAuthorBuilder().Build()
	expectedAuthors := []model.Author{
		firstAuthor,
		secondAuthor,
	}
	mockRepo.On("FindAll").Return(expectedAuthors, nil)

	result, err := app.GetAuthorsUseCase()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedAuthors, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenAppGetAuthorsThenReturnError(t *testing.T) {
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	expectedError := errors.New("error fetching authors")
	mockRepo.On("FindAll").Return(nil, expectedError)

	result, err := app.GetAuthorsUseCase()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenAppGetAuthorThenReturnAuthor(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	expectedAuthor := builder.NewAuthorBuilder().Build()
	mockRepo.On("FindById", authorID).Return(expectedAuthor, nil)

	result, err := app.GetAuthorUseCase(authorID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedAuthor, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenAppGetAuthorThenReturnError(t *testing.T) {
	authorID := uuid.New()
	mockRepo := new(mmockAuthorRepo.AuthorRepoMock)
	svc := service.NewAuthorService(mockRepo)
	app := application.NewAuthorUseCase(svc)
	expectedError := errors.New("author not found")
	mockRepo.On("FindById", authorID).Return(nil, expectedError)

	result, err := app.GetAuthorUseCase(authorID)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
