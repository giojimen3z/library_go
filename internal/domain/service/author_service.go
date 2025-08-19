package service

import (
	"library/internal/domain/model"
	"library/internal/domain/port"
)

type AuthorService struct {
	repo port.AuthorPort
}

func NewAuthorService(repo port.AuthorPort) *AuthorService {
	return &AuthorService{repo}
}

// CreateAuthor save a new author using the repository.
func (s *AuthorService) CreateAuthor(author *model.Author) error {
	return s.repo.Save(author)
}

// GetAuthors get all authors, using the repository.
func (s *AuthorService) GetAuthors() ([]model.Author, error) {
	return s.repo.FindAll()
}

// GetAuthor get a author according to de ID received.
func (s *AuthorService) GetAuthor(id uint64) (*model.Author, error) {
	return s.repo.FindById(id)
}

// UpdateAuthor update an existing author in the database.
func (s *AuthorService) UpdateAuthor(id uint64, patch *model.Author) (*model.Author, error) {
	return s.repo.Update(id, patch)
}
