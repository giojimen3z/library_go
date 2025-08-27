package model

import (
	"time"

	"github.com/google/uuid"
)

// Loan represents the loans table
// Schema: library.loans

type Loan struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	MemberID   uuid.UUID  `json:"member_id" gorm:"type:uuid;not null"`
	Member     *Member    `json:"member,omitempty" gorm:"foreignKey:MemberID;references:ID;constraint:OnDelete:RESTRICT"`
	CopyID     uuid.UUID  `json:"copy_id" gorm:"type:uuid;not null"`
	Copy       *BookCopy  `json:"copy,omitempty" gorm:"foreignKey:CopyID;references:ID;constraint:OnDelete:RESTRICT"`
	LoanedAt   time.Time  `json:"loaned_at" gorm:"autoCreateTime"`
	DueDate    time.Time  `json:"due_date"`
	ReturnedAt *time.Time `json:"returned_at"` // nullable
	FineCents  int        `json:"fine_cents" gorm:"not null;default:0"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
}
