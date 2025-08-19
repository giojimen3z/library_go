package service

import (
	"library/internal/domain/model"
	"library/internal/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	mockRepo := new(mock.MockAuthorRepo)
	service := NewAuthorService(mockRepo)

	author := &model.Author{
		FirstName: "John",
		LastName:  "Doe",
	}

	// Definir comportamiento del mock
	mockRepo.On("Save", author).Return(nil)

	err := service.CreateAuthor(author)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
