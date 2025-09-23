package port

import (
	"context"
	"library/internal/domain/model"

	"github.com/google/uuid"
)

type LoanPort interface {
	Save(ctx context.Context, loan *model.Loan) error
	FindAll(ctx context.Context) ([]model.Loan, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Loan, error)
	FindByMemeberId(ctx context.Context, id uuid.UUID) ([]model.Loan, error)
	Update(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error)
	UpdateFineCents(ctx context.Context, id uuid.UUID, fineCents int) (*model.Loan, error)
	CantLoansByMember(ctx context.Context, id uuid.UUID) (int64, error)
}
