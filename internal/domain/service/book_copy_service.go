package service

import "github.com/google/uuid"

type BookCopyServiceInterface interface {
	IsAvailable(id uuid.UUID) (bool, error)
	IncrementQuantity(id uuid.UUID) error
	DecrementQuantity(id uuid.UUID) error
}

type BookCopyService struct {
}

func NewBookCopyService() *BookCopyService {
	return &BookCopyService{}
}

func (m *BookCopyService) IsAvailable(id uuid.UUID) (bool, error) {
	return true, nil
}

func (m *BookCopyService) IncrementQuantity(id uuid.UUID) error {
	return nil
}

func (m *BookCopyService) DecrementQuantity(id uuid.UUID) error {
	return nil
}
