package mock

import (
	"context"
	"library/internal/domain/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type LoanRepoMock struct {
	mock.Mock
}

func (m *LoanRepoMock) Save(ctx context.Context, loan *model.Loan) error {
	args := m.Called(ctx, loan)
	return args.Error(0)
}

func (m *LoanRepoMock) FindAll(ctx context.Context) ([]model.Loan, error) {
	args := m.Called()
	return args.Get(0).([]model.Loan), args.Error(1)
}

func (m *LoanRepoMock) FindById(ctx context.Context, id uuid.UUID) (*model.Loan, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Loan), args.Error(1)
}

func (m *LoanRepoMock) FindByMemeberId(ctx context.Context, id uuid.UUID) ([]model.Loan, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]model.Loan), args.Error(1)
}

func (m *LoanRepoMock) Update(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error) {
	args := m.Called(ctx, id, patch)
	return args.Get(0).(*model.Loan), args.Error(1)
}

func (m *LoanRepoMock) UpdateFineCents(ctx context.Context, id uuid.UUID, fineCents int) (*model.Loan, error) {
	args := m.Called(ctx, id, fineCents)
	return args.Get(0).(*model.Loan), args.Error(1)
}

func (m *LoanRepoMock) CantLoansByMember(ctx context.Context, id uuid.UUID) (int64, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(int64), args.Error(1)
}
