package model

import (
	"time"

	"github.com/google/uuid"
)

// Reservation represents the reservations table
// Schema: library.reservations

type Reservation struct {
	ID              uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	MemberID        uuid.UUID  `json:"member_id" gorm:"type:uuid;not null"`
	Member          *Member    `json:"member,omitempty" gorm:"foreignKey:MemberID;references:ID;constraint:OnDelete:RESTRICT"`
	BookID          uuid.UUID  `json:"book_id" gorm:"type:uuid;not null"`
	Book            *Book      `json:"book,omitempty" gorm:"foreignKey:BookID;references:ID;constraint:OnDelete:RESTRICT"`
	ReservedAt      time.Time  `json:"reserved_at" gorm:"autoCreateTime"`
	CanceledAt      *time.Time `json:"canceled_at"`                        // nullable
	FulfilledLoanID *uuid.UUID `json:"fulfilled_loan_id" gorm:"type:uuid"` // nullable
	FulfilledLoan   *Loan      `json:"fulfilled_loan,omitempty" gorm:"foreignKey:FulfilledLoanID;references:ID;constraint:OnDelete:SET NULL"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
}
