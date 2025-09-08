package service

import (
	"library/internal/domain/model"
	"library/internal/domain/port"

	"github.com/google/uuid"
)

type BookServiceInterface interface {
	CreateBook(book *model.Book) (*model.Book, error)
	GetAllBook() ([]model.Book, error)
	GetByIDBook(id uuid.UUID) (*model.Book, error)
	UpdateBook(id uuid.UUID, patch *model.Book) (*model.Book, error)
	DeleteBook(id uuid.UUID) error
}

type BookService struct {
	repo port.BookPort
}

func NewBookService(repo port.BookPort) *BookService { return &BookService{repo: repo} }

func (s *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	book.ID = uuid.New()
	if err := s.repo.Save(book); err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) GetAllBook() ([]model.Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) GetByIDBook(id uuid.UUID) (*model.Book, error) {
	return s.repo.FindById(id)
}

func (s *BookService) UpdateBook(id uuid.UUID, patch *model.Book) (*model.Book, error) {
	if _, err := s.repo.FindById(id); err != nil {
		return nil, err
	}
	if _, err := s.repo.Update(id, patch); err != nil {
		return nil, err
	}
	return s.repo.FindById(id)
}

func (s *BookService) DeleteBook(id uuid.UUID) error {
	return s.repo.Delete(id)
}
