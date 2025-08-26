package model

import (
	"time"

	"github.com/google/uuid"
)

// Book represents the books table
// Schema: library.books
// CREATE TABLE books (
//   id UUID PK,
//   title TEXT NOT NULL,
//   isbn TEXT UNIQUE,
//   description TEXT,
//   published_year INT,
//   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
//   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
// )

type Book struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title         string    `json:"title" gorm:"not null"`
	ISBN          string    `json:"isbn" gorm:"uniqueIndex"`
	Description   string    `json:"description"`
	PublishedYear int       `json:"published_year"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relations
	Authors []Author `json:"authors,omitempty" gorm:"many2many:book_authors;"`
}
