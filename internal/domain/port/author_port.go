package port

import (
	"library/internal/domain/model"

	"github.com/google/uuid"
)

type AuthorPort interface {
	Save(author *model.Author) error
	FindAll() ([]model.Author, error)
	FindById(id uuid.UUID) (*model.Author, error)
	Update(id uuid.UUID, patch *model.Author) (*model.Author, error)
}
