package builder

import (
	"library/internal/domain/model"

	"github.com/google/uuid"
)

type BookAuthorBuilder struct {
	bookID   uuid.UUID
	authorID uuid.UUID
}

func NewBookAuthorBuilder() *BookAuthorBuilder {
	return &BookAuthorBuilder{bookID: uuid.New(), authorID: uuid.New()}
}

func (b *BookAuthorBuilder) Build() *model.BookAuthor {
	return &model.BookAuthor{
		BookID:   b.bookID,
		AuthorID: b.authorID,
	}
}
