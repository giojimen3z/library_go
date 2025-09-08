package application

import (
	"library/internal/domain/model"
	"library/internal/domain/service"

	"github.com/google/uuid"
)

type BookUseCaseInterface interface {
	CreateBookUseCase(book *model.Book) (*model.Book, error)
	GetBooksUseCase() ([]model.Book, error)
	GetBookByIdUseCase(id uuid.UUID) (*model.Book, error)
	UpdateBookUseCase(id uuid.UUID, patch *model.Book) (*model.Book, error)
}

type BookUseCase struct {
	service service.BookServiceInterface
}

func NewBookUseCase(srv service.BookServiceInterface) *BookUseCase {
	return &BookUseCase{service: srv}
}

func (b *BookUseCase) CreateBookUseCase(book *model.Book) (*model.Book, error) {
	return b.service.CreateBook(book)
}

func (b *BookUseCase) GetBooksUseCase() ([]model.Book, error) {
	return b.service.GetAllBook()
}
func (b *BookUseCase) GetBookByIdUseCase(id uuid.UUID) (*model.Book, error) {
	return b.service.GetByIDBook(id)
}

func (b *BookUseCase) UpdateBookUseCase(id uuid.UUID, patch *model.Book) (*model.Book, error) {
	return b.service.UpdateBook(id, patch)
}
