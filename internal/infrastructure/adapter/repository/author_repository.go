package repository

import (
	"fmt"
	"library/internal/domain/model"
	"library/internal/domain/port"

	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) port.AuthorPort {
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
	if err != nil {
		fmt.Printf("Failed to find authors: %v", err)
		return nil, err
	}

	return authors, err
}

func (r *AuthorRepositoryImpl) FindById(id uint64) (*model.Author, error) {
	var author model.Author
	err := r.db.First(&author, id).Error
	if err != nil {
		fmt.Printf("Failed to find author: %v", err)
		return nil, err
	}
	return &author, err
}

func (r *AuthorRepositoryImpl) Update(id uint64, patch *model.Author) (*model.Author, error) {
	var existing model.Author
	if err := r.db.First(&existing, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&existing).Updates(patch).Error; err != nil {
		return nil, err
	}

	updated, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
