package port

import "library/internal/domain/model"

type AuthorPort interface {
	SaveRepository(author *model.Author) error
	FindAllRepository() ([]model.Author, error)
	FindByIdRepository(id uint64) (*model.Author, error)
	UpdateRepository(id uint64, patch *model.Author) (*model.Author, error)
}
