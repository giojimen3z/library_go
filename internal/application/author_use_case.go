package application

import (
	"context"

	"library/internal/domain/model"
	"library/internal/domain/service"

	"github.com/google/uuid"
)

type AuthorUseCaseInterface interface {
	CreateAuthorUseCase(ctx context.Context, author *model.Author) error
	GetAuthorsUseCase(ctx context.Context) ([]model.Author, error)
	GetAuthorUseCase(ctx context.Context, id uuid.UUID) (*model.Author, error)
	UpdateAuthorUseCase(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error)
}

type AuthorUseCase struct {
	service service.AuthorServiceInterface
}

func NewAuthorUseCase(service service.AuthorServiceInterface) *AuthorUseCase {
	return &AuthorUseCase{service}
}

func (a *AuthorUseCase) CreateAuthorUseCase(ctx context.Context, author *model.Author) error {
	return a.service.CreateAuthor(ctx, author)
}

func (a *AuthorUseCase) GetAuthorsUseCase(ctx context.Context) ([]model.Author, error) {
	return a.service.GetAuthors(ctx)
}

func (a *AuthorUseCase) GetAuthorUseCase(ctx context.Context, id uuid.UUID) (*model.Author, error) {
	return a.service.GetAuthor(ctx, id)
}

func (a *AuthorUseCase) UpdateAuthorUseCase(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error) {
	return a.service.UpdateAuthor(ctx, id, patch)
}
