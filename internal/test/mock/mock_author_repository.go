package mock

import (
	"library/internal/domain/model"

	"github.com/google/uuid"
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
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Author), args.Error(1)
}

func (m *AuthorRepoMock) FindById(id uuid.UUID) (*model.Author, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Author), args.Error(1)
}

func (m *AuthorRepoMock) Update(id uuid.UUID, patch *model.Author) (*model.Author, error) {
	args := m.Called(id, patch)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Author), args.Error(1)
}
