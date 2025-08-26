package builder

import (
	"time"

	"library/internal/domain/model"

	"github.com/google/uuid"
)

type BookBuilder struct {
	id            uuid.UUID
	title         string
	isbn          string
	description   string
	publishedYear int
	createdAt     *time.Time
	updatedAt     *time.Time
}

func NewBookBuilder() *BookBuilder {
	date := time.Now()
	return &BookBuilder{
		id:            uuid.New(),
		title:         "A Great Book",
		isbn:          "978-1-23456-789-7",
		description:   "Sample description",
		publishedYear: date.Year(),
		createdAt:     &date,
		updatedAt:     &date,
	}
}

func (b *BookBuilder) WithID(id uuid.UUID) *BookBuilder         { b.id = id; return b }
func (b *BookBuilder) WithTitle(t string) *BookBuilder          { b.title = t; return b }
func (b *BookBuilder) WithISBN(s string) *BookBuilder           { b.isbn = s; return b }
func (b *BookBuilder) WithDescription(d string) *BookBuilder    { b.description = d; return b }
func (b *BookBuilder) WithPublishedYear(y int) *BookBuilder     { b.publishedYear = y; return b }
func (b *BookBuilder) WithCreatedAt(t time.Time) *BookBuilder   { b.createdAt = &t; return b }
func (b *BookBuilder) WithUpdatedAt(t time.Time) *BookBuilder   { b.updatedAt = &t; return b }

func (b *BookBuilder) Build() *model.Book {
	return &model.Book{
		ID:            b.id,
		Title:         b.title,
		ISBN:          b.isbn,
		Description:   b.description,
		PublishedYear: b.publishedYear,
		CreatedAt:     *b.createdAt,
		UpdatedAt:     *b.updatedAt,
	}
}
