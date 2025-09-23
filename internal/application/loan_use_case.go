package application

import (
	"context"
	"library/internal/domain/model"
	"library/internal/domain/service"

	"github.com/google/uuid"
)

type LoanUseCaseInteface interface {
	CreateLoanUseCase(ctx context.Context, loan *model.Loan) error
	GetLoansUseCase(ctx context.Context) ([]model.Loan, error)
	GetLoanUseCase(ctx context.Context, id uuid.UUID) (*model.Loan, error)
	GetLoanByMemeberIdUseCase(ctx context.Context, id uuid.UUID) ([]model.Loan, error)
	UpdateLoanUseCase(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error)
}

type LoanUseCase struct {
	service service.LoanServiceInterface
}

func NewLoanUseCase(service service.LoanServiceInterface) *LoanUseCase {
	return &LoanUseCase{service}
}

func (l *LoanUseCase) CreateLoanUseCase(ctx context.Context, loan *model.Loan) error {
	return l.service.CreateLoan(ctx, loan)
}

func (l *LoanUseCase) GetLoansUseCase(ctx context.Context) ([]model.Loan, error) {
	return l.service.GetLoans(ctx)
}

func (l *LoanUseCase) GetLoanByMemeberIdUseCase(ctx context.Context, id uuid.UUID) ([]model.Loan, error) {
	return l.service.GetLoanByMemeberId(ctx, id)
}

func (l *LoanUseCase) GetLoanUseCase(ctx context.Context, id uuid.UUID) (*model.Loan, error) {
	return l.service.GetLoanById(ctx, id)
}

func (l *LoanUseCase) UpdateLoanUseCase(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error) {
	return l.service.UpdateLoan(ctx, id, patch)
}
