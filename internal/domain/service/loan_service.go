package service

import (
	"context"
	"errors"
	"library/internal/domain/model"
	"library/internal/domain/port"
	"log/slog"

	"github.com/google/uuid"
)

type LoanServiceInterface interface {
	CreateLoan(ctx context.Context, loan *model.Loan) error
	GetLoans(ctx context.Context) ([]model.Loan, error)
	GetLoanById(ctx context.Context, id uuid.UUID) (*model.Loan, error)
	GetLoanByMemeberId(ctx context.Context, id uuid.UUID) ([]model.Loan, error)
	UpdateLoan(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error)
}

type LoanService struct {
	repo            port.LoanPort
	memberService   MemberService
	bookCopyService BookCopyService
}

func NewLoanService(repo port.LoanPort, memeberService *MemberService, bookCopyService *BookCopyService) *LoanService {
	return &LoanService{
		repo:            repo,
		memberService:   *memeberService,
		bookCopyService: *bookCopyService,
	}
}

func (l *LoanService) CreateLoan(ctx context.Context, loan *model.Loan) error {
	count, err := l.repo.CantLoansByMember(ctx, loan.MemberID)
	if err != nil {
		textError := "failed to count loans"
		slog.Error(textError)
		return errors.New(textError)
	}
	if count >= 5 {
		textError := "Max amount of allow loans to this members exceeded"
		slog.Error(textError)
		return errors.New(textError)
	}

	isActive, err := l.memberService.IsActive(loan.MemberID)
	if err != nil {
		return err
	}
	if !isActive {
		textError := "This members is not active"
		slog.Error(textError)
		return errors.New(textError)
	}

	isAvailable, err := l.bookCopyService.IsAvailable(loan.CopyID)
	if err != nil {
		return err
	}
	if !isAvailable {
		textError := "This book is unavailable in this moment"
		slog.Error(textError)
		return errors.New(textError)
	}

	loans, err := l.GetLoanByMemeberId(ctx, loan.MemberID)
	if err != nil {
		textError := "It was not possible evaluate the member"
		slog.Error(textError)
		return errors.New(textError)
	}

	for _, loanInternal := range loans {
		if loanInternal.CopyID == loan.CopyID {
			textError := "There is a loan with the same book copy"
			slog.Error(textError)
			return errors.New(textError)
		}
	}

	id := uuid.New()
	loan.ID = id
	if err := l.repo.Save(ctx, loan); err != nil {
		return err
	}

	err = l.bookCopyService.DecrementQuantity(loan.CopyID)
	if err != nil {
		return err
	}

	return nil
}

func (l *LoanService) GetLoans(ctx context.Context) ([]model.Loan, error) {
	return l.repo.FindAll(ctx)
}

func (l *LoanService) GetLoanById(ctx context.Context, id uuid.UUID) (*model.Loan, error) {
	return l.repo.FindById(ctx, id)
}

func (l *LoanService) GetLoanByMemeberId(ctx context.Context, id uuid.UUID) ([]model.Loan, error) {
	return l.repo.FindByMemeberId(ctx, id)
}

func (l *LoanService) UpdateLoan(ctx context.Context, id uuid.UUID, patch *model.Loan) (*model.Loan, error) {
	return l.repo.Update(ctx, id, patch)
}
