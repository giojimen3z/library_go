package model

import (
	"time"

	"github.com/google/uuid"
)

// Member represents the members table
// Schema: library.members

type Member struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FullName  string    `json:"full_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
