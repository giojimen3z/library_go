package mock

import (
	"library/internal/domain/model"

	"github.com/stretchr/testify/mock"
)

type MockAuthorRepo struct {
	mock.Mock
}

func (m *MockAuthorRepo) Save(author *model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *MockAuthorRepo) FindAll() ([]model.Author, error) {
	args := m.Called()
	return args.Get(0).([]model.Author), args.Error(1)
}

func (m *MockAuthorRepo) FindById(id uint64) (*model.Author, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Author), args.Error(1)
}

func (m *MockAuthorRepo) Update(id uint64, patch *model.Author) (*model.Author, error) {
	args := m.Called(id, patch)
	return args.Get(0).(*model.Author), args.Error(1)
}
