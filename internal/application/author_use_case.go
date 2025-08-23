package application

import (
	"library/internal/domain/model"
	"library/internal/domain/service"

	"github.com/google/uuid"
)

type AuthorUseCaseInterface interface {
	CreateAuthorUseCase(author *model.Author) error
	GetAuthorsUseCase() ([]model.Author, error)
	GetAuthorUseCase(id uuid.UUID) (*model.Author, error)
	UpdateAuthorUseCase(id uuid.UUID, patch *model.Author) (*model.Author, error)
}

type AuthorUseCase struct {
	service service.AuthorServiceInterface
}

func NewAuthorUseCase(service service.AuthorServiceInterface) *AuthorUseCase {
	return &AuthorUseCase{service}
}

func (a *AuthorUseCase) CreateAuthorUseCase(author *model.Author) error {
	return a.service.CreateAuthor(author)
}

func (a *AuthorUseCase) GetAuthorsUseCase() ([]model.Author, error) {
	return a.service.GetAuthors()
}

func (a *AuthorUseCase) GetAuthorUseCase(id uuid.UUID) (*model.Author, error) {
	return a.service.GetAuthor(id)
}

func (a *AuthorUseCase) UpdateAuthorUseCase(id uuid.UUID, patch *model.Author) (*model.Author, error) {
	return a.service.UpdateAuthor(id, patch)
}
