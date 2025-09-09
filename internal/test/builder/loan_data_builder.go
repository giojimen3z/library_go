package builder

import (
	"time"

	"github.com/google/uuid"

	"library/internal/domain/model"
)

type LoanBuilder struct {
	id         uuid.UUID
	memberID   uuid.UUID
	copyID     uuid.UUID
	loanedAt   *time.Time
	dueDate    *time.Time
	returnedAt *time.Time
	fineCents  int
	createdAt  *time.Time
}

func NewLoanBuilder() *LoanBuilder {
	now := time.Now()
	due := now.AddDate(0, 0, 14)
	return &LoanBuilder{
		id:         uuid.New(),
		memberID:   uuid.New(),
		copyID:     uuid.New(),
		loanedAt:   &now,
		dueDate:    &due,
		returnedAt: nil,
		fineCents:  0,
		createdAt:  &now,
	}
}

func (b *LoanBuilder) Build() *model.Loan {
	return &model.Loan{
		ID:         b.id,
		MemberID:   b.memberID,
		CopyID:     b.copyID,
		LoanedAt:   *b.loanedAt,
		DueDate:    *b.dueDate,
		ReturnedAt: b.returnedAt,
		FineCents:  b.fineCents,
		CreatedAt:  *b.createdAt,
	}
}
