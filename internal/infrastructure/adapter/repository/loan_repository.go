package repository

import (
	"context"
	"library/internal/domain/model"
	"library/internal/domain/port"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoanRepositoryImpl struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) port.LoanPort {
	return &LoanRepositoryImpl{db: db}
}

func (l *LoanRepositoryImpl) Save(ctx context.Context, loan *model.Loan) error {
	err := l.db.WithContext(ctx).Create(loan).Error
	if err != nil {
		return err
	}
	slog.Info("Loan save successful", "loan", loan)
	return nil
}

func (l *LoanRepositoryImpl) FindAll(ctx context.Context) ([]model.Loan, error) {
	var loans []model.Loan
	err := l.db.WithContext(ctx).Find(&loans).Error
	if err != nil {
		slog.Error("Failed to find loans", "error", err)
		return nil, err
	}

	return loans, err
}

func (l *LoanRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*model.Loan, error) {
	var loan model.Loan
	err := l.db.WithContext(ctx).First(&loan, id).Error
	if err != nil {
		slog.Error("Failed to find loan", "error", err)
		return nil, err
	}

	return &loan, err
}

func (l *LoanRepositoryImpl) FindByMemeberId(ctx context.Context, id uuid.UUID) ([]model.Loan, error) {
	var loans []model.Loan
	err := l.db.WithContext(ctx).Model(&model.Loan{}).Where("member_id = ?", id).Find(&loans).Error
	if err != nil {
		slog.Error("Failed to find loan", "error", err)
		return nil, err
	}

	return loans, err
}

func (l *LoanRepositoryImpl) Update(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error) {
	if err := l.db.WithContext(ctx).Model(&model.Loan{}).Where("id = ?", id).Updates(patch).Error; err != nil {
		slog.Error("Failed to update loan", "error", err)
		return nil, err
	}

	return nil, nil
}

func (l *LoanRepositoryImpl) UpdateFineCents(ctx context.Context, id uuid.UUID, fineCents int) (*model.Loan, error) {

	if err := l.db.WithContext(ctx).Model(&model.Loan{}).Where("id = ?", id).Update("fine_cents", fineCents).Error; err != nil {
		slog.Error("Failed to update loan", "error", err)
		return nil, err
	}

	return nil, nil
}

func (l *LoanRepositoryImpl) CantLoansByMember(ctx context.Context, id uuid.UUID) (int64, error) {
	var count int64
	err := l.db.WithContext(ctx).Model(&model.Loan{}).Where("member_id = ?", id).Count(&count).Error
	if err != nil {
		slog.Error("Failed to find loan", "error", err)
		return 0, err
	}

	return count, err
}
