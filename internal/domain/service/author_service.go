package service

import (
	"library/internal/domain/model"
	"library/internal/domain/port"

	"github.com/google/uuid"
)

type AuthorServiceInterface interface {
	CreateAuthorService(author *model.Author) error
	GetAuthorsService() ([]model.Author, error)
	GetAuthorService(id uint64) (*model.Author, error)
	UpdateAuthorService(id uint64, patch *model.Author) (*model.Author, error)
}

type AuthorService struct {
	repo port.AuthorPort
}

func NewAuthorService(repo port.AuthorPort) *AuthorService {
	return &AuthorService{repo}
}

func (s *AuthorService) CreateAuthorService(author *model.Author) error {
	id := uuid.New()
	author.ID = id
	return s.repo.SaveRepository(author)
}

func (s *AuthorService) GetAuthorsService() ([]model.Author, error) {
	return s.repo.FindAllRepository()
}

func (s *AuthorService) GetAuthorService(id uint64) (*model.Author, error) {
	return s.repo.FindByIdRepository(id)
}

func (s *AuthorService) UpdateAuthorService(id uint64, patch *model.Author) (*model.Author, error) {
	return s.repo.UpdateRepository(id, patch)
}
