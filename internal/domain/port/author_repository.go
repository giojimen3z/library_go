package port

import "hexagonal_base/cmd/api/app/domain/model"

type AuthorRepository interface {
	Save(author *model.Author) error
	FindAll() ([]model.Author, error)
	FindById(id uint64) (*model.Author, error)
}
