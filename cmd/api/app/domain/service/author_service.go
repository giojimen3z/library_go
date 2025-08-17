package service

import (
	"hexagonal_base/cmd/api/app/domain/model"
	"hexagonal_base/cmd/api/app/domain/port"
)

type AuthorService struct {
	repo port.AuthorRepository
}

func NewAuthorService(repo port.AuthorRepository) *AuthorService {
	return &AuthorService{repo}
}

func (s *AuthorService) CreateAuthor(author *model.Author) error {
	return s.repo.Save(author)
}

func (s *AuthorService) GetAuthors() ([]model.Author, error) {
	return s.repo.FindAll()
}
