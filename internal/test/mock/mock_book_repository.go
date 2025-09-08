package mock

import (
	"library/internal/domain/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookRepoMock struct {
	mock.Mock
}

func (m *BookRepoMock) Save(book *model.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *BookRepoMock) FindAll() ([]model.Book, error) {
	args := m.Called()

	var list []model.Book
	if v := args.Get(0); v != nil {
		list = v.([]model.Book)
	}
	return list, args.Error(1)
}

func (m *BookRepoMock) FindById(id uuid.UUID) (*model.Book, error) {
	args := m.Called(id)

	var b *model.Book
	if v := args.Get(0); v != nil {
		b = v.(*model.Book)
	}
	return b, args.Error(1)
}

func (m *BookRepoMock) Update(id uuid.UUID, book *model.Book) (*model.Book, error) {
	args := m.Called(id, book)

	var updated *model.Book
	if v := args.Get(0); v != nil {
		updated = v.(*model.Book)
	}
	return updated, args.Error(1)
}

func (m *BookRepoMock) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
