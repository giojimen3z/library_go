package application_test

import (
	"errors"
	"library/internal/application"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/builder"
	mmockBookRepo "library/internal/test/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenBookWhenCreateBookThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	book := builder.NewBookBuilder().Build()
	mockRepo.On("Save", mock.AnythingOfType("*model.Book")).Return(nil)
	err := app.CreateBookUseCase(book)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongBookWhenAppCreateBookThenReturnError(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	book := builder.NewBookBuilder().Build()
	errorExpected := errors.New("error saving DB")
	mockRepo.On("Save", mock.Anything).Return(errorExpected).Once()

	_, err := app.CreateBookUseCase(book)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenBookWhenAppUpdateInDataBaseThenReturnNilError(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	//patch := builder.UpdateBookBuilder().Build()
	patch := &model.Book{Description: "updated desc"}
	existing := builder.NewBookBuilder().Build()
	updated := *existing
	updated.Description = patch.Description
	mockRepo.On("Save", bookID).Return(existing, nil).Once()
	mockRepo.On("Update", bookID, patch).Return(nil, nil).Once()
	mockRepo.On("Save", bookID).Return(&updated, nil).Once()
	result, err := app.UpdateBookUseCase(bookID, patch)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, updated.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestGivenWrongBookWhenAppUpdateInDataBaseThenReturnError(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	patch := &model.Book{Description: "updated desc"}
	//patch := builder.UpdateBookBuilder().Build()
	existing := builder.NewBookBuilder().Build()
	errorExpected := errors.New("error updating into DB")
	mockRepo.On("Save", bookID).Return(existing, nil).Once()
	mockRepo.On("Update", bookID, mock.AnythingOfType("*model.Book")).Return(nil, errorExpected).Once()

	result, err := app.UpdateBookUseCase(bookID, patch)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenBooksInDBWhenAppGetBooksThenReturnList(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	firstBook := *builder.NewBookBuilder().Build()
	secondBook := *builder.NewBookBuilder().
		WithID(uuid.New()).
		WithTitle("Another Great Book").
		Build()
	expectedBooks := []model.Book{
		firstBook,
		secondBook,
	}
	mockRepo.On("FindAll").Return(expectedBooks, nil)

	result, err := app.GetBooksUseCase()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBooks, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenAppGetBooksThenReturnError(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	expectedError := errors.New("error fetching Books")
	mockRepo.On("FindAll").Return(nil, expectedError)

	result, err := app.GetBooksUseCase()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenAppGetBookThenReturnBook(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	expectedBook := builder.NewBookBuilder().Build()
	mockRepo.On("FindAll", bookID).Return(expectedBook, nil)

	result, err := app.GetBookByIdUseCase(bookID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBook, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenInvalidIDWhenAppGetBookThenReturnError(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	svc := service.NewBookService(mockRepo)
	app := application.NewBookUseCase(svc)
	expectedError := errors.New("Book not found")
	mockRepo.On("FindAll", bookID).Return(nil, expectedError)

	result, err := app.GetBookByIdUseCase(bookID)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
