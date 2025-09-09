package service

import (
	"context"

	"github.com/google/uuid"

	"library/internal/domain/model"
	"library/internal/domain/port"
)

type AuthorServiceInterface interface {
	CreateAuthor(ctx context.Context, author *model.Author) error
	GetAuthors(ctx context.Context) ([]model.Author, error)
	GetAuthor(ctx context.Context, id uuid.UUID) (*model.Author, error)
	UpdateAuthor(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error)
}

type AuthorService struct {
	repo port.AuthorPort
}

func NewAuthorService(repo port.AuthorPort) *AuthorService {
	return &AuthorService{repo}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author *model.Author) error {
	id := uuid.New()
	author.ID = id
	return s.repo.Save(ctx, author)
}

func (s *AuthorService) GetAuthors(ctx context.Context) ([]model.Author, error) {
	return s.repo.FindAll(ctx)
}

func (s *AuthorService) GetAuthor(ctx context.Context, id uuid.UUID) (*model.Author, error) {
	return s.repo.FindById(ctx, id)
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error) {
	return s.repo.Update(ctx, id, patch)
}
