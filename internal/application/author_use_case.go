package application

import (
	"library/internal/domain/model"
	"library/internal/domain/service"
)

type AuthorUseCaseInterface interface {
	CreateAuthorUseCase(author *model.Author) error
	GetAuthorsUseCase() ([]model.Author, error)
	GetAuthorUseCase(id uint64) (*model.Author, error)
	UpdateAuthorUseCase(id uint64, patch *model.Author) (*model.Author, error)
}

type AuthorUseCase struct {
	service service.AuthorServiceInterface
}

func NewAuthorUseCase(service service.AuthorServiceInterface) *AuthorUseCase {
	return &AuthorUseCase{service}
}

func (a *AuthorUseCase) CreateAuthorUseCase(author *model.Author) error {
	return a.service.CreateAuthorService(author)
}

func (a *AuthorUseCase) GetAuthorsUseCase() ([]model.Author, error) {
	return a.service.GetAuthorsService()
}

func (a *AuthorUseCase) GetAuthorUseCase(id uint64) (*model.Author, error) {
	return a.service.GetAuthorService(id)
}

func (a *AuthorUseCase) UpdateAuthorUseCase(id uint64, patch *model.Author) (*model.Author, error) {
	return a.service.UpdateAuthorService(id, patch)
}
