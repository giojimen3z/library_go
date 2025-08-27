package model

import (
	"time"

	"github.com/google/uuid"
)

// BookCopy represents the book_copies table
// Schema: library.book_copies

type BookCopy struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BookID    uuid.UUID `json:"book_id" gorm:"type:uuid;not null"`
	Book      *Book     `json:"book,omitempty" gorm:"foreignKey:BookID;references:ID;constraint:OnDelete:CASCADE"`
	Barcode   string    `json:"barcode" gorm:"uniqueIndex"`
	Condition string    `json:"condition"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
