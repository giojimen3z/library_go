package builder

import (
	"time"

	"github.com/google/uuid"

	"library/internal/domain/model"
)

type ReservationBuilder struct {
	id              uuid.UUID
	memberID        uuid.UUID
	bookID          uuid.UUID
	reservedAt      *time.Time
	canceledAt      *time.Time
	fulfilledLoanID *uuid.UUID
	createdAt       *time.Time
}

func NewReservationBuilder() *ReservationBuilder {
	now := time.Now()
	return &ReservationBuilder{
		id:              uuid.New(),
		memberID:        uuid.New(),
		bookID:          uuid.New(),
		reservedAt:      &now,
		canceledAt:      nil,
		fulfilledLoanID: nil,
		createdAt:       &now,
	}
}

func (b *ReservationBuilder) Build() *model.Reservation {
	return &model.Reservation{
		ID:              b.id,
		MemberID:        b.memberID,
		BookID:          b.bookID,
		ReservedAt:      *b.reservedAt,
		CanceledAt:      b.canceledAt,
		FulfilledLoanID: b.fulfilledLoanID,
		CreatedAt:       *b.createdAt,
	}
}
