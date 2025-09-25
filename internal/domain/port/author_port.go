package port

import (
	"context"

	"github.com/google/uuid"

	"library/internal/domain/model"
)

type AuthorPort interface {
	Save(ctx context.Context, author *model.Author) error
	FindAll(ctx context.Context) ([]model.Author, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Author, error)
	Update(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error)
}
