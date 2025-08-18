package repository

import (
	"fmt"
	"library/cmd/api/app/domain/model"
	"library/cmd/api/app/domain/port"

	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) port.AuthorRepository {
	return &AuthorRepositoryImpl{db: db}
}

func (r *AuthorRepositoryImpl) Save(author *model.Author) error {
	err := r.db.Create(author).Error
	if err != nil {
		return err
	}
	fmt.Printf("Author save successful: %+v\n", author)
	return nil
}

func (r *AuthorRepositoryImpl) FindAll() ([]model.Author, error) {
	var authors []model.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *AuthorRepositoryImpl) FindById(id uint64) (*model.Author, error) {
	var author model.Author
	err := r.db.First(&author, id).Error
	return &author, err
}
