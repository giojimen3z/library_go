package port

import "library/internal/domain/model"

// AuthorPort define all methods that must to implemented by repository.
type AuthorPort interface {
	Save(author *model.Author) error
	FindAll() ([]model.Author, error)
	FindById(id uint64) (*model.Author, error)
	Update(id uint64, patch *model.Author) (*model.Author, error)
}
