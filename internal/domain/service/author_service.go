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

func (s *AuthorService) CreateAuthor(author *model.Author) error {
	return s.repo.Save(author)
}

func (s *AuthorService) GetAuthors() ([]model.Author, error) {
	return s.repo.FindAll()
}

func (s *AuthorService) GetAuthor(id uint64) (*model.Author, error) {
	return s.repo.FindById(id)
}

func (s *AuthorService) UpdateAuthor(id uint64, patch *model.Author) (*model.Author, error) {
	return s.repo.Update(id, patch)
}
