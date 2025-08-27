package builder

import (
	"time"

	"library/internal/domain/model"

	"github.com/google/uuid"
)

type BookCopyBuilder struct {
	id        uuid.UUID
	bookID    uuid.UUID
	barcode   string
	condition string
	isActive  bool
	createdAt *time.Time
}

func NewBookCopyBuilder() *BookCopyBuilder {
	date := time.Now()
	return &BookCopyBuilder{
		id:        uuid.New(),
		bookID:    uuid.New(),
		barcode:   uuid.NewString(),
		condition: "NEW",
		isActive:  true,
		createdAt: &date,
	}
}

func (b *BookCopyBuilder) Build() *model.BookCopy {
	return &model.BookCopy{
		ID:        b.id,
		BookID:    b.bookID,
		Barcode:   b.barcode,
		Condition: b.condition,
		IsActive:  b.isActive,
		CreatedAt: *b.createdAt,
	}
}
