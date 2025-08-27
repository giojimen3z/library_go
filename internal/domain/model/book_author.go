package model

import "github.com/google/uuid"

// BookAuthor represents the pivot table book_authors for many-to-many between books and authors
// Primary key is composite (book_id, author_id)

type BookAuthor struct {
	BookID   uuid.UUID `json:"book_id" gorm:"type:uuid;primaryKey"`
	AuthorID uuid.UUID `json:"author_id" gorm:"type:uuid;primaryKey"`
}

func (BookAuthor) TableName() string { return "book_authors" }
