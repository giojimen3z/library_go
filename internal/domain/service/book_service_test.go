package service_test

import (
	"errors"
	"library/internal/domain/model"
	"library/internal/domain/service"
	"library/internal/test/builder"
	mmockBookRepo "library/internal/test/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGivenAnBookWhenSaveIdInDBThenReturnNilError(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	book := builder.NewBookBuilder().Build()
	mockRepo.On("Save", mock.AnythingOfType("*model.Book")).Return(nil)
	_, err := serviceBook.CreateBook(book)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenWronkBookWhenSaveIdInDBThenReturnError(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	book := builder.NewBookBuilder().Build()
	errorExpected := errors.New("error saving into DB")
	mockRepo.On("Save", mock.Anything).Return(errorExpected)

	_, err := serviceBook.CreateBook(book)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenAnBookWhenUpdateInDbThenReturnUpdateEntity(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	patch := &model.Book{Description: "updated desc"}
	existing := builder.NewBookBuilder().Build()
	updated := *existing
	updated.Description = patch.Description

	mockRepo.On("FindById", bookID, mock.AnythingOfType("*model.Book")).Return(updated, nil)
	mockRepo.On("Update", bookID, mock.AnythingOfType("*model.Book")).Return(updated, nil).Once()

	result, err := serviceBook.UpdateBook(bookID, patch)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updated.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestGivenUpdateBookErrorWhenUpdateInDBThenReturnError(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	patch := &model.Book{Description: "updated desc"}
	existing := builder.NewBookBuilder().Build()
	errorExpected := errors.New("error updating into DB")
	mockRepo.On("FindById", bookID, mock.AnythingOfType("*model.Book")).Return(existing, nil)
	mockRepo.On("Update", bookID, mock.Anything).Return(nil, errorExpected)

	result, err := serviceBook.UpdateBook(bookID, patch)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, errorExpected, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenBooksInDBWhenGetBooksThenReturnList(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	firstBook := *builder.NewBookBuilder().Build()
	secondBook := *builder.NewBookBuilder().Build()
	expectedBooks := []model.Book{
		firstBook,
		secondBook,
	}
	mockRepo.On("FindAll").Return(expectedBooks, nil)

	result, err := serviceBook.GetAllBook()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBooks, result)
	mockRepo.AssertExpectations(t)
}

func TestGivenErrorWhenGetBooksThenReturnError(t *testing.T) {
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	expectedError := errors.New("error fetching books")
	mockRepo.On("FindAll").Return(nil, expectedError)

	result, err := serviceBook.GetAllBook()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGivenValidIDWhenGetBookThenReturnBook(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	expectedBook := builder.NewBookBuilder().Build()
	mockRepo.On("FindAll", bookID).Return(expectedBook, nil)

	result, err := serviceBook.GetByIDBook(bookID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBook, result)
	mockRepo.AssertExpectations(t)

}

func TestGivenInvalidIDWhenGetBookThenReturnError(t *testing.T) {
	bookID := uuid.New()
	mockRepo := new(mmockBookRepo.BookRepoMock)
	serviceBook := service.NewBookService(mockRepo)
	expectedError := errors.New("error fetching book")
	mockRepo.On("FindAll", bookID).Return(nil, expectedError)

	result, err := serviceBook.GetByIDBook(bookID)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
