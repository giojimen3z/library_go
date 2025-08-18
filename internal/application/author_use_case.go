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

// Métodos que exponen la funcionalidad al Controller
func (a *AuthorUseCase) CreateAuthor(author *model.Author) error {
	return a.service.CreateAuthor(author)
}

func (a *AuthorUseCase) GetAuthors() ([]model.Author, error) {
	return a.service.GetAuthors()
}

func (a *AuthorUseCase) GetAuthor(id uint64) (*model.Author, error) {
	return a.service.GetAuthor(id)
}

func (a *AuthorUseCase) UpdateAuthor(id uint64, patch *model.Author) (*model.Author, error) {
	return a.service.UpdateAuthor(id, patch)
}
