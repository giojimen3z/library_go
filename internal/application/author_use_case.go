package application

import (
	"library/internal/domain/model"
	"library/internal/domain/service"
)

type AuthorUseCase struct {
	service *service.AuthorService
}

func NewAuthorApp(service *service.AuthorService) *AuthorUseCase {
	return &AuthorUseCase{service}
}

// CreateAuthor save a new author using the service.
func (a *AuthorUseCase) CreateAuthor(author *model.Author) error {
	return a.service.CreateAuthor(author)
}

// GetAuthors get all authors, using the service.
func (a *AuthorUseCase) GetAuthors() ([]model.Author, error) {
	return a.service.GetAuthors()
}

// GetAuthor get a author according to de ID received.
func (a *AuthorUseCase) GetAuthor(id uint64) (*model.Author, error) {
	return a.service.GetAuthor(id)
}

// UpdateAuthor update an existing author in the database.
func (a *AuthorUseCase) UpdateAuthor(id uint64, patch *model.Author) (*model.Author, error) {
	return a.service.UpdateAuthor(id, patch)
}
