package port

import (
	"library/internal/domain/model"

	"github.com/google/uuid"
)

type BookPort interface {
	Save(book *model.Book) error
	FindAll() ([]model.Book, error)
	FindById(id uuid.UUID) (*model.Book, error)
	Update(id uuid.UUID, book *model.Book) (*model.Book, error)
	Delete(id uuid.UUID) error
}
