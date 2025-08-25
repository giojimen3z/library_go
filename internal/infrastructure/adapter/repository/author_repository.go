package repository

import (
	"log/slog"

	"library/internal/domain/model"
	"library/internal/domain/port"

	"github.com/google/uuid"
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
	slog.Info("Author save successful", "author", author)
	return nil
}

func (r *AuthorRepositoryImpl) FindAll() ([]model.Author, error) {
	var authors []model.Author
	err := r.db.Find(&authors).Error
	if err != nil {
		slog.Error("Failed to find authors", "error", err)
		return nil, err
	}

	return authors, err
}

func (r *AuthorRepositoryImpl) FindById(id uuid.UUID) (*model.Author, error) {
	var author model.Author
	err := r.db.First(&author, id).Error
	if err != nil {
		slog.Error("Failed to find author", "error", err)
		return nil, err
	}
	return &author, err
}

func (r *AuthorRepositoryImpl) Update(id uuid.UUID, patch *model.Author) (*model.Author, error) {
	if err := r.db.Model(&model.Author{}).Where("id = ?", id).Updates(patch).Error; err != nil {
		slog.Error("Failed to update author", "error", err)
		return nil, err
	}
	return nil, nil
}
