package service

import (
	"library/internal/domain/model"
	"library/internal/domain/port"

	"github.com/google/uuid"
)

type AuthorServiceInterface interface {
	CreateAuthor(author *model.Author) error
	GetAuthors() ([]model.Author, error)
	GetAuthor(id uuid.UUID) (*model.Author, error)
	UpdateAuthor(id uuid.UUID, patch *model.Author) (*model.Author, error)
}

type AuthorService struct {
	repo port.AuthorPort
}

func NewAuthorService(repo port.AuthorPort) *AuthorService {
	return &AuthorService{repo}
}

func (s *AuthorService) CreateAuthor(author *model.Author) error {
	id := uuid.New()
	author.ID = id
	return s.repo.Save(author)
}

func (s *AuthorService) GetAuthors() ([]model.Author, error) {
	return s.repo.FindAll()
}

func (s *AuthorService) GetAuthor(id uuid.UUID) (*model.Author, error) {
	return s.repo.FindById(id)
}

func (s *AuthorService) UpdateAuthor(id uuid.UUID, patch *model.Author) (*model.Author, error) {
	if _, err := s.repo.FindById(id); err != nil {
		return nil, err
	}
	if _, err := s.repo.Update(id, patch); err != nil {
		return nil, err
	}
	return s.repo.FindById(id)
}
