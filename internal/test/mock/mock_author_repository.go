package mock

import (
	"library/internal/domain/model"

	"github.com/stretchr/testify/mock"
)

type AuthorRepoMock struct {
	mock.Mock
}

func (m *AuthorRepoMock) Save(author *model.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

func (m *AuthorRepoMock) FindAll() ([]model.Author, error) {
	args := m.Called()
	return args.Get(0).([]model.Author), args.Error(1)
}

func (m *AuthorRepoMock) FindById(id uint64) (*model.Author, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Author), args.Error(1)
}

func (m *AuthorRepoMock) Update(id uint64, patch *model.Author) (*model.Author, error) {
	args := m.Called(id, patch)
	return args.Get(0).(*model.Author), args.Error(1)
}
