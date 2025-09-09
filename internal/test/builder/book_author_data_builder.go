package builder

import (
	"github.com/google/uuid"

	"library/internal/domain/model"
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
