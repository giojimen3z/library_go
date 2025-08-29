package port

import (
	"context"
	"library/internal/domain/model"

	"github.com/google/uuid"
)

type AuthorPort interface {
	Save(ctx context.Context, author *model.Author) error
	FindAll(ctx context.Context) ([]model.Author, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Author, error)
	Update(ctx context.Context, id uuid.UUID, patch *model.Author) (*model.Author, error)
}
